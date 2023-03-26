package main

import "fmt"

// decalring a pubic varibale
const JWT = "AXDNSFUISJBUSDHjnsdasdshdsadi"

func main() {
	var name = "Ankit"
	fmt.Println(name)
	fmt.Printf("Variable is type:: %T \n", name)

	var isAdmin bool = false
	fmt.Println(isAdmin)
	fmt.Printf("Variable is type:: %T \n", isAdmin)

	var Id uint8 = 255
	fmt.Println(Id)
	fmt.Printf("Variable is type:: %T \n", Id)

	var fId float32 = 255.2324344
	fmt.Println(fId)
	fmt.Printf("Variable is type:: %T \n", fId)

	var fIds float64 = 255.1234567890
	fmt.Println(fIds)
	fmt.Printf("Variable is type:: %T \n", fIds)

	// default value and aliases
	var Idd uint8
	fmt.Println(Idd)
	fmt.Printf("Variable is type:: %T \n", Idd)

	var ndefault string
	fmt.Println(ndefault)
	fmt.Printf("Variable is type:: %T \n", ndefault)

	var dbool bool
	fmt.Println(dbool)
	fmt.Printf("Variable is type:: %T \n", dbool)

	var dcomplex complex64
	fmt.Println(dcomplex)
	fmt.Printf("Variable is type:: %T \n", dcomplex)

	// another way of variable declaring
	// u cannot declare this type of vairbale outside the methods
	msg := "GM"
	fmt.Println(msg)

	fmt.Println(JWT)

}
