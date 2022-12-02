package metermaid

import (
  "sync/atomic"
)

func (m *Metermaid) Calc() *Metrics {
  metrics := &Metrics{}
  if atomic.LoadUint64(&m.Count) == 0 {
		return metrics
	}
  m.Lock()

  // Set our start and end times
  metrics.StartTime = m.Times[0]
  metrics.EndTime = metrics.StartTime

  for _, ts := range(m.Times) {

    // break out if the ts is the default of 0 time
    if ts.UnixNano() < 0 {
      break
    }

    // Set the Endtime
    metrics.EndTime = ts
    metrics.Samples++
  }

  m.Unlock()
  return metrics
}
