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

type MonthlyPayment = udecimal.Decimal

type RemainingPrincipal = udecimal.Decimal

// for level payment
func CalculateMonthlyPayment(p Principal, r InterestRate, t Term) MonthlyPayment {
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
	// 元金 * 数値の月利
	temp1 := p.Mul(rateMonth)
	// ( 1 + 数値の月利 )^ 期間
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

// for level payment
func CalculateRemainingPrincipal(p Principal, r InterestRate, x MonthlyPayment, currentMonth Term) RemainingPrincipal {
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
	// ( 1 + 数値の月値 )^ 経過月数
	temp1, err := rateMonth.Add64(1).PowInt32(currentMonth)
	if err != nil {
		panic(err)
	}
	// 元金 * ( 1 + 数値の月利 )^ 経過月数
	temp2 := p.Mul(temp1)
	// 月の支払い額 - 月の支払い額 * ( 1 + 数値の月利 )^ 経過月数 )
	temp3 := x.Sub(x.Mul(temp1))
	// temp3 / 数値の月利
	temp4, err := temp3.Div(rateMonth)
	if err != nil {
		panic(err)
	}
	return temp2.Add(temp4)
}
