package histoutils

import (
	"math"
)

type PercentileInfo struct {
	p int
	i int
}

func GetPercentiles(h *Histogram, percentiles []int) map[int]float64 {
	finalResults := make(map[int]float64)
	totalItems := 0
	for _, bin := range h.Bins {
		totalItems += bin.Y
	}

	percentileMapper := make([]PercentileInfo, len(percentiles))
	for _, percentile := range percentiles {
		percentileMapper[0].p = percentile
		percentileMapper[0].i = int(math.Ceil(float64(percentile) / 100.0 * float64(totalItems)))
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
				finalResults[percentileInfo.p] = float64(bin.X)
			} else {
				percentileMapper = append([]PercentileInfo{percentileInfo}, percentileMapper...)
				break
			}
		}
	}

	return finalResults
}
