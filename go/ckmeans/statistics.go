package ckmeans

func centers(data, weights []float64, k int, index []int) []float64 {
	sum := make([]float64, k)
	sumw := make([]float64, k)
	centers := make([]float64, k)

	if weights == nil || len(weights) == 0 {
		for i, ix := range index {
			sum[ix] += data[i]
			sumw[ix] += 1.0
		}
	} else {
		for i, ix := range index {
			sum[ix] += data[i] * weights[i]
			sumw[ix] += weights[i]
		}
	}

	for i := 0; i < k; i++ {
		centers[i] = sum[i] / sumw[i]
	}

	return centers
}

func withinss(data, weights []float64, k int, index []int) []float64 {
	withinss := make([]float64, k)
	centers := centers(data, weights, k, index)

	if weights == nil || len(weights) == 0 {
		for i, ix := range index {
			withinss[ix] += (data[i] - centers[ix]) * (data[i] - centers[ix])
		}
	} else {
		for i, ix := range index {
			withinss[ix] += weights[i] * (data[i] - centers[ix]) * (data[i] - centers[ix])
		}
	}

	return withinss
}
