package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/try-out/models"
)

func ListNotificationService(c *gin.Context, status string) (int, models.ApiModels) {

	//TODO: validate clientid

	notifications, err := theMailProvider.FetchAllNotification(c)
	if err != nil {
		return http.StatusInternalServerError, models.ApiModels{
			Error: "Failed to Fetch Notification",
		}
	}

	var resp []*models.MailStatus

	for _, v := range notifications {
		if v.Status == status {
			resp = append(resp, v)
		}
	}

	return http.StatusOK, models.ApiModels{
		Data: resp,
	}
}
