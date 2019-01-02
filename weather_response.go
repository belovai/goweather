package main

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

func (w *WeatherResponse) Render(outputWriter OutputWriterInterface) {
	outputWriter.Render(w)
}
