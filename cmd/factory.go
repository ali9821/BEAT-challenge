package cmd

import (
	cfg "BEAT/config"
	"BEAT/pkg/model"
	"BEAT/pkg/service"
	"log"
)

type Factory struct {
	Config         *cfg.Config
	PipelineStages []model.Runner
	Chan           chan bool
}

func NewFactory() (*Factory, error) {
	config := cfg.NewConfig()
	travelChannel := make(chan model.Travel, 1000)
	validRoutes := make(chan model.RouteInfo, 1000)
	done := make(chan bool)

	resposeWriterService := setWriterService(config)

	serviceGenerator, err := service.GenerateTravelService(travelChannel)
	if err != nil {
		log.Fatal(err)
	}

	serviceValidator, err := service.GenerateValidationService(travelChannel, validRoutes)
	if err != nil {
		log.Fatal(err)
	}

	serviceCalculator, err := service.GeneratePriceCalcService(validRoutes, resposeWriterService, done)
	if err != nil {
		log.Fatal(err)
	}

	pipelineStages := setPipelineStages(serviceGenerator, serviceValidator, serviceCalculator)

	return &Factory{
		Config:         config,
		Chan:           done,
		PipelineStages: pipelineStages,
	}, nil
}

func setWriterService(config *cfg.Config) model.Writer {
	return service.NewFileResposeWriter(config)
}

func setPipelineStages(stages ...model.Runner) []model.Runner {
	pipelineStages := append([]model.Runner{}, stages...)
	return pipelineStages
}
