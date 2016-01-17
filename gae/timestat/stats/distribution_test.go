package stats

import (
	"testing"

	m "timestat/model"
)

func TestNormalize(t *testing.T) {
	dist := newDistribution(0.0)
	p := dist.MinuteProbabilities
	p[0] = 99
	Normalize(dist)
	if p[0] != 1.0 {
		t.Errorf("Expected 1.0 but p[0] was %v", p[0])
	}
	if p[1439] != 0.0 {
		t.Errorf("Expected 0.0 but p[1439] was %v", p[1439])
	}
	p[1439] = 1.0
	Normalize(dist)
	if p[0] != 0.5 {
		t.Errorf("Expected 0.5 but p[0] was %v", p[0])
	}
	if p[1439] != 0.5 {
		t.Errorf("Expected 0.5 but p[1439] was %v", p[1439])
	}
}

func newDistribution(base float64) *m.Distribution {
	dist := &m.Distribution{
		MinuteProbabilities: make([]float64, 1440),
	}
	for i := range dist.MinuteProbabilities {
		dist.MinuteProbabilities[i] = base
	}
	return dist
}
