package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 1; i <= 10; i++ {
		a := make([]*int, 10000000*i)
		for j := 1; j <= 101; j++ {
			start := time.Now()
			runtime.GC()
			if j != 1 {
				if j != 101 {
					fmt.Print("", time.Since(start))
					fmt.Print(",")
				} else {
					fmt.Print("", time.Since(start))
				}
			}
		}
		runtime.KeepAlive(a)
		fmt.Println("")
	}
}

func printMemStat(ms runtime.MemStats) {
	runtime.ReadMemStats(&ms)
	fmt.Print(ms.Alloc)
	fmt.Print(",")
	fmt.Print(ms.TotalAlloc)
	fmt.Print(",")
	fmt.Print(ms.NumGC)
	fmt.Print(",")
	fmt.Print(ms.PauseTotalNs)
	fmt.Println("")
}
