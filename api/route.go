package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/update_pod", UpdatePodConfigurtion)
}
