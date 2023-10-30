package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/msssrp/SportEquipmentBorrowing/pkg/app"
	"github.com/msssrp/SportEquipmentBorrowing/pkg/app/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	Username     string   `json:"username"`
	Password     string   `json:"password"`
	Email        string   `json:"email" `
	F_name       string   `json:"f_name"`
	L_name       string   `json:"l_name"`
	Phone_number string   `json:"phone_number"`
	Address      string   `json:"address"`
	Roles        []string `json:"roles"`
}

type UserSignInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//Get
func (h *UserHandler) HandlerGetUsers(c *gin.Context) {
	users, err := h.app.UserService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) HandlerGetUserByIDFromToken(c *gin.Context) {
	userID, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cant get user id on token"})
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user ID is not a string"})
		return
	}

	userIDHex, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		fmt.Println("Error converting to ObjectID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	user, err := h.app.UserService.GetUserByID(userIDHex)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) HandlerGetUserRolesByID(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}
	roles, err := h.app.UserService.GetUserRolesById(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"roles": roles})
}

func (h *UserHandler) HandlersGetUserByID(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := primitive.ObjectIDFromHex(userIDStr)
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

func (h *UserHandler) HandlerVerifySession(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "successfully authenticate"})
}

//Post
func (h *UserHandler) HandlerCreateUser(c *gin.Context) {
	var userInput UserInput

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(userInput.Roles) == 0 {
		userInput.Roles = []string{"user"}
	}

	user := &user.User{
		Username:     userInput.Username,
		Password:     userInput.Password,
		Email:        userInput.Email,
		F_name:       userInput.F_name,
		L_name:       userInput.L_name,
		Phone_number: userInput.Phone_number,
		Address:      userInput.Address,
		Roles:        userInput.Roles,
	}

	err := h.app.UserService.CreateUser(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (h *UserHandler) HandlerSignIn(c *gin.Context) {
	var signInInput UserSignInInput
	if err := c.ShouldBindJSON(&signInInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := h.app.UserService.UserSignIn(signInInput.Username, signInInput.Password)
	if err != nil {
		var errInvalidCredentials = errors.New("Invalid username or password")
		if err == errInvalidCredentials {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// Include both access and refresh tokens in the response
	c.JSON(http.StatusOK, gin.H{"access_token": accessToken, "refresh_token": refreshToken})
}

func (h *UserHandler) HanlderNewAccessToken(c *gin.Context) {
	userID, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cant get user id on token"})
		return
	}
	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user ID is not a string"})
		return
	}
	newAccesToken, err := h.app.UserService.NewAccessToken(userIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"newAccessToken": newAccesToken})
}

//Put
func (h *UserHandler) HandlerUpdateUser(c *gin.Context) {

	userIDsrt := c.Param("id")

	userID, err := primitive.ObjectIDFromHex(userIDsrt)
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

//Delete
func (h *UserHandler) HandlerDeleteUser(c *gin.Context) {

	userIDsrt := c.Param("id")

	userID, err := primitive.ObjectIDFromHex(userIDsrt)
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
