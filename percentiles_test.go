package histoutils

import (
	"testing"
)

func TestSinglePercentile(t *testing.T) {
	h := Histogram{BinSize: 10}
	h.AddBin(Bin{0, 5})
	h.AddBin(Bin{10, 20})
	h.AddBin(Bin{20, 5})

	percentiles := GetPercentiles(&h, []int{50}, false)
	if percentiles[50] != 10 {
		t.Error("Wrong percentile")
	}
}

func TestMultiplePercentiles(t *testing.T) {
	h := Histogram{BinSize: 10}
	h.AddBin(Bin{0, 5})
	h.AddBin(Bin{10, 20})
	h.AddBin(Bin{20, 5})

	percentiles := GetPercentiles(&h, []int{50, 95}, false)
	if percentiles[50] != 10 {
		t.Error("Wrong percentile")
	}

	if percentiles[95] != 20 {
		t.Error("Wrong percentile")
	}
}

func TestUnorderedMultiplePercentiles(t *testing.T) {
	h := Histogram{BinSize: 10}
	h.AddBin(Bin{0, 5})
	h.AddBin(Bin{10, 20})
	h.AddBin(Bin{20, 5})

	percentiles := GetPercentiles(&h, []int{95, 50}, false)
	if percentiles[50] != 10 {
		t.Error("Wrong percentile")
	}

	if percentiles[95] != 20 {
		t.Error("Wrong percentile")
	}
}

func TestMultiplePercentilesWithPerbin(t *testing.T) {
	h := Histogram{BinSize: 100}
	h.AddBin(Bin{100, 10})
	h.AddBin(Bin{200, 80})
	h.AddBin(Bin{300, 8})
	h.AddBin(Bin{400, 2})

	percentiles := GetPercentiles(&h, []int{50, 95, 99}, true)
	if percentiles[50] != 250 {
		t.Error("Wrong percentile")
	}

	if percentiles[95] != 362.5 {
		t.Error("Wrong percentile")
	}

	if percentiles[99] != 450 {
		t.Error("Wrong percentile")
	}
}
