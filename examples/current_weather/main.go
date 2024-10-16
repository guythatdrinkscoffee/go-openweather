package main

import (
	"context"
	"github.com/guythatdrinkscoffee/go-openweather"
	"log"
	"os"
	"time"
)

func main() {
	t := os.Getenv("OPENWEATHER_API_KEY")

	c := openweather.NewClient(t)

	r := openweather.Request{
		Lat: 33.10,
		Lon: -117.28,
	}

	w, err := c.Weather(context.Background(), r, openweather.CurrentWeatherDataSet)
	if err != nil {
		log.Fatalln(err)
	}

	curr := w.CurrentWeather

	log.Printf("Current Location: %s, Temp: %.2fC, Sunrise: %v, Sunset: %v", curr.Name, curr.Main.Temp, time.Unix(curr.Sys.Sunrise, 0), time.Unix(curr.Sys.Sunset, 0))
}
