package api

import (
	"net/http"

	queue "github.com/Vaibhavsahu2810/email-queue-implementation/internal/redisQueue"
	"github.com/gin-gonic/gin"
)

type SendEmailRequest struct {
	To           string                 `json:"to" binding:"required,email"`
	Subject      string                 `json:"subject" binding:"required"`
	TemplateName string                 `json:"templateName" binding:"required"`
	Data         map[string]interface{} `json:"data" binding:"required"`
}

func RegisterHandlers(router *gin.Engine, redisQueue *queue.RedisQueue) {
	router.GET("/health", healthCheck)

	api := router.Group("/api")
	{
		api.POST("/send", sendEmailHandler(redisQueue))
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func sendEmailHandler(redisQueue *queue.RedisQueue) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req SendEmailRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		task := queue.EmailTask{
			To:           req.To,
			Subject:      req.Subject,
			TemplateName: req.TemplateName,
			Data:         req.Data,
		}

		if err := redisQueue.EnqueueEmail(c.Request.Context(), task); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error enqueueing email: " + err.Error(),
			})
			return
		}

		c.JSON(http.StatusAccepted, gin.H{
			"message": "email queued",
		})
	}
}
