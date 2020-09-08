package main

import (
	"FundSearch/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Add logging

// Add return codes in header

// Output as json rather than struct

func setupRoutes(router *mux.Router) {
	router.HandleFunc("/", controllers.GetHomePage)
	router.HandleFunc("/api/v1/fund/{fundName}", controllers.GetFundHoldings).Methods("Get")
	router.HandleFunc("/api/v1/fund/{fundName}", controllers.GetFundHoldings).Methods("Post")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	setupRoutes(router)
	// exits if fatal raised
	log.Fatal(http.ListenAndServe(":10000", router))

}
