# goweather

It's a simple project for geting current weather from terminal via openweathermap api.

## API key (APPID)

You can generate your own API key on https://openweathermap.org

## Build

```shell
make all
```

## Usage

```shell
./goweather -h
```

### Options

#### -a, --appid=value
Your APPID from https://openweathermap.org. Default value will be your GOWEATHER_APPID environment variable.

#### -c, --city=value
City name and country code separated by comma. Use ISO 3166 country codes. Example: London,gb Default value will be your GOWEATHER_CITY environment varible.

#### -f, --format=value
Output format. Possible values: pretty, json. Default value is pretty

#### -h, --help
Shows the help

#### -u, --units=value
Temperature is available in Fahrenheit and Celsius units. Possible values: imperial, metric. Default value will be metric if your GOWEATHER_UNITS not set.

### Example

```shell
./goweather -a YOUR_APP_ID -c London,gb
```