package ckmeans

import ()

func backtrack3(x []float64, J [][]int, k int) []int {
	count := make([]int, k)
	N := len(J[0])
	cluster_right := N - 1
	var cluster_left int

	for q := k - 1; q >= 0; q-- {
		cluster_left = J[q][cluster_right]
		count[q] = cluster_right - cluster_left + 1
		if q > 0 {
			cluster_right = cluster_left - 1
		}
	}

	return count
}

func backtrack6(x []float64, J [][]int, cluster []int, centers []float64, withinss []float64, count []float64) {
	K := len(J)
	N := len(J[0])
	var cluster_right = N - 1

	// Backtrack the clusters from the dynamic programming matrix
	for q := K - 1; q >= 0; q-- {
		cluster_left := J[q][cluster_right]

		for i := cluster_left; i <= cluster_right; i++ {
			cluster[i] = q
		}

		sum := 0.0

		for i := cluster_left; i <= cluster_right; i++ {
			sum += x[i]
		}

		centers[q] = sum / float64(cluster_right-cluster_left+1)

		for i := cluster_left; i <= cluster_right; i++ {
			withinss[q] += (x[i] - centers[q]) * (x[i] - centers[q])
		}

		count[q] = float64(cluster_right - cluster_left + 1)

		if q > 0 {
			cluster_right = cluster_left - 1
		}
	}
}
