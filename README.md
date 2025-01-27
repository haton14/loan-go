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

## 今のところは元利金等返済での、毎月の支払額, 毎月の元金額, 利息額を求める機能だけ対応しています
$ loan --principal 10000000 --interestRate 0.625 --term 420
月の支払額: 26514.7336202080067668863
元金, 利息
21306.4002868746737668309, 5208.3333333333330000554

## 住宅ローン金利の上がり幅に対して125%ルール内に収まっているか確認するなどに使えます
$ loan --principal 10000000 --interestRate 1.625 --term 420
31234.4320858842827750994

$ python
>>> 31234.4320858842827750994 / 26514.7336202080067668863
1.1780028618533507
```
