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

// result <- Ckmeans.1d.dp(c(-1,2,-1,2,4,5,6,-1,2,-1),3,method="linear")
func TestCKMeansWithGivenK(t *testing.T) {
	x := []float64{-1, 2, -1, 2, 4, 5, 6, -1, 2, -1}
	expected := []int{1, 2, 1, 2, 3, 3, 3, 1, 2, 1}

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

// result <- Ckmeans.1d.dp(c(-1,2,4,5,6),3,method="linear")
func TestCKMeansWithWeightedInput(t *testing.T) {
	{
		x := []float64{-1, 2, 4, 5, 6}
		w := []float64{4, 3, 1, 1, 1}
		expected := []int{1, 2, 3, 3, 3}

		k, clusters, err := ck.CKMeans(x, w, 3, 3)
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

	{
		x := []float64{-0.9, 1, 1.1, 1.9, 2, 2.1}
		w := []float64{3, 1, 2, 2, 1, 1}
		expected := []int{1, 2, 2, 3, 3, 3}

		k, clusters, err := ck.CKMeans(x, w, 1, 6)
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
}

// result <- Ckmeans.1d.dp(c(0.9,1,1.1,1.9,2,2.1),1,6,method="linear")
// result <- Ckmeans.1d.dp(c(2.1,2,1.9,1.1,1,0.9),1,6,method="linear")
// result <- Ckmeans.1d.dp(c(2.1,2,1.9,1.1,1,0.9),1,10,method="linear")
func TestCKMeansEstimateKExampleSet1(t *testing.T) {
	{
		x := []float64{0.9, 1, 1.1, 1.9, 2, 2.1}
		expected := []int{1, 1, 1, 2, 2, 2}

		k, clusters, err := ck.CKMeans(x, nil, 1, 6)
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

	{
		x := []float64{2.1, 2, 1.9, 1.1, 1, 0.9}
		expected := []int{2, 2, 2, 1, 1, 1}

		k, clusters, err := ck.CKMeans(x, nil, 1, 6)
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

	{
		x := []float64{2.1, 2, 1.9, 1.1, 1, 0.9}
		expected := []int{2, 2, 2, 1, 1, 1}

		k, clusters, err := ck.CKMeans(x, nil, 1, 10)
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
}

// result <- Ckmeans.1d.dp(c(-0.8390715,-0.9111303,-0.1455000,0.7539023,0.9601703,0.2836622,
//                           -0.6536436,-0.9899925,-0.4161468,0.5403023,1.0000000,0.5403023,
//                           -0.4161468,-0.9899925,-0.6536436,0.2836622,0.9601703,0.7539023,
//                           -0.1455000,-0.9111303,-0.8390715),1,21,method="linear")
func TestCKMeansEstimateKExampleSet2(t *testing.T) {
	x := []float64{
		-0.8390715, -0.9111303, -0.1455000, 0.7539023, 0.9601703, 0.2836622,
		-0.6536436, -0.9899925, -0.4161468, 0.5403023, 1.0000000, 0.5403023,
		-0.4161468, -0.9899925, -0.6536436, 0.2836622, 0.9601703, 0.7539023,
		-0.1455000, -0.9111303, -0.8390715,
	}
	expected := []int{1, 1, 1, 2, 2, 2, 1, 1, 1, 2, 2, 2, 1, 1, 1, 2, 2, 2, 1, 1, 1}

	k, clusters, err := ck.CKMeans(x, nil, 1, 21)
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

// result <- Ckmeans.1d.dp(dgamma(seq(1,10, by=0.5), shape=2, rate=1),1,19,method="linear")
func TestCKMeansEstimateKExampleSet3(t *testing.T) {
	x := []float64{
		0.3678794412, 0.3346952402, 0.2706705665, 0.2052124966, 0.1493612051,
		0.1056908420, 0.0732625556, 0.0499904844, 0.0336897350, 0.0224772429,
		0.0148725131, 0.0097723548, 0.0063831738, 0.0041481328, 0.0026837010,
		0.0017294811, 0.0011106882, 0.0007110924, 0.0004539993,
	}
	expected := []int{3, 3, 3, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	k, clusters, err := ck.CKMeans(x, nil, 1, 19)
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

// result <- Ckmeans.1d.dp(dgamma(seq(1,10, by=0.5), shape=2, rate=1),1,19,method="linear")
func TestCKMeansTaps(t *testing.T) {
	x := []float64{
		4.570271991, 5.063594027, 5.603539973, 6.102690998, 6.642708943, 7.141796968, 7.710649857, 8.192470916,
		4.506176116, 5.045971061, 5.591722996, 6.114172975, 6.619153989, 7.13578898, 7.693071891, 8.203885893,
		4.52956007, 5.057670039, 5.591721996, 6.13742393, 6.630941966, 7.1766839, 7.69897488, 8.227207848,
		4.52956007, 5.069284016, 5.603428973, 6.102591998, 6.613455, 7.147644957, 7.69912088, 8.215609871,
		4.517865093, 5.022782107, 5.580101018, 6.096715009, 6.654118921, 7.1763719, 7.681405914, 8.215537871,
		5.133092891, 5.545395086, 6.067721066, 6.578564068, 7.130096991, 7.652464971, 8.13427303,
		4.494581138, 5.040234073, 5.562732052, 6.079333043, 6.624973977, 7.141650968, 7.664070948, 8.198270905,
		4.52940807, 5.040295073, 5.556940064, 6.131584941, 6.654145921, 7.193876866, 7.722112835, 8.244539814,
		4.523631082, 5.046071061, 5.586102007, 6.09099502, 6.596029034, 7.130224991, 7.652501971, 8.180805939,
		4.517979093, 5.046165061, 5.551068075, 6.073547054, 6.607636011, 7.165018923, 7.687334903, 8.238953825,
		4.517911093, 5.069403016, 5.586174007, 6.108568986, 6.578649068, 7.147523957, 7.681606914, 8.26211078,
	}
	expected := []int{
		1, 2, 3, 4, 5, 6, 7, 8,
		1, 2, 3, 4, 5, 6, 7, 8,
		1, 2, 3, 4, 5, 6, 7, 8,
		1, 2, 3, 4, 5, 6, 7, 8,
		1, 2, 3, 4, 5, 6, 7, 8,
		2, 3, 4, 5, 6, 7, 8,
		1, 2, 3, 4, 5, 6, 7, 8,
		1, 2, 3, 4, 5, 6, 7, 8,
		1, 2, 3, 4, 5, 6, 7, 8,
		1, 2, 3, 4, 5, 6, 7, 8,
		1, 2, 3, 4, 5, 6, 7, 8,
	}

	k, clusters, err := ck.CKMeans(x, nil, 1, 87)
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
