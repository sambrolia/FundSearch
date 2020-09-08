package controllers

import (
	"FundSearch/fund"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetHomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, welcome to fund search, please use api!"))
	fmt.Fprintln(w, "at /api/v1/fund/{FundName}!")
	fmt.Println("Endpoint Hit: Home Page")
}

func GetFundHoldings(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fundName := vars["fundName"]

	// Read json from file, passing in a default file path,
	// which will be replaced if a file exists in the request
	var funds []fund.Fund

	funds = fund.ReadFromFile("examples/sample-funds.json", r)

	companyList, err := fund.GetAllCompanies(funds, fundName)
	if err != nil {
		fmt.Println(err)
	}

	jsonCompanyList, err := json.Marshal(companyList)
	if err != nil {
		log.Println("Unable to marshal struct into json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		fmt.Fprintln(w, string(jsonCompanyList))
		log.Println("Endpoint Hit: Get Fund Holdings for ", fundName)
	}
}
