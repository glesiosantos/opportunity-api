package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "Get Opening",
			})
		})
		v1.POST("/openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "Post Opening",
			})
		})
		v1.DELETE("/openings/{id}", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "Delete Opening",
			})
		})
		v1.PUT("/openings/edit/{id}", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "Put Opening",
			})
		})

	}
}
