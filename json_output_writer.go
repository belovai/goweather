package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

type JsonOutputWriter struct {
}

func (j *JsonOutputWriter) Render(w *WeatherResponse) {
	tempSign := "°C"
	speedSign := "m/s"
	if *Units == "imperial" {
		tempSign = "°F"
		speedSign = "mph"
	}

	temp := fmt.Sprintf("%.0f%s", w.Main.Temp, tempSign)

	wind := ""
	if w.Wind.Speed > 0 {
		wind = fmt.Sprintf(", %.1f %s (%s)", w.Wind.Speed, speedSign, CalculateDirections(w.Wind.Deg))
	}

	transformer := map[string]string{
		"city":        w.Name,
		"description": w.Weather[0].Description,
		"temp":        temp,
		"wind":        wind,
		"pressure":    fmt.Sprintf("%d hPa", w.Main.Pressure),
		"humidity":    fmt.Sprintf("%d%%", w.Main.Humidity),
		"sunrise":     strconv.Itoa(w.Sys.Sunrise),
		"sunset":      strconv.Itoa(w.Sys.Sunset),
	}

	jsonString, err := json.Marshal(transformer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonString))
}
