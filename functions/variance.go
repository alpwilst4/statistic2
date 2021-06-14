package functions

func CalculateVariance(dataSet []float64, ortalama float64) float64 {
	sonuc := 0.0
	for i := 0; i < len(dataSet); i++ {

		if dataSet[i] > ortalama {
			sonuc = sonuc + ((dataSet[i] - ortalama) * (dataSet[i] - ortalama))

		} else {
			sonuc = sonuc + ((ortalama - dataSet[i]) * (ortalama - dataSet[i]))

		}

	}

	return sonuc / (float64(len(dataSet) - 1.0))

}
