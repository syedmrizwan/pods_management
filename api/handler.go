package api

import (
	"github.com/gin-gonic/gin"
	"github.com/syedmrizwan/pods_management/database"
	"github.com/syedmrizwan/pods_management/model"
	"net/http"
	"context"
)

// UpdatePodConfigurtion godoc
// @Tags API
// @Summary Update Pod Configurtion
// @Accept json
// @Produce json
// @Param payload body []model.PodConfiguration true "description"
// @Router /api/update_pod [post]
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
