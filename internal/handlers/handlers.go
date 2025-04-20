package handlers

import (
	"awesomeProject2/internal/service"
	"awesomeProject2/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	_ "strings"
)

type ItemHandler struct {
	Service *service.ItemService
}

func NewItemHandler(service *service.ItemService) *ItemHandler {
	return &ItemHandler{Service: service}
}

// CreateItem godoc
// @Summary Create a new item
// @Description Create a new item
// @Tags items
// @Accept json
// @Produce json
// @Param item body models.Item true "Item object to be created"
// @Success 201 {object} models.Item
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /items [post]
func (h *ItemHandler) CreateItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.CreateItem(c.Request.Context(), &item); err != nil {
		log.Printf("❌ Error creating item: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item"})
		return
	}
	c.JSON(http.StatusCreated, item)
}

// GetItemByID godoc
// @Summary Get item by ID
// @Description Get item by ID
// @Tags items
// @Produce json
// @Param id path string true "Item ID"
// @Success 200 {object} models.Item
// @Failure 404 {object} models.ErrorResponse
// @Router /items/{id} [get]
func (h *ItemHandler) GetItemByID(c *gin.Context) {
	id := c.Param("id")
	item, err := h.Service.GetItemByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(http.StatusOK, item)
}

// GetItemByQRCode godoc
// @Summary Get item by QR code
// @Description Get item by QR code
// @Tags items
// @Produce json
// @Param qr query string true "QR code data"
// @Success 200 {object} models.Item
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /items/qr [get]
func (h *ItemHandler) GetItemByQRCode(c *gin.Context) {
	qrCodeData := c.Query("qr")
	if qrCodeData == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing QR code data"})
		return
	}
	item, err := h.Service.GetItemByQRCode(c.Request.Context(), qrCodeData)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found for this QR code"})
		return
	}
	c.JSON(http.StatusOK, item)
}

// UpdateItem godoc
// @Summary Update an existing item
// @Description Update an existing item
// @Tags items
// @Accept json
// @Produce json
// @Param id path string true "Item ID to be updated"
// @Param item body models.Item true "Updated item object"
// @Success 200 {object} models.Item
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /items/{id} [put]
func (h *ItemHandler) UpdateItem(c *gin.Context) {
	var item models.Item
	id := c.Param("id")
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item.ID = id
	if err := h.Service.UpdateItem(c.Request.Context(), &item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item"})
		return
	}
	c.JSON(http.StatusOK, item)
}

// DeleteItem godoc
// @Summary Delete item by ID
// @Description Delete item by ID
// @Tags items
// @Produce json // Remove or adjust this line
// @Param id path string true "Item ID to be deleted"
// @Success 204 "No Content"
// @Failure 500 {object} models.ErrorResponse
// @Router /items/{id} [delete]
func (h *ItemHandler) DeleteItem(c *gin.Context) {
	id := c.Param("id")
	if err := h.Service.DeleteItem(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item"})
		return
	}
	c.Status(http.StatusNoContent)
}

// DownloadQRCode godoc
// @Summary Download QR code for an item
// @Description Download QR code for an item as PNG
// @Tags items
// @Produce image/png
// @Param id path string true "Item ID"
// @Success 200 {strings} Image data
// @Failure 404 {object} models.ErrorResponse
// @Router /items/{id}/qrcode [get]
func (h *ItemHandler) DownloadQRCode(c *gin.Context) {
	id := c.Param("id")
	h.Service.DownloadQRCode(c.Request.Context(), id, c)
}
