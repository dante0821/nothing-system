package globals

import "testing"

func TestPriceMulAmount(t *testing.T) {
	var a int64 = 20 * exp
	var b int64 = 20 * exp
	c := PriceMulAmount(a, b, "")

	d := AllPriceDivPrice(c, b, "")
	t.Log(d)
}

func TestPriceToString(t *testing.T) {
	a := 1000
	t.Log(PriceToString(uint64(a)))
	t.Log(AmountToString(uint64(a)))
}

func TestPriceToString2(t *testing.T) {
	a := "50.1112"
	t.Log(StringToPrice(a))
	t.Log(StringToAmount(a))
}
