package utils

import (
	"fmt"
	"os"
	"strings"
)

func ReadData(model string) ([]string, error) {
	urlData := fmt.Sprintf("../../datos/%v.txt", model)
	data, err := os.ReadFile(urlData)
	if err == nil {
		arrStrings := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
		if len(arrStrings) > 0 {
			arrStrings = arrStrings[:len(arrStrings)-1]
		}
		return arrStrings, nil

	} else {
		return nil, fmt.Errorf("Error Read %v", err)
	}
}
