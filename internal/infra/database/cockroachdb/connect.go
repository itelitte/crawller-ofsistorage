package cockroachdb

import (
	"context"
	"log"

	"github.com/itelitte/ofsistorage-work/internal/config"
	"github.com/vingarcia/ksql"
	"github.com/vingarcia/ksql/adapters/kpgx"
)

type CockroachDB struct {
	DB *ksql.DB
}

func NewCockroachDB() *CockroachDB {
	ctx := context.Background()
	db, err := kpgx.New(ctx, config.GetDataBaseDNS(), ksql.Config{})
	if err != nil {
		log.Fatalf("unable connect to database: %s", err)
	}

	return &CockroachDB{
		DB: &db,
	}
}
