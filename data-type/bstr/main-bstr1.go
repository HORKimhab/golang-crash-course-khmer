package main

import (
	"fmt"
	"strings"
)

func mainBooleanStr() {
	var isMain bool    // false by default
	isCambodia := true // true
	isBitcoin := true  // true

	fmt.Println("\nLearning golang Boolean---")
	fmt.Printf("Type: %T, value: %v\n", isMain, isMain)
	fmt.Println("isCambodia:", isCambodia)
	fmt.Println("isBitcoin:", isBitcoin)

	fmt.Println("\n--- Common operations---")
	x, y := true, false

	// logical ops
	var boolean1 = x && y    // false
	var boolean2 = x || y    // true
	var boolean3 = !(x || y) // false

	fmt.Println("boolean1:", boolean1)
	fmt.Println("boolean2:", boolean2)
	fmt.Println("boolean3:", boolean3)

	fmt.Println("\n---Boolean in control flow---")
	if isBitcoin {
		fmt.Println("Btc is a bitcoin")
	} else {
		fmt.Println("It is a other coin")
	}

	fmt.Println("\n---Boolean in loop---")
	for n := 0; n < 28; n++ {
		if n%2 == 0 {
			fmt.Println("even", n)
		}
	}

	fmt.Println("\n---Golang data type string---")
	name := "HKimhab"
	string2 := `Raw string\nBitcoin` // backtickes - no escapes
	string3 := "Hello" + " " + "Bitcoin"

	fmt.Println("name: ", name)
	fmt.Println("string2: ", string2)
	fmt.Println("string3: ", string3)

	// check type and value
	fmt.Printf("Tye: %T, value: %v\n", name, name)
	// len of string
	fmt.Println("length of string3:", len(string3))

	// Tolowercase and ToUppercase
	fmt.Println("Uppercase of string3: ", strings.ToUpper(string3))
	fmt.Println("Lowercase of string2 + name: ", strings.ToLower(string2+name))

	// Find Bitcoin in string3
	containBitcoin := strings.Contains(string3, "Bitcoin")
	// find value or substring of string
	fmt.Println(containBitcoin)
	// Thank you and bye. 

}
