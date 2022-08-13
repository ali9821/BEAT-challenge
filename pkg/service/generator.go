package service

import (
	"BEAT/pkg/model"
	data "BEAT/tools"
	"context"
)

type travelService struct {
	travelChannel chan model.Travel
}

func GenerateTravelService(channel chan model.Travel) (model.Runner, error) {
	return &travelService{
		travelChannel: channel,
	}, nil
}

func (ts *travelService) Run(context.Context) error {
	for _, route := range data.Data {
		ts.travelChannel <- route
	}
	close(ts.travelChannel)
	return nil
}
