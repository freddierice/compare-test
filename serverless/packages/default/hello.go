package main

import (
	"fmt"
	"strconv"
)

// Main represents the function that will run.
func Main(args map[string]interface{}) map[string]interface{} {
	aStr, ok := args["a"].(string)
	if !ok {
		aStr = "0"
	}
	bStr, ok := args["b"].(string)
	if !ok {
		bStr = "0"
	}

	a, _ := strconv.Atoi(aStr)
	b, _ := strconv.Atoi(bStr)

	msg := make(map[string]interface{})
	msg["body"] = fmt.Sprintf("%v", a+b)
	return msg
}
