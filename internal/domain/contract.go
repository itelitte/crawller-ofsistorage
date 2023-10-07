package domain

import (
	"time"

	"github.com/vingarcia/ksql"
)

var UsersTable = ksql.NewTable("countryrestriction")

type SanctionEntry struct {
	Names             string
	AKA               string
	Title             string
	NameNonLatin      string
	DOB               string
	POB               string
	Nationality       string
	PassportNumber    string
	PassportDetails   string
	Position          string
	OtherInformation  string
	ListedOn          string
	SanctionsListDate string
	LastUpdated       string
	GroupID           string
}

type OrganizationEntry struct {
	OrganizationName  string
	NameNonLatin      string
	AKA               []string
	Address           []string
	OtherInformation  string
	ListedOn          string
	SanctionsListDate string
	LastUpdated       string
	GroupID           string
}

type Coutry struct {
	Name     string
	UpdateAt string
}

type CompanyRestriction struct {
	ID       int       `ksql:"id" json:"id" `
	Name     string    `ksql:"name" json:"name" `
	List     string    `ksql:"list" json:"list" `
	Type     string    `ksql:"type" json:"type" `
	Alias    string    `ksql:"alias" json:"alias" `
	UpdateAt time.Time `ksql:"update_at" json:"update_at" `
}

type PersonRestrict struct {
	ID           int       `ksql:"id" json:"id"`
	Name         string    `ksql:"name" json:"name"`
	OriginalName string    `ksql:"original_name" json:"original_name"`
	List         string    `ksql:"list" json:"list"`
	Type         string    `ksql:"type" json:"type"`
	AlsoKnown    string    `ksql:"also_known" json:"also_known"`
	UpdateAt     time.Time `ksql:"update_at" json:"update_at"`
}
