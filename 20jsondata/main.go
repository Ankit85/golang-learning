package main

import (
	"encoding/json"
	"fmt"
)

type AccountDetails struct {
	AccountName    string   `json:"accountname"`
	AccountNumber  int      `json:"accountnumber"`
	AccountBalance int      `json:"-"`
	Address        string   `json:"address"`
	MutualFunds    []string `json:"mutualfunds,omitempty"`
}

func main() {

	fmt.Println("Learning Consuming JSON Data")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	accountData := []AccountDetails{
		{"Sam", 1456, 400, "mumbai", []string{"ICICI Prud", "Kotak Direct"}},
		{"Dam", 1234, 1000, "Delhi", []string{"SBI Prud", "Paras Direct"}},
		{"Con", 7896, 7860, "Kolkata", []string{"Cred Prud", "Bandan Direct"}},
		{"Fon", 5544, 47562, "Banglore", nil},
	}

	formatedData, _ := json.MarshalIndent(accountData, "", "\t")
	fmt.Printf("%s \n", formatedData)
}

func DecodeJson() {

	jsonData := []byte(
		`{
		"AccountName": "Fon",
		"AccountNumber": 5544,
		"AccountBalance": 47562,
		"Address": "Banglore",
		"MutualFunds": ["Adani Cho Prud","Niyka scam Direct"]
	}`)

	var mylocalData AccountDetails

	checkValid := json.Valid(jsonData)
	if checkValid {
		fmt.Println("JSON valid h ji")
		json.Unmarshal(jsonData, &mylocalData)
		fmt.Printf("%#v\n", mylocalData)
	} else {
		fmt.Println("JSON is not valid h ji")
	}

	var mapData map[string]interface{}
	json.Unmarshal(jsonData, &mapData)
	fmt.Printf("%#v\n", mapData)

	for k, v := range mapData {
		fmt.Printf("Key is %v , vaau is %v and type is %T\n", k, v, v)
	}
}
