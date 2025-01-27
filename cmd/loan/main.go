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
	monthlyPayment := loango.CalculateMonthlyPayment(principal, interestRate, loango.Term(termVar))
	fmt.Println("月の支払額:", monthlyPayment)
	currentPrincipal := principal
	fmt.Println("元金, 利息")
	for i := loango.Term(1); i <= loango.Term(termVar); i++ {
		remaingPrincipal := loango.CalculateRemainingPrincipal(principal, interestRate, monthlyPayment, i)
		principalPayment := currentPrincipal.Sub(remaingPrincipal)
		InterestPayment := monthlyPayment.Sub(principalPayment)
		fmt.Printf("%s, %s\n", principalPayment, InterestPayment)
		currentPrincipal = remaingPrincipal
	}
}
