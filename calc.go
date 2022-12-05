package metermaid

import (
	"sync/atomic"

	"github.com/montanaflynn/stats"
)

func (m *Metermaid) Calc() *Metrics {
	var data stats.Float64Data
	metrics := &Metrics{}
	metrics.DataPoints = make(map[int64]int)
	if atomic.LoadUint64(&m.Count) == 0 {
		return metrics
	}
	m.Lock()

	// Set our start and end times
	metrics.StartTime = m.Times[0]
	metrics.EndTime = metrics.StartTime

	for _, ts := range m.Times {

		// break out if the ts is the default of 0 time
		if ts.UnixNano() < 0 {
			break
		}

		// Set the Endtime
		metrics.EndTime = ts
		metrics.Samples++

		// Bucket this
		bucket := int64(ts.UnixNano()/int64(1000000000*m.SampleSeconds)) * int64(1000000000*m.SampleSeconds)
		metrics.DataPoints[bucket]++
	}

	for _, rate := range metrics.DataPoints {
		data = append(data, float64(rate))
	}

	metrics.MinRate, _ = stats.Min(data)
	metrics.MaxRate, _ = stats.Max(data)
	metrics.MedianRate, _ = stats.Median(data)
	metrics.P95Rate, _ = stats.Percentile(data, 95)
	metrics.P99Rate, _ = stats.Percentile(data, 99)
	metrics.P999Rate, _ = stats.Percentile(data, 99.9)

	m.Unlock()
	return metrics
}
