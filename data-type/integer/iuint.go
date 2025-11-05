package main

import "fmt"

func main() {

	// Go integer - data type
	var bitcoinPrice = 104000

  // int refer to negative and positive 
  // unit refer to positive

	// signed Integer: can be negative or positive
	// e.g -92, OR 92 | 12
	// unsigned Integer: must be positive
	// e.g 0 - 838
	var numberCoin2010 int8 = 10

  var numberCountry uint8 = 255

	var bitcoinPriceFuture int64 = 2938183482838

	fmt.Println("\n--- Golang integer---")
	fmt.Printf("Type: %T, value: %v\n", bitcoinPrice, bitcoinPrice)

	fmt.Printf("Type: %T, value: %v\n", bitcoinPriceFuture, bitcoinPriceFuture)

	fmt.Printf("Type: %T, value: %v\n", numberCoin2010, numberCoin2010)

  fmt.Printf("Type: %T, value: %v\n", numberCountry, numberCountry)

  // Thank you and bye.

}
