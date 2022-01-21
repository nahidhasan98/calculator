package validator

func firstAndLastChar(s string) bool {
	if string(s[0]) != "." && (string(s[0]) < "0" && string(s[0]) > "9") {
		return false
	}
	if string(s[len(s)-1]) != "." && (string(s[len(s)-1]) < "0" && string(s[len(s)-1]) > "9") {
		return false
	}
	return true
}

func singleOperator(s string) bool {
	counter := 0

	for i := 0; i < len(s); i++ {
		if string(s[i]) == "+" || string(s[i]) == "-" || string(s[i]) == "*" || string(s[i]) == "/" {
			counter++
		}
	}

	return counter == 1
}

func validDigit(s string) bool {
	for i := 0; i < len(s); i++ {
		if string(s[i]) != "+" && string(s[i]) != "-" && string(s[i]) != "*" && string(s[i]) != "/" && string(s[i]) != "." && (string(s[i]) < "0" || string(s[i]) > "9") {
			return false
		}
	}

	return true
}

func IsValid(s string) bool {
	if firstAndLastChar(s) && singleOperator(s) && validDigit(s) {
		return true
	}
	return false
}
