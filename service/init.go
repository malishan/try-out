package service

import (
	"github.com/try-out/external/database"
	"github.com/try-out/external/mail"
)

var (
	theDbProvider   database.Database
	theMailProvider mail.Mail
)

func InitService(db database.Database, mail mail.Mail) {
	theDbProvider = db
	theMailProvider = mail
}
