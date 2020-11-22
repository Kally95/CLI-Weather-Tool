package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const APIurl = "http://api.openweathermap.org/data/2.5/weather?q="

type APIResp struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

var location string
var WeatherInfo APIResp
var dotenv string

// GetAPIkey -> string
// Retrieves the API key stored within the .env file
func GetAPIkey(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}

// KtoC -> String
// Convert a given Kelvin metric to Celsius
func KtoC(k float64) string {
	return fmt.Sprintf("%.2f", k-273.15)
}

// GetAPI -> *http.Response, error
// Uses the HTTP GET method to make a GET request to the API
func GetAPI(loc, key string) (*http.Response, error) {
	resp, err := http.Get(APIurl + location + "&appid=" + dotenv)
	if err != nil {
		log.Fatalln(err)
	}
	return resp, err
}

// ReadAPI -> []byte, error
// Reads the response returned from GetApi's GET method call
func ReadAPI(resp *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
	return body, err
}

func main() {

	dotenv = GetAPIkey("APIKEY")
	flag.StringVar(&location, "location", "London", "Flagname should specify a location e.g. London")
	flag.Parse()

	r, err := GetAPI(location, dotenv)
	if err != nil {
		fmt.Println(err)
	}

	rb, err := ReadAPI(r)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(rb, &WeatherInfo)

	Desc := WeatherInfo.Weather[0].Description
	Temp := WeatherInfo.Main.Temp

	fmt.Println("The weather in " + location + " is currently " + KtoC(Temp) + " degrees Celsius. With a description of " + Desc)

}
