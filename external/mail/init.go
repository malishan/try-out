package mail

import (
	"context"

	"github.com/try-out/models"
)

type Mail interface {
	SendMail(c context.Context, client string, obj *models.Notification) (string, error)
	StoreMailStatus(c context.Context, client string, obj *models.MailStatus) error
	FetchNotification(c context.Context, client string) (*models.MailStatus, error)
	FetchAllNotification(c context.Context) ([]*models.MailStatus, error)
}

type mail struct {
}

var client mail

func InitializeMail() {
	//TODO: initialize actual database client
	// it will panic if connection error

	store = make(map[string]*models.MailStatus)
}

func GetClient() Mail {
	return client
}
