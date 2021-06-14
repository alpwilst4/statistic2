package main

import (
	"colleration/functions"
	"fmt"
)

func main() {

	var x = []float64{100, 150, 200, 250, 300}
	var y = []float64{8, 12, 14, 15, 16}
	ortalamaX := (functions.Ortalama(x[:]))
	ortalamaY := (functions.Ortalama(y[:]))

	sdX := functions.CalculateSd(x[:], ortalamaX)
	sdY := functions.CalculateSd(y[:], ortalamaY)

	r := functions.Coleration(x[:], y[:], ortalamaX, ortalamaY, sdX, sdY)
	fmt.Println(r)

	a, b := functions.Regression(x[:], y[:], ortalamaX, ortalamaY)

	fmt.Println(a + b*180)

}
