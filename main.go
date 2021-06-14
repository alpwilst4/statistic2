package main

import (
	"colleration/functions"
	"fmt"
)

func main() {

	var x = []float64{100, 110, 112, 115, 117, 116, 118, 120, 121, 120, 117, 123}
	var y = []float64{5.5, 5.8, 6, 5.9, 6.2, 6.3, 6.5, 6.6, 6.4, 6.5, 6.7, 6.8}
	ortalamaX := (functions.Ortalama(x[:]))
	ortalamaY := (functions.Ortalama(y[:]))

	sdX := functions.CalculateSd(x[:], ortalamaX)
	sdY := functions.CalculateSd(y[:], ortalamaY)

	r := functions.Coleration(x[:], y[:], ortalamaX, ortalamaY, sdX, sdY)
	fmt.Println(r)

	functions.Regression(x[:], y[:], ortalamaX, ortalamaY, sdX, sdY)

}
