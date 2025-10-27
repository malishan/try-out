package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/try-out/models"
	"github.com/try-out/service"
)

func listNotification(c *gin.Context) {

	req := c.Query("status")

	if !(req == "sent" || req == "outstanding" || req == "failed") {
		log.Error().Str("clientId", "").Str("api", c.Request.URL.Path).Err(errors.New("invalid request-" + req)).
			Msg("api request validation")
		c.JSON(http.StatusBadRequest, &models.ApiModels{
			Error: "Invalid Request",
			Data:  nil,
		})
		return
	}

	status, resp := service.ListNotificationService(c, req)
	c.JSON(status, resp)
}
