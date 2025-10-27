package database

import (
	"context"
	"errors"
	"sync"

	"github.com/try-out/models"
)

var (
	rwMutex sync.RWMutex
	store   map[int]*models.Notification
)

func (d database) StoreNotification(c context.Context, obj *models.Notification) error {

	if obj == nil || obj.Id == 0 {
		return errors.New("invalid object")
	}

	rwMutex.Lock()
	defer rwMutex.Unlock()
	store[int(obj.Id)] = obj //create/update

	return nil
}

func (d database) FetchNotification(c context.Context, id int) (*models.Notification, error) {

	if id == 0 {
		return nil, errors.New("invalid object")
	}

	rwMutex.RLock()
	defer rwMutex.RUnlock()

	return store[id], nil
}

func (d database) DeletehNotification(c context.Context, id int) error {

	if id == 0 {
		return errors.New("invalid object")
	}

	rwMutex.Lock()
	defer rwMutex.Unlock()
	delete(store, id)

	return nil
}
