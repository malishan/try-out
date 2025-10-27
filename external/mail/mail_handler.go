package mail

import (
	"context"
	"encoding/json"
	"errors"
	"sync"

	"github.com/try-out/models"
)

var (
	rwMutex sync.RWMutex
	store   map[string]*models.MailStatus
)

func (d mail) SendMail(c context.Context, client string, obj *models.Notification) (string, error) {

	if obj == nil || client == "" {
		return "failed", errors.New("invalid object")
	}

	_, err := json.Marshal(obj)
	if err != nil {
		return "failed", errors.New("marshal failed-" + err.Error())
	}

	return "sent", nil
}

func (d mail) StoreMailStatus(c context.Context, client string, obj *models.MailStatus) error {

	if obj == nil || client == "" {
		return errors.New("invalid object")
	}

	rwMutex.Lock()
	defer rwMutex.Unlock()
	store[client] = obj //create/update

	return nil
}

func (d mail) FetchNotification(c context.Context, client string) (*models.MailStatus, error) {

	if client == "" {
		return nil, errors.New("invalid object")
	}

	rwMutex.RLock()
	defer rwMutex.RUnlock()

	return store[client], nil
}

func (d mail) FetchAllNotification(c context.Context) ([]*models.MailStatus, error) {

	//TODO: improve by pagination

	rwMutex.RLock()
	defer rwMutex.RUnlock()

	var notif []*models.MailStatus

	for _, v := range store {
		notif = append(notif, v)
	}

	return notif, nil
}
