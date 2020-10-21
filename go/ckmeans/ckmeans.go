package ckmeans

import ()

type Cluster struct {
	Cluster  []int
	Centers  []float64
	Withinss []float64
	Size     []float64
}

func CKMeans(data, weights []float64) ([]Cluster, error) {
	clusters := []Cluster{}

	// single unique element
	N := len(data)
	cluster := make([]int, N)
	centers := make([]float64, 1)
	withinss := make([]float64, 1)
	size := make([]float64, 1)

	for i := range cluster {
		cluster[i] = 1
	}

	centers[0] = data[0]
	withinss[0] = 0.0

	if weights == nil {
		size[0] = float64(N) * 1.0
	} else {
		size[0] = float64(N) * weights[0]
	}

	clusters = append(clusters, Cluster{
		Cluster:  cluster,
		Centers:  centers,
		Withinss: withinss,
		Size:     size,
	})

	return clusters, nil
}
