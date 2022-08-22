package main

import (
	"fmt"
)

func main() {
	fmt.Println("starting...")
	fmt.Println("getting current gps coordinates...")
	//note: the gps position you will see in the code below will
	//not be real as i still need to programm the arduino to get
	//those and get them live into the script
	currentGps_lon, currentGps_lat := 50.45635, 50.45635
	fmt.Print("current longitude: ")
	fmt.Println(currentGps_lon)
	fmt.Print("current latitude:  ")
	fmt.Println(currentGps_lat)
}
