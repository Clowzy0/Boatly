//important note: all values below are only placeholders used for testing purposes
// 				  as i didn't yet programm the arduino because im waiting for parts
//
// 				  as long as this note will be here all of the code below won't be
//                even remotly functional

package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/fogleman/gg"
)

var pi float64 = 3.141592653589793238462643383279502884197169399375105820974944592307816406286208998628034825342117067

var goal_Gps_lat float64 = 0.0
var goal_Gps_lon float64 = 0.0
var goal_yaw float64 = 0.0
var goal_mag_heading float64 = 0.0

var current_Gps_lat float64 = 0.0
var current_Gps_lon float64 = 0.0
var current_mag_heading float64 = 0.0
var current_yaw float64 = 0.0
var current_speed float64 = 0.4 //in m/s

var angle_to_rotate float64 = 0.0

func main() {
	i := 0
	for i < 1 {
		angle_to_rotate = 0.0
		i = i + 1
		fmt.Println("starting...")

		im, err := gg.LoadPNG("boat.png")
		if err != nil {
			panic(err)
		}

		rand.Seed(time.Now().Local().UnixNano())
		min := 20
		max := 100
		col := rand.Intn(max-min+1) + min
		col1 := float64(col) / 10.0
		col2 := float64(col) / 19.0
		col3 := float64(col) / 12.0
		goal_Gps_lat, goal_Gps_lon = float64(rand.Intn(1999-1+1)+1), float64(rand.Intn(1999-1+1)+1)
		current_Gps_lat, current_Gps_lon = float64(rand.Intn(1999-1+1)+1), float64(rand.Intn(1999-1+1)+1)

		const S = 2000
		dc := gg.NewContext(S, S)
		dc.SetRGB(col1, col2, col3)
		dc.Clear()

		dc.SetRGB(0, 0, 0)
		dc.DrawCircle((2000 - float64(goal_Gps_lat)), float64(goal_Gps_lon), 15.0)
		dc.Fill()
		dc.Stroke()

		if (goal_Gps_lat < current_Gps_lat) && (goal_Gps_lon > current_Gps_lon) { //tar1
			offset := float64(270)
			alpha := math.Atan((float64(current_Gps_lat) - float64(goal_Gps_lat)) / (float64(goal_Gps_lon) - float64(current_Gps_lon)))
			beta := 180 - (90 + conv_rad_deg(alpha))
			gamma := conv_deg_rad(beta) + conv_deg_rad(offset)
			angle_to_rotate = gamma

			fmt.Println("gamma:        ", gamma)
			fmt.Println("gamma in deg: ", conv_rad_deg(gamma))
		}

		if (goal_Gps_lat > current_Gps_lat) && (goal_Gps_lon > current_Gps_lon) { //tar2
			offset := float64(0)
			alpha := math.Atan((float64(goal_Gps_lat) - float64(current_Gps_lat)) / (float64(goal_Gps_lon) - float64(current_Gps_lon)))
			beta := 180 - (90 + conv_rad_deg(alpha))
			gamma := conv_deg_rad(beta) + conv_deg_rad(offset)
			angle_to_rotate = gamma

			fmt.Println("gamma:        ", gamma)
			fmt.Println("gamma in deg: ", conv_rad_deg(gamma))
		}

		if (goal_Gps_lat > current_Gps_lat) && (goal_Gps_lon < current_Gps_lon) { //tar3
			offset := float64(90)
			alpha := math.Atan((float64(goal_Gps_lat) - float64(current_Gps_lat)) / (float64(current_Gps_lon) - float64(goal_Gps_lon)))
			beta := 180 - (90 + conv_rad_deg(alpha))
			gamma := conv_deg_rad(beta) + conv_deg_rad(offset)
			angle_to_rotate = gamma

			fmt.Println("gamma:        ", gamma)
			fmt.Println("gamma in deg: ", conv_rad_deg(gamma))
		}

		if (goal_Gps_lat < current_Gps_lat) && (goal_Gps_lon < current_Gps_lon) { //tar4
			offset := float64(0)
			alpha := math.Atan((float64(current_Gps_lat) - float64(goal_Gps_lat)) / (float64(current_Gps_lon) - float64(goal_Gps_lon)))
			beta := 180 - (90 + conv_rad_deg(alpha))
			gamma := conv_deg_rad(beta) + conv_deg_rad(offset)
			angle_to_rotate = gamma

			fmt.Println("gamma:        ", gamma)
			fmt.Println("gamma in deg: ", conv_rad_deg(gamma))
		}

		//dc.RotateAbout(rot_atan(goal_Gps_lat, goal_Gps_lon, current_Gps_lat, current_Gps_lon), float64(current_Gps_lat), float64(current_Gps_lon))
		dc.RotateAbout((angle_to_rotate + conv_deg_rad(180)), float64((2000 - current_Gps_lat)), float64(current_Gps_lon))
		dc.DrawImageAnchored(im, int((2000 - current_Gps_lat)), int(current_Gps_lon), 0.5, 0.5)
		//name := fmt.Sprint(time.Now().Local().UnixMilli())
		name := "out"
		suffix := ".png"
		dc.SavePNG(name + suffix)
		fmt.Println("done...")
	}
}

/*






 */

func conv_rad_deg(radians float64) float64 {
	degress := radians * (180 / pi)
	return degress
}

func conv_deg_rad(degress float64) float64 {
	radians := degress * (pi / 180)
	return radians
}
