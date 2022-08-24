//important note: all values below are only placeholders used for testing purposes
// 				  as i didn't yet programm the arduino because im waiting for parts

package main

import (
	"fmt"
	"math"
	//"time"
)

var current_Gps_lat float64 = 50.69512
var current_Gps_lon float64 = 56.85762
var current_mag_heading float64 = 0.0
var current_yaw float64 = 0.0
var current_speed float64 = 0.4 //in m/s

func main() {
	fmt.Println(math.Atan(5.5 / 2.3))
	fmt.Println(conv_rad_deg(math.Atan(5.5 / 2.3)))
	/*fmt.Println("starting...")

		fmt.Println("starting goroutine get_current_gps...")

		go get_current_gps()

		for 1 == 1 {
			fmt.Println("waiting...")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("current_Gps_lat:     ", current_Gps_lat)
			fmt.Println("current_Gps_lon:     ", current_Gps_lon)
			fmt.Println("current_Mag_heading: ", current_mag_heading)
			fmt.Println("current_yaw:         ", current_yaw)
			fmt.Println("current_speed:       ", current_speed)
		}

		fmt.Println("current_Gps_lat:     ", current_Gps_lat)
		fmt.Println("current_Gps_lon:     ", current_Gps_lon)
		fmt.Println("current_Mag_heading: ", current_mag_heading)
		fmt.Println("current_yaw:         ", current_yaw)
		fmt.Println("current_speed:       ", current_speed)

		fmt.Println("done...")
	}

	// testing goroutines
	func get_current_gps() {
		sum := 1
		for i := 1; i < 500; i++ {
			current_Gps_lat = current_Gps_lat + 2.5
			sum = sum + sum
			fmt.Println(sum)
			time.Sleep(200 * time.Millisecond)
		}
	*/
}

func conv_rad_deg(rad float64) float64 {
	deg := rad * (180 / math.Pi)
	return deg
}
