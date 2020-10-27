package ckmeans

import (
	"math"
)

func selectLevelsBIC(x []float64, J [][]int, kmin, kmax int) int {
	N := len(x)

	if kmin > kmax || N < 2 {
		if kmin < kmax {
			return kmin
		}
		return kmax
	}

	kopt := kmin
	maxBIC := 0.0

	lambda := make([]float64, kmax)
	mu := make([]float64, kmax)
	sigma2 := make([]float64, kmax)
	coeff := make([]float64, kmax)

	for K := kmin; K <= kmax; K++ {
		// Backtrack the matrix to determine boundaries between the bins
		size := backtrack(x, J, K)

		indexLeft := 0
		var indexRight int

		for k := 0; k < K; k++ {
			// Estimate GMM parameters first
			lambda[k] = float64(size[k]) / float64(N)
			indexRight = indexLeft + size[k] - 1

			mu[k], sigma2[k] = shiftedDataVariance(x, indexLeft, indexRight)
			if sigma2[k] == 0 || size[k] == 1 {
				dmin := 0.0

				if indexLeft > 0 && indexRight < N-1 {
					dmin = math.Min(x[indexLeft]-x[indexLeft-1], x[indexRight+1]-x[indexRight])
				} else if indexLeft > 0 {
					dmin = x[indexLeft] - x[indexLeft-1]
				} else {
					dmin = x[indexRight+1] - x[indexRight]
				}

				if sigma2[k] == 0 {
					sigma2[k] = dmin * dmin / 4.0 / 9.0
				}

				if size[k] == 1 {
					sigma2[k] = dmin * dmin
				}
			}

			coeff[k] = lambda[k] / math.Sqrt(2.0*math.Pi*sigma2[k])

			indexLeft = indexRight + 1
		}

		loglikelihood := 0.0

		for i := 0; i < N; i++ {
			L := 0.0
			for k := 0; k < K; k++ {
				L += coeff[k] * math.Exp(-(x[i]-mu[k])*(x[i]-mu[k])/(2.0*sigma2[k]))
			}

			loglikelihood += math.Log(L)
		}

		// compute the Bayesian information criterion
		bic := 2*loglikelihood - float64(3*K-1)*math.Log(float64(N)) //(K*3-1)

		if K == kmin {
			maxBIC = bic
			kopt = kmin
		} else {
			if bic > maxBIC {
				maxBIC = bic
				kopt = K
			}
		}
	}

	return kopt
}

func shiftedDataVariance(x []float64, left, right int) (float64, float64) {
	sum := 0.0
	sumsq := 0.0
	mean := 0.0
	variance := 0.0

	n := right - left + 1

	if right >= left {
		median := x[(left+right)/2]

		for i := left; i <= right; i++ {
			sum += x[i] - median
			sumsq += (x[i] - median) * (x[i] - median)
		}

		mean = sum/float64(n) + median

		if n > 1 {
			variance = (sumsq - sum*sum/float64(n)) / float64(n-1)
		}
	}

	return mean, variance
}
