package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type RequestBody struct {
	Water *int `json:"water"`
	Wind  *int `json:"wind"`
}

func main() {
	http.HandleFunc("/loop", func(w http.ResponseWriter, r *http.Request) {
		var water, wind int

		for {
			water = rand.Intn(100)
			wind = rand.Intn(100)

			var statWater, statWind string

			if water <= 5 {
				statWater = "aman"
			} else if water <= 8 {
				statWater = "siaga"
			} else {
				statWater = "bahaya"
			}

			if wind <= 6 {
				statWind = "aman"
			} else if wind <= 15 {
				statWind = "siaga"
			} else {
				statWind = "bahaya"
			}

			var BodyJson RequestBody
			BodyJson.Water = &water
			BodyJson.Wind = &wind
			BodyBytes, err := json.MarshalIndent(BodyJson, "", "    ")
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			url := "https://jsonplaceholder.typicode.com/posts"

			resp, err := http.Post(url, "application/json", bytes.NewBuffer(BodyBytes))
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			defer resp.Body.Close()

			_, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println(string(BodyBytes))
			fmt.Printf("status_water : %s\n", statWater)
			fmt.Printf("status_wind : %s\n\n", statWind)
			// Loop every 15s
			time.Sleep(15 * time.Second)
		}

	})
	server := new(http.Server)
	server.Addr = ":9000"

	fmt.Println("Server berjalan di localhost:9000")
	server.ListenAndServe()
}
