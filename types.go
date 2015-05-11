package histoutils

type Bin struct {
	X int
	Y int
}

type Histogram struct {
	BinSize int
	Bins    []Bin
}

func (h *Histogram) AddBin(b Bin) {
	h.Bins = append(h.Bins, b)
}
