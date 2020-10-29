package EWL2

/*
   x: One dimension vector to be clustered, must be sorted (in any order).
   S: K x N matrix. S[q][i] is the sum of squares of the distance from
   each x[i] to its cluster mean when there are exactly x[i] is the
   last point in cluster q
   J: K x N backtrack matrix
*/
func FillDPMatrix(x, w []float64, S [][]float64, J [][]int) {
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
		// REDUCE

		js_odd := make([]int, len(js))

		reduce_in_place(imin, imax, istep, q, js, js_odd, S, J, sum_x, sum_x_sq)

		istepx2 := istep << 1
		imin_odd := imin + istep
		imax_odd := imin_odd + (imax-imin_odd)/istepx2*istepx2

		// Recursion on odd rows (0-based):
		smawk(imin_odd, imax_odd, istepx2, q, js_odd, S, J, sum_x, sum_x_sq)

		fill_even_positions(imin, imax, istep, q, js, S, J, sum_x, sum_x_sq)
	}
}

func fill_even_positions(imin, imax, istep, q int, js []int, S [][]float64, J [][]int, sum_x, sum_x_sq []float64) {
	panic("NOT IMPLEMENTED")
	// #   // Derive j for even rows (0-based)
	// #   size_t n = (js.size());
	// #   int istepx2 = (istep << 1);
	// #   size_t jl = (js[0]);
	// #   for(int i=(imin), r(0); i<=imax; i+=istepx2) {
	// #
	// #     // auto jmin = (i == imin) ? js[0] : J[q][i - istep];
	// #
	// #     while(js[r] < jl) {
	// #       // Increase r until it points to a value of at least jmin
	// #       r ++;
	// #     }
	// #
	// #     // Initialize S[q][i] and J[q][i]
	// #     S[q][i] = S[q-1][js[r]-1] +
	// #       dissimilarity(js[r], i, sum_x, sum_x_sq);
	// #     // ssq(js[r], i, sum_x, sum_x_sq, sum_w);
	// #     J[q][i] = js[r]; // rmin
	// #
	// #     // Look for minimum S upto jmax within js
	// #     int jh = (int) ( (i + istep <= imax) ? J[q][i + istep] : js[n-1] );
	// #
	// #     int jmax = std::min((int)jh, (int)i);
	// #
	// #     ldouble sjimin(
	// #         dissimilarity(jmax, i, sum_x, sum_x_sq)
	// #       // ssq(jmax, i, sum_x, sum_x_sq, sum_w)
	// #     );
	// #
	// #     for(++ r; r < n && js[r]<=jmax; r++) {
	// #
	// #       const size_t & jabs = js[r];
	// #
	// #       if(jabs > i) break;
	// #
	// #       if(jabs < J[q-1][i]) continue;
	// #
	// #       ldouble s =
	// #         dissimilarity(jabs, i, sum_x, sum_x_sq);
	// #       // (ssq(jabs, i, sum_x, sum_x_sq, sum_w));
	// #       ldouble Sj = (S[q-1][jabs-1] + s);
	// #
	// #       if(Sj <= S[q][i]) {
	// #         S[q][i] = Sj;
	// #         J[q][i] = js[r];
	// #       } else if(S[q-1][jabs-1] + sjimin > S[q][i]) {
	// #         break;
	// #       } /*else if(S[q-1][js[rmin]-1] + s > S[q][i]) {
	// #  break;
	// #       } */
	// #     }
	// #     r --;
	// #     jl = jh;
	// #   }
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

func reduce_in_place(imin, imax, istep, q int, js, js_red []int, S [][]float64, J [][]int, sum_x, sum_x_sq []float64) {
	N := (imax-imin)/istep + 1

	copy(js_red, js)

	if N >= len(js) {
		return
	}

	// Two positions to move candidate j's back and forth
	left := -1 // points to last favorable position / column
	right := 0 // points to current position / column

	m := len(js_red)

	for m > N { // js_reduced has more than N positions / columns
		p := left + 1
		i := imin + p*istep
		j := js_red[right]
		Sl := S[q-1][j-1] + dissimilarity(j, i, sum_x, sum_x_sq)

		// ssq(j, i, sum_x, sum_x_sq, sum_w));
		jplus1 := js_red[right+1]
		Slplus1 := S[q-1][jplus1-1] + dissimilarity(jplus1, i, sum_x, sum_x_sq)

		println("??", m, Sl, Slplus1)

		if Sl < Slplus1 && p < N-1 {
			left++
			js_red[left] = j // i += istep;
			right++          // move on to next position / column p+1
		} else if Sl < Slplus1 && p == N-1 {
			right++
			js_red[right] = j // delete position / column p+1
			m--
		} else { // (Sl >= Slplus1)
			if p > 0 { // i > imin
				// delete position / column p and move back to previous position / column p-1:
				js_red[right] = js_red[left]
				left--
				// p --; // i -= istep;
			} else {
				right++ // delete position / column 0
			}
			m--
		}
	}

	for r := left + 1; r < m; r++ {
		js_red[r] = js_red[right]
		right++
	}

	tmp := make([]int, m)
	copy(tmp, js_red)
	js_red = tmp
}
