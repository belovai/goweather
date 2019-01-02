package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/belovai/goopenweathermapapi"
	"github.com/pborman/getopt/v2"
)

const apiURL = "https://api.openweathermap.org/data/2.5/weather"

type ErrorResponse struct {
	Cod     int    `json:"cod"`
	Message string `json:"message"`
}

type OutputWriterInterface interface {
	Render(w *WeatherResponse)
}

var Help *bool
var City *string
var Units *string
var AppID *string
var Format *string
var Lang *string

func main() {
	SetOptions()

	if *Help {
		ShowHelp("")
	}

	if *City == "" {
		ShowHelp("You must set the city")
	}

	var params map[string]string
	params = map[string]string{
		"q":     *City,
		"units": *Units,
		"APPID": *AppID,
	}

	GetCurrentWerather(params)

}

func SetOptions() {
	defaultUnits := os.Getenv("GOWEATHER_UNITS")
	if defaultUnits == "" {
		defaultUnits = "metric"
	}

	Help = getopt.BoolLong("help", 'h', "Shows this help")
	City = getopt.StringLong("city", 'c', os.Getenv("GOWEATHER_CITY"), "City name and country code separated by comma. Use ISO 3166 country codes. Example: London,gb Default value will be your GOWEATHER_CITY environment varible.")
	Units = getopt.EnumLong("units", 'u', []string{"imperial", "metric"}, defaultUnits, "Temperature is available in Fahrenheit and Celsius units. Possible values: imperial, metric. Default value will be metric if your GOWEATHER_UNITS not set.")
	AppID = getopt.StringLong("appid", 'a', os.Getenv("GOWEATHER_APPID"), "Your APPID from https://openweathermap.org. Default value will be your GOWEATHER_APPID environment variable.")
	Format = getopt.EnumLong("format", 'f', []string{"pretty", "json"}, "pretty", "Output format. Possible values: pretty, json. Default value is pretty")
	Lang = getopt.StringLong("lang", 'l', os.Getenv("GOWEATHER_LANG"), "API language")
	getopt.Parse()
}

func ShowHelp(message string) {
	if message != "" {
		fmt.Println(message)
	}
	getopt.Usage()
	os.Exit(0)
}

func GetCurrentWerather(params map[string]string) {

	client := goopenweathermapapi.NewClient(*AppID)

	weatherJson, err := client.GetWeatherByCityName(*City, *Units, *Lang)

	if err != nil {
		log.Println("Error on request: ", err)
		var errorResponse ErrorResponse
		if err := json.NewDecoder(strings.NewReader(weatherJson)).Decode(&errorResponse); err != nil {
			log.Fatal("Decode:", err)
		}
		log.Println("Code:", errorResponse.Cod)
		log.Fatal("Message: ", errorResponse.Message)
		return
	}

	var currentWeather WeatherResponse

	if err := json.NewDecoder(strings.NewReader(weatherJson)).Decode(&currentWeather); err != nil {
		log.Println(err)
	}

	switch *Format {
	case "pretty":
		currentWeather.Render(&PrettyOutputWriter{})
	case "json":
		currentWeather.Render(&JsonOutputWriter{})
	}

}

func CalculateDirections(deg int) string {
	directions := []string{"N", "NE", "NE", "E", "E", "SE", "SE", "S", "S", "SW", "SW", "W", "W", "NW", "NW", "N"}

	return directions[int(float64(deg)/22.5)]
}
