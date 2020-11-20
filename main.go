package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const APIkey = "09ec83df87c30f753e96bff8681f03af"

func main() {

	var location string
	fmt.Println("Starting application...")
	flag.StringVar(&location, "location", "London", "Flagname should specify a location e.g. London")
	flag.Parse()

	type APIResp struct {
		Weather []struct {
			Description string `json:"description"`
		} `json:"weather"`
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + location + "&appid=" + APIkey)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var WeatherInfo APIResp
	json.Unmarshal(body, &WeatherInfo)

	Desc := WeatherInfo.Weather[0].Description
	Temp := WeatherInfo.Main.Temp

	fmt.Println("The weather in " + location + " is currently " + KtoC(Temp) + " degrees Celsius. With a description of " + Desc)

}

// KtoC -> String
// Convert a given Kelvin metric to Celsius
func KtoC(k float64) string {
	return fmt.Sprintf("%.2f", k-273.15)
}
