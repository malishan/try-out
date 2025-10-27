package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/try-out/models"
)

func DeleteNotificationService(c *gin.Context, id int) (int, models.ApiModels) {

	//TODO: validate clientid

	err := theDbProvider.DeletehNotification(c, id)
	if err != nil {
		return http.StatusInternalServerError, models.ApiModels{
			Error: "Failed to Delete Notification",
		}
	}

	return http.StatusOK, models.ApiModels{
		Data: map[string]int{
			"id": id,
		},
	}
}
