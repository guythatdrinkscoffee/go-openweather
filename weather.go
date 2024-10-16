package openweather

type WeatherDataSet string

const (
	CurrentWeatherDataSet  WeatherDataSet = "weather"
	ForecastWeatherDataSet WeatherDataSet = "forecast"
)

type WeatherResponse struct {
	CurrentWeather  *CurrentWeatherResponse
	ForecastWeather *ForecastWeatherResponse
}

type WeatherResponseData struct {
	Coord               Coord     `json:"coord"`
	Weather             []Weather `json:"weather"`
	Base                string    `json:"base"`
	Main                Main      `json:"main"`
	Visibility          int       `json:"visibility"`
	PrecipitationChance float64   `json:"pop"`
	Wind                Wind      `json:"wind"`
	Rain                Rain      `json:"rain"`
	Snow                Snow      `json:"snow"`
	Timestamp           int64     `json:"dt"`
	Sys                 Sys       `json:"sys"`
	Timezone            int64     `json:"timezone"`
	ID                  int64     `json:"id"`
	Name                string    `json:"name"`
}

type CurrentWeatherResponse = WeatherResponseData

type ForecastWeatherResponse struct {
	List []WeatherResponseData `json:"list"`
	City City                  `json:"city"`
}
