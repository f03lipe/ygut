package globals

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type globals struct {
	Db *sqlx.DB
}

var G *globals

func Setup() *globals {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Puts.\n")
		}
	}()

	G = new(globals)
	G.Db = setupDb()

	return G
}

func Close(g *globals) {
	closeDb(g.Db)
}
