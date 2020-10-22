package ckmeans

import ()

func CKMeans(data, weights []float64, k int) (int, []int, error) {
	// single unique element
	N := len(data)
	clusters := make([]int, N)

	for i := range clusters {
		clusters[i] = 1
	}

	return 1, clusters, nil
}

// type Cluster struct {
// 	Cluster  []int
// 	Centers  []float64
// 	Withinss []float64
// 	Size     []float64
// }
//
//func CKMeans(data, weights []float64) ([]int, error) {
//	clusters := []Cluster{}
//
//	// single unique element
//	N := len(data)
//	cluster := make([]int, N)
//	centers := make([]float64, 1)
//	withinss := make([]float64, 1)
//	size := make([]float64, 1)
//
//	for i := range cluster {
//		cluster[i] = 1
//	}
//
//	centers[0] = data[0]
//	withinss[0] = 0.0
//
//	if weights == nil {
//		size[0] = float64(N) * 1.0
//	} else {
//		size[0] = float64(N) * weights[0] // as per the 'R' code but seems somewhat arbitrary
//	}
//
//	clusters = append(clusters, Cluster{
//		Cluster:  cluster,
//		Centers:  centers,
//		Withinss: withinss,
//		Size:     size,
//	})
//
//	return clusters, nil
//}
