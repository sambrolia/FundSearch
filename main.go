package main

import (
	"FundSearch/fund"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func readFile() []fund.Fund {
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

func buildFundTree(funds []fund.Fund) {

	// Give each fund a pointer back to the array
	for i, _ := range funds {
		funds[i].Root = &funds
	}

	// Retreive the complete list of companies and
	// their weightings
	var holdings []fund.Holding
	for _, fund := range funds {
		holdings = append(holdings, fund.GetHoldings()...)
	}

	var totalWeight float32
	for _, holding := range holdings {
		totalWeight += holding.Weight
	}
	// Print result
	fmt.Println("Holdings = ", holdings)
	fmt.Println("Total Weight = ", totalWeight)

}

func main() {
	funds := readFile()
	buildFundTree(funds)
}
