package helpers

import (
	"github.com/michey/gokkan/data"
	"time"
)

func ReadByTimeWindow(d time.Duration, input <-chan data.Response, output chan<- []data.Response) {
	run := true
	cache := make([]data.Response, 0, 10)

	go func() {
		for run {
			m := <-input
			cache = append(cache, m)
		}
	}()
	time.Sleep(d)
	output <- cache
}
