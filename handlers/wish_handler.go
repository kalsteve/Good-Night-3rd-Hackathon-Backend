package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kalsteve/Good-Night-3rd-Hackathon-Backend/models"
	"github.com/kalsteve/Good-Night-3rd-Hackathon-Backend/services"
)

type WishHandler struct {
	service services.WishService
}

func NewWishHandler(service services.WishService) *WishHandler {
	return &WishHandler{service: service}
}

/**
 * TODO: 소원 핸들러 구현
 * 1. 소원 등록
 * 2. 소원 조회
 * 3. 소원 삭제
 * 4. 소원 승인 여부
 */

func (h *WishHandler) CreateWish(c *gin.Context) {
	var wishInput struct {
		Title    string `json:"title"`
		Content  string `json:"content"`
		Category string `json:"category"`
	}

	if err := c.BindJSON(&wishInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := models.StringToCategory(wishInput.Category)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category"})
		return
	}

	newWish := &models.Wish{
		Title:    wishInput.Title,
		Content:  wishInput.Content,
		Category: category,
	}

	if err := h.service.CreateWish(newWish); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create wish"})
		return
	}

	c.JSON(http.StatusOK, newWish)
}

func (h *WishHandler) GetWish(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	wish, err := h.service.GetWishByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wish not found"})
		return
	}

	c.JSON(http.StatusOK, wish)
}

func (h *WishHandler) GetWishList(c *gin.Context) {
	confirm := c.Query("confirm")
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))

	conf, err := models.StringToConfirm(confirm)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid confirm"})
		return
	}

	wishes, err := h.service.GetWishListByConfirm(conf, page, size)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get wish list"})
		return
	}

	c.JSON(http.StatusOK, wishes)
}

func (h *WishHandler) DeleteWish(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.service.DeleteWish(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wish not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Wish deleted"})
}

func (h *WishHandler) UpdateWish(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var wishInput struct {
		Confirm string `json:"confirm"`
	}

	if err := c.BindJSON(&wishInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	conf, err := models.StringToConfirm(wishInput.Confirm)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid confirm"})
		return
	}

	if h.service.UpdateWish(uint(id), conf) != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wish not found"})
		return
	}
}
