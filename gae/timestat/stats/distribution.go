package stats

import (
	m "timestat/model"
)

// Normalize ensures a distribution of probabilities sums to 1.0.
func Normalize(dist *m.Distribution) {
	var sum float64
	for _, v := range dist.MinuteProbabilities {
		sum += v
	}
	for i, v := range dist.MinuteProbabilities {
		dist.MinuteProbabilities[i] = v / sum
	}
}
