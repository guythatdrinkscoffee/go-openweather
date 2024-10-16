package openweather

type WeatherResponseData struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility int       `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Timestamp  int64     `json:"dt"`
	Sys        Sys       `json:"sys"`
	Timezone   int64     `json:"timezone"`
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
}

type CurrentWeatherResponse = WeatherResponseData

type ForecastWeatherResponse struct {
	List []WeatherResponseData `json:"list"`
	City City                  `json:"city"`
}
