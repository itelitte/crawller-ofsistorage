package crawller

import (
	"fmt"
	"testing"
)

func TestGetNames(t *testing.T) {
	entries := GetNamesEntries()

	if len(entries) == 0 {
		t.Errorf("Expected result to be greater than 0")
	}

}

func TestGetNamesOrganization(t *testing.T) {

	organization := GetNamesOrganization()

	fmt.Println("total: ", len(organization))
	if len(organization) == 0 {
		t.Errorf("Expected result to be greater than 0")
	}
}
