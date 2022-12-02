package metermaid

import (
	"sync"
	"sync/atomic"
	"time"
)

type Config struct {
	Size  int
	HBins int
}

type tsSlice []time.Time

type Metrics struct {
  Samples int
  StartTime time.Time
  EndTime time.Time
}

type Metermaid struct {
  sync.Mutex
  Size uint64
  Times tsSlice
  Count uint64
  HBins int
}

// Functions for sorting
func (p tsSlice) Len() int           { return len(p) }
func (p tsSlice) Less(i, j int) bool { return p[i].UnixNano() < p[j].UnixNano() }
func (p tsSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Startup
func New(c *Config) *Metermaid {
	var hSize int
	if c.HBins != 0 {
		hSize = c.HBins
	} else {
		hSize = 10
	}

	return &Metermaid{
		Size:  uint64(c.Size),
		Times: make([]time.Time, c.Size),
		HBins: hSize,
	}
}

// Add at Time 
func (m *Metermaid) AddTS(t time.Time){
  m.Times[(atomic.AddUint64(&m.Count, 1)-1)%m.Size] = t
}

// Add at Now 
func (m *Metermaid) Add(){
  m.Times[(atomic.AddUint64(&m.Count, 1)-1)%m.Size] = time.Now()
}
