package core

import (
	"log"
	"strings"

	"github.com/davevurby/shieldwall"
	"github.com/davevurby/shieldwall-server/core/config"
	"github.com/davevurby/shieldwall/store/postgres"
)

func CreateStore(cfg *config.Config) shieldwall.Store {
	dsn := cfg.GetDSN()
	if strings.HasPrefix(dsn, "postgres") {
		store, err := postgres.NewPgStoreFromConnString(dsn)
		if err != nil {
			log.Fatalf("connection to postgres failed - %s\n", err.Error())
		}

		err = store.RunMigrations()
		if err != nil {
			log.Fatalf("postgres migrations failed - %s\n", err.Error())
		}

		return store

	} else {
		log.Fatalf("your DSN setting is not supported")
	}

	return nil
}
