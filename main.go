package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/pborman/getopt/v2"
)

const apiURL = "https://api.openweathermap.org/data/2.5/weather"

type Coord struct {
	Lon float64
	Lat float64
}

type Weather struct {
	Id          int
	Main        string
	Description string
	Icon        string
}

type Main struct {
	Temp     float64
	Pressure int
	Humidity int
	Temp_min float64
	Temp_max float64
}

type Wind struct {
	Speed float64
	Deg   int
}

type Clouds struct {
	All int
}

type Sys struct {
	Type    int
	Id      int
	Message float32
	Country string
	Sunrise int
	Sunset  int
}

type WeatherResponse struct {
	Coord      Coord          `json:"coord"`
	Weather    []Weather      `json:"weather"`
	Base       string         `json:"base"`
	Main       Main           `json:"main"`
	Visibility int            `json:"visibility"`
	Wind       Wind           `json:"wind"`
	Clouds     Clouds         `json:"clouds"`
	Rain       map[string]int `json:"rain"`
	Snow       map[string]int `json:"snow"`
	Dt         int            `json:"dt"`
	Sys        Sys            `json:"sys"`
	Id         int            `json:"id"`
	Name       string         `json:"name"`
	Cod        int            `json:"cod"`
}

type ErrorResponse struct {
	Cod     int    `json:"cod"`
	Message string `json:"message"`
}

var Help *bool
var City *string
var Units *string
var AppID *string

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

func (w *WeatherResponse) Print() {
	tempSign := "°C"
	speedSign := "m/s"
	if *Units == "imperial" {
		tempSign = "°F"
		speedSign = "mph"
	}

	temp := fmt.Sprintf("%.0f%s", w.Main.Temp, tempSign)

	wind := ""
	if w.Wind.Speed > 0 {
		wind = fmt.Sprintf(", %.1f %s (%s) wind", w.Wind.Speed, speedSign, CalculateDirections(w.Wind.Deg))
	}

	sunset := time.Unix(int64(w.Sys.Sunset), 0)
	sunrise := time.Unix(int64(w.Sys.Sunrise), 0)
	fmt.Printf("Current weather in %s:\n", w.Name)
	fmt.Printf("%s, %s%s\n", w.Weather[0].Description, temp, wind)
	fmt.Printf("Pressure: %d hPa\n", w.Main.Pressure)
	fmt.Printf("Humidity: %d%%\n", w.Main.Humidity)
	fmt.Printf("Sunset: %02d:%02d\n", sunset.Hour(), sunset.Minute())
	fmt.Printf("Sunrise: %02d:%02d\n", sunrise.Hour(), sunrise.Minute())
}

func SetOptions() {
	defaultUnits := os.Getenv("GOWEATHER_UNITS")
	if defaultUnits == "" {
		defaultUnits = "metric"
	}
	Help = getopt.BoolLong("help", 'h', "Shows this help")
	City = getopt.StringLong("city", 'c', os.Getenv("GOWEATHER_CITY"), "City name and country code separated by comma. Use ISO 3166 country codes. Example: London,gb Default value will be your GOWEATHER_CITY environment varible.")
	Units = getopt.EnumLong("units", 'u', []string{"imperial", "metric"}, defaultUnits, "Temperature is available in Fahrenheit and Celsius units. Values could be imperial or metric. Default value will be metric if your GOWEATHER_UNITS not set.")
	AppID = getopt.StringLong("appid", 'a', os.Getenv("GOWEATHER_APPID"), "Your APPID from https://openweathermap.org. Default value will be your GOWEATHER_APPID environment variable.")
	getopt.Parse()
}

func ShowHelp(message string) {
	if message != "" {
		fmt.Println(message)
	}
	getopt.Usage()
	os.Exit(0)
}

func BuildParams(params map[string]string) string {
	p := url.Values{}
	for key, value := range params {
		p.Add(key, value)
	}

	return p.Encode()
}

func GetCurrentWerather(params map[string]string) {
	url := fmt.Sprintf("%s?%s", apiURL, BuildParams(params))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		log.Println("Error on request: ", resp.Status)
		var errorResponse ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			log.Fatal("Decode:", err)
		}
		log.Println("Code:", errorResponse.Cod)
		log.Fatal("Message: ", errorResponse.Message)
		return
	}

	var currentWeather WeatherResponse

	if err := json.NewDecoder(resp.Body).Decode(&currentWeather); err != nil {
		log.Println(err)
	}

	currentWeather.Print()
}

func CalculateDirections(deg int) string {
	directions := []string{"N", "NE", "NE", "E", "E", "SE", "SE", "S", "S", "SW", "SW", "W", "W", "NW", "NW", "N"}

	return directions[int(float64(deg)/22.5)]
}
