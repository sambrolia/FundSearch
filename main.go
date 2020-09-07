package main

import (
	"FundSearch/fund"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	funds := ReadFromFile("fund.json")
	companyList, err := fund.GetAllCompanies(funds, "Ethicallobal Fund")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(companyList)
}
