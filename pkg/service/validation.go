package service

import (
	"BEAT/pkg/model"
	"context"
	"github.com/umahmood/haversine"
	"sync"
)

type calcPrice struct {
	channel     chan model.Travel
	validRoutes chan model.RouteInfo
}

func GenerateValidationService(channel chan model.Travel, validRoutes chan model.RouteInfo) (model.Runner, error) {
	return &calcPrice{
		channel:     channel,
		validRoutes: validRoutes,
	}, nil
}

func (cp *calcPrice) Run(context.Context) error {
	var wg sync.WaitGroup
	wg.Add(1)
	go cp.worker(&wg)
	wg.Wait()
	return nil
}

func (cp *calcPrice) worker(wg *sync.WaitGroup) {
	defer wg.Done()
	routes := cp.channel
	var coordA haversine.Coord
	var timeStampA int
	var id string
	var validRoute model.RouteInfo
	for route := range routes {
		if id == "" {
			id = route.Id
			coordA.Lat = route.Lat
			coordA.Lon = route.Lng
			timeStampA = route.TimeStamp
		} else {
			if route.Id == id {
				coordB := haversine.Coord{Lat: route.Lat, Lon: route.Lng}
				_, km := haversine.Distance(coordA, coordB)
				time := route.TimeStamp - timeStampA
				speed := (km * 1000) / float64(time) * 3.6
				if speed > 100 {

				} else {
					coordA = coordB
					timeStampA = route.TimeStamp
					validRoute.ID = route.Id
					validRoute.SPEED = speed
					validRoute.KM = km
					validRoute.Time = time
					cp.validRoutes <- validRoute
				}

			} else {
				id = route.Id
				coordA.Lat = route.Lat
				coordA.Lon = route.Lng
				timeStampA = route.TimeStamp
			}
		}
	}
}
