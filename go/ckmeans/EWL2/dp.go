package EWL2

/*
   x: One dimension vector to be clustered, must be sorted (in any order).
   S: K x N matrix. S[q][i] is the sum of squares of the distance from
   each x[i] to its cluster mean when there are exactly x[i] is the
   last point in cluster q
   J: K x N backtrack matrix
*/
func FillDP(x, w []float64, S [][]float64, J [][]int) {
	K := len(S)
	N := len(S[0])

	sum_x := make([]float64, N)
	sum_x_sq := make([]float64, N)
	//	jseq := make([]int, N) // ??
	shift := x[N/2] // median. used to shift the values of x to numerical stability

	sum_x[0] = x[0] - shift
	sum_x_sq[0] = (x[0] - shift) * (x[0] - shift)

	S[0][0] = 0
	J[0][0] = 0

	for i := 1; i < N; i++ {
		dx := x[i] - shift
		sum_x[i] = sum_x[i-1] + dx
		sum_x_sq[i] = sum_x_sq[i-1] + dx*dx

		// initialize for q=0
		S[0][i] = dissimilarity(0, i, sum_x, sum_x_sq)
		J[0][i] = 0
	}

	for q := 1; q < K; q++ {
		imin := 1
		if q < K-1 {
			if q > 1 {
				imin = q
			}
		} else {
			// No need to compute S[K-1][0] ... S[K-1][N-2]
			imin = N - 1
		}

		//                if(method == "linear") {
		fill_row_q_SMAWK(imin, N-1, q, S, J, sum_x, sum_x_sq)
		//                } else if(method == "loglinear") {
		//                        fill_row_q_log_linear(imin, N-1, q, q, N-1, S, J, sum_x, sum_x_sq);
		//                } else if(method == "quadratic") {
		//                        fill_row_q(imin, N-1, q, S, J, sum_x, sum_x_sq);
		//                } else {
		//                        throw std::string("ERROR: unknown method") + method + "!";
		//                }
	}
}

func dissimilarity(j, i int, sum_x, sum_x_sq []float64) float64 {
	return ssq(j, i, sum_x, sum_x_sq)
}

func ssq(j, i int, sum_x, sum_x_sq []float64) float64 {
	var sji float64 = 0.0

	if j >= i {
		sji = 0.0
	} else if j > 0 {
		muji := (sum_x[i] - sum_x[j-1]) / float64(i-j+1)
		sji = sum_x_sq[i] - sum_x_sq[j-1] - float64(i-j+1)*muji*muji
	} else {
		sji = sum_x_sq[i] - sum_x[i]*sum_x[i]/float64(i+1)
	}

	if sji < 0.0 {
		return 0.0
	}

	return sji
}

func fill_row_q_SMAWK(imin, imax, q int, S [][]float64, J [][]int, sum_x, sum_x_sq []float64) {
	// ASSUMPTION: each cluster must have at least one point

	js := make([]int, imax-q+1)
	abs := q

	for i := range js {
		js[i] = abs
		abs += 1
	}

	smawk(imin, imax, 1, q, js, S, J, sum_x, sum_x_sq)
}

func smawk(imin, imax, istep, q int, js []int, S [][]float64, J [][]int, sum_x, sum_x_sq []float64) {
	if imax-imin <= 0*istep {
		// base case only one element left
		find_min_from_candidates(imin, imax, istep, q, js, S, J, sum_x, sum_x_sq)
	} else {
		panic("NOT IMPLEMENTED")
		//    // REDUCE
		//
		//    std::vector<size_t> js_odd;
		//
		//    reduce_in_place(imin, imax, istep, q, js, js_odd,
		//                    S, J, sum_x, sum_x_sq);
		//
		//    int istepx2 = (istep << 1);
		//    int imin_odd = (imin + istep);
		//    int imax_odd = (imin_odd + (imax - imin_odd) / istepx2 * istepx2);
		//
		//    // Recursion on odd rows (0-based):
		//    SMAWK(imin_odd, imax_odd, istepx2,
		//          q, js_odd, S, J, sum_x, sum_x_sq);
		//
		//    fill_even_positions(imin, imax, istep, q, js,
		//                        S, J, sum_x, sum_x_sq);
	}
}

func find_min_from_candidates(imin, imax, istep, q int, js []int, S [][]float64, J [][]int, sum_x, sum_x_sq []float64) {
	rmin_prev := 0

	for i := imin; i <= imax; i += istep {

		rmin := rmin_prev

		// Initialization of S[q][i] and J[q][i]
		S[q][i] = S[q-1][js[rmin]-1] + dissimilarity(js[rmin], i, sum_x, sum_x_sq)

		// ssq(js[rmin], i, sum_x, sum_x_sq, sum_w);
		J[q][i] = js[rmin]

		for r := rmin + 1; r < len(js); r++ {

			j_abs := js[r]

			if j_abs < J[q-1][i] {
				continue
			}

			if j_abs > i {
				break
			}

			Sj := (S[q-1][j_abs-1] + dissimilarity(j_abs, i, sum_x, sum_x_sq))
			// ssq(j_abs, i, sum_x, sum_x_sq, sum_w));
			if Sj <= S[q][i] {
				S[q][i] = Sj
				J[q][i] = js[r]
				rmin_prev = r
			}
		}
	}
}
