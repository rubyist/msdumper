package msdumper

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"sync/atomic"
	"time"
)

var (
	running           = int32(0)
	wait              = make(chan int)
	errAlreadyRunning = errors.New("Already running")
)

func Start(file string, period time.Duration) error {
	if !atomic.CompareAndSwapInt32(&running, 0, 1) {
		return errAlreadyRunning
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}

	go func() {
		var s runtime.MemStats
		for {
			select {
			case <-wait:
				return
			default:
				runtime.ReadMemStats(&s)
				fmt.Fprintf(f, "%d %d %d %d %d %d %d %d\n",
					s.Alloc,
					s.TotalAlloc,
					s.Sys,
					s.HeapAlloc,
					s.HeapSys,
					s.HeapIdle,
					s.HeapInuse,
					s.HeapObjects,
				)
				time.Sleep(period)
			}
		}
	}()

	return nil
}

func Stop() {
	if atomic.CompareAndSwapInt32(&running, 1, 0) {
		wait <- 1
	}
}
