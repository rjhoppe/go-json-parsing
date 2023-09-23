package main

import (
	"encoding/json"
	"fmt"
)

type SensorReading struct {
	Name        string `json:"name"`
	Capacity    string `json:"name"`
	Time        string `json:"time"`
	Information `json:"info"`
}

type Information struct {
	Description string `json:"desc"`
}

func main() {

	jsonString := `{"name": "battery sensor", "capacity": 40, "time":
	"2019-01-21T19:07:282", "info": {
		"desc": "a sensor reading"
	}}`

	var reading SensorReading
	// Another way you could do this
	// var reading map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &reading)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", reading)

}
