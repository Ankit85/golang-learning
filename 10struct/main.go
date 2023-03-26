package main

import "fmt"

func main() {

	fmt.Println("Learning Struct")

	shopee := BankAccount{"Ankit", 11110001456, "India", true, "SalaryAccount"}
	fmt.Printf("User Account %v:\n", shopee)
	fmt.Printf("User Account details %+v\n", shopee)

}

type BankAccount struct {
	AccountHolderName string
	AccountNumber     uint64
	Address           string
	isZeroBalance     bool
	AccountType       string
}
