package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nahidhasan98/calculator/helper"
	"github.com/nahidhasan98/calculator/validator"
)

func main() {
	fmt.Println("==========================================================================================")
	fmt.Println("===> Calculator app is running. A very simple console calculator app.                 <===")
	fmt.Println("===> It only supports integer and float value and 4 opertor i.e. +, -, *, /.          <===")
	fmt.Println("===> Acceptable character list: 0123456789.+-*/                                       <===")
	fmt.Println("===> Just type (i.e. 8+5.7, two number & one operator) and press enter to get output. <===")
	fmt.Println("==========================================================================================")

	for {
		fmt.Print("Input: ")

		reader := bufio.NewReader(os.Stdin)
		s, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Something went wrong while taking input. Try again...")
			continue
		}
		s = helper.RemoveNewLine(s)
		s = helper.IgnoreSpace(s)

		if !validator.IsValid(s) {
			fmt.Println("Input not acceptable. Try again...")
			continue
		}

		ans, err := helper.GetResult(s)
		if err != nil {
			fmt.Println("Something went wrong while doing calculation. Try again...")
			continue
		}

		fmt.Println(s, "=", ans)
	}
}
