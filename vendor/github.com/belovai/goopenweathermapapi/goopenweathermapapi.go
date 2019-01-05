package goopenweathermapapi

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

//Client api client
type Client struct {
	APPID string
}

const apiURL = "https://api.openweathermap.org/data/2.5/"

//NewClient appid should be the openweathermap APPID
func NewClient(appid string) *Client {
	return &Client{APPID: appid}
}

//GetWeatherByCityName You can call by city name or city name and country code
//separated by comma. Use ISO 3166 country codes.
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

	url := fmt.Sprintf("%sweather?%s", apiURL, params.Encode())

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

//GetWeatherByCityID You can call by city id.
//List of city ID city.list.json.gz can be downloaded here http://bulk.openweathermap.org/sample/
//Units possible values are: metric, imperial or empty string.
//Lang possible values are: ar, bg, ca, cz, de, el, en, fa, fi, fr, gl, hr, hu, it,
//ja, kr, la, lt, mk, nl, pl, pt, ro, ru, se, sk, sl, es, tr, ua, vi, zh_cn, zh_tw
func (c *Client) GetWeatherByCityID(cityID int, units, lang string) (jsonString string, err error) {
	params := url.Values{}

	params.Add("APPID", c.APPID)
	params.Add("id", strconv.Itoa(cityID))
	params.Add("lang", lang)
	params.Add("units", units)

	url := fmt.Sprintf("%sweather?%s", apiURL, params.Encode())

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

//GetWeatherByCoordinates You can call By geographic coordinates.
//lat, lon coordinates of the location of your interest.
//Units possible values are: metric, imperial or empty string.
//Lang possible values are: ar, bg, ca, cz, de, el, en, fa, fi, fr, gl, hr, hu, it,
//ja, kr, la, lt, mk, nl, pl, pt, ro, ru, se, sk, sl, es, tr, ua, vi, zh_cn, zh_tw
func (c *Client) GetWeatherByCoordinates(lat, lon float64, units, lang string) (jsonString string, err error) {
	params := url.Values{}

	params.Add("APPID", c.APPID)
	params.Add("lat", strconv.FormatFloat(lat, 'f', 2, 64))
	params.Add("lon", strconv.FormatFloat(lon, 'f', 2, 64))
	params.Add("lang", lang)
	params.Add("units", units)

	url := fmt.Sprintf("%sweather?%s", apiURL, params.Encode())

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

//GetWeatherByZipCode You can call by zip code or zip code and country code seprated
//by comma. Use ISO 3166 country codes.
//Please note if country is not specified then the search works for USA as a default.
//Units possible values are: metric, imperial or empty string.
//Lang possible values are: ar, bg, ca, cz, de, el, en, fa, fi, fr, gl, hr, hu, it,
//ja, kr, la, lt, mk, nl, pl, pt, ro, ru, se, sk, sl, es, tr, ua, vi, zh_cn, zh_tw
func (c *Client) GetWeatherByZipCode(zip, units, lang string) (jsonString string, err error) {
	params := url.Values{}

	params.Add("APPID", c.APPID)
	params.Add("zip", zip)
	params.Add("lang", lang)
	params.Add("units", units)

	url := fmt.Sprintf("%sweather?%s", apiURL, params.Encode())

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

//GetForecastByCityName You can search weather forecast for 5 days with data every 3 hours by city name.
//Units possible values are: metric, imperial or empty string.
//Lang possible values are: ar, bg, ca, cz, de, el, en, fa, fi, fr, gl, hr, hu, it,
//ja, kr, la, lt, mk, nl, pl, pt, ro, ru, se, sk, sl, es, tr, ua, vi, zh_cn, zh_tw
func (c *Client) GetForecastByCityName(city, units, lang string) (jsonString string, err error) {
	params := url.Values{}

	params.Add("APPID", c.APPID)
	params.Add("q", city)
	params.Add("lang", lang)
	params.Add("units", units)

	url := fmt.Sprintf("%sforecast?%s", apiURL, params.Encode())

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
