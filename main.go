package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type AccountNumbers struct {
	AccountNumber [9]Display
}

type Display struct {
	DisplayRow1 [3]string
	DisplayRow2 [3]string
	DisplayRow3 [3]string
}

var (
	Display0 = Display{
		DisplayRow1: [3]string{" ", "_", " "},
		DisplayRow2: [3]string{"|", " ", "|"},
		DisplayRow3: [3]string{"|", "_", "|"},
	}
	Display1 = Display{
		DisplayRow1: [3]string{" ", " ", " "},
		DisplayRow2: [3]string{" ", " ", "|"},
		DisplayRow3: [3]string{" ", " ", "|"},
	}
	Display2 = Display{
		DisplayRow1: [3]string{" ", "_", " "},
		DisplayRow2: [3]string{" ", "_", "|"},
		DisplayRow3: [3]string{"|", "_", " "},
	}
	Display3 = Display{
		DisplayRow1: [3]string{" ", "_", " "},
		DisplayRow2: [3]string{" ", "_", "|"},
		DisplayRow3: [3]string{" ", "_", "|"},
	}
	Display4 = Display{
		DisplayRow1: [3]string{" ", " ", " "},
		DisplayRow2: [3]string{"|", "_", "|"},
		DisplayRow3: [3]string{" ", " ", "|"},
	}
	Display5 = Display{
		DisplayRow1: [3]string{" ", "_", " "},
		DisplayRow2: [3]string{"|", "_", " "},
		DisplayRow3: [3]string{" ", "_", "|"},
	}
	Display6 = Display{
		DisplayRow1: [3]string{" ", "_", " "},
		DisplayRow2: [3]string{"|", "_", " "},
		DisplayRow3: [3]string{"|", "_", "|"},
	}
	Display7 = Display{
		DisplayRow1: [3]string{" ", "_", " "},
		DisplayRow2: [3]string{" ", " ", "|"},
		DisplayRow3: [3]string{" ", " ", "|"},
	}
	Display8 = Display{
		DisplayRow1: [3]string{" ", "_", " "},
		DisplayRow2: [3]string{"|", "_", "|"},
		DisplayRow3: [3]string{"|", "_", "|"},
	}
	Display9 = Display{
		DisplayRow1: [3]string{" ", "_", " "},
		DisplayRow2: [3]string{"|", "_", "|"},
		DisplayRow3: [3]string{" ", "_", "|"},
	}
	accountNumber AccountNumbers
)

func main() {
	r := gin.Default()
	r.GET("/number", func(c *gin.Context) {
		PutInStruct()
		c.JSON(200, gin.H{
			"number": PrintNumber(GetNumber(accountNumber)),
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func PrintNumber(input []string) string {
	return strings.Join(input, "")
}

func PutInStruct() {

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
	}

}

func GetNumber(input AccountNumbers) []string {

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

func ValidateDisplay(input Display) string {

	switch input {
	case Display0:
		return "0"
	case Display1:
		return "1"
	case Display2:
		return "2"
	case Display3:
		return "3"
	case Display4:
		return "4"
	case Display5:
		return "5"
	case Display6:
		return "6"
	case Display7:
		return "7"
	case Display8:
		return "8"
	case Display9:
		return "9"
	default:
		return "?"
	}

}
