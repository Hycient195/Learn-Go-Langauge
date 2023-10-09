package main;

import (
	"fmt"
)

func main() {
	var days map[string]string = make(map[string]string);

	days["mon"] = "Monday";
	days["tue"] = "Tuesday";
	days["wed"] = "Wednesday";
	days["thu"] = "Thursday";

	fmt.Println(days);
	fmt.Println(days["tue"]);

	delete(days, "wed");
	fmt.Println(days);

	/* Initializing a map using map Literal */
	var age map[string]int = map[string]int{
		"john": 29,
		"jane": 34,
		"alice": 15,
	}
	fmt.Println(age);
	fmt.Println(len(age));

	// A demonstration to show that maps are displayed in ascending order (alphabetically or numerically) of their keys
	score := map[int]string{
		0: "Z",
		1: "Y",
		3: "X",
	}
	fmt.Println(score);

	scoreCopy := score;
	scoreCopy[1] = "D";
	fmt.Println(score);
}