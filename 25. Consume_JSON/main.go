package main;

import (
	"encoding/json"
	"fmt"
)


type Player struct {
	Name string `json:"playerName"`
	Password string `json:"-"`
	Roles []string `json:"tags,omitempty"`
	Address struct {
		Lng float64 
		Lat float64
	} `json:"address"`
}

func main() {
	fmt.Println("Comsuming JSON in Go Langauge")

	var jsonData = []byte(`{
		"playerName": "Faith",
		"tags": ["defend","attach"],
		"address": {
			"lng": 25223.2535,
			"lat": 24235.23423
		}
	}`)

	var player Player;
	var isJSONValid = json.Valid(jsonData)

	if isJSONValid {
		json.Unmarshal(jsonData, &player);
		fmt.Printf("%+v\n", player);
	} else {
		fmt.Println("JSON Provided is invalid")
	}

	/* Consuming JSON data as a Map */
	var playerMap map[string]interface{};
	json.Unmarshal(jsonData, &playerMap);
	fmt.Printf("%+v\n\n", playerMap);

	var data map[string]interface{} = map[string]interface{}{
		"name": "Hycient",
		"age": 15,
		"isActive": false,
		"address": map[string]float64{ // Map
			"lng": 2345.2345,
			"lat": 2143.2344,
		},
		"addressly": struct{ // Object
			Lng string
			Lat string
		}{
			Lng: "long",
			Lat: "lati",
		},
		"hobbies": []string{"Jumping", "Flying", "Crawling", "Travelling"}, // Slice
	}
	fmt.Printf("%+v\n",data)
}


/*
	1. Define the struct pattern the converterd JSON would conform to;

	2. Create a variable instance of the struct

	2. Check if the JSON is valid by using "json.Valid";

	3. If the data to be converted is valid JSON, "Unmarahal" it, passing
			the data to be converd and the reference of an instance of the 
			structure the JSON is to be converted into to "json.Unmarshal"

	4. The unmarshalled value can now be obtained from the pointer reference
			variable passed to the JSON unmarshall function, as it stores the
			result in the struct instance variable passed to it.
			json.Unmarshal returns an error, which is populated, if an error is
			encountered in parsing the JSON.
*/