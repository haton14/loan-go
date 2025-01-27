package loango

import (
	"github.com/quagmt/udecimal"
)

// 元金
type Principal = udecimal.Decimal

// 金利(年利)(%)
type InterestRate = udecimal.Decimal

// 期間(年)
type Term = int32

// for level payment
func CalculateMonthlyPayment(p Principal, r InterestRate, t Term) udecimal.Decimal {
	// 月利
	ratePerMonth, err := r.Div64(12)
	if err != nil {
		panic(err)
	}
	// %から数値に変換
	rateMonth, err := ratePerMonth.Div64(100)
	if err != nil {
		panic(err)
	}
	// 元利金等の月々の支払い額求めるぞ
	// 元金 * 数値の月値
	temp1 := p.Mul(rateMonth)
	// ( 1 + 数値の月値 )^ 期間
	temp2, err := rateMonth.Add64(1).PowInt32(t)
	if err != nil {
		panic(err)
	}
	result, err := temp1.Mul(temp2).Div(temp2.Sub64(1))
	if err != nil {
		panic(err)
	}
	return result
}
