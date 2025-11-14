// This file models the CurrentWeatherResponse found at https://docs.apilayer.com/weatherstack/docs/weatherstack-api-v-1-0-0#/CurrentWeatherResponse
package api

type Current struct {
	ObservationTime     string     `json:"observation_time"`
	Temperature         float64    `json:"temperature"`
	WeatherCode         int        `json:"weather_code"`
	WeatherIcons        []string   `json:"weather_icons"`
	WeatherDescriptions []string   `json:"weather_descriptions"`
	Astro               Astro      `json:"astro"`
	AirQuality          AirQuality `json:"air_quality"`
	WindSpeed           float64    `json:"wind_speed"`
	WindDegree          float64    `json:"wind_degree"`
	WindDirection       string     `json:"wind_dir"`
	Pressure            float64    `json:"pressure"`
	Precip              float64    `json:"precip"`
	Humidity            int        `json:"humidity"`
	CloudCover          int        `json:"cloudcover"`
	FeelsLike           float64    `json:"feelslike"`
	UVIndex             float64    `json:"uv_index"`
	Visibility          float64    `json:"visibility"`
	IsDay               string     `json:"is_day"`
}

type Astro struct {
	Sunrise          string `json:"sunrise"`
	Sunset           string `json:"sunset"`
	Moonrise         string `json:"moonrise"`
	Moonset          string `json:"moonset"`
	MoonPhase        string `json:"moon_phase"`
	MoonIllumination int    `json:"moon_illumination"`
}

type AirQuality struct {
	CO           string `json:"co"`
	NO2          string `json:"no2"`
	O3           string `json:"o3"`
	SO2          string `json:"so2"`
	PM2_5        string `json:"pm2_5"`
	PM10         string `json:"pm10"`
	UsEpaIndex   int    `json:"us_epa_index"`
	GBDefraIndex int    `json:"gb_defra_index"`
}
