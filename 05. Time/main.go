package main;

import (
	"fmt"
	"time"
)

func main() { // This gives the current time in the Go Langauge standard time format
	currentTime := time.Now();
	fmt.Println(currentTime);

	formattedCurrentTime := currentTime.Format("01/02/2006, 15:04:05 Monday"); // All time formatted in Go must make use of the exact digits specified here of which each has a unit value. How they are to be presented however, can be tweaked.
	fmt.Println(formattedCurrentTime);

	createdTime := time.Date(2020, time.January, 12, 20, 45, 20, 0, time.UTC);
	fmt.Println(createdTime);
}