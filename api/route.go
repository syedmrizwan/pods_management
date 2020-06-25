package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/create_pod", CreatePod)
	router.POST("/update_pod", UpdatePodConfigurtion)
	router.DELETE("delete_pod", DeletePods)
}
