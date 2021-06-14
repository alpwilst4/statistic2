package functions

import (
	"fmt"
)

func Regression(dataSetX []float64, dataSetY []float64, ortalamaX float64, ortalamaY float64, sdX float64, sdY float64) float64 {
	sonuc1 := 0.0
	sonucToplamX := 0.0
	sonucToplamY := 0.0
	sonucKareToplamX := 0.0

	for i := 0; i < len(dataSetX); i++ {
		sonuc1 = sonuc1 + (dataSetX[i] * dataSetY[i])
		sonucToplamX = sonucToplamX + dataSetX[i]
		sonucToplamY = sonucToplamY + dataSetY[i]
		sonucKareToplamX = sonucKareToplamX + (dataSetX[i])*(dataSetX[i])

	}
	sonucKareToplamX = float64(len(dataSetX)) * sonucKareToplamX
	sonuc1 = float64(len(dataSetX)) * sonuc1

	b := (sonuc1 - (sonucToplamX * sonucToplamY)) / (sonucKareToplamX - (sonucToplamX * sonucToplamX))

	a := ortalamaY - (b * ortalamaX)

	fmt.Println(b)
	return a

}
