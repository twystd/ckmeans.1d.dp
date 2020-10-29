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
	L2 Criterion = iota + 1
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

	if weights != nil {
		panic("Weights not implemented")
	}

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
	index := make([]struct {
		i int
		x float64
	}, len(data))

	for i := range index {
		index[i] = struct {
			i int
			x float64
		}{i, data[i]}
	}

	copy(x, data)

	for i := range w {
		w[i] = 1.0
	}

	sort.Float64s(x)
	sort.Float64s(w)
	sort.SliceStable(index, func(i, j int) bool { return index[i].x < index[j].x })

	for i := range order {
		order[i] = index[i].i
	}

	// construct DP matrix
	N := len(data)
	S := make([][]float64, kmax)
	J := make([][]int, kmax)

	for i := range S {
		S[i] = make([]float64, N)
		J[i] = make([]int, N)
	}

	EWL2.FillDPMatrix(x, w, S, J)

	kopt := selectLevelsBIC(x, J, kmin, kmax)
	if kopt < kmax { // Reform the dynamic programming matrix S and J
		panic("ooops")
		//     J.erase(J.begin() + Kopt, J.end());
	}

	cluster_sorted := make([]int, N)
	centers := make([]float64, N)
	withinss := make([]float64, N)
	size := make([]float64, kmax)
	// Backtrack to find the clusters beginning and ending indices
	//     if(is_equally_weighted && criterion == L1) {
	//         backtrack_L1(x_sorted, J, &cluster_sorted[0], centers, withinss, size);
	//     } else if (is_equally_weighted && criterion == L2) {
	backtrack6(x, J, cluster_sorted, centers, withinss, size)
	//     } else if(criterion == L2Y) {
	//       backtrack_L2Y(x_sorted, y_sorted, J, &cluster_sorted[0], centers, withinss, size);
	//     } else {
	//       backtrack_weighted(x_sorted, y_sorted, J, &cluster_sorted[0], centers, withinss, size);
	//     }

	clusters := make([]int, N)
	for i := 0; i < N; i++ {
		clusters[order[i]] = cluster_sorted[i] + 1 // '1' based clustering a la 'R' implementation
	}

	return kopt, clusters, nil
}
