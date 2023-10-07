package companyrestrict

import (
	"context"

	"github.com/itelitte/ofsistorage-work/internal/domain"
	"github.com/itelitte/ofsistorage-work/internal/infra/database/cockroachdb"
	"github.com/sirupsen/logrus"
	"github.com/vingarcia/ksql"
)

var CountryTable = ksql.NewTable("company_restricted")

type Repository struct {
	DB *ksql.DB
}

func NewRepository(db cockroachdb.CockroachDB) *Repository {
	return &Repository{
		DB: db.DB,
	}
}

func (r *Repository) SaveList(company *domain.CompanyRestriction) error {
	ctx := context.Background()

	err := r.DB.Insert(ctx, CountryTable, company)
	if err != nil {
		logrus.Errorf("unable to save country: %#v", err)
		return err
	}

	return nil

}

func (r *Repository) ClearTable() error {
	ctx := context.Background()

	q := `DELETE FROM company_restricted where list = 'ofsistorage'`

	_, err := r.DB.Exec(ctx, q)

	if err != nil {
		logrus.Errorf("unable to clear table: %#v", err)
		return err
	}

	return nil
}
