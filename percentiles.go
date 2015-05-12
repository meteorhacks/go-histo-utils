package histoutils

import (
	"math"
	"sort"
)

type PercentileInfo struct {
	p int
	i int
}

func GetPercentiles(h *Histogram, percentiles []int, perbinCalculation bool) map[int]float64 {
	finalResults := make(map[int]float64)
	totalItems := 0
	for _, bin := range h.Bins {
		totalItems += bin.Y
	}

	sort.Ints(percentiles)
	percentileMapper := make([]PercentileInfo, len(percentiles))
	for lc, percentile := range percentiles {
		percentileMapper[lc].p = percentile
		percentileMapper[lc].i = int(math.Ceil(float64(percentile) / 100.0 * float64(totalItems)))
	}

	itemsUpto := 0
	for _, bin := range h.Bins {
		beginItems := itemsUpto
		endItems := itemsUpto + bin.Y
		itemsUpto = endItems

		for len(percentileMapper) > 0 {
			percentileInfo := percentileMapper[0]
			percentileMapper = percentileMapper[1:]

			if percentileInfo.i > beginItems && percentileInfo.i <= endItems {
				percentile := float64(bin.X)
				if perbinCalculation {
					diff := float64(percentileInfo.i - beginItems)
					binsPerItem := float64(h.BinSize) / float64(bin.Y)
					percentile = float64(bin.X) + (diff * binsPerItem)
				}
				finalResults[percentileInfo.p] = percentile
			} else {
				percentileMapper = append([]PercentileInfo{percentileInfo}, percentileMapper...)
				break
			}
		}
	}

	return finalResults
}
