package functions

func Ortalama(dataSet []float64) float64 {
	x := 0.0
	for i := 0; i < len(dataSet); i++ {
		x = dataSet[i] + x
	}
	return x / (float64(len(dataSet)))

}
