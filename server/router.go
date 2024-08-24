package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kalsteve/Good-Night-3rd-Hackathon-Backend/handlers"
)

// 라우터 설정
func setupRouter(handler *handlers.Handler) *gin.Engine {
	router := gin.Default()
	apiGroup := router.Group("/api/v1")
	{
		// Wish routes
		/**
		 * TODO: 소원 관련 라우터 설정
		 * 1. 소원 등록
		 * 2. 소원 삭제
		 * 3. 소원 승인 여부 변경
		 * 4. 소원 조회
		 */
		wishGroup := apiGroup.Group("/wish")
		{
			wishGroup.POST("/", handler.WishHandler.CreateWish)
			wishGroup.DELETE("/:id", handler.WishHandler.DeleteWish)
			wishGroup.PUT("/:id", handler.WishHandler.UpdateWish)
			wishGroup.GET("/:id", handler.WishHandler.GetWish)
			wishGroup.GET("/", handler.WishHandler.GetWishList)
		}

		// Comment routes
		/**
		 * TODO: 댓글 관련 라우터 설정
		 * 1. 댓글 등록
		 * 2. 댓글 조회
		 * 3. 댓글 삭제
		 */
		commentGroup := apiGroup.Group("/comment")
		{
			commentGroup.POST("/", handler.CommentHandler.CreateComment)
			commentGroup.GET("/:id", handler.CommentHandler.GetComments)
			commentGroup.DELETE("/:id", handler.CommentHandler.DeleteComment)
		}
	}
	return router
}
