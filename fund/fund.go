package fund

// import (
// 	"fmt"
// )

type Holding struct {
	Name   string
	Weight float32
}

type Fund struct {
	Name     string
	Holdings []Holding
	Root     *[]Fund
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
