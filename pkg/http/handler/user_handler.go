package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/msssrp/SportEquipmentBorrowing/pkg/app"
	"github.com/msssrp/SportEquipmentBorrowing/pkg/app/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type UserHandler struct {
	app *app.App
}

func NewUserHandler(app *app.App) *UserHandler {
	return &UserHandler{
		app: app,
	}
}

type UserInput struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email" `
	F_name       string `json:"f_name"`
	L_name       string `json:"l_name"`
	Phone_number string `json:"phone_number"`
	Address      string `json:"address"`
}

func (h *UserHandler) HandlerGetUsers(c *gin.Context) {
	users, err := h.app.UserService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) HandlerGetUserByID(c *gin.Context) {
	userIDsrt := c.Param("id")

	userID, err := primitive.ObjectIDFromHex(userIDsrt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
	}
	user, err := h.app.UserService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) HandlerCreateUser(c *gin.Context) {
	var userInput UserInput

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &user.User{
		Username:     userInput.Username,
		Password:     userInput.Password,
		Email:        userInput.Email,
		F_name:       userInput.F_name,
		L_name:       userInput.L_name,
		Phone_number: userInput.Phone_number,
		Address:      userInput.Address,
	}

	err := h.app.UserService.CreateUser(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})

}

func (h *UserHandler) HandlerUpdateUser(c *gin.Context) {
	userIDStr := c.Param("id")

	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var updateUserInput UserInput
	if err := c.ShouldBindJSON(&updateUserInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &user.User{
		Id:           userID,
		Username:     updateUserInput.Username,
		Password:     updateUserInput.Password,
		Email:        updateUserInput.Email,
		F_name:       updateUserInput.F_name,
		L_name:       updateUserInput.L_name,
		Phone_number: updateUserInput.Phone_number,
		Address:      updateUserInput.Address,
	}

	err = h.app.UserService.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) HandlerDeleteUser(c *gin.Context) {
	userIDStr := c.Param("id")

	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = h.app.UserService.DeleteUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
