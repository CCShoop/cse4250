/*
{-
 - Author: Cael Shoop, cshoop2018@my.fit.edu
 - Course: CSE 4250, Fall 2020
 - Project: Proj1, Projection Please
 - Language implementation: go version go1.15 linux/amd64
 -}
*/

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"strconv"
	//for string to float conversion
)

func main() {
	var mollweide = 1 //variable to check whether image should be mollweide or cylindrical
	/*
	 * Commandline input code adapted from
	 * https://golangdocs.com/command-line-arguments-in-golang
	 */

	if len(os.Args[1:]) < 2 {
		fmt.Println("Not enough commandline arguments, exiting.")
		os.Exit(1)
	}
	img, err := os.Open(os.Args[1])
	/*
	 * Studied the GoLang png library at
	 * https://golang.org/pkg/image/png/
	 * Image importing and creation tips were found on
	 * https://www.devdungeon.com/content/working-images-go
	 */
	if err != nil {
		fmt.Println("Error opening image.")
	}

	projection, err := png.Decode(img) //"projection" is the decoded image
	if err != nil {
		fmt.Println("Error decoding png.")
	}
	defer img.Close()

	outputIMG := image.NewRGBA(image.Rect(0, 0, 0, 0))
	res := projection.Bounds()

	if len(os.Args[1:]) > 2 {
		if os.Args[3] == "Cylindrical" {
			mollweide = 0 //if "Cylindrical" is in the commandline input, mollweide is false
		}
	}
	if mollweide == 1 {
		outputIMG = image.NewRGBA(image.Rect(0, 0, res.Dx(), res.Dy())) //creates new image
		for i := 0; i <= res.Dx(); i++ {
			for j := 0; j <= res.Dy(); j++ {
				outputIMG.Set(i, j, color.RGBA{255, 255, 255, 255}) //sets every pixel to be white
			}
		}
		radius := math.Sqrt(float64(res.Dx()) * float64(res.Dy()) / 16.0) //sets radius
		for i := 0; i <= res.Dy(); i++ {
			coords1 := float64(res.Dy())/2 - float64(i)         //sets coords
			x := math.Asin(coords1 / (radius * math.Sqrt(2)))   //determines theta value
			y := math.Asin(((2 * x) + math.Sin(2*x)) / math.Pi) //determines phi value
			ydeg := (y * 180) / math.Pi                         //determines degrees for phi
			for j := 0; j <= res.Dx(); j++ {
				coords2 := float64(j) - float64(res.Dx())/2                             //sets coords
				z := (math.Pi * coords2) / (2 * radius * math.Sqrt(2) * math.Cos(x))    //determines lambda
				zdeg := (z * 180) / math.Pi                                             //determines degrees for lambda
				yIMG := int((ydeg - 90) * float64(res.Dy()) / -180)                     //determines y coordinates for new location
				zIMG := int((zdeg + 180) * float64(res.Dx()) / 360)                     //determines x coordinates for new location
				if (yIMG >= 0 && yIMG <= res.Dy()) && (zIMG >= 0 && zIMG <= res.Dx()) { //checks to make sure the coordinates are within the image boundaries
					outputIMG.Set(j, i, projection.At(zIMG, yIMG)) //replaces pixel in output image with corresponding pixel from input image
				}
			}
		}
	} else {
		var stdlat = 0.00
		if len(os.Args[1:]) > 3 {
			stdlat, err = strconv.ParseFloat(os.Args[4], 32) //converts commandline input to float
			if err != nil {
				fmt.Println("Error converting commandline argument to float.")
			}
		}
		outputIMG = image.NewRGBA(image.Rect(0, 0, res.Dx(), res.Dy())) //creates new image
		for i := 0; i <= res.Dx(); i++ {
			for j := 0; j <= res.Dy(); j++ {
				outputIMG.Set(i, j, color.RGBA{255, 255, 255, 255}) //sets every pixel to be white
			}
		}
		distort := int(float64((stdlat / 90)) * float64((res.Dy() / 2)))
		for j := 0; j <= res.Dy(); j++ {
			for i := 0; i <= res.Dx(); i++ {
				if (j < distort) || (j > (res.Dy() - distort)) {
					x := int(math.Abs(float64(res.Dx()/2) - float64(i)))
					y := int(math.Abs(float64(res.Dy()/2) - float64(j)))
					outputIMG.Set(i, j, projection.At(x, y))
				} else {
					outputIMG.Set(i, j, projection.At(i, j))
				}
			}
		}
	}

	completed, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Println("Error reading new image name.")
	}
	png.Encode(completed, outputIMG)
	defer completed.Close()
}
