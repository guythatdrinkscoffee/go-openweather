package main

import (
	"context"
	"github.com/guythatdrinkscoffee/go-openweather"
	"log"
	"os"
)

func main() {
	t := os.Getenv("OPENWEATHER_API_KEY")

	c := openweather.NewClient(t)

	r := openweather.Request{
		Lat: 17.07,
		Lon: -96.72,
	}

	w, err := c.Weather(context.Background(), r, openweather.CurrentWeatherDataSet)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("\n City: %s \n Temp: %.2f \n Feels Like: %.2f", w.CurrentWeather.Name, w.CurrentWeather.Main.Temp, w.CurrentWeather.Main.FeelsLike)
}
