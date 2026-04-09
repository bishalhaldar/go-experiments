# Go Experiments

Small Go experiments for learing cloud-native development.

---

## Weather Client

A simple CLI program written in Go that calls a weather API and prints the response.

### Features

- Calls REST API using Go `net/http`
- Parses JSON response
- CLI argument for city
- Proper error handling
- Uses only Go standard library

---

### Run the program

First start the weather API container:

- docker run -p 3000:3000 bishalhaldar/weather-api

Then run the Go client:

- go run weather/main.go Delhi

---

### Example Output

City: Delhi Temperature: dummy data
