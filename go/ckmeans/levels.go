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

	//	maxBIC := 0.0
	//
	//	lambda := make([]float64, Kmax)
	//	mu := make([]float64, Kmax)
	//	sigma2 := make([]float64, Kmax)
	//	coeff := make([]float64, Kmax)
	//	counts := make([]int, Kmax)
	//	weights := make([]float64, Kmax)

	for K := Kmin; K <= Kmax; K++ {
		// #
		// #     // std::vector< std::vector< size_t > > JK(J.begin(), J.begin()+K);
		// #
		// #     // Backtrack the matrix to determine boundaries between the bins.
		// #     backtrack_weighted(x, y, J, counts, weights, (int)K);
		// #
		// #
		// #     // double totalweight = std::accumulate(weights.begin(), weights.begin() + K, 0, std::plus<double>());
		// #
		// #     double totalweight;
		// #
		// #     totalweight = 0;
		// #     for(size_t k=0; k<K; k++) {
		// #       totalweight += weights[k];
		// #     }
		// #
		// #
		// #     size_t indexLeft = 0;
		// #     size_t indexRight;
		// #
		// #     for (size_t k = 0; k < K; ++k) { // Estimate GMM parameters first
		// #
		// #       lambda[k] = weights[k] / totalweight;
		// #
		// #       indexRight = indexLeft + counts[k] - 1;
		// #
		// #       shifted_data_variance_weighted(x, y, weights[k], indexLeft, indexRight, mu[k], sigma2[k]);
		// #
		// #       if(sigma2[k] == 0 || counts[k] == 1) {
		// #
		// #         double dmin;
		// #
		// #         if(indexLeft > 0 && indexRight < N-1) {
		// #           dmin = std::min(x[indexLeft] - x[indexLeft-1], x[indexRight+1] - x[indexRight]);
		// #         } else if(indexLeft > 0) {
		// #           dmin = x[indexLeft] - x[indexLeft-1];
		// #         } else {
		// #           dmin = x[indexRight+1] - x[indexRight];
		// #         }
		// #
		// #         // std::cout << "sigma2[k]=" << sigma2[k] << "==>";
		// #         if(sigma2[k] == 0) sigma2[k] = dmin * dmin / 4.0 / 9.0 ;
		// #         if(counts[k] == 1) sigma2[k] = dmin * dmin;
		// #         // std::cout << sigma2[k] << std::endl;
		// #       }
		// #
		// #       /*
		// #        if(sigma2[k] == 0) sigma2[k] = variance_min;
		// #        if(size[k] == 1) sigma2[k] = variance_max;
		// #        */
		// #
		// #       coeff[k] = lambda[k] / std::sqrt(2.0 * M_PI * sigma2[k]);
		// #
		// #       indexLeft = indexRight + 1;
		// #     }
		// #
		// #     long double loglikelihood = 0;
		// #
		// #     for (size_t i=0; i<N; ++i) {
		// #       long double L=0;
		// #       for (size_t k = 0; k < K; ++k) {
		// #         L += coeff[k] * std::exp(- (x[i] - mu[k]) * (x[i] - mu[k]) / (2.0 * sigma2[k]));
		// #       }
		// #       loglikelihood += y[i] * std::log(L);
		// #     }
		// #
		// #     // double & bic = BIC[K-Kmin];
		// #     long double bic;
		// #
		// #     // Compute the Bayesian information criterion
		// #     bic = 2 * loglikelihood - (3 * K - 1) * std::log(totalweight);  //(K*3-1)
		// #
		// #     // std::cout << "k=" << K << ": Loglh=" << loglikelihood << ", BIC=" << BIC << std::endl;
		// #
		// #     if (K == Kmin) {
		// #       maxBIC = bic;
		// #       Kopt = Kmin;
		// #     } else {
		// #       if (bic > maxBIC) {
		// #         maxBIC = bic;
		// #         Kopt = K;
		// #       }
		// #     }
		// #
		// #     BIC[K-Kmin] = (double)bic;
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
