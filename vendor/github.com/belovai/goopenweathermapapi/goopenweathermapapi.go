package goopenweathermapapi

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

//Client api client
type Client struct {
	APPID string
}

const apiURL = "https://api.openweathermap.org/data/2.5/weather"

//NewClient appid should be the openweathermap APPID
func NewClient(appid string) *Client {
	return &Client{APPID: appid}
}

//GetWeatherByCityName You can call by city name or city name and country code.
//API responds with a list of results that match a searching word.
//Units possible values are: metric, imperial or empty string.
//Lang possible values are: ar, bg, ca, cz, de, el, en, fa, fi, fr, gl, hr, hu, it,
//ja, kr, la, lt, mk, nl, pl, pt, ro, ru, se, sk, sl, es, tr, ua, vi, zh_cn, zh_tw
func (c *Client) GetWeatherByCityName(city, units, lang string) (jsonString string, err error) {
	params := url.Values{}

	params.Add("APPID", c.APPID)
	params.Add("q", city)
	params.Add("lang", lang)
	params.Add("units", units)

	url := fmt.Sprintf("%s?%s", apiURL, params.Encode())

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	buff := new(bytes.Buffer)
	buff.ReadFrom(resp.Body)
	jsonString = buff.String()

	if resp.StatusCode >= 300 {
		err = fmt.Errorf("API returned with: %s", resp.Status)
	}

	return
}
