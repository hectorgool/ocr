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

	accountNumber.AccountNumber[0].DisplayRow1[0] = "a"
	accountNumber.AccountNumber[0].DisplayRow1[1] = "b"
	accountNumber.AccountNumber[0].DisplayRow1[2] = "c"

	accountNumber.AccountNumber[1].DisplayRow1[0] = "a"
	accountNumber.AccountNumber[1].DisplayRow1[1] = "b"
	accountNumber.AccountNumber[1].DisplayRow1[2] = "c"

	accountNumber.AccountNumber[2].DisplayRow1[0] = "a"
	accountNumber.AccountNumber[2].DisplayRow1[1] = "b"
	accountNumber.AccountNumber[2].DisplayRow1[2] = "c"

	accountNumber.AccountNumber[3].DisplayRow1[0] = "a"
	accountNumber.AccountNumber[3].DisplayRow1[1] = "b"
	accountNumber.AccountNumber[3].DisplayRow1[2] = "c"

	accountNumber.AccountNumber[8].DisplayRow1[0] = "a"
	accountNumber.AccountNumber[8].DisplayRow1[1] = "b"
	accountNumber.AccountNumber[8].DisplayRow1[2] = "c"

	fmt.Printf("\nLa estructura es:\n %v\n", accountNumber)

}
