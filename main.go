package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const APIkey = "09ec83df87c30f753e96bff8681f03af"

func main() {
	input := os.Args[1]

	fmt.Println("City name:")
	fmt.Println(input)

	type APIResp struct {
		Weather []struct {
			Description string `json:"description"`
		} `json:"weather"`
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + string(input) + "&appid=" + APIkey)
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

	log.Println(": The weather in " + string(input) + " is currently " + KtoC(Temp) + " degrees Celsius. With a description of " + Desc)

}

// Converts Kelvin to Celsius
func KtoC(k float64) string {
	return fmt.Sprintf("%.2f", k-273.15)
}
