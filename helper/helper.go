package helper

import (
	"strconv"
	"strings"

	"github.com/nahidhasan98/calculator/calculation"
)

func RemoveNewLine(s string) string {
	ss := ""

	for i := 0; i < len(s); i++ {
		if s[i] != 10 && s[i] != 13 { // newline character \r=13 \n=10
			ss += string(s[i])
		}
	}

	return ss
}

func IgnoreSpace(s string) string {
	ss := strings.ReplaceAll(s, " ", "")
	return ss
}

func GetFirstVal(s string) (float64, error) {
	var value string
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "+" || string(s[i]) == "-" || string(s[i]) == "*" || string(s[i]) == "/" {
			break
		}
		value += string(s[i])
	}

	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}

	return val, nil
}

func GetSecondVal(s string) (float64, error) {
	var value string
	flag := false
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "+" || string(s[i]) == "-" || string(s[i]) == "*" || string(s[i]) == "/" {
			flag = true
			continue
		}
		if flag {
			value += string(s[i])
		}
	}

	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}

	return val, nil
}

func GetOperator(s string) string {
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "+" || string(s[i]) == "-" || string(s[i]) == "*" || string(s[i]) == "/" {
			return string(s[i])
		}
	}
	return ""
}

func GetResult(s string) (float64, error) {
	firstVal, err := GetFirstVal(s)
	if err != nil {
		return 0, err
	}

	secondVal, err := GetSecondVal(s)
	if err != nil {
		return 0, err
	}

	operater := GetOperator(s)

	res := calculation.Calculate(firstVal, secondVal, operater)

	return res, nil
}
