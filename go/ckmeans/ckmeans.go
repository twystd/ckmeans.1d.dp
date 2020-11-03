package ckmeans

import (
	"sort"

	"github.com/twystd/ckmeans.1d.dp/go/ckmeans/EWL2"
)

type CKMEANS struct {
	Method    Method
	EstimateK EstimateK
	Criterion Criterion
}

type Method int
type Criterion int
type EstimateK int

const (
	Linear Method = iota + 1
)

const (
	L1 Criterion = iota + 1
	L2
	L2Y
)

const (
	BIC EstimateK = iota + 1
)

func (ck *CKMEANS) CKMeans(data, weights []float64, kmin, kmax int) (int, []int, error) {
	// TODO data nil/empty
	if data == nil {
		panic("Invalid data")
	}

	// TODO len(weights) != len(data)
	if weights != nil && len(weights) != len(data) {
		panic("Invalid weights")
	}

	// TODO kmin > kmax
	if kmin > kmax {
		panic("Invalid K")
	}

	// TODO kmax > len(data)
	if kmax > len(data) {
		panic("Invalid Kmax")
	}

	// TODO move this to ckmeans (and eliminate redundant sort)
	// edge case: single unique value
	if kmax > 1 {
		sorted := make([]float64, len(data))
		copy(sorted, data)
		sort.Float64s(sorted)

		unique := 1
		p := data[0]
		for _, q := range data[1:] {
			if q != p {
				p = q
				unique += 1
			}
		}

		if unique == 1 {
			kmax = 1
		}
	}

	// special case: K=1
	if kmax == 1 {
		N := len(data)
		clusters := make([]int, N)

		for i := range clusters {
			clusters[i] = 1
		}

		return 1, clusters, nil
	}

	// K > 1

	return ck.ckmeans(data, weights, kmin, kmax)
}

// FIXME: assumes equally weighted, L2, BIC
func (ck *CKMEANS) ckmeans(data, weights []float64, kmin, kmax int) (int, []int, error) {
	// ... validate

	if ck.Method != Linear {
		panic("Only implements 'linear' method")
	}

	if ck.EstimateK != BIC {
		panic("Only implements BIC estimate-k")
	}

	if ck.Criterion != L2 {
		panic("Only implements L2 criterion")
	}

	// sort and order data
	x := make([]float64, len(data))
	w := make([]float64, len(data))
	order := make([]int, len(data))

	for i := range order {
		order[i] = i
	}

	isEquallyWeighted := true
	if weights != nil {
		for i, v := range weights[1:] {
			if weights[i] != v {
				isEquallyWeighted = false
				break
			}
		}
	}

	sort.SliceStable(order, func(i, j int) bool { return data[order[i]] < data[order[j]] })
	for i := range data {
		x[i] = data[order[i]]
		if weights != nil {
			w[i] = weights[order[i]]
		} else {
			w[i] = 1.0
		}
	}

	// construct DP matrix
	var kopt int
	N := len(data)
	S := make([][]float64, kmax)
	J := make([][]int, kmax)

	for i := range S {
		S[i] = make([]float64, N)
		J[i] = make([]int, N)
	}

	if isEquallyWeighted {
		if ck.Criterion == L2 {
			EWL2.FillDPMatrix(x, w, S, J)
		} else {
			panic("NOT IMPLEMENTED")
			// fill_dp_matrix(x_sorted, y_sorted, S, J, method, criterion);
		}

		if ck.EstimateK == BIC {
			kopt = selectLevelsBIC(x, J, kmin, kmax)
		} else {
			panic("NOT IMPLEMENTED")
			// Kopt = select_levels_3_4_12(x_sorted, J, Kmin, Kmax, BIC);
		}
	} else {
		fill_dp_matrix(x, w, S, J, ck.Method, ck.Criterion)

		switch ck.Criterion {
		case L2Y:
			panic("NOT IMPLEMENTED")
		//      if (estimate_k=="BIC") {
		//         Kopt = select_levels(y_sorted, J, Kmin, Kmax, BIC);
		//      } else {
		//         Kopt = select_levels_3_4_12(y_sorted, J, Kmin, Kmax, BIC);
		//      }
		//      break;

		default:
			if ck.EstimateK == BIC {
				bic := []float64{}
				// Choose an optimal number of levels between Kmin and Kmax
				kopt = select_levels_weighted(x, w, J, kmin, kmax, bic)
			} else {
				panic("NOT IMPLEMENTED")
				//         Kopt = select_levels_weighted_3_4_12(x_sorted, y_sorted, J, Kmin, Kmax, BIC);
			}
		}
	}

	if kopt < kmax { // Reform the dynamic programming matrix S and J
		J = J[0:kopt]
	}

	cluster_sorted := make([]int, N)
	centers := make([]float64, N)
	withinss := make([]float64, N)
	size := make([]float64, kmax)
	// Backtrack to find the clusters beginning and ending indices
	if isEquallyWeighted && ck.Criterion == L1 {
		panic("NOT IMPLEMENTED")
		//         backtrack_L1(x_sorted, J, &cluster_sorted[0], centers, withinss, size);
	} else if isEquallyWeighted && ck.Criterion == L2 {
		backtrack6(x, J, cluster_sorted, centers, withinss, size)
	} else if ck.Criterion == L2Y {
		panic("NOT IMPLEMENTED")
		//       backtrack_L2Y(x_sorted, y_sorted, J, &cluster_sorted[0], centers, withinss, size);
	} else {
		panic("NOT IMPLEMENTED")
		//       backtrack_weighted(x_sorted, y_sorted, J, &cluster_sorted[0], centers, withinss, size);
	}

	clusters := make([]int, N)
	for i := 0; i < N; i++ {
		clusters[order[i]] = cluster_sorted[i] + 1 // '1' based clustering a la 'R' implementation
	}

	return kopt, clusters, nil
}
