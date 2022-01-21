package main

func calculate(a, b float64, op string) float64 {
	var value float64

	switch op {
	case "+":
		value = a + b
	case "-":
		value = a - b
	case "*":
		value = a * b
	case "/":
		value = a / b
	}

	return value
}
