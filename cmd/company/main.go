package main

import (
	"time"

	"github.com/go-co-op/gocron"
	companyrestrict "github.com/itelitte/ofsistorage-work/internal/domain/company_restrict"
	"github.com/itelitte/ofsistorage-work/internal/infra/database/cockroachdb"
	"github.com/sirupsen/logrus"
)

func main() {
	//set time america/sao_paulo
	_, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		logrus.Errorf("error to load location: %s", err)
		return
	}

	logrus.Infof("start to run ofsistorage ofsistorage")

	s := gocron.NewScheduler(time.UTC)

	logrus.Infof("start to run crawller ofsistorage every 5 hours")
	//run every day
	_, err = s.Every(5).Hours().Tag("crawller-ofsistorage").Do(run)

	//_, err = s.Every(4).Minutes().Tag("crawller-ofsistorage").Do(run)
	if err != nil {
		return
	}

	s.StartBlocking()

}

func run() {
	db := cockroachdb.NewCockroachDB()
	repo := companyrestrict.NewRepository(*db)
	repo.ClearTable()
	service := companyrestrict.NewService(*repo)
	service.SyncronyzeDB()

}
