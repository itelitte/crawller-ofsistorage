package personrestrict

import (
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

	logrus.Info("Starting Syncronize")

	entries := crawller.GetNamesEntries()

	logrus.Info(fmt.Sprintf("Total of %v person", len(entries)))

	if len(entries) > 0 {

		logrus.Info("Start Insert in database")
		for _, e := range entries {

			person := domain.PersonRestrict{
				Name:         e.Names,
				OriginalName: e.NameNonLatin,
				List:         "ofsistorage",
				Type:         "Black",
				AlsoKnown:    e.AKA,
				UpdateAt:     time.Now(),
			}

			err := s.repo.SaveList(&person)

			if err != nil {
				logrus.Error("Error SaveList Person ", err)
			}
		}

		logrus.Info("Finished Syncronyze")
	}

}
