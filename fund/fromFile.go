package fund

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ReadFromFile(path string, r *http.Request) []Fund {
	file, _, err := r.FormFile("file")
	var byteValue []byte
	if err != nil {
		fmt.Println("Cannot detect valid file, using default")
		byteValue, err = ioutil.ReadFile(path)
	} else {
		byteValue, err = ioutil.ReadAll(file)
		defer file.Close()
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully read file")
	}

	//byteValue, _ := ioutil.ReadAll(jsonFile)

	var funds []Fund
	err = json.Unmarshal(byteValue, &funds)
	if err != nil {
		fmt.Println("error:", err)
	}

	return funds
}
