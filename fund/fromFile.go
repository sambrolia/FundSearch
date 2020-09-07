package fund

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ReadFromFile(defaultPath string, r *http.Request) []Fund {
	// Read file to Byte array
	file, _, err := r.FormFile("file")
	var byteValue []byte
	if err != nil {
		fmt.Println("Cannot detect valid file, using default")
		byteValue, err = ioutil.ReadFile(defaultPath)
		fmt.Println("Successfully read file")
	} else {
		fmt.Println("Using file passed by user request")
		byteValue, err = ioutil.ReadAll(file)
		fmt.Println("Successfully read file")
		defer file.Close()
	}

	// Unmarshal Byte array into a list of funds
	var funds []Fund
	err = json.Unmarshal(byteValue, &funds)
	if err != nil {
		fmt.Println("error:", err)
	}

	return funds
}
