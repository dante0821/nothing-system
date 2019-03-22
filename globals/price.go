package globals

import (
	"github.com/shopspring/decimal"
	"math/big"
)

const (
	exp = 100 * 10000 * 10000
)

var (
	e = decimal.New(exp, 0)
)

func PriceToString(i uint64) string {
	bi := big.NewInt(0).SetUint64(i)
	di := decimal.NewFromBigInt(bi, 0).Div(e)
	return di.String()
}

func AmountToString(i uint64) string {
	bi := big.NewInt(0).SetUint64(i)
	di := decimal.NewFromBigInt(bi, 0).Div(e)
	return di.String()
}

func StringToAmount(s string) (uint64, error) {
	v1, err := decimal.NewFromString(s)
	if err != nil {
		return 0, err
	}
	return uint64(v1.Mul(e).IntPart()), nil
}

func StringToPrice(s string) (uint64, error) {
	v1, err := decimal.NewFromString(s)
	if err != nil {
		return 0, err
	}
	return uint64(v1.Mul(e).IntPart()), nil
}

func PriceMulAmount(price int64, amount int64, ratio string) int64 {
	priceBig := decimal.New(price, 0)
	amountBig := decimal.New(amount, 0)

	result := priceBig.Mul(amountBig)
	if ratio != "" {
		ratioBig, _ := decimal.NewFromString(ratio)
		result = result.Mul(ratioBig)
	}
	result = result.Div(e)
	return result.IntPart()
}

func AllPriceDivPrice(allprice int64, amount int64, ratio string) int64 {
	allpriceBig := decimal.New(allprice, 0)
	amountBig := decimal.New(amount, 0)

	result := allpriceBig.Mul(e)
	result = result.Div(amountBig)
	if ratio != "" {
		ratioBig, _ := decimal.NewFromString(ratio)
		result = result.Mul(ratioBig)
	}
	return result.IntPart()
}
