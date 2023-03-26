package main

import "fmt"

func main() {

	fmt.Println("Learning Methods")

	shopee := BankAccount{"Ankit", 11110001456, "India", true, "SalaryAccount", true}
	fmt.Printf("User Account %v:\n", shopee)
	fmt.Printf("User Account details %+v\n", shopee)
	shopee.GetLoanStatus()
	shopee.changeName()
	fmt.Printf("User AccountName %v:\n", shopee.AccountHolderName)

}

type BankAccount struct {
	AccountHolderName string
	AccountNumber     uint64
	Address           string
	isZeroBalance     bool
	AccountType       string
	isLoanApproved    bool
}

func (b BankAccount) GetLoanStatus() {
	if b.isLoanApproved {
		fmt.Println("Your Loan request for 400k is approved")
	} else {
		fmt.Println("Your Loan request for 400k is Awaiting for Approval")
	}
}

func (b BankAccount) changeName() {
	b.AccountHolderName = "Sam"
	fmt.Printf("User account name is changed to this: %v \n", b.AccountHolderName)
}
