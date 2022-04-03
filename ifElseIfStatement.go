package main

import (
	"fmt"
	"strconv"
)

func evaluateValue(value int64) string {
	replyMessage := ""
	if value%5 == 0 && value%6 == 0 {
		replyMessage = "The number is divisible by both 5 and 6."
	} else {
		replyMessage = "The number is not divisible by both 5 and 6."
	}

	//Do not remove these lines
	fmt.Println((replyMessage))
	return replyMessage

}

func main() {
	var compareValue string
	fmt.Println("Enter an integer: ")
	fmt.Scanln(&compareValue)

	//conversion of value
	valueInt, _ := strconv.ParseInt(compareValue, 10, 0)
	evaluateValue(valueInt)

}
