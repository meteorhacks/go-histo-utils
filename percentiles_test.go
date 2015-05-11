package histoutils

import (
	"testing"
)

func TestSinglePercentile(t *testing.T) {
	h := Histogram{BinSize: 10}
	h.AddBin(Bin{0, 5})
	h.AddBin(Bin{10, 20})
	h.AddBin(Bin{20, 5})

	pecentiles := GetPercentiles(&h, []int{50})
	if pecentiles[50] != 10 {
		t.Error("Wrong percentile")
	}
}
