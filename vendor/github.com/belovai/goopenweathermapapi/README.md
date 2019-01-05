# goopenweathermapapi
Simple API client for openweathermap. 

## API key (APPID)
You can generate your own API key on https://openweathermap.org

## Usage
```go
package main

import (
	"fmt"
	"log"

	gowma "github.com/belovai/goopenweathermapapi"
)

func main() {
	client := gowma.NewClient("YOUR_APPID")

	jsonString, err := client.GetWeatherByCityName("London,gb", "metric", "en")
	if err != nil {
		log.Fatal(err, jsonString)
	}

	fmt.Println(jsonString)
}
```