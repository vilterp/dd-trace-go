package main

import (
	"time"

	"gopkg.in/DataDog/dd-trace-go.v1/internal/log"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
)

func main() {
	err := profiler.Start(
		profiler.WithProfileTypes(profiler.ExternalProfile),
		profiler.WithService("rai-server"),
		profiler.WithEnv("laptop"), // TODO: make configurable
	)
	defer profiler.Stop()

	if err != nil {
		log.Error("profiler couldn't start", err)
	}

	for {
		time.Sleep(10 * time.Second)
		log.Info("tick")
	}
}
