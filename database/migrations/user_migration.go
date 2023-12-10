package migrations

import (
	"helpdesk/app/user"
	"helpdesk/database"
)

func MigrateUser() {
	if !database.DBConn.Migrator().HasTable(&user.User{}) {
		database.DBConn.Migrator().CreateTable(&user.User{})
	}
}
