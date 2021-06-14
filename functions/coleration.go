package functions

func Coleration(dataSetX []float64, dataSetY []float64, ortalamaX float64, ortalamaY float64, sdX float64, sdY float64) float64 {
	sonucX := 0.0
	sonucY := 0.0
	sonucBind := 0.0

	for i := 0; i < len(dataSetX); i++ {
		sonucX = (dataSetX[i] - ortalamaX)
		sonucY = (dataSetY[i] - ortalamaY)
		sonucBind = sonucBind + (sonucX * sonucY)
	}

	return sonucBind / (float64(len(dataSetX)-1) * sdX * sdY)

}
