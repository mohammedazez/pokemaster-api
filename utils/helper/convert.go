package helper

import (
	"strconv"
	"strings"
)

func StringToInt64(sValue string, defValue int64) int64 {
	value := defValue
	iValue, err := strconv.ParseInt(sValue, 10, 64)
	if err == nil {
		value = iValue
	}
	return value
}

func StringToInt(sValue string, defValue int) int {
	value := StringToInt64(sValue, int64(defValue))
	return int(value)
}

func StringToBool(sValue string, defValue bool) bool {
	value := defValue
	if sValue == "1" || strings.ToLower(sValue) == "true" {
		value = true
	} else if sValue == "0" || strings.ToLower(sValue) == "false" {
		value = false
	}
	return value
}
