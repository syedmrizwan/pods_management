package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/syedmrizwan/pods_management/database"
	"github.com/syedmrizwan/pods_management/model"
	"net/http"
	"time"
)

// CreatePod godoc
// @Tags API
// @Summary Create Pod
// @Accept json
// @Produce json
// @Param payload body model.PodBody true "description"
// @Success 200 {object} model.Pod
// @Router /api/v1/create_pod [post]
func CreatePod(c *gin.Context) {
	var podBody model.PodBody
	ctx, cancel := context.WithTimeout(c, 50*time.Second)
	defer cancel()
	if err := c.ShouldBindJSON(&podBody); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	db := database.GetConnection()
	pod := model.Pod{
		Name:      podBody.Name,
		IpAddress: podBody.IpAddress,
		Status:    podBody.Status,
	}
	if _, err := db.ModelContext(ctx, &pod).Insert(); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, pod)

}

// UpdatePodConfigurtion godoc
// @Tags API
// @Summary Update Pod Configurtion
// @Accept json
// @Produce json
// @Param payload body []model.PodConfiguration true "description"
// @Router /api/v1/update_pod [post]
func UpdatePodConfigurtion(c *gin.Context) {
	db := database.GetConnection()
	var podConfigurations []model.PodConfiguration

	if err := c.ShouldBindJSON(&podConfigurations); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Couldn't parse the request"})
		return
	}
	var pods []*model.Pod
	for _, podConfiguration := range podConfigurations {
		pod := &model.Pod{
			ID:     podConfiguration.PodID,
			Name:   podConfiguration.PodName,
			Status: podConfiguration.Status,
		}
		db.ModelContext(context.Background(), &pods).Column("status").Where("").Update()
		if podConfiguration.Status != "failed" {
			pod.ClusterID = podConfiguration.Configuration.ClusterID
			pod.DatastoreID = podConfiguration.Configuration.DatastoreID
		}
		pods = append(pods, pod)
	}
	//Bulk update
	db.ModelContext(context.Background(), &pods).Where("status = ? ", "pending").
		Column("status", "datastore_id", "cluster_id").Update()

	c.JSON(http.StatusOK, gin.H{"Message": "Pods updated"})
}
