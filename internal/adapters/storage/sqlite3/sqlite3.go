package sqlite3

import (
	"database/sql"
	"errors"

	"github.com/AndrusGerman/fumiko/internal/core/domain"
	"github.com/AndrusGerman/fumiko/internal/core/ports"
	"go.uber.org/fx"
)

type sqlite3 struct {
	db      *sql.DB
	dialect string
}

// GetDB implements ports.Storage.
func (s *sqlite3) GetDB() *sql.DB {
	return s.db
}

// GetDialect implements ports.Storage.
func (s *sqlite3) GetDialect() string {
	return s.dialect
}

func New(lc fx.Lifecycle) (ports.Storage, error) {
	var dialect = "sqlite3"
	var uri = "file:whatsapp.db?_foreign_keys=on"
	var db *sql.DB
	var err error

	if db, err = sql.Open(dialect, uri); err != nil {
		return nil, errors.Join(err, domain.ErrFailedOpenDatabase)
	}

	lc.Append(fx.StopHook(db.Close))
	return &sqlite3{db: db, dialect: dialect}, nil
}
