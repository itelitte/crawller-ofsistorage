package main

import (
	personrestrict "github.com/itelitte/ofsistorage-work/internal/domain/person_restrict"
	"github.com/itelitte/ofsistorage-work/internal/infra/database/cockroachdb"
)

func main() {
	db := cockroachdb.NewCockroachDB()
	repo := personrestrict.NewRepository(*db)

	service := personrestrict.NewService(*repo)

	service.SyncronyzeDB()

}
