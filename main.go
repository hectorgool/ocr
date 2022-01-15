package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type AccountNumbers struct {
	AccountNumber [9]Display
}

type Display struct {
	DisplayRow1 [3]string
	DisplayRow2 [3]string
	DisplayRow3 [3]string
}

func main() {

	/*
		var accountNumber AccountNumbers

		for row := 0; row < 3; row++ {
			for display := 0; display < 9; display++ {
				for segment := 0; segment < 3; segment++ {
					if row == 0 {
						accountNumber.AccountNumber[display].DisplayRow1[segment] = "a"
					}
					if row == 1 {
						accountNumber.AccountNumber[display].DisplayRow2[segment] = "b"
					}
					if row == 2 {
						accountNumber.AccountNumber[display].DisplayRow3[segment] = "c"
					}
				}
			}
		}

		fmt.Printf("\nLa estructura es:\n %v\n", accountNumber)
	*/

	rows := ReadFileNuberAccounts()

	for i, row := range rows {
		fmt.Printf("%v %v\n", i, row)
	}

}

func ReadFileNuberAccounts() []string {

	file, err := os.Open("test_numbers.txt")
	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()
	return text
}
