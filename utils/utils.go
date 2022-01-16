package utils

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strings"

	"ocr/schema"
)

func PrintNumber(input []string) string {
	return strings.Join(input, "")
}

func PutInStruct() {

	countAccountNumber := 0
	countDisplayRow := 0

	rows := ReadFileNuberAccounts()

	for rowCount, row := range rows {

		whiteLine := len(row)
		segments := bytes.Runes([]byte(row))

		if whiteLine != 0 {

			for _, segment := range segments {

				if rowCount == 0 {
					schema.ACTNumber.AccountNumber[countAccountNumber].DisplayRow1[countDisplayRow] = string(segment)
				}
				if rowCount == 1 {
					schema.ACTNumber.AccountNumber[countAccountNumber].DisplayRow2[countDisplayRow] = string(segment)
				}
				if rowCount == 2 {
					schema.ACTNumber.AccountNumber[countAccountNumber].DisplayRow3[countDisplayRow] = string(segment)
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

		}
	}

}

func GetNumber(input schema.AccountNumbers) []string {

	var str []string

	for _, DisplayInputNumber := range input.AccountNumber {
		str = append(str, ValidateDisplay(DisplayInputNumber))
	}
	return str

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

func ValidateDisplay(input schema.Display) string {

	switch input {
	case schema.Display0:
		return "0"
	case schema.Display1:
		return "1"
	case schema.Display2:
		return "2"
	case schema.Display3:
		return "3"
	case schema.Display4:
		return "4"
	case schema.Display5:
		return "5"
	case schema.Display6:
		return "6"
	case schema.Display7:
		return "7"
	case schema.Display8:
		return "8"
	case schema.Display9:
		return "9"
	default:
		return "?"
	}

}
