package main

import (
	"FundSearch/fund"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func getHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Fund Search!")
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

	fmt.Fprintln(w, companyList)
	fmt.Println("Endpoint Hit: Get Fund Holdings for ", fundName)

}

func setupRoutes(router *mux.Router) {
	router.HandleFunc("/", getHomePage)
	router.HandleFunc("/api/v1/fund/{fundName}", GetFundHoldings).Methods("Get")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	setupRoutes(router)
	log.Fatal(http.ListenAndServe(":10000", router))
}
