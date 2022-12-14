package main

import (
	"fmt"
	currency2 "golang/internal/03_data_processing/currency"
)

func main() {
	// 15달러 93센트의 사용자 입력 값에서 시작한다.
	userInput := "15.93"

	pennies, err := currency2.ConvertStringDollarsToPennies(userInput)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User input converted to %d pennies\n", pennies)

	// 15센트를 더한다.
	pennies += 15

	dollars := currency2.ConvertPenniesToDollarString(pennies)
	fmt.Printf("Added 15 cents, new values is %s dollars\n", dollars)
}
