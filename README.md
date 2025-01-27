# loan-go
ローンの計算

## usage

```bash
$ go install github.com/haton14/loan-go/cmd/loan@v0.0.1

$ loan -help     
Usage of loan:
  -interestRate float
        金利(年利)(%) (default 1)
  -principal int
        元金 - 借りる金額 (単位: 円) (default 10000000)
  -term int
        期間(月) (default 420)

## 今のところは毎月の支払額を求める機能だけ対応しています
$ loan --principal 10000000 --interestRate 0.625 --term 420
26514.7336202080067668863
```
