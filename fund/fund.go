package fund

import (
	"errors"
)

type Holding struct {
	Name   string
	Weight float32
}

type Fund struct {
	Name     string
	Holdings []Holding
	Root     *[]Fund
}

func ReadFromFile() []fund.Fund {
	jsonFile, err := os.Open("fund.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully opened fund.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var funds []fund.Fund
	err = json.Unmarshal(byteValue, &funds)
	if err != nil {
		fmt.Println("error:", err)
	}

	return funds
}

func (f Fund) GetHoldings() []Holding {
	var holdings []Holding
	for _, holding := range f.Holdings {
		var fundFound *Fund
		for i, fund := range *f.Root {
			if holding.Name == fund.Name {
				fundFound = &(*f.Root)[i]
			}
		}
		if fundFound != nil {
			holdings = append(holdings,
				DiluteHoldings(holding.Weight, fundFound.GetHoldings()...)...)
		} else {
			holdings = append(holdings, holding)
		}
	}

	return holdings
}

func DiluteHoldings(parentWeight float32, holdings ...Holding) []Holding {
	for i, _ := range holdings {
		holdings[i].Weight = holdings[i].Weight * parentWeight
	}
	return holdings
}

func GetAllCompanies(funds []Fund, fundName string) ([]Holding, error) {
	// Give each fund a pointer back to the array
	for i, _ := range funds {
		funds[i].Root = &funds
	}

	// Retreive the complete list of companies and
	// their weightings
	var holdings []Holding
	fundFound := false
	for _, fund := range funds {
		if fund.Name == fundName {
			holdings = append(holdings, fund.GetHoldings()...)
			fundFound = true
		}
	}
	if !fundFound {
		return []Holding{}, errors.New("Fund not found: " + fundName)
	}

	/* Where multiple funds have returned the same company
	combine them into a single total holding of each one */
	return combineCompanyHoldings(holdings), nil
}

func combineCompanyHoldings(holdings []Holding) []Holding {
	seen := make(map[string]float32)
	combinedHoldings := []Holding{}
	for _, holding := range holdings {
		seen[holding.Name] += holding.Weight
	}

	for k, v := range seen {
		combinedHoldings = append(combinedHoldings, Holding{k, v})
	}

	return combinedHoldings
}
