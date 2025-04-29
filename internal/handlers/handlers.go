package handlers

import (
	"awesomeProject2/internal/service"
	"awesomeProject2/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ItemHandler struct {
	Service *service.ItemService
}

func NewItemHandler(service *service.ItemService) *ItemHandler {
	return &ItemHandler{Service: service}
}

func (h *ItemHandler) CreateItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.CreateItem(c, &item); err != nil {
		log.Printf("❌ Error creating item: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item"})
		return
	}
	c.JSON(http.StatusCreated, item)
}

func (h *ItemHandler) GetItemByID(c *gin.Context) {
	id := c.Param("id")
	item, err := h.Service.GetItemByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *ItemHandler) UpdateItem(c *gin.Context) {
	var item models.Item
	id := c.Param("id")
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item.ID = id
	if err := h.Service.UpdateItem(c, &item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item"})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *ItemHandler) DeleteItem(c *gin.Context) {
	id := c.Param("id")
	if err := h.Service.DeleteItem(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item"})
		return
	}
	c.Status(http.StatusNoContent)
}
