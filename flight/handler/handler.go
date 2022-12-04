package handler

import (
	"errors"
	"flightBooking/common/util"
	"flightBooking/flight/dao"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	v "github.com/go-playground/validator/v10"
)

type Handler struct {
	validator     *v.Validate
	flightService dao.IFlightService
}

func NewHandler(service dao.IFlightService) Handler {
	return Handler{validator: v.New(), flightService: service}
}

func (h *Handler) HandleHealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (h *Handler) HandleGetList(ctx *gin.Context) {
	list, err := h.flightService.GetList()
	if err != nil {
		log.Println("Error occurred", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, list)
}

func (h *Handler) HandleGetByCode(ctx *gin.Context) {
	code := ctx.Param("code")
	if len(code) < 1 {
		log.Println("no username path code provided")
		util.SendValidationError(ctx, errors.New("no username path code provided"))
		return
	}
	result, err := h.flightService.Get(code)
	if err != nil {
		log.Println("Error occurred", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}
