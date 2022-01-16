package main

import (
	"bufio"
	"bytes"
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

	var accountNumber AccountNumbers

	countAccountNumber := 0
	countDisplayRow := 0

	rows := ReadFileNuberAccounts()

	for rowCount, row := range rows {

		segments := bytes.Runes([]byte(row))

		for _, segment := range segments {

			if rowCount == 0 {
				accountNumber.AccountNumber[countAccountNumber].DisplayRow1[countDisplayRow] = string(segment)
			}
			if rowCount == 1 {
				accountNumber.AccountNumber[countAccountNumber].DisplayRow2[countDisplayRow] = string(segment)
			}
			if rowCount == 2 {
				accountNumber.AccountNumber[countAccountNumber].DisplayRow3[countDisplayRow] = string(segment)
			}

			countDisplayRow++

			if countDisplayRow == 3 {
				countAccountNumber++
				countDisplayRow = 0
			}
			if countAccountNumber == 9 {
				countAccountNumber = 0
			}

		}
		fmt.Println()
	}

	fmt.Printf("\nLa estructura es:\n %v\n", accountNumber)

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
