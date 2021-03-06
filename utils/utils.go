package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"encoding/base64"
	"encoding/json"
	"ocr/config"
	"ocr/schema"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PrintNumber(input []string) string {
	return strings.Join(input, "")
}

func GetArrayAccounts(inputs []string) []schema.AccountNumbers {

	var (
		accounts           []schema.AccountNumbers
		countAccountNumber int = 0
		countDisplayRow    int = 0
		rowCount           int = 0
	)

	rows := inputs

	for _, row := range rows {

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
			rowCount++
			if rowCount == 3 {
				rowCount = 0
				accounts = append(accounts, schema.ACTNumber)
			}
		}
	}
	return accounts
}

func GetArrayAccountsFromFile() []schema.AccountNumbers {

	var (
		accounts           []schema.AccountNumbers
		countAccountNumber int = 0
		countDisplayRow    int = 0
		rowCount           int = 0
	)

	rows := ReadFileNuberAccounts()

	for _, row := range rows {

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
			rowCount++
			if rowCount == 3 {
				rowCount = 0
				accounts = append(accounts, schema.ACTNumber)
			}
		}
	}
	return accounts
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

func ArrayStringsValidateWithInput(inputs []string) []string {

	var (
		inputNumbers []string = ArrayDisplaysToArrayStrings()
		output       []string
	)

	for _, number := range inputNumbers {
		character := CheckCharacterXInString(number)
		if character {
			flag := fmt.Sprintf("%v ILL", number)
			output = append(output, flag)
		} else {
			flag := ValidateMod(GetMod(SumArrayString(StringToArray(number))))
			output = append(output, fmt.Sprintf("%v %v", number, flag))
		}
	}
	return output
}

func ArrayStringsValidate() []string {

	var (
		inputNumbers []string = ArrayDisplaysToArrayStrings()
		output       []string
	)

	for _, number := range inputNumbers {
		character := CheckCharacterXInString(number)
		if character {
			flag := fmt.Sprintf("%v ILL", number)
			output = append(output, flag)
		} else {
			flag := ValidateMod(GetMod(SumArrayString(StringToArray(number))))
			output = append(output, fmt.Sprintf("%v %v", number, flag))
		}
	}
	return output
}

func ArrayDisplaysToArrayStringsWithInput(inputs []string) []string {

	var (
		validate, output []string
	)

	for _, account := range GetArrayAccounts(inputs) {
		for n := 0; n < 9; n++ {
			validate = append(validate, ValidateDisplay(account.AccountNumber[n]))
		}
		output = append(output, strings.Join(validate, ""))
		validate = nil
	}
	return output
}

func ArrayDisplaysToArrayStrings() []string {

	var (
		validate, output []string
	)

	for _, account := range GetArrayAccountsFromFile() {
		for n := 0; n < 9; n++ {
			validate = append(validate, ValidateDisplay(account.AccountNumber[n]))
		}
		output = append(output, strings.Join(validate, ""))
		validate = nil
	}
	return output
}

func CheckCharacterXInString(input string) bool {
	var output bool = false
	if strings.Contains(input, "?") {
		output = true
	}
	return output
}

func StringToArray(input string) []string {
	output := strings.Split(input, "")
	return output
}

func SumArrayString(input []string) int {
	n := 10
	sum := 0
	for _, digit := range input {
		n--
		sum += stringToInt(digit) * n
	}
	return sum
}

func stringToInt(input string) int {
	output, _ := strconv.Atoi(input)
	return output
}

func GetMod(input int) int {
	output := input % 11
	return output
}

func ValidateMod(input int) string {
	output := ""
	if input == 0 {
		output = "OK"
	} else {
		output = "ERR"
	}
	return output
}

func DecodeBase64String(input string) string {
	output, _ := base64.StdEncoding.DecodeString(input)
	fmt.Println(string(output))
	return string(output)
}

func SplitStringByCharToArray(input string, splitChar string) []string {
	output := strings.Split(input, splitChar)
	return output
}

func CreateLog(endpoint string, method string, input gin.H, output gin.H) {

	id := uuid.New()
	logDB := []schema.LogDB{
		{
			ID:       id,
			EndPoint: endpoint,
			Method:   method,
			Input:    StructToJson(input),
			Output:   StructToJson(output),
		},
	}

	_, err := config.GetDB().NamedExec(`
		INSERT
		INTO log_db
			(id, endpoint, method, json_input, json_output)
		VALUES
			(:id, :endpoint, :method, :json_input, :json_output)`, logDB)
	if err != nil {
		panic(err.Error())
	}

}

func StructToJson(input interface{}) []byte {

	output, err := json.Marshal(input)
	if err != nil {
		panic(err.Error())
	}
	return output

}

func GetLogRows() []schema.LogDB {

	logDB := []schema.LogDB{}
	err := config.GetDB().Select(&logDB, "SELECT id, endpoint, method, json_input, json_output, created_on FROM log_db")
	if err != nil {
		log.Fatalln(err)
	}
	return logDB

}
