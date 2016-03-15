package globals

import (
	"os"

	"github.com/f03lipe/ygut/models"
	"github.com/jmoiron/sqlx"
)

const (
	defaultDbCreds = "user=foo dbname=bar sslmode=disable"
)

func setupDb() *sqlx.DB {
	addr := os.Getenv("DEFAULT_CREDENTIALS")
	if addr == "" {
		addr = defaultDbCreds
	}

	db := sqlx.MustConnect("postgres", defaultDbCreds)

	db.MustExec(models.Schema)

	if db.Ping() == nil {
		panic("ping you.")
	}

	return db
}

func closeDb(db *sqlx.DB) {
	// There's nothing to be done.
}
