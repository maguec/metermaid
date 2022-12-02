package metermaid_test

import (
	"github.com/maguec/metermaid"
	"testing"
	"time"
)

// TestAddingTimesShort: Make sure that if we don't add enough times that the slice is set to 0 time
func TestAddingTimesShort(t *testing.T) {
	mm := metermaid.New(&metermaid.Config{Size: 3})
	for i := 0; i < 2; i++ {
		mm.AddTS(time.Now())
		time.Sleep(100 * time.Millisecond)
	}

  if mm.Times[2].Unix() > 0 {
    t.Fatalf("The last timestamp should be beginning of time but says: %+v\n", mm.Times[2])
  }

  if mm.Calc().Samples !=2 {
    t.Fatalf("We don't have the right number of samples: got %d - want %d\n", mm.Calc().Samples, 2)
  }
}

// TestAddingTimesComplete: Last time needs to be bigger than now
func TestAddingTimesComplete(t *testing.T) {
	mm := metermaid.New(&metermaid.Config{Size: 3})
  now := time.Now()
	for i := 0; i < 3; i++ {
		mm.AddTS(time.Now())
		time.Sleep(100 * time.Millisecond)
	}

  if mm.Times[2].Unix() > now.Unix() {
    t.Fatalf("The last timestamp should be after %+v but is %+v\n", now, mm.Times[2])
  }

  if mm.Calc().Samples !=3 {
    t.Fatalf("We don't have the right number of samples: got %d - want %d\n", mm.Calc().Samples, 3)
  }
}
