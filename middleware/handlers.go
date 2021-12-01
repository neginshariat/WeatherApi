package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func WeatherUrl(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	city := r.URL.Query()["q"]
	fmt.Println(city)
	if len(city) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - no city was found!"))
		return
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=dbc56333b42c09c5f16dba4cc8462396", city[0])

	res, err := http.Get(url)
	if err != nil {
		log.Fatal("Unable to read url", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Unable to read body", err)
	}

	var data map[string]interface{}
	var temprature interface{}

	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		panic(err)
	}

	messages := data["main"].(map[string]interface{})
	for i, message := range messages {
		if i == "temp" {
			temprature = message
		}
		fmt.Println("range over Main", i, message)
	}

	fmt.Println(data["main"])
	json.NewEncoder(w).Encode(temprature)

}
