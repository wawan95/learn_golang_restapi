package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"golang-user-api/models"
	"golang-user-api/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	Service *services.UserService
}

func NewUserController(
	service *services.UserService,
) *UserController {
	return &UserController{Service: service}
}

// GET ALL USERS
func (ctrl *UserController) GetAll(c *gin.Context) {
	users, err := ctrl.Service.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"message": "Failed to get users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GET USER BY ID
func (ctrl *UserController) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest,
			gin.H{"message": "Invalid user ID"})
		return
	}

	user, err := ctrl.Service.GetByID(uint(id))

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound,
			gin.H{"message": "User not found"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"message": "Failed to get user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// CREATE USER
func (ctrl *UserController) Create(c *gin.Context) {
	var req models.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"message": "Invalid request", "error": err.Error()})
		return
	}

	user, err := ctrl.Service.Create(req)

	if err != nil {
		c.JSON(http.StatusConflict,
			gin.H{"message": "Could not create user. Check email uniqueness."})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// UPDATE USER
func (ctrl *UserController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest,
			gin.H{"message": "Invalid user ID"})
		return
	}

	var req models.UpdateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"message": "Invalid request", "error": err.Error()})
		return
	}

	user, err := ctrl.Service.Update(uint(id), req)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound,
			gin.H{"message": "User not found"})
		return
	}

	if err != nil {
		c.JSON(http.StatusConflict,
			gin.H{"message": "Could not update user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DELETE USER
func (ctrl *UserController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest,
			gin.H{"message": "Invalid user ID"})
		return
	}

	err = ctrl.Service.Delete(uint(id))

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound,
			gin.H{"message": "User not found"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"message": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK,
		gin.H{"message": "User deleted successfully"})
}
