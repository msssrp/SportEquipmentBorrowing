package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/msssrp/SportEquipmentBorrowing/pkg/app"
	"github.com/msssrp/SportEquipmentBorrowing/pkg/app/borrowing"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BorrowingHandler struct {
	app *app.App
}

func NewBorrowingHandler(app *app.App) *BorrowingHandler {
	return &BorrowingHandler{
		app: app,
	}
}

type BorrowingInput struct {
	User_id      primitive.ObjectID `json:"user_id"`
	Equipment_id primitive.ObjectID `json:"equipment_id" `
	Borrow_date  time.Time          `json:"borrow_date" `
	Return_date  time.Time          `json:"return_date" `
	DaysLeft     int                `json:"days_left" form:"-"`
	Status       string             `json:"status"`
}

type Approve struct {
	Id           primitive.ObjectID `json:"borrowing_id"`
	Equipment_id primitive.ObjectID `json:"equipment_id"`
}

//Get
func (h *BorrowingHandler) HandlerGetAllBorrowings(c *gin.Context) {

	borrowings, err := h.app.BorrowingService.GetAllBorrowings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, borrowings)
}

func (h *BorrowingHandler) HandlerGetBorrowingByID(c *gin.Context) {
	borrowingIDStr := c.Param("id")

	borrowingID, err := primitive.ObjectIDFromHex(borrowingIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid borrowing ID"})
		return
	}

	borrowing, err := h.app.BorrowingService.GetBorrowingByID(borrowingID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if borrowing == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Borrowing not found"})
		return
	}

	c.JSON(http.StatusOK, borrowing)
}

func (h *BorrowingHandler) HandlerGetBorrowingsByUserID(c *gin.Context) {
	userIDStr := c.Param("id")

	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	borrowings, err := h.app.BorrowingService.GetBorrowingsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, borrowings)
}

func (h *BorrowingHandler) HandlerGetBorrowingByEquipmentID(c *gin.Context) {
	equipmentIDStr := c.Param("id")

	equipmentID, err := primitive.ObjectIDFromHex(equipmentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	borrowing, err := h.app.BorrowingService.GetBorrowingByEquipmentID(equipmentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, borrowing)
}

//Post
func (h *BorrowingHandler) HandlerCreateBorrowing(c *gin.Context) {
	var borrowingInput BorrowingInput

	if err := c.ShouldBindJSON(&borrowingInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	daysLeft := int(borrowingInput.Return_date.Sub(borrowingInput.Borrow_date).Hours() / 24)

	borrowing := &borrowing.Borrowing{
		Id:           primitive.NewObjectID(),
		User_id:      borrowingInput.User_id,
		Equipment_id: borrowingInput.Equipment_id,
		Borrow_date:  borrowingInput.Borrow_date,
		Return_date:  borrowingInput.Return_date,
		DayLeft:      daysLeft,
		Status:       borrowingInput.Status,
	}

	err := h.app.BorrowingService.CreateBorrowing(borrowing)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.app.EquipmentService.UpdateQuantity_available(borrowing.Equipment_id, "pending")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, borrowing)
}

func (h *BorrowingHandler) HandlerApproveBorrowing(c *gin.Context) {
	var borrowingApproveInput Approve

	if err := c.ShouldBindJSON(&borrowingApproveInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.app.BorrowingService.ApproveBorrow(borrowingApproveInput.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.app.EquipmentService.UpdateQuantity_available(borrowingApproveInput.Equipment_id, "In use")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Borrowing Approved"})
}

//Put
func (h *BorrowingHandler) HandlerUpdateBorrowing(c *gin.Context) {
	borrowingIDStr := c.Param("id")

	borrowingID, err := primitive.ObjectIDFromHex(borrowingIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid borrowing ID"})
		return
	}

	var updateBorrowingInput BorrowingInput
	if err := c.ShouldBindJSON(&updateBorrowingInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	borrowing := &borrowing.Borrowing{
		Id:           borrowingID,
		User_id:      updateBorrowingInput.User_id,
		Equipment_id: updateBorrowingInput.Equipment_id,
		Borrow_date:  updateBorrowingInput.Borrow_date,
		Return_date:  updateBorrowingInput.Return_date,
		Status:       updateBorrowingInput.Status,
	}

	err = h.app.BorrowingService.UpdateBorrowing(borrowing)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, borrowing)
}

//Delete
func (h *BorrowingHandler) HandlerDeleteBorrowing(c *gin.Context) {
	borrowingIDStr := c.Param("id")
	equipmentIDStr := c.Param("equipmentID")
	borrowingID, err := primitive.ObjectIDFromHex(borrowingIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid borrowing ID"})
		return
	}

	equipmentID, err := primitive.ObjectIDFromHex(equipmentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid borrowing ID"})
		return
	}

	err = h.app.BorrowingService.DeleteBorrowingByID(borrowingID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.app.EquipmentService.UpdateQuantity_available(equipmentID, "available")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Borrowing deleted successfully"})
}
