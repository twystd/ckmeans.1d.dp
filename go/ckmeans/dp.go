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

// x: One dimension vector to be clustered, must be sorted (in any order).
// S: K x N matrix. S[q][i] is the sum of squares of the distance from
// each x[i] to its cluster mean when there are exactly x[i] is the
// last point in cluster q
// J: K x N backtrack matrix
//
// NOTE: All vector indices in this program start at position 0
func fill_dp_matrix(x, w []float64, S [][]float64, J [][]int, method Method, criterion Criterion) {
	K := len(S)
	N := len(S[0])

	sum_x := make([]float64, N)
	sum_x_sq := make([]float64, N)
	sum_w := make([]float64, len(w))
	sum_w_sq := make([]float64, len(w))

	//jseq := []int{}

	shift := x[N/2] // median. used to shift the values of x to improve numerical stability

	if len(w) == 0 { // equally weighted
		sum_x[0] = x[0] - shift
		sum_x_sq[0] = (x[0] - shift) * (x[0] - shift)
	} else { // unequally weighted
		sum_x[0] = w[0] * (x[0] - shift)
		sum_x_sq[0] = w[0] * (x[0] - shift) * (x[0] - shift)
		sum_w[0] = w[0]
		sum_w_sq[0] = w[0] * w[0]
	}

	S[0][0] = 0
	J[0][0] = 0

	for i := 1; i < N; i++ {
		if len(w) == 0 { // equally weighted
			sum_x[i] = sum_x[i-1] + x[i] - shift
			sum_x_sq[i] = sum_x_sq[i-1] + (x[i]-shift)*(x[i]-shift)
		} else { // unequally weighted
			sum_x[i] = sum_x[i-1] + w[i]*(x[i]-shift)
			sum_x_sq[i] = sum_x_sq[i-1] + w[i]*(x[i]-shift)*(x[i]-shift)
			sum_w[i] = sum_w[i-1] + w[i]
			sum_w_sq[i] = sum_w_sq[i-1] + w[i]*w[i]
		}

		// Initialize for q = 0
		S[0][i] = dissimilarity(criterion, 0, i, sum_x, sum_x_sq, sum_w, sum_w_sq) // ssq(0, i, sum_x, sum_x_sq, sum_w);
		J[0][i] = 0
	}

	for q := 1; q < K; q++ {
		var imin int
		if q < K-1 {
			imin = 1
			if q > imin {
				imin = q
			}
		} else {
			// No need to compute S[K-1][0] ... S[K-1][N-2]
			imin = N - 1
		}

		// fill_row_k_linear_recursive(imin, N-1, 1, q, jseq, S, J, sum_x, sum_x_sq);
		// fill_row_k_linear(imin, N-1, q, S, J, sum_x, sum_x_sq);
		if method == Linear {
			// #       fill_row_q_SMAWK(imin, N-1, q, S, J, sum_x, sum_x_sq, sum_w, sum_w_sq, criterion);
			//		} else if method == "loglinear" {
			//			panic("NOT IMPLEMENTED")
			//			// #       fill_row_q_log_linear(imin, N-1, q, q, N-1, S, J, sum_x, sum_x_sq, sum_w, sum_w_sq, criterion);
			//		} else if method == "quadratic" {
			//			panic("NOT IMPLEMENTED")
			//			// #       fill_row_q(imin, N-1, q, S, J, sum_x, sum_x_sq, sum_w, sum_w_sq, criterion);
		} else {
			panic("NOT IMPLEMENTED")
			// #       throw std::string("ERROR: unknown method") + method + "!";
		}
	}
}

func dissimilarity(dis Criterion, j, i int, sum_x, sum_x_sq, sum_w, sum_w_sq []float64) float64 {
	d := 0.0

	switch dis {
	case L1:
		d = sabs(j, i, sum_x, sum_w)
		break
	case L2:
		d = ssq(j, i, sum_x, sum_x_sq, sum_w)
		break
	case L2Y:
		d = ssq(j, i, sum_w, sum_w_sq, nil)
		break
	}
	return d
}

func sabs(j, i int, sum_x, sum_w []float64) float64 {
	sji := 0.0

	if sum_w == nil { // equally weighted version
		if j >= i {
			sji = 0.0
		} else if j > 0 {
			l := (i + j) >> 1 // l is the index to the median of the cluster

			if ((i - j + 1) % 2) == 1 {
				// If i-j+1 is odd, we have
				//   sum (x_l - x_m) over m = j .. l-1
				//   sum (x_m - x_l) over m = l+1 .. i
				sji = -sum_x[l-1] + sum_x[j-1] + sum_x[i] - sum_x[l]
			} else {
				// If i-j+1 is even, we have
				//   sum (x_l - x_m) over m = j .. l
				//   sum (x_m - x_l) over m = l+1 .. i
				sji = -sum_x[l] + sum_x[j-1] + sum_x[i] - sum_x[l]
			}
		} else { // j==0
			l := i >> 1 // l is the index to the median of the cluster

			if ((i + 1) % 2) == 1 {
				// If i-j+1 is odd, we have
				//   sum (x_m - x_l) over m = 0 .. l-1
				//   sum (x_l - x_m) over m = l+1 .. i
				sji = -sum_x[l-1] + sum_x[i] - sum_x[l]
			} else {
				// If i-j+1 is even, we have
				//   sum (x_m - x_l) over m = 0 .. l
				//   sum (x_l - x_m) over m = l+1 .. i
				sji = -sum_x[l] + sum_x[i] - sum_x[l]
			}
		}
	} else { // unequally weighted version
		// #     // no exact solutions are known.
	}

	if sji < 0.0 {
		sji = 0.0
	}

	return sji
}

func ssq(j, i int, sum_x, sum_x_sq, sum_w []float64) float64 {
	sji := 0.0

	if sum_w == nil { // equally weighted version
		if j >= i {
			sji = 0.0
		} else if j > 0 {
			muji := (sum_x[i] - sum_x[j-1]) / float64(i-j+1)
			sji = sum_x_sq[i] - sum_x_sq[j-1] - float64(i-j+1)*muji*muji
		} else {
			sji = sum_x_sq[i] - sum_x[i]*sum_x[i]/float64(i+1)
		}
	} else { // unequally weighted version
		if sum_w[j] >= sum_w[i] {
			sji = 0.0
		} else if j > 0 {
			muji := (sum_x[i] - sum_x[j-1]) / (sum_w[i] - sum_w[j-1])
			sji = sum_x_sq[i] - sum_x_sq[j-1] - (sum_w[i]-sum_w[j-1])*muji*muji
		} else {
			sji = sum_x_sq[i] - sum_x[i]*sum_x[i]/sum_w[i]
		}
	}

	if sji < 0.0 {
		sji = 0.0
	}

	return sji
}
