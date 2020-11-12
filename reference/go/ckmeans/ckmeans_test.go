package ckmeans

import (
	"math"
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
	expected := Clusters{
		K:       3,
		Index:   []int{1, 2, 1, 2, 3, 3, 3, 1, 2, 1},
		Centers: []float64{-1, 2, 5},
	}

	//	withinss := []float64{0, 0, 2}
	//	size := []float64{4, 3, 3}

	testCKMeans(x, nil, 3, 3, expected, t)
}

// result <- Ckmeans.1d.dp(c(3,2,-5.4,0.1),4,method="linear")
func TestCKNlteK(t *testing.T) {
	x := []float64{3, 2, -5.4, 0.1}
	expected := Clusters{
		K:       4,
		Index:   []int{4, 3, 1, 2},
		Centers: []float64{-5.4, 0.1, 2, 3},
		//{0, 0, 0, 0},{1, 1, 1, 1}};
	}

	testCKMeans(x, nil, 4, 4, expected, t)
}

// result <- Ckmeans.1d.dp(c(-2.5,-2.5,-2.5,-2.5),1,method="linear")
func TestCKMeansWithSingleUniqueValue(t *testing.T) {
	x := []float64{-2.5, -2.5, -2.5, -2.5}
	expected := Clusters{
		K:       1,
		Index:   []int{1, 1, 1, 1},
		Centers: []float64{-2.5},
		// {0}, {4}};

	}

	testCKMeans(x, nil, 1, 4, expected, t)
}

// result <- Ckmeans.1d.dp(c(1,2,3,4,5,6,7,8,9,10),2,method="linear")
func TestCKMeansK2(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := Clusters{
		K:       2,
		Index:   []int{1, 1, 1, 1, 1, 2, 2, 2, 2, 2},
		Centers: []float64{3, 8},
		//,{10,10},{5,5}};

	}

	testCKMeans(x, nil, 2, 2, expected, t)
}

// result <- Ckmeans.1d.dp(rep(1,100),1,method="linear")
func TestCKMeansK1(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	x := make([]float64, 100)
	expected := Clusters{
		K:     1,
		Index: make([]int, 100),
	}

	for i := range x {
		x[i] = r.Float64()
	}

	for i := range expected.Index {
		expected.Index[i] = 1
	}

	testCKMeans(x, nil, 1, 1, expected, t)
}

// result <- Ckmeans.1d.dp(c(-2.5,-2.5,-2.5,-2.5),1,c(1.2,1.1,0.9,0.8), method="linear")
func TestCKMeansWeightedK1(t *testing.T) {
	x := []float64{-2.5, -2.5, -2.5, -2.5}
	w := []float64{1.2, 1.1, 0.9, 0.8}
	expected := Clusters{
		K:       1,
		Index:   []int{1, 1, 1, 1},
		Centers: []float64{-2.5},
	}

	testCKMeans(x, w, 1, 1, expected, t)
}

// result <- Ckmeans.1d.dp(c(3,3,3,3,1,1,1,2,2,2),3,method="linear")
func TestCKMeansN10K3(t *testing.T) {
	x := []float64{3, 3, 3, 3, 1, 1, 1, 2, 2, 2}
	expected := Clusters{
		K:       3,
		Index:   []int{3, 3, 3, 3, 1, 1, 1, 2, 2, 2},
		Centers: []float64{1, 2, 3},
		//{0, 0, 0},{3, 3, 4}};

	}

	testCKMeans(x, nil, 3, 3, expected, t)
}

// result <- Ckmeans.1d.dp(c(-3,2.2,-6,7,9,11,-6.3,75,82.6,32.3,-9.5,62.5,7,95.2),3,method="linear")
func TestCKMeansN14K8(t *testing.T) {
	x := []float64{-3, 2.2, -6, 7, 9, 11, -6.3, 75, 82.6, 32.3, -9.5, 62.5, 7, 95.2}
	expected := Clusters{
		K:       8,
		Index:   []int{2, 2, 1, 3, 3, 3, 1, 6, 7, 4, 1, 5, 3, 8},
		Centers: []float64{-7.266666667, -0.4, 8.5, 32.3, 62.5, 75.0, 82.6, 95.2},
		//                        {7.526666667, 13.52, 11.0, 0.0, 0.0, 0.0, 0.0, 0.0},
		//                      {3, 2, 4, 1, 1, 1, 1, 1}};
	}

	testCKMeans(x, nil, 8, 8, expected, t)
}

