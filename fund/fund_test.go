package fund

import (
	"errors"
	"net/http"
	"sort"
	"testing"
)

// ByAge implements sort.Interface based on the Age field.
type ByName []Holding

func (a ByName) Len() int           { return len(a) }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func RunGetAllCompanies(correctHoldings []Holding, fundName string) error {
	var r http.Request
	funds := ReadFromFile("../examples/sample-funds.json", &r)
	companyList, err := GetAllCompanies(funds, fundName)
	if err != nil {
		return err
	}

	// Sort by Holding Name
	sort.Sort(ByName(companyList))

	for i, _ := range companyList {
		if companyList[i].Name != correctHoldings[i].Name {
			return errors.New("Holding Name is incorrect")
		} else if companyList[i].Weight != correctHoldings[i].Weight {
			return errors.New("Holding Name is incorrect")
		}
	}
	return nil
}
func TestGetAllCompanies_EthicalGlobalFund(t *testing.T) {
	fundName := "Ethical Global Fund"
	correctHoldings := []Holding{
		Holding{Name: "BeanzRUS", Weight: 0.21},
		Holding{Name: "GoldenGadgets", Weight: 0.15},
		Holding{Name: "GrapeCo", Weight: 0.347},
		Holding{Name: "GreenCo", Weight: 0.06},
		Holding{Name: "MicroFit", Weight: 0.1},
		Holding{Name: "SolarCorp", Weight: 0.028},
		Holding{Name: "SpaceY", Weight: 0.105},
	}

	err := RunGetAllCompanies(correctHoldings, fundName)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGetAllCompanies_FundB(t *testing.T) {
	fundName := "Fund B"
	correctHoldings := []Holding{
		Holding{Name: "GrapeCo", Weight: 0.2},
		Holding{Name: "GreenCo", Weight: 0.3},
		Holding{Name: "MicroFit", Weight: 0.5},
	}

	err := RunGetAllCompanies(correctHoldings, fundName)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGetAllCompanies_FundC(t *testing.T) {
	fundName := "Fund C"
	correctHoldings := []Holding{
		Holding{Name: "BeanzRUS", Weight: 0.24},
		Holding{Name: "GoldenGadgets", Weight: 0.3},
		Holding{Name: "GrapeCo", Weight: 0.308},
		Holding{Name: "SolarCorp", Weight: 0.032},
		Holding{Name: "SpaceY", Weight: 0.12},
	}

	err := RunGetAllCompanies(correctHoldings, fundName)
	if err != nil {
		t.Errorf(err.Error())
	}
}
func TestGetAllCompanies_FundD(t *testing.T) {
	fundName := "Fund D"
	correctHoldings := []Holding{
		Holding{Name: "BeanzRUS", Weight: 0.6},
		Holding{Name: "GrapeCo", Weight: 0.02},
		Holding{Name: "SolarCorp", Weight: 0.08},
		Holding{Name: "SpaceY", Weight: 0.3},
	}

	err := RunGetAllCompanies(correctHoldings, fundName)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGetAllCompanies_FundE(t *testing.T) {
	fundName := "Fund E"
	correctHoldings := []Holding{
		Holding{Name: "GrapeCo", Weight: 0.2},
		Holding{Name: "SolarCorp", Weight: 0.8},
	}

	err := RunGetAllCompanies(correctHoldings, fundName)
	if err != nil {
		t.Errorf(err.Error())
	}
}
