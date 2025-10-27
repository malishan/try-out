package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/try-out/models"
	"github.com/try-out/service"
)

func createNotification(c *gin.Context) {
	var request models.CreateNotification

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error().Str("clientId", "").Str("api", c.Request.URL.Path).Err(err).Msg("api request binding")
		c.JSON(http.StatusInternalServerError, &models.ApiModels{
			Error: "Parsing Request Error",
			Data:  nil,
		})
		return
	}

	if err := request.Validation(); err != nil {
		log.Error().Str("clientId", "").Str("api", c.Request.URL.Path).Err(err).Msg("api request validation")
		c.JSON(http.StatusBadRequest, &models.ApiModels{
			Error: "Invalid Request",
			Data:  nil,
		})
		return
	}

	status, resp := service.CreateService(c, request)
	c.JSON(status, resp)
}
