package main

import "fmt"

type AccountNumbers struct {
	AccountNumber [9]Display
}

type Display struct {
	DisplayRow1 [3]string
	DisplayRow2 [3]string
	DisplayRow3 [3]string
}

func main() {

	var accountNumber AccountNumbers

	for display := 0; display < 9; display++ {
		for segment := 0; segment < 3; segment++ {
			accountNumber.AccountNumber[display].DisplayRow1[segment] = "a"
		}
	}

	fmt.Printf("\nLa estructura es:\n %v\n", accountNumber)

}
