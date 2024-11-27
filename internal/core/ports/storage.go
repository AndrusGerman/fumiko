package ports

import "database/sql"

type Storage interface {
	GetDB() *sql.DB
	GetDialect() string
}
