package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// os.Args provides access to raw command-line arguments.
// Note that the first value in this slice is the path to
// the program, and os.Args[1:] holds the arguments to the program.

const APIkey = "09ec83df87c30f753e96bff8681f03af"

func main() {
	input := os.Args[1]

	fmt.Println("City name:")
	fmt.Println(input)

	type Weather struct {
		Description string `json:"description"`
		Temp        int    `json:"temp"`
	}

	type Wrapperv2 struct {
		Temp string `json:"temp"`
	}

	type Wrapper struct {
		Weather []Weather                 `json:"weather"`
		Main    map[string]map[string]int `json:"main"`
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

	// Create an instance of type Weather (Struct)
	var asd Wrapper
	json.Unmarshal(body, &asd)
	fmt.Print(asd)

	// Currently body returns a slice of BYTES

}

// http.Handle("/foo", fooHandler)

// http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("api.openweathermap.org/data/2.5/weather?q=input&appid=APIkey")
// })
