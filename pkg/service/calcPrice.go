package service

import (
	"BEAT/pkg/model"
	"context"
	"fmt"
	"io"
	"sync"
)

type calc struct {
	validRoutes chan model.RouteInfo
	routePrice  chan model.TravelPrice
	writer      io.Writer
	done        chan bool
}

func GeneratePriceCalcService(validRoutes chan model.RouteInfo, writer io.Writer, done chan bool) (model.Runner, error) {
	channelPrice := make(chan model.TravelPrice, 1000)
	return &calc{
		validRoutes: validRoutes,
		routePrice:  channelPrice,
		writer:      writer,
		done:        done,
	}, nil
}

func (c *calc) Run(ctx context.Context) error {
	ctx, done := context.WithCancel(ctx)
	data := make(map[string]model.TravelPrice)
	var wg sync.WaitGroup
	wg.Add(2)
	go c.adder(&wg, c.routePrice)
	go c.worker(&wg, c.routePrice, data, done)
	wg.Wait()
	return nil
}

func (c *calc) adder(wg *sync.WaitGroup, channel chan<- model.TravelPrice) {
	defer wg.Done()
	for i := range c.validRoutes {
		if i.SPEED > 10 {
			var route model.TravelPrice
			price := i.KM * 1.5
			route.Id = i.ID
			route.Km = i.KM
			route.Price = price
			channel <- route
		} else {
			var route model.TravelPrice
			price := float64(i.Time) * 11.90 / 3600
			route.Id = i.ID
			route.Km = i.KM
			route.Price = price
			channel <- route
		}
	}
	close(channel)
}
func (c *calc) worker(wg *sync.WaitGroup, channel <-chan model.TravelPrice, data map[string]model.TravelPrice, done context.CancelFunc) {
	defer wg.Done()
	var a model.TravelPrice
	for i := range channel {
		if a.Id == "" {
			a.Id = i.Id
			a.Km = i.Km
			a.Price = i.Price
			data[a.Id] = a
		} else {
			if a.Id == i.Id {
				a.Price += i.Price
				a.Km += i.Km
				data[a.Id] = a
				fmt.Println("")
				if len(channel) == 0 {
					fmt.Println(data)
					for _, item := range data {
						if item.Price/item.Km < 3.47 {
							item.Price = item.Km * 3.47
							c.writer.Write([]byte(fmt.Sprintf("%s; %v\n", item.Id, item.Price)))
						} else {
							c.writer.Write([]byte(fmt.Sprintf("%s; %v\n", item.Id, item.Price)))
						}
					}
					done()
					c.done <- true
				}

			} else {
				a.Id = i.Id
				a.Km = i.Km
				a.Price = i.Price
				data[a.Id] = a
			}
		}
	}
}
