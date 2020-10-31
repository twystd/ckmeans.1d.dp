package ckmeans

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

var ck = CKMEANS{
	Method:    Linear,
	EstimateK: BIC,
	Criterion: L2,
}

// 'R':
// library(Ckmeans.1d.dp)
//
// result <- Ckmeans.1d.dp(x, k, w, method="linear")
//
// print(result$cluster)
// print(result$centers)
// print(result$size)
// print(result)

// result <- Ckmeans.1d.dp(c(-2.5,-2.5,-2.5,-2.5),1,method="linear")
func TestCKMeansWithSingleUniqueValue(t *testing.T) {
	x := []float64{-2.5, -2.5, -2.5, -2.5}
	expected := []int{1, 1, 1, 1}

	k, clusters, err := ck.CKMeans(x, nil, 1, 4)
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

// result <- Ckmeans.1d.dp(rep(1,100),1,method="linear")
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

	k, clusters, err := ck.CKMeans(x, nil, 1, 1)
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

// result <- Ckmeans.1d.dp(c(-2.5,-2.5,-2.5,-2.5),1,c(1.2,1.1,0.9,0.8), method="linear")
func TestCKMeansWeightedK1(t *testing.T) {
	x := []float64{-2.5, -2.5, -2.5, -2.5}
	w := []float64{1.2, 1.1, 0.9, 0.8}
	expected := []int{1, 1, 1, 1}

	k, clusters, err := ck.CKMeans(x, w, 1, 1)
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

// result <- Ckmeans.1d.dp(c(1,2,3,4,5,6,7,8,9,10),2,method="linear")
func TestCKMeansK2(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{1, 1, 1, 1, 1, 2, 2, 2, 2, 2}

	k, clusters, err := ck.CKMeans(x, nil, 2, 2)
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

// result <- Ckmeans.1d.dp(c(3,3,3,3,1,1,1,2,2,2),3,method="linear")
func TestCKMeansN10K3(t *testing.T) {
	x := []float64{3, 3, 3, 3, 1, 1, 1, 2, 2, 2}
	expected := []int{3, 3, 3, 3, 1, 1, 1, 2, 2, 2}

	k, clusters, err := ck.CKMeans(x, nil, 3, 3)
	if err != nil {
		t.Fatalf("Unexpected error (%v)", err)
	}

	if k != 3 {
		t.Errorf("Expected K=%v, got: %v\n", 3, k)
	}

	if !reflect.DeepEqual(clusters, expected) {
		t.Errorf("Returned invalid clusters:\n   expected: %v\n   got:      %v\n", expected, clusters)
	}
}

// result <- Ckmeans.1d.dp(c(-3,2.2,-6,7,9,11,-6.3,75,82.6,32.3,-9.5,62.5,7,95.2),3,method="linear")
func TestCKMeansN14K8(t *testing.T) {
	x := []float64{-3, 2.2, -6, 7, 9, 11, -6.3, 75, 82.6, 32.3, -9.5, 62.5, 7, 95.2}
	expected := []int{2, 2, 1, 3, 3, 3, 1, 6, 7, 4, 1, 5, 3, 8}

	k, clusters, err := ck.CKMeans(x, nil, 8, 8)
	if err != nil {
		t.Fatalf("Unexpected error (%v)", err)
	}

	if k != 8 {
		t.Errorf("Expected K=%v, got: %v\n", 8, k)
	}

	if !reflect.DeepEqual(clusters, expected) {
		t.Errorf("Returned invalid clusters:\n   expected: %v\n   got:      %v\n", expected, clusters)
	}
}

// result <- Ckmeans.1d.dp(c(3,2,-5.4,0.1),4,method="linear")
func TestCKNlteK(t *testing.T) {
	x := []float64{3, 2, -5.4, 0.1}
	expected := []int{4, 3, 1, 2}

	k, clusters, err := ck.CKMeans(x, nil, 4, 4)
	if err != nil {
		t.Fatalf("Unexpected error (%v)", err)
	}

	if k != 4 {
		t.Errorf("Expected K=%v, got: %v\n", 4, k)
	}

	if !reflect.DeepEqual(clusters, expected) {
		t.Errorf("Returned invalid clusters:\n   expected: %v\n   got:      %v\n", expected, clusters)
	}
}
