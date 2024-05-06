package fileHandler

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(balance float64, fileName string) {
	balanceText := fmt.Sprint(balance)

	os.WriteFile(fileName, []byte(balanceText), 0644)
}

func GetFloatFromFile(fileName string) (floatVal float64, err error) {
	byteData, err := os.ReadFile(fileName)
	amountAsString := string(byteData)

	if err != nil {
		return getDefaultFloatAndError("failed to find file")
	}

	floatVal, err = strconv.ParseFloat(amountAsString, 64)
	if err != nil {
		return getDefaultFloatAndError("failed to parse stored value")
	}
	return
}

func getDefaultFloatAndError(errorMessage string) (defaultFloat float64, err error) {
	err = errors.New(errorMessage)
	defaultFloat = 0
	return
}
