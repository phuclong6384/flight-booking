package handler

import (
	"flightBooking/booking/client"
	"flightBooking/booking/dao"
	"flightBooking/booking/dto"
	"flightBooking/common/database"
	"flightBooking/common/util"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	v "github.com/go-playground/validator/v10"
)

type Handler struct {
	validator      *v.Validate
	bookingService dao.IBookingService
	grpcClient     client.GrpcClient
}

func NewHandler(service dao.IBookingService) Handler {
	return Handler{validator: v.New(), bookingService: service, grpcClient: client.NewUserGrpcClient()}
}

func (h *Handler) HandleHealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (h *Handler) HandleReserveBooking(ctx *gin.Context) {
	req := dto.BookFlightRequest{}
	err := ctx.ShouldBind(&req)
	if err != nil {
		log.Println("Error binding request", err)
		util.SendDefaultBadRequestError(ctx)
		return
	}
	err = h.validator.Struct(&req)
	if err != nil {
		util.SendValidationError(ctx, err)
		return
	}

	success := h.grpcClient.ValidatePasswordUserService(req.Username, req.Password)
	if !success {
		log.Println("Invalid cred", err)
		ctx.AbortWithStatusJSON(http.StatusForbidden, err)
		return
	}

	detail, err := h.grpcClient.GetFlightDetail(req.FlightId)
	if err != nil {
		log.Println("Error get result from grpc flight service", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	reservedFlights, err := h.bookingService.GetByFlightId(req.FlightId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	seatsAllocated := 0
	for _, f := range reservedFlights {
		seatsAllocated += f.ReservedSlot
	}
	if req.NumberOfSlot+seatsAllocated > int(detail.TotalSlot) {
		log.Printf("Cannot allocate %v seats; Now: %v/%v", req.NumberOfSlot, seatsAllocated, detail.TotalSlot)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}
	create, err := h.bookingService.Create(&database.Booking{
		CustomerUsername: req.Username,
		FlightId:         req.FlightId,
		Code:             detail.Code,
		Status:           database.BookingStatusCreated,
		ReservedSlot:     req.NumberOfSlot,
		BookedDate:       time.Now(),
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	})
	if err != nil {
		log.Println("Cannot save booking", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, create)
}

func (h *Handler) HandlerGetListBooking(ctx *gin.Context) {
	req := dto.GetListBookingRequest{}
	err := ctx.ShouldBind(&req)
	if err != nil {
		log.Println("Error binding request", err)
		util.SendDefaultBadRequestError(ctx)
		return
	}
	err = h.validator.Struct(&req)
	if err != nil {
		util.SendValidationError(ctx, err)
		return
	}
	list, err := h.bookingService.GetList(req.Username)
	if err != nil {
		log.Println("Error occurred", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, list)
}

func (h *Handler) HandleGetBookingById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println("ID cannot be parsed", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}
	byId, err := h.bookingService.GetById(id)
	if err != nil {
		log.Println("Error getting booking by ID", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, byId)
}
