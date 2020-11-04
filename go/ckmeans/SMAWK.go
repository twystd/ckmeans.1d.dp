package ckmeans

import ()

func SMAWK(imin, imax, istep, q int, js []int, S [][]float64, J [][]int, sum_x, sum_x_sq, sum_w, sum_w_sq []float64, criterion Criterion) {
	if imax-imin <= 0*istep { // base case only one element left
		find_min_from_candidates(imin, imax, istep, q, js, S, J, sum_x, sum_x_sq, sum_w, sum_w_sq, criterion)
	} else {
		// REDUCE

		js_odd := make([]int, len(js))

		reduce_in_place(imin, imax, istep, q, js, js_odd, S, J, sum_x, sum_x_sq, sum_w, sum_w_sq, criterion)

		istepx2 := istep << 1
		imin_odd := imin + istep
		imax_odd := imin_odd + (imax-imin_odd)/istepx2*istepx2

		// Recursion on odd rows (0-based):
		SMAWK(imin_odd, imax_odd, istepx2, q, js_odd, S, J, sum_x, sum_x_sq, sum_w, sum_w_sq, criterion)

		fill_even_positions(imin, imax, istep, q, js, S, J, sum_x, sum_x_sq, sum_w, sum_w_sq, criterion)
	}
}

func fill_even_positions(imin, imax, istep, q int, js []int, S [][]float64, J [][]int, sum_x, sum_x_sq, sum_w, sum_w_sq []float64, criterion Criterion) {
	// Derive j for even rows (0-based)
	n := len(js)
	istepx2 := istep << 1
	jl := js[0]
	r := 0

	for i := imin; i <= imax; i += istepx2 {
		for js[r] < jl {
			// Increase r until it points to a value of at least jmin
			r++
		}

		// Initialize S[q][i] and J[q][i]
		S[q][i] = S[q-1][js[r]-1] + dissimilarity(criterion, js[r], i, sum_x, sum_x_sq, sum_w, sum_w_sq)
		// ssq(js[r], i, sum_x, sum_x_sq, sum_w);
		J[q][i] = js[r] // rmin

		// Look for minimum S upto jmax within js
		var jh int

		if i+istep <= imax {
			jh = J[q][i+istep]
		} else {
			jh = js[n-1]
		}

		jmax := jh
		if i < jmax {
			jmax = i
		}

		sjimin := dissimilarity(criterion, jmax, i, sum_x, sum_x_sq, sum_w, sum_w_sq) // ssq(jmax, i, sum_x, sum_x_sq, sum_w)

		r++
		for ; r < n && js[r] <= jmax; r++ {
			jabs := js[r]

			if jabs > i {
				break
			}

			if jabs < J[q-1][i] {
				continue
			}

			s := dissimilarity(criterion, jabs, i, sum_x, sum_x_sq, sum_w, sum_w_sq)
			// (ssq(jabs, i, sum_x, sum_x_sq, sum_w));
			Sj := S[q-1][jabs-1] + s

			if Sj <= S[q][i] {
				S[q][i] = Sj
				J[q][i] = js[r]
			} else if S[q-1][jabs-1]+sjimin > S[q][i] {
				break
			}
		}
		r--
		jl = jh
	}
}

func find_min_from_candidates(imin, imax, istep, q int, js []int, S [][]float64, J [][]int, sum_x, sum_x_sq, sum_w, sum_w_sq []float64, criterion Criterion) {
	rmin_prev := 0

	for i := imin; i <= imax; i += istep {
		rmin := rmin_prev

		// Initialization of S[q][i] and J[q][i]
		S[q][i] = S[q-1][js[rmin]-1] + dissimilarity(criterion, js[rmin], i, sum_x, sum_x_sq, sum_w, sum_w_sq)
		// ssq(js[rmin], i, sum_x, sum_x_sq, sum_w);
		J[q][i] = js[rmin]

		for r := (rmin + 1); r < len(js); r++ {
			j_abs := js[r]

			if j_abs < J[q-1][i] {
				continue
			}
			if j_abs > i {
				break
			}

			Sj := (S[q-1][j_abs-1] + dissimilarity(criterion, j_abs, i, sum_x, sum_x_sq, sum_w, sum_w_sq))
			// ssq(j_abs, i, sum_x, sum_x_sq, sum_w));
			if Sj <= S[q][i] {
				S[q][i] = Sj
				J[q][i] = js[r]
				rmin_prev = r
			}
		}
	}
}

func reduce_in_place(imin, imax, istep, q int, js, js_red []int, S [][]float64, J [][]int, sum_x, sum_x_sq, sum_w, sum_w_sq []float64, criterion Criterion) {
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
		Sl := S[q-1][j-1] + dissimilarity(criterion, j, i, sum_x, sum_x_sq, sum_w, sum_w_sq)
		// ssq(j, i, sum_x, sum_x_sq, sum_w));

		jplus1 := js_red[right+1]
		Slplus1 := S[q-1][jplus1-1] + dissimilarity(criterion, jplus1, i, sum_x, sum_x_sq, sum_w, sum_w_sq)
		// ssq(jplus1, i, sum_x, sum_x_sq, sum_w));

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
				// delete position / column p and
				//   move back to previous position / column p-1:
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

func fill_row_q_SMAWK(imin, imax, q int, S [][]float64, J [][]int, sum_x, sum_x_sq, sum_w, sum_w_sq []float64, criterion Criterion) {
	// Assumption: each cluster must have at least one point.
	js := make([]int, imax-q+1)
	abs := q

	for i := range js {
		js[i] = abs
		abs++
	}

	SMAWK(imin, imax, 1, q, js, S, J, sum_x, sum_x_sq, sum_w, sum_w_sq, criterion)
}
