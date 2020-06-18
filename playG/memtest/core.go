package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {

	log.Println("At begin:")
	PrintMemUsage()

	log.Println("Forece GC:")
	var overall1 [][]int
	for i := 0; i < 999; i++ {
		a := make([]int, 0, 999999)
		overall1 = append(overall1, a)
		a = nil
	}

	// Force GC to clear up, should see a memory drop
	runtime.GC()
	PrintMemUsage()
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
