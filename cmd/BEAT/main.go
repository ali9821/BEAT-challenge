package main

import (
	"BEAT/cmd"
	"BEAT/pkg/model"
	"BEAT/tools"
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("Process started")
	startTime := time.Now()
	ctx := context.Background()
	factory, err := cmd.NewFactory()
	if err != nil {
		log.Panic("error creating factory", err)
	}
	data.CsvReader(factory.Config)
	doneChannel := factory.Chan

	for _, runner := range factory.PipelineStages {
		go func(runner model.Runner) {
			err := runner.Run(ctx)
			if err != nil {
				log.Fatal(err)
			}
		}(runner)
	}
	<-doneChannel
	fmt.Println("Process ended")
	fmt.Println(time.Since(startTime))

}
