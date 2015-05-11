package histoutils

import (
	"testing"
)

func TestCreateHistogram(t *testing.T) {
	h := Histogram{BinSize: 10, Bins: make([]Bin, 5)}
	if size := len(h.Bins); size != 5 {
		t.Error("Bin size needs to be five")
	}
}

func TestAddBins(t *testing.T) {
	h := Histogram{BinSize: 10}
	if size := len(h.Bins); size != 0 {
		t.Error("Bin size needs to zero")
	}

	h.Bins = append(h.Bins, Bin{100, 20})
	h.Bins = append(h.Bins, Bin{200, 40})

	if size := len(h.Bins); size != 2 {
		t.Error("Bin should have appended 2 items")
	}

	if h.Bins[1].Y != 40 {
		t.Error("Bin appending failed")
	}
}

func TestAddBinsWithMethod(t *testing.T) {
	h := Histogram{BinSize: 10}
	if size := len(h.Bins); size != 0 {
		t.Error("Bin size needs to zero")
	}

	h.AddBin(Bin{100, 20})
	h.AddBin(Bin{200, 40})

	if size := len(h.Bins); size != 2 {
		t.Error("Bin should have appended 2 items")
	}

	if h.Bins[1].Y != 40 {
		t.Error("Bin appending failed")
	}
}
