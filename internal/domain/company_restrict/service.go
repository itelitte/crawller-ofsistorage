package companyrestrict

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/itelitte/ofsistorage-work/internal/domain"
	"github.com/itelitte/ofsistorage-work/internal/domain/crawller"
	"github.com/sirupsen/logrus"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return Service{
		repo,
	}
}

func (s Service) SyncronyzeDB() {
	logrus.Info("Start Insert in database")
	organizations := crawller.GetNamesOrganization()

	if len(organizations) > 0 {

		for _, o := range organizations {

			organization := domain.CompanyRestriction{
				Name:     o.OrganizationName,
				Type:     "Black",
				List:     "ofsistorage",
				UpdateAt: time.Now(),
			}
			company, _ := json.Marshal(organization)
			err := s.repo.SaveList(&organization)

			if err != nil {
				logrus.Error(string(company))
				logrus.Error("Error SaveList Company ", err)
			}
		}
		logrus.Info(fmt.Sprintf("Total of %v companyes", len(organizations)))
		logrus.Info("Finished Syncronyze")
	}
}
