package ckmeans

import (
	"reflect"
	"testing"
)

// library(Ckmeans.1d.dp)
//
// res <- Ckmeans.1d.dp(c(-2.5, -2.5, -2.5, -2.5), 1, method="linear")
//
// print(res$cluster)
// print(res$centers)
// print(res$size)
// print(res)

// res <- Ckmeans.1d.dp(c(-2.5, -2.5, -2.5, -2.5), 1, method="linear")
func TestCKMeansK1(t *testing.T) {
	x := []float64{-2.5, -2.5, -2.5, -2.5}

	expected := []Cluster{
		Cluster{
			Cluster:  []int{1, 1, 1, 1},
			Centers:  []float64{-2.5},
			Withinss: []float64{0.0},
			Size:     []float64{4.0},
		},
	}

	clusters, err := CKMeans(x, nil)
	if err != nil {
		t.Fatalf("Unexpected error (%v)", err)
	}

	if !reflect.DeepEqual(clusters, expected) {
		t.Errorf("Returned invalid clusters:\n   expected: %v\n   got:      %v\n", expected, clusters)
	}
}

// res <- Ckmeans.1d.dp(c(-2.5, -2.5, -2.5, -2.5), 1, c(1.2, 1.1, 0.9, 0.8), method="linear")
func TestCKMeansWeightedK1(t *testing.T) {
	x := []float64{-2.5, -2.5, -2.5, -2.5}
	w := []float64{1.2, 1.1, 0.9, 0.8}

	expected := []Cluster{
		Cluster{
			Cluster:  []int{1, 1, 1, 1},
			Centers:  []float64{-2.5},
			Withinss: []float64{0.0},
			Size:     []float64{4.8},
		},
	}

	clusters, err := CKMeans(x, w)
	if err != nil {
		t.Fatalf("Unexpected error (%v)", err)
	}

	if !reflect.DeepEqual(clusters, expected) {
		t.Errorf("Returned invalid clusters:\n   expected: %v\n   got:      %v\n", expected, clusters)
	}
}
