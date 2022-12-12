package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"time"
)

func main() {

	var ms runtime.MemStats
	var gc [50]int
	var gcCalc [50]int

	const threads = 8                  // Saturate CPU
	const totalMem = 1000 * 1000 * 256 // 256MB
	const memSize = totalMem / threads // Divide memory
	const threadTime = 50              // ms

	// Create threads that do concurrent work
	for i := 0; i < threads; i++ {
		go func() {
			for {
				mem := make([]byte, memSize)
				time.Sleep(time.Millisecond * threadTime)
				runtime.KeepAlive(mem) // Keep alive beyond scope
			}
		}()
	}

	for i := 0; i < 50; i++ {
		print("\rPercentage: ", i*2)
		print("%")
		debug.SetGCPercent((i + 1) * 10)
		time.Sleep(time.Millisecond * 5000) // Do 5 seconds of work
		runtime.ReadMemStats(&ms)
		gc[i] = int(ms.NumGC)
	}

	println("\rPercentage: 100%")

	for i := 0; i < 50; i++ {
		if i != 0 {
			gcCalc[i] = gc[i] - gc[i-1]
		} else {
			gcCalc[0] = gc[0]
		}
		fmt.Println(gcCalc[i])
	}
}
