package main

import (
	"fmt"
	"log"
	"sort"

	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
)

func main() {
	err := profiler.Start(
		profiler.WithService("go-sorter"),
		profiler.WithEnv("example"),
		profiler.WithVersion("1.0"),
		profiler.WithProfileTypes(
			profiler.CPUProfile,
			profiler.HeapProfile,
			// The profiles below are disabled by default to keep overhead
			// low, but can be enabled as needed.

			// profiler.BlockProfile,
			// profiler.MutexProfile,
			// profiler.GoroutineProfile,
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer profiler.Stop()

	for {
		x := []int{}
		for i := 0; i < 100000000; i++ {
			x = append(x, i)
		}
		sort.Ints(x)
		fmt.Println("sorted")
	}
}
