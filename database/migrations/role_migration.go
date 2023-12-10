package migrations

import (
	"helpdesk/app/role"
	"helpdesk/database"
)

func MigrateRole() {
	if !database.DBConn.Migrator().HasTable(&role.Role{}) {
		database.DBConn.Migrator().CreateTable(&role.Role{})
	}
}