// result <- Ckmeans.1d.dp(c(-1,2,4,5,6),3,method="linear")
// result <- Ckmeans.1d.dp(c(-0.9, 1, 1.1, 1.9, 2, 2.1),c(3, 1, 2, 2, 1, 1),3,method="linear")
func TestCKMeansWithWeightedInput(t *testing.T) {
	{
		x := []float64{-1, 2, 4, 5, 6}
		w := []float64{4, 3, 1, 1, 1}
		expected := Clusters{
			K:       3,
			Index:   []int{1, 2, 3, 3, 3},
			Centers: []float64{-1, 2, 5},
			//,{0,0,2},{4,3,3}};

		}

		testCKMeans(x, w, 3, 3, expected, t)
	}

	{
		x := []float64{-0.9, 1, 1.1, 1.9, 2, 2.1}
		w := []float64{3, 1, 2, 2, 1, 1}
		expected := Clusters{
			K:       3,
			Index:   []int{1, 2, 2, 3, 3, 3},
			Centers: []float64{-0.9, 1.0666666666666667, 1.975},
			// Centers: []float64{-0.9, (1 + 2.2) / 3, (1.9*2 + 2 + 2.1) / 4},
			//{0,1.00666667,0.0275},{3,3,4}};

		}

		testCKMeans(x, w, 1, 6, expected, t)
	}
}

// result <- Ckmeans.1d.dp(c(0.9,1,1.1,1.9,2,2.1),1,6,method="linear")
// result <- Ckmeans.1d.dp(c(2.1,2,1.9,1.1,1,0.9),1,6,method="linear")
// result <- Ckmeans.1d.dp(c(2.1,2,1.9,1.1,1,0.9),1,10,method="linear")
func TestCKMeansEstimateKExampleSet1(t *testing.T) {
	{
		x := []float64{0.9, 1, 1.1, 1.9, 2, 2.1}
		expected := Clusters{
			K:       2,
			Index:   []int{1, 1, 1, 2, 2, 2},
			Centers: []float64{1, 2},
			//{0.02,0.02},{3,3}};
		}

		testCKMeans(x, nil, 1, 6, expected, t)
	}

	{
		x := []float64{2.1, 2, 1.9, 1.1, 1, 0.9}
		expected := Clusters{
			K:       2,
			Index:   []int{2, 2, 2, 1, 1, 1},
			Centers: []float64{1, 2},
			//{0.02,0.02},{3,3}};
		}

		testCKMeans(x, nil, 1, 6, expected, t)
	}

	{
		x := []float64{2.1, 2, 1.9, 1.1, 1, 0.9}
		expected := Clusters{
			K:       2,
			Index:   []int{2, 2, 2, 1, 1, 1},
			Centers: []float64{1, 2},
			//,{0.02,0.02},{3,3}};
		}

		testCKMeans(x, nil, 1, 6, expected, t)
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

	expected := Clusters{
		K:       2,
		Index:   []int{1, 1, 1, 2, 2, 2, 1, 1, 1, 2, 2, 2, 1, 1, 1, 2, 2, 2, 1, 1, 1},
		Centers: []float64{-0.6592474631, 0.6751193405},
		//{1.0564793100, 0.6232976959},{12,9}};

	}

	testCKMeans(x, nil, 1, 21, expected, t)
}

// result <- Ckmeans.1d.dp(dgamma(seq(1,10, by=0.5), shape=2, rate=1),1,19,method="linear")
func TestCKMeansEstimateKExampleSet3(t *testing.T) {
	x := []float64{
		0.3678794412, 0.3346952402, 0.2706705665, 0.2052124966, 0.1493612051,
		0.1056908420, 0.0732625556, 0.0499904844, 0.0336897350, 0.0224772429,
		0.0148725131, 0.0097723548, 0.0063831738, 0.0041481328, 0.0026837010,
		0.0017294811, 0.0011106882, 0.0007110924, 0.0004539993,
	}

	expected := Clusters{
		K:       3,
		Index:   []int{3, 3, 3, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		Centers: []float64{0.01702193495, 0.15342151455, 0.32441508262},
		// {0.006126754998,0.004977009034,0.004883305120},{13,3,3}};
	}

	testCKMeans(x, nil, 1, 19, expected, t)
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

	expected := Clusters{
		K: 8,
		Index: []int{
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
		},
		Centers: []float64{4.52369438, 5.0576875, 5.5780842, 6.1004859, 6.6182161, 7.1533345, 7.6857560, 8.2103333},
	}

	testCKMeans(x, nil, 1, 87, expected, t)
}

func testCKMeans(x, w []float64, kmin, kmax int, expected Clusters, t *testing.T) {
	clusters, err := ck.CKMeans(x, w, kmin, kmax)
	if err != nil {
		t.Fatalf("Unexpected error (%v)", err)
	}

	if clusters.K != expected.K {
		t.Errorf("Incorrect K - expected:%v, got:%v", expected.K, clusters.K)
	}

	if !reflect.DeepEqual(clusters.Index, expected.Index) {
		t.Errorf("Returned invalid clusters:\n   expected: %v\n   got:      %v\n", expected.Index, clusters.Index)
	}

	if expected.Centers != nil {
		for i := range expected.Centers {
			if math.IsNaN(clusters.Centers[i]) || math.Abs(clusters.Centers[i]-expected.Centers[i]) > 0.0000001 {
				t.Errorf("Returned invalid centers:\n   expected: %v\n   got:      %v\n", expected.Centers, clusters.Centers)
				break
			}
		}
	}
}
