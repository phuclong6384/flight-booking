package handler

import (
	"errors"
	"flightBooking/common/database"
	"flightBooking/common/util"
	"flightBooking/user/dao"
	"flightBooking/user/dto"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	v "github.com/go-playground/validator/v10"
)

type Handler struct {
	validator   *v.Validate
	userService dao.IUserService
}

func NewHandler(service dao.IUserService) Handler {
	return Handler{validator: v.New(), userService: service}
}

func (h *Handler) HandleHealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (h *Handler) HandleQuery(ctx *gin.Context) {
	username := ctx.Param("username")
	if len(username) < 1 {
		log.Println("no username path parameter provided")
		util.SendValidationError(ctx, errors.New("no username path parameter provided"))
		return
	}
	userFound, err := h.userService.GetByUsername(username)
	if err != nil {
		log.Println("Not found user with username: ", username)
		util.SendDefaultBadRequestError(ctx)
		return
	}
	userFound.Password = "******"
	ctx.JSON(http.StatusOK, userFound)
}

func (h *Handler) HandleRegister(ctx *gin.Context) {
	req := &dto.RegisterUserRequest{}
	err := ctx.ShouldBind(&req)
	if err != nil {
		log.Println("Error binding request", err)
		util.SendDefaultBadRequestError(ctx)
		return
	}
	if err = h.validator.Struct(req); err != nil {
		util.SendValidationError(ctx, err)
		return
	}

	encrypted, err := util.Encrypt(req.Password)
	if err != nil {
		log.Println("Encryption error occurred", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	create, err := h.userService.Create(&database.User{
		Username:  req.Username,
		Password:  encrypted,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Gender:    req.Gender,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		log.Println("Error occurred", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	create.Password = "******"
	ctx.JSON(http.StatusOK, create)
}

func (h *Handler) HandleUpdate(ctx *gin.Context) {
	req := dto.UpdateUserRequest{}
	err := ctx.ShouldBind(&req)
	if err != nil {
		log.Println("Error binding request", err)
		util.SendDefaultBadRequestError(ctx)
		return
	}
	username := ctx.Param("username")
	if len(username) < 1 {
		log.Println("no username path parameter provided")
		util.SendValidationError(ctx, errors.New("no username path parameter provided"))
		return
	}
	userFound, err := h.userService.GetByUsername(username)
	if err != nil {
		log.Println("Not found user with username: ", username)
		util.SendDefaultBadRequestError(ctx)
		return
	}
	userFound.FirstName = req.FirstName
	userFound.LastName = req.LastName
	userFound.Gender = req.Gender
	userFound.UpdatedAt = time.Now()
	updated, err := h.userService.Update(userFound)
	if err != nil {
		log.Println("Error occurred when update user")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	updated.Password = "*****"
	ctx.JSON(http.StatusOK, updated)
}

func (h *Handler) ValidatePassword(ctx *gin.Context) {
	req := dto.ValidatePasswordRequest{}
	err := ctx.ShouldBind(&req)
	if err != nil {
		log.Println("Error binding request", err)
		util.SendDefaultBadRequestError(ctx)
		return
	}
	username := req.Username
	userFound, err := h.userService.GetByUsername(username)
	if err != nil {
		log.Println("Not found user with username: ", username)
		util.SendDefaultBadRequestError(ctx)
		return
	}
	success := h.userService.ValidatePassword(userFound, req.Password)
	ctx.JSON(http.StatusOK, gin.H{
		"success": success,
	})
}
