package main;

import (
	"encoding/json"
	"fmt"
)

type Player struct {
	Name string `json:"playerName"`
	Password string `json:"-"`
	Roles []string `json:"tags,omitempty"`
}

func main() {	
	fmt.Println("Conversion of structs to JSON in Go");

	players := []Player{
		{"Faith", "pass123", []string{"defend", "attach"}},
		{"Precious", "pass4321", []string{"referee", "keeper"}},
		{"Rita", "test12", nil},
	}

	jsonData, _ := json.MarshalIndent(players, "", "\t");

	fmt.Println(string(jsonData));
	// fmt.Printf("%s\n", jsonData);
}