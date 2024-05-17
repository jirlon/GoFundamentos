package main

import (
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementar o contador 3 vezes resulta no valor 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		if counter.Value() != 3 {
			t.Errorf("got %d, want %d", counter.Value(), 3)
		}
	})
}
