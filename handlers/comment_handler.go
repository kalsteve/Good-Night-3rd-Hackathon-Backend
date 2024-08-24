package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/kalsteve/Good-Night-3rd-Hackathon-Backend/models"
	"github.com/kalsteve/Good-Night-3rd-Hackathon-Backend/services"
	"net/http"
	"strconv"
)

type CommentHandler struct {
	service services.CommentService
}

func NewCommentHandler(service services.CommentService) *CommentHandler {
	return &CommentHandler{service: service}
}

func (h *CommentHandler) CreateComment(c *gin.Context) {
	var commentInput struct {
		Content string `json:"content"`
	}

	if err := c.BindJSON(&commentInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newComment := &models.Comment{
		Content: commentInput.Content,
	}

	if err := h.service.CreateComment(newComment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusOK, newComment)
}

func (h *CommentHandler) GetComments(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	comments, err := h.service.GetCommentListByWishID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}
	c.JSON(http.StatusOK, comments)
}

func (h *CommentHandler) DeleteComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.DeleteComment(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted"})
}
