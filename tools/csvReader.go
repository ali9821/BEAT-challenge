package data

import (
	cfg "BEAT/config"
	"BEAT/pkg/model"
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

var Data []model.Travel

func CsvReader(config *cfg.Config) []model.Travel {
	var Travels []model.Travel
	file, err := os.Open(config.PathFile)
	if err != nil {
		log.Fatal(err)
	}
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var travel model.Travel

	for _, t := range data {
		travel.Id = t[0]
		lat, _ := strconv.ParseFloat(t[1], 64)
		travel.Lat = lat
		lng, _ := strconv.ParseFloat(t[2], 64)
		travel.Lng = lng
		ts, _ := strconv.Atoi(t[3])
		travel.TimeStamp = ts
		Travels = append(Travels, travel)
		Data = append(Data, travel)
	}
	return Travels
}
