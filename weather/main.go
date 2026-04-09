package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type WeatherResponse struct {
  City string `json: "city"`
  Temp string `json: "temp"`
}

func main(){
  if len(os.Args) < 2 {
	fmt.Println("Usage: go run main.go <city>")
	os.Exit(1)
  }
  
  city := os.Args[1]	//getting the city name

  url := "http://localhost:3000/weather?city=" + city	//calling the API with city as query parameter
  
  resp, err := http.Get(url)	//HTTP request to the API
  if err != nil {
	fmt.Println("Error calling API:", err)
	os.Exit(1)
  }

  defer resp.Body.Close()

  if resp.StatusCode != http.StatusOK {
	fmt.Println("Unexpected status:", resp.Status)
	os.Exit(1)
  }

  body, err := io.ReadAll(resp.Body)	//reading the response body
  if err != nil {
	fmt.Println("Error reading response:", err)
	os.Exit(1)
  }

  var weather WeatherResponse
  
  err =json.Unmarshal(body, &weather)	//parsing the JSON response
  if err != nil {
	fmt.Println("Error parsing JSON:", err)
	os.Exit(1)
  }
  
  fmt.Println("City:", weather.City)
  fmt.Println("Temperature:", weather.Temp)
}
