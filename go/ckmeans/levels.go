package ckmeans

import (
	"fmt"
	"math"
)

func selectLevelsBIC(x []float64, J [][]int, kmin, kmax int) int {
	fmt.Printf("X: %+v\n", x)
	fmt.Printf("J: %+v\n", J)
	fmt.Printf("K: %v,%v\n", kmin, kmax)

	N := len(x)

	if kmin > kmax || N < 2 {
		if kmin < kmax {
			return kmin
		}
		return kmax
	}

	kopt := kmin
	maxBIC := 0.0

	println(kopt, maxBIC)
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
			fmt.Printf("MU: %v\n", mu)
			fmt.Printf("SIGMA: %v\n", sigma2)
			//MU: 3,0 SIGMA: 2.5,0
			//MU: 3,8 SIGMA: 2.5,2.5
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

		fmt.Printf(">> MU: %v\n", mu)
		fmt.Printf(">> SIGMA: %v\n", sigma2)
		//>> MU: 3,8
		//>> SIGMA: 2.5,2.5
	}

	return 0
}

// # // Choose an optimal number of levels between Kmin and Kmax
// # size_t select_levels(const std::vector<double> & x,
// #                      const std::vector< std::vector< size_t > > & J,
// #                      size_t Kmin, size_t Kmax,
// #                      double * BIC)
// # {
// #   const size_t N = x.size();
// #
// #   if (Kmin > Kmax || N < 2) {
// #     return std::min(Kmin, Kmax);
// #   }
// #
// #   /*
// #   if(BIC.size() != Kmax - Kmin + 1) {
// #     BIC.resize(Kmax - Kmin + 1);
// #   }
// #   */
// #
// #   // double variance_min, variance_max;
// #   // range_of_variance(x, variance_min, variance_max);
// #
// #   size_t Kopt = Kmin;
// #
// #   double maxBIC = (0.0);
// #
// #   std::vector<double> lambda(Kmax);
// #   std::vector<double> mu(Kmax);
// #   std::vector<double> sigma2(Kmax);
// #   std::vector<double> coeff(Kmax);
// #
// #   for(size_t K = Kmin; K <= Kmax; ++K) {
// #
// #     std::vector<size_t> size(K);
// #
// #     // Backtrack the matrix to determine boundaries between the bins.
// #     backtrack(x, J, size, (int)K);
// #
// #     size_t indexLeft = 0;
// #     size_t indexRight;
// #
// #     for (size_t k = 0; k < K; ++k) { // Estimate GMM parameters first
// #       lambda[k] = size[k] / (double) N;
// #
// #       indexRight = indexLeft + size[k] - 1;
// #
// #     shifted_data_variance(x, indexLeft, indexRight, mu[k], sigma2[k]);
// #
// #     if(sigma2[k] == 0 || size[k] == 1) {
// #
// #       double dmin;
// #
// #       if(indexLeft > 0 && indexRight < N-1) {
// #         dmin = std::min(x[indexLeft] - x[indexLeft-1], x[indexRight+1] - x[indexRight]);
// #       } else if(indexLeft > 0) {
// #         dmin = x[indexLeft] - x[indexLeft-1];
// #       } else {
// #         dmin = x[indexRight+1] - x[indexRight];
// #       }
// #
// #        // std::cout << "sigma2[k]=" << sigma2[k] << "==>";
// #        if(sigma2[k] == 0) sigma2[k] = dmin * dmin / 4.0 / 9.0 ;
// #        if(size[k] == 1) sigma2[k] = dmin * dmin;
// #        // std::cout << sigma2[k] << std::endl;
// #      }
// #
// #      /*
// #       if(sigma2[k] == 0) sigma2[k] = variance_min;
// #      if(size[k] == 1) sigma2[k] = variance_max;
// #      */
// #
// #      coeff[k] = lambda[k] / std::sqrt(2.0 * M_PI * sigma2[k]);
// #
// #     indexLeft = indexRight + 1;
// #    }
//
//    double loglikelihood = 0;
//
//    for (size_t i=0; i<N; ++i) {
//      double L=0;
//      for (size_t k = 0; k < K; ++k) {
//        L += coeff[k] * std::exp(- (x[i] - mu[k]) * (x[i] - mu[k]) / (2.0 * sigma2[k]));
//      }
//      loglikelihood += std::log(L);
//    }
//
//    double & bic = BIC[K-Kmin];
//
//    // Compute the Bayesian information criterion
//    bic = 2 * loglikelihood - (3 * K - 1) * std::log((double)N);  //(K*3-1)
//
//    // std::cout << "k=" << K << ": Loglh=" << loglikelihood << ", BIC=" << BIC << std::endl;
//
//    if (K == Kmin) {
//      maxBIC = bic;
//      Kopt = Kmin;
//    } else {
//      if (bic > maxBIC) {
//        maxBIC = bic;
//        Kopt = K;
//      }
//    }
//  }
//  return Kopt;
//}

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
