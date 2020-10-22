package ckmeans

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

// 'R':
// library(Ckmeans.1d.dp)
//
// res <- Ckmeans.1d.dp(x, k, w, method="linear")
//
// print(res$cluster)
// print(res$centers)
// print(res$size)
// print(res)

// res <- Ckmeans.1d.dp(c(-2.5, -2.5, -2.5, -2.5), 1, method="linear")
func TestCKMeansWithSingleUniqueValue(t *testing.T) {
	x := []float64{-2.5, -2.5, -2.5, -2.5}
	expected := []int{1, 1, 1, 1}

	k, clusters, err := CKMeans(x, nil, 4)
	if err != nil {
		t.Fatalf("Unexpected error (%v)", err)
	}

	if k != 1 {
		t.Errorf("Expected K=%v, got: %v\n", 1, k)
	}

	if !reflect.DeepEqual(clusters, expected) {
		t.Errorf("Returned invalid clusters:\n   expected: %v\n   got:      %v\n", expected, clusters)
	}
}

func TestCKMeansK1(t *testing.T) {
	x := make([]float64, 100)
	expected := make([]int, 100)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range x {
		x[i] = r.Float64()
	}

	for i := range expected {
		expected[i] = 1
	}

	k, clusters, err := CKMeans(x, nil, 1)
	if err != nil {
		t.Fatalf("Unexpected error (%v)", err)
	}

	if k != 1 {
		t.Errorf("Expected K=%v, got: %v\n", 1, k)
	}

	if !reflect.DeepEqual(clusters, expected) {
		t.Errorf("Returned invalid clusters:\n   expected: %v\n   got:      %v\n", expected, clusters)
	}
}

// res <- Ckmeans.1d.dp(c(-2.5, -2.5, -2.5, -2.5), 1, c(1.2, 1.1, 0.9, 0.8), method="linear")
func TestCKMeansWeightedK1(t *testing.T) {
	x := []float64{-2.5, -2.5, -2.5, -2.5}
	w := []float64{1.2, 1.1, 0.9, 0.8}
	expected := []int{1, 1, 1, 1}

	k, clusters, err := CKMeans(x, w, 1)
	if err != nil {
		t.Fatalf("Unexpected error (%v)", err)
	}

	if k != 1 {
		t.Errorf("Expected K=%v, got: %v\n", 1, k)
	}

	if !reflect.DeepEqual(clusters, expected) {
		t.Errorf("Returned invalid clusters:\n   expected: %v\n   got:      %v\n", expected, clusters)
	}
}

// res <- Ckmeans.1d.dp(c(1,2,3,4,5,6,7,8,9,10), 2, method="linear")
func TestCKMeansK2(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{1, 1, 1, 1, 1, 2, 2, 2, 2, 2}

	k, clusters, err := CKMeans(x, nil, 2)
	if err != nil {
		t.Fatalf("Unexpected error (%v)", err)
	}

	if k != 2 {
		t.Errorf("Expected K=%v, got: %v\n", 2, k)
	}

	if !reflect.DeepEqual(clusters, expected) {
		t.Errorf("Returned invalid clusters:\n   expected: %v\n   got:      %v\n", expected, clusters)
	}
}

// // res <- Ckmeans.1d.dp(c(-2.5, -2.5, -2.5, -2.5), 1, method="linear")
// func TestCKMeansK1(t *testing.T) {
// 	x := []float64{-2.5, -2.5, -2.5, -2.5}
//
// 	expected := []Cluster{
// 		Cluster{
// 			Cluster:  []int{1, 1, 1, 1},
// 			Centers:  []float64{-2.5},
// 			Withinss: []float64{0.0},
// 			Size:     []float64{4.0},
// 		},
// 	}
//
// 	clusters, err := CKMeans(x, nil)
// 	if err != nil {
// 		t.Fatalf("Unexpected error (%v)", err)
// 	}
//
// 	if !reflect.DeepEqual(clusters, expected) {
// 		t.Errorf("Returned invalid clusters:\n   expected: %v\n   got:      %v\n", expected, clusters)
// 	}
// }
//
// // res <- Ckmeans.1d.dp(c(-2.5, -2.5, -2.5, -2.5), 1, c(1.2, 1.1, 0.9, 0.8), method="linear")
// func TestCKMeansWeightedK1(t *testing.T) {
// 	x := []float64{-2.5, -2.5, -2.5, -2.5}
// 	w := []float64{1.2, 1.1, 0.9, 0.8}
//
// 	expected := []Cluster{
// 		Cluster{
// 			Cluster:  []int{1, 1, 1, 1},
// 			Centers:  []float64{-2.5},
// 			Withinss: []float64{0.0},
// 			Size:     []float64{4.8},
// 		},
// 	}
//
// 	clusters, err := CKMeans(x, w)
// 	if err != nil {
// 		t.Fatalf("Unexpected error (%v)", err)
// 	}
//
// 	if !reflect.DeepEqual(clusters, expected) {
// 		t.Errorf("Returned invalid clusters:\n   expected: %v\n   got:      %v\n", expected, clusters)
// 	}
// }
