package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/msssrp/SportEquipmentBorrowing/pkg/app"

	"github.com/msssrp/SportEquipmentBorrowing/pkg/app/equipment"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EquipmentHandler struct {
	app *app.App
}

func NewEquipmentHandler(app *app.App) *EquipmentHandler {
	return &EquipmentHandler{
		app: app,
	}
}

type EquipmentInput struct {
	Name               string `json:"name"`
	Category           string `json:"category"`
	Description        string `json:"description"`
	Quantity_available string `json:"quantity_available"`
	Condition          string `json:"condition"`
	Image_url          string `json:"image_url"`
}

//Get
func (h *EquipmentHandler) HandlerGetEquipments(c *gin.Context) {
	equipments, err := h.app.EquipmentService.GetAllEquipments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []gin.H

	// Loop through each equipment
	for _, equipment := range equipments {
		// Get borrowing information by equipment ID
		borrowing, err := h.app.BorrowingService.GetBorrowingByEquipmentID(equipment.Id)
		if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
			// Handle errors other than "not found"
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Prepare the equipment response
		equipmentResponse := gin.H{
			"equipment": equipment,
		}

		// If borrowing is found, include it in the response
		if borrowing != nil {
			equipmentResponse["borrowing"] = borrowing
		}

		// Append the equipment response to the overall response slice
		response = append(response, equipmentResponse)
	}

	c.JSON(http.StatusOK, response)
}

func (h *EquipmentHandler) HandlerGetEquipmentByID(c *gin.Context) {
	equipmentIDsrt := c.Param("id")

	equipmentID, err := primitive.ObjectIDFromHex(equipmentIDsrt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Equipment ID"})
	}
	equipment, err := h.app.EquipmentService.GetEquipmentByID(equipmentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, equipment)
}

//Post
func (h *EquipmentHandler) HandlerCreateEquipment(c *gin.Context) {
	var equipmentInput EquipmentInput

	if err := c.ShouldBindJSON(&equipmentInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	equipment := &equipment.Equipment{
		Name:               equipmentInput.Name,
		Category:           equipmentInput.Category,
		Description:        equipmentInput.Description,
		Quantity_available: equipmentInput.Quantity_available,
		Condition:          equipmentInput.Condition,
		Image_url:          equipmentInput.Image_url,
	}

	err := h.app.EquipmentService.CreateEquipment(equipment)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Equipment created successfully"})

}

//Put
func (h *EquipmentHandler) HandlerUpdateEquipment(c *gin.Context) {
	equipmentIDsrt := c.Param("id")

	equipmentID, err := primitive.ObjectIDFromHex(equipmentIDsrt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var updateEquipementInput EquipmentInput
	if err := c.ShouldBindJSON(&updateEquipementInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	equipment := &equipment.Equipment{
		Id:                 equipmentID,
		Name:               updateEquipementInput.Name,
		Category:           updateEquipementInput.Category,
		Description:        updateEquipementInput.Description,
		Quantity_available: updateEquipementInput.Quantity_available,
		Condition:          updateEquipementInput.Condition,
		Image_url:          updateEquipementInput.Image_url,
	}

	err = h.app.EquipmentService.UpdateEquipment(equipment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, equipment)
}

//Delete
func (h *EquipmentHandler) HandlerDeleteEquipment(c *gin.Context) {
	equipmentIDStr := c.Param("id")

	equipmentID, err := primitive.ObjectIDFromHex(equipmentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid equipment ID"})
		return
	}

	err = h.app.EquipmentService.DeleteEquipment(equipmentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Equipment deleted successfully"})
}
