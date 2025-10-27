package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/try-out/models"
	"github.com/try-out/utils/helper"
)

func CreateService(c *gin.Context, req models.CreateNotification) (int, models.ApiModels) {

	obj := &models.Notification{
		Id:                helper.GenerateRandomInteger(),
		CurrentPrice:      req.CurrentPrice,
		MarketTradeVolume: req.MarketTradeVolume,
		IntraDayHighPrice: req.IntraDayHighPrice,
		MarketCap:         req.MarketCap,
	}

	if err := theDbProvider.StoreNotification(c, obj); err != nil {
		return http.StatusInternalServerError, models.ApiModels{
			Error: "Failed to Store Record",
		}
	}

	return http.StatusOK, models.ApiModels{
		Data: map[string]int64{
			"id": obj.Id,
		},
	}
}
