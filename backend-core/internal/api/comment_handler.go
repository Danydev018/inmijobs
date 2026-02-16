package api

import (
	"net/http"
	"strconv"

	// Verifica que esta ruta sea exactamente la de tu go.mod
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/core"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/dto"
	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	// Aquí es donde el compilador busca core.CommentService
	service *core.CommentService
}

func NewCommentHandler(service *core.CommentService) *CommentHandler {
	return &CommentHandler{service: service}
}

// Create maneja POST /comments
func (h *CommentHandler) Create(c *gin.Context) {
	// Extraer userID del contexto (inyectado por el middleware de Auth)
	val, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "usuario no autenticado"})
		return
	}
	userID := val.(uint)

	var req dto.CreateCommentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "datos inválidos: " + err.Error()})
		return
	}

	comment, err := h.service.CreateComment(userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, comment)
}

// GetByPost maneja GET /posts/:postId/comments
func (h *CommentHandler) GetByPost(c *gin.Context) {
	postIDStr := c.Param("postId")
	postID, err := strconv.ParseUint(postIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de post inválido"})
		return
	}

	comments, err := h.service.GetCommentsByPost(uint(postID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

// Update maneja PUT /comments/:id
func (h *CommentHandler) Update(c *gin.Context) {
	val, _ := c.Get("userID")
	userID := val.(uint)

	commentIDStr := c.Param("id")
	commentID, err := strconv.ParseUint(commentIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de comentario inválido"})
		return
	}

	var req dto.UpdateCommentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment, err := h.service.UpdateComment(userID, uint(commentID), req)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comment)
}

// Delete maneja DELETE /comments/:id
func (h *CommentHandler) Delete(c *gin.Context) {
	val, _ := c.Get("userID")
	userID := val.(uint)

	commentIDStr := c.Param("id")
	commentID, err := strconv.ParseUint(commentIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de comentario inválido"})
		return
	}

	err = h.service.DeleteComment(userID, uint(commentID))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "comentario eliminado exitosamente"})
}