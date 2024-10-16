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
		Lat: 33.10,
		Lon: -117.28,
	}

	w, err := c.Weather(context.Background(), r, openweather.ForecastWeatherDataSet)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Current Location: %s, %s", w.ForecastWeather.City.Name, w.ForecastWeather.City.Country)
	for _, weather := range w.ForecastWeather.List {
		log.Println(fmt.Sprintf("Date: %v, Temp: %.2f", time.Unix(weather.Timestamp, 0), weather.Main.Temp))
		log.Println(fmt.Sprintf("Precipitation Chance: %.2f, Rain: %.2f", weather.PrecipitationChance, weather.Rain.ThreeHours))
	}
}
