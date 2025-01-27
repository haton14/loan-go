package main

import (
	"flag"
	"fmt"

	loango "github.com/haton14/loan-go"
	"github.com/quagmt/udecimal"
)

func main() {
	var (
		principalVar    int64
		interestRateVar float64
		termVar         int
	)
	flag.Int64Var(&principalVar, "principal", 1000*10000, "元金 - 借りる金額 (単位: 円)")
	flag.Float64Var(&interestRateVar, "interestRate", 1.0, "金利(年利)(%)")
	flag.IntVar(&termVar, "term", 35*12, "期間(月)")
	flag.Parse()
	principal, err := udecimal.NewFromInt64(principalVar, 0)
	if err != nil {
		panic(err)
	}
	interestRate, err := udecimal.NewFromFloat64(interestRateVar)
	if err != nil {
		panic(err)
	}
	result := loango.CalculateMonthlyPayment(principal, interestRate, loango.Term(termVar))
	fmt.Println(result)
}
