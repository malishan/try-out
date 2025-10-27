package v1

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/try-out/models"
	"github.com/try-out/service"
)

func sendNotification(c *gin.Context) {

	id := c.Query("id")
	client := c.Query("client")

	if client == "" || id == "" {
		log.Error().Str("clientId", client).Str("api", c.Request.URL.Path).Err(errors.New("invalid request")).
			Msg("api request validation")
		c.JSON(http.StatusBadRequest, &models.ApiModels{
			Error: "Invalid Request",
			Data:  nil,
		})
		return
	}

	//TODO: handle error
	notificationId, _ := strconv.Atoi(id)

	status, resp := service.SendEmailService(c, client, notificationId)
	c.JSON(status, resp)
}
