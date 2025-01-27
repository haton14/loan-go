package main

import (
	"fmt"

	loango "github.com/haton14/loan-go"
	"github.com/quagmt/udecimal"
)

func main() {
	principal, err := udecimal.NewFromUint64(8000*10000, 0)
	if err != nil {
		panic(err)
	}
	interestRate, err := udecimal.NewFromFloat64(0.625)
	if err != nil {
		panic(err)
	}
	result := loango.CalculateMonthlyPayment(principal, interestRate, 40*12)
	fmt.Println(result)
}
