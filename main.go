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

func main() {
	funds := readFile()
	companyList, err := fund.GetAllCompanies(funds, "Ethicallobal Fund")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(companyList)
}
