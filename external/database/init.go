package database

import (
	"context"

	"github.com/try-out/models"
)

type Database interface {
	StoreNotification(c context.Context, obj *models.Notification) error
	FetchNotification(c context.Context, id int) (*models.Notification, error)
	DeletehNotification(c context.Context, id int) error
}

type database struct {
}

var db database

func InitializeDatabase() {
	//TODO: initialize actual database client
	// it will panic if connection error

	store = make(map[int]*models.Notification)
}

func GetClient() Database {
	return db
}
