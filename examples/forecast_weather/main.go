package main

import (
	"context"
	"fmt"
	"github.com/guythatdrinkscoffee/go-openweather"
	"log"
	"os"
	"time"
)

func main() {
	t := os.Getenv("OPENWEATHER_API_KEY")

	c := openweather.NewClient(t, openweather.WithUserAgent("myweatherapp"))

	r := openweather.Request{
		Lat: 17.07,
		Lon: -96.72,
	}

	w, err := c.Weather(context.Background(), r, openweather.ForecastWeatherDataSet)
	if err != nil {
		log.Fatalln(err)
	}

	for _, weather := range w.ForecastWeather.List {
		log.Println(fmt.Sprintf("Date: %v, Temp: %.2f", time.Unix(weather.Timestamp, 0), weather.Main.Temp))
	}
}
