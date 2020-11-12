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
		size := backtrack3(x, J, K)

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

// Choose an optimal number of levels between Kmin and Kmax
func select_levels_weighted(x, y []float64, J [][]int, Kmin, Kmax int, bic []float64) int {
	N := len(x)

	if Kmin > Kmax || N < 2 {
		if Kmax < Kmin {
			return Kmax
		}
		return Kmin
	}

	//	variance_min, variance_max := rangeOfVariance(x)

	Kopt := Kmin
	maxBIC := 0.0

	lambda := make([]float64, Kmax)
	mu := make([]float64, Kmax)
	sigma2 := make([]float64, Kmax)
	coeff := make([]float64, Kmax)
	counts := make([]int, Kmax)
	weights := make([]float64, Kmax)

	for K := Kmin; K <= Kmax; K++ {
		// Backtrack the matrix to determine boundaries between the bins.
		backtrackWeighted6(x, y, J, counts, weights, K)

		totalweight := 0.0
		for k := 0; k < K; k++ {
			totalweight += weights[k]
		}

		indexLeft := 0
		var indexRight int

		for k := 0; k < K; k++ { // Estimate GMM parameters first
			lambda[k] = weights[k] / totalweight

			indexRight = indexLeft + counts[k] - 1

			mu[k], sigma2[k] = shiftedDataVarianceWeighted(x, y, weights[k], indexLeft, indexRight)

			if sigma2[k] == 0 || counts[k] == 1 {
				var dmin float64

				if indexLeft > 0 && indexRight < N-1 {
					dmin = x[indexLeft] - x[indexLeft-1]
					if dmin > (x[indexRight+1] - x[indexRight]) {
						dmin = x[indexRight+1] - x[indexRight]
					}
				} else if indexLeft > 0 {
					dmin = x[indexLeft] - x[indexLeft-1]
				} else {
					dmin = x[indexRight+1] - x[indexRight]
				}

				if sigma2[k] == 0 {
					sigma2[k] = dmin * dmin / 4.0 / 9.0
				}
				if counts[k] == 1 {
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

			loglikelihood += y[i] * math.Log(L)
		}

		// Compute the Bayesian information criterion
		bicx := 2*loglikelihood - float64(3*K-1)*math.Log(totalweight) //(K*3-1)

		if K == Kmin {
			maxBIC = bicx
			Kopt = Kmin
		} else {
			if bicx > maxBIC {
				maxBIC = bicx
				Kopt = K
			}
		}

		bic[K-Kmin] = bicx
	}

	return Kopt
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

func rangeOfVariance(x []float64) (float64, float64) {
	dposmin := x[len(x)-1] - x[0]
	dposmax := 0.0

	for n := 1; n < len(x); n++ {
		d := x[n] - x[n-1]
		if d > 0 && dposmin > d {
			dposmin = d
		}

		if d > dposmax {
			dposmax = d
		}
	}

	variance_min := dposmin * dposmin / 3.0
	variance_max := dposmax * dposmax

	return variance_min, variance_max
}

func shiftedDataVarianceWeighted(x, y []float64, total_weight float64, left, right int) (float64, float64) {
	sum := 0.0
	sumsq := 0.0

	mean := 0.0
	variance := 0.0

	n := right - left + 1

	if right >= left {

		median := x[(left+right)/2]

		for i := left; i <= right; i++ {
			sum += (x[i] - median) * y[i]
			sumsq += (x[i] - median) * (x[i] - median) * y[i]
		}
		mean = sum/total_weight + median

		if n > 1 {
			variance = (sumsq - sum*sum/total_weight) / (total_weight - 1)
		}
	}
	return mean, variance
}
