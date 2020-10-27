package ckmeans

import (
	"fmt"
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
	// TODO len(weights) != len(data)
	// TODO kmin > kmax
	// TODO kmax > len(data)

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

	// FIXME: assumes equally weighted, L2, BIC
	var clusters []int
	var k int

	// sort data
	x := make([]float64, len(data))
	w := make([]float64, len(data))

	copy(x, data)
	fill(w, 1.0)

	sort.Float64s(x)
	sort.Float64s(w)

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

	fmt.Printf("KOPT: %v\n", kopt)

	return k, clusters, nil
}

func fill(v []float64, f float64) {
	for i := range v {
		v[i] = f
	}
}

// type Cluster struct {
// 	Cluster  []int
// 	Centers  []float64
// 	Withinss []float64
// 	Size     []float64
// }
//
//func CKMeans(data, weights []float64) ([]int, error) {
//	clusters := []Cluster{}
//
//	// single unique element
//	N := len(data)
//	cluster := make([]int, N)
//	centers := make([]float64, 1)
//	withinss := make([]float64, 1)
//	size := make([]float64, 1)
//
//	for i := range cluster {
//		cluster[i] = 1
//	}
//
//	centers[0] = data[0]
//	withinss[0] = 0.0
//
//	if weights == nil {
//		size[0] = float64(N) * 1.0
//	} else {
//		size[0] = float64(N) * weights[0] // as per the 'R' code but seems somewhat arbitrary
//	}
//
//	clusters = append(clusters, Cluster{
//		Cluster:  cluster,
//		Centers:  centers,
//		Withinss: withinss,
//		Size:     size,
//	})
//
//	return clusters, nil
//}
