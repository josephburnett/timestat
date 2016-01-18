package stats

import (
	"errors"

	m "timestat/model"

	s "github.com/josephburnett/stats"
)

// Update records the new minute and updates summary statistics.
func Update(dist *m.Distribution, minute int32) {
	counts := dist.MinuteCounts
	if minute < int32(0) || minute > int32(len(counts)-1) {
		panic("Minute out of bounds.")
	}
	counts[minute]++
	dist.SampleCount++
	updateDerived(dist)
}

func updateDerived(dist *m.Distribution) {
	samples := make([]float64, 0, dist.SampleCount)
	for minute, count := range dist.MinuteCounts {
		for i := int32(0); i < count; i++ {
			samples = append(samples, float64(minute))
		}
	}
	dist.Mean, _ = s.Mean(samples)
	dist.Median, _ = s.Median(samples)
	dist.StandardDeviation, _ = s.StandardDeviationPopulation(samples)
}

// Minute extracts the minute value from a running timer.
func Minute(timer *m.RunningTimer) (int32, error) {
	if timer.End.Before(timer.Start) {
		return 0, errors.New("End is before start.")
	}
	duration := timer.End.Sub(timer.Start)
	minutes := int32(duration.Minutes())
	if minutes > 1439 {
		return 0, errors.New("Timer took too long.")
	}
	return minutes, nil
}
