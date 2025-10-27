package service

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/try-out/models"
)

func SendEmailService(c *gin.Context, clientId string, id int) (int, models.ApiModels) {

	//TODO: validate clientid

	notification, err := theDbProvider.FetchNotification(c, id)
	if err != nil {
		return http.StatusInternalServerError, models.ApiModels{
			Error: "Failed to Fetch Notification",
		}
	}

	if notification == nil {
		return http.StatusBadRequest, models.ApiModels{
			Error: "Invalid Request Notification",
		}
	}

	mailStatus, err := theMailProvider.SendMail(c, clientId, notification)
	if err != nil {
		//TODO: loggin error
	}

	if err := theMailProvider.StoreMailStatus(c, clientId, &models.MailStatus{
		Client:   clientId,
		Status:   mailStatus,
		SentTime: time.Now().Local(),
	}); err != nil {
		return http.StatusInternalServerError, models.ApiModels{
			Error: "Failed to Store Mail Record",
		}
	}

	return http.StatusOK, models.ApiModels{
		Data: map[string]string{
			"status": mailStatus,
		},
	}
}
