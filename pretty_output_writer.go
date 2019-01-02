package main

import (
	"fmt"
	"time"
)

type PrettyOutputWriter struct {
}

func (p *PrettyOutputWriter) Render(w *WeatherResponse) {
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
