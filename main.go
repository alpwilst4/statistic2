package main

import (
	"colleration/functions"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var y_egitim []float64
	var y_test []float64
	var x_egitim []float64
	var x_test []float64
	var x1 = []float64{100, 110, 112, 115, 117, 116, 118, 120, 121, 120, 117, 123}
	var x2 = []float64{30, 35, 22, 29, 6, 38, 40, 43, 32, 45, 50, 20}
	var y = []float64{5.5, 5.8, 6, 5.9, 6.2, 6.3, 6.5, 6.6, 6.4, 6.5, 6.7, 6.8}
	var y1_egitim []float64
	var x1_egitim []float64
	var x1_test []float64
	var y1_test []float64
	var y2_egitim []float64
	var x2_egitim []float64
	var x2_test []float64
	var y2_test []float64
	egitimBoyutu := int(float64(len(x1)) / 100 * 70)
	testBoyutu := len(x1) - egitimBoyutu

	for i := 0; i < egitimBoyutu; i++ {
		rand.Seed(time.Now().UnixNano())
		min := 0
		max := len(x1) - 1
		randomNumber := rand.Intn(max-min+1) + min

		x1_egitim = append(x1_egitim, x1[randomNumber])
		y1_egitim = append(y1_egitim, y[randomNumber])
		x2_egitim = append(x2_egitim, x2[randomNumber])
		y2_egitim = append(y2_egitim, y[randomNumber])

	}
	for i := 0; i < testBoyutu; i++ {
		rand.Seed(time.Now().UnixNano())
		min := 0
		max := len(x1) - 1
		randomNumber := rand.Intn(max-min+1) + min

		x1_test = append(x1_test, x1[randomNumber])
		y1_test = append(y1_test, y[randomNumber])
		x2_test = append(x2_test, x2[randomNumber])
		y2_test = append(y2_test, y[randomNumber])

	}

	ortalamaX2 := (functions.Ortalama(x2_egitim[:]))
	ortalamaX := (functions.Ortalama(x1_egitim[:]))
	ortalamaY := (functions.Ortalama(y1_egitim[:]))
	ortalamaY2 := (functions.Ortalama(y2_egitim[:]))

	sdX2 := functions.CalculateSd(x2_egitim[:], ortalamaX2)
	sdX := functions.CalculateSd(x1_egitim[:], ortalamaX)
	sdY := functions.CalculateSd(y1_egitim[:], ortalamaY)
	sdY2 := functions.CalculateSd(y2_egitim[:], ortalamaY2)

	r1 := functions.Coleration(x1_egitim[:], y1_egitim[:], ortalamaX, ortalamaY, sdX, sdY)
	r2 := functions.Coleration(x2_egitim[:], y2_egitim[:], ortalamaX2, ortalamaY2, sdX2, sdY2)
	if r1 > r2 {
		x_egitim = x1_egitim
		x_test = x1_test
		y_egitim = y1_egitim
		y_test = y1_test
	} else {
		x_egitim = x2_egitim
		x_test = x2_test
		y_egitim = y2_egitim
		y_test = y2_test
	}
	ssh_egitim := 0.0
	ssh_test := 0.0
	a, b := functions.Regression(x1_egitim[:], y1_egitim[:], ortalamaX, ortalamaY)
	fmt.Println("Eğitim")
	for i := 0; i < len(x_egitim); i++ {
		yi := a + b*x_egitim[i]
		ssh_egitim = ssh_egitim + (y_egitim[i]-yi)*(y_egitim[i]-yi)
		fmt.Print(x_egitim[i], yi, "\n")
	}
	fmt.Println(ssh_egitim)

	fmt.Println("Test")
	for i := 0; i < len(x_test); i++ {
		yi := a + b*x_test[i]
		ssh_test = ssh_test + (y_test[i]-yi)*(y_test[i]-yi)
		fmt.Print(x_test[i], yi, "\n")
	}
	fmt.Println(ssh_test)
}
