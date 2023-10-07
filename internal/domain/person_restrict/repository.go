package personrestrict

import (
	"context"
	"encoding/json"

	"github.com/itelitte/ofsistorage-work/internal/domain"
	"github.com/itelitte/ofsistorage-work/internal/infra/database/cockroachdb"
	"github.com/sirupsen/logrus"
	"github.com/vingarcia/ksql"
)

var PersonTable = ksql.NewTable("person_restricted")

type Repository struct {
	DB *ksql.DB
}

func NewRepository(db cockroachdb.CockroachDB) *Repository {
	return &Repository{
		DB: db.DB,
	}
}

func (r *Repository) SaveList(person *domain.PersonRestrict) error {
	ctx := context.Background()
	p, _ := json.Marshal(person)

	err := r.DB.Insert(ctx, PersonTable, person)
	if err != nil {
		logrus.Errorf("unable to save country: %#v", err)
		logrus.Info(string(p))
		return err
	}

	return nil

}

func (r *Repository) ClearTable() error {
	ctx := context.Background()

	q := `DELETE FROM person_restrict where list = 'ofsistorage'`

	_, err := r.DB.Exec(ctx, q)

	if err != nil {
		logrus.Errorf("unable to clear table: %#v", err)
		return err
	}

	return nil
}
