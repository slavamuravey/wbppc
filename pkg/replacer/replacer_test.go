package replacer_test

import (
	"testing"

	"github.com/slavamuravey/wbppc/pkg/assert"
	"github.com/slavamuravey/wbppc/pkg/replacer"
)

var calcDiscountPercentTests = []struct {
	fromPrice float64
	toPrice   float64
	exp       int
}{
	{240, 148, 39},
	{367, 138, 63},
	{8500, 3700, 57},
}

func TestCalcDiscountPercent(t *testing.T) {
	for _, e := range calcDiscountPercentTests {
		discount := replacer.CalcDiscountPercent(e.fromPrice, e.toPrice)
		assert.Equal(t, e.exp, discount, "discounts should be equal")
	}
}
