package ckmeans

import (
	"bytes"
	"fmt"
	"math"
	"strings"
	"testing"
)

type test = struct {
	ID       string
	data     []float64
	weights  []float64
	k        int
	index    []int
	centers  []float64
	withinss []float64
}

var tests = []test{
	{
		ID:       "given K",
		data:     []float64{-1, 2, -1, 2, 4, 5, 6, -1, 2, -1},
		weights:  nil,
		k:        3,
		index:    []int{0, 1, 0, 1, 2, 2, 2, 0, 1, 0},
		centers:  []float64{-1, 2, 5},
		withinss: []float64{0, 0, 2}},
	{
		ID:       "N <= K",
		data:     []float64{3, 2, -5.4, 0.1},
		weights:  nil,
		k:        4,
		index:    []int{3, 2, 0, 1},
		centers:  []float64{-5.4, 0.1, 2, 3},
		withinss: []float64{0, 0, 0, 0}},
	{
		ID:       "single unique value,unweighted",
		data:     []float64{-2.5, -2.5, -2.5, -2.5},
		weights:  nil,
		k:        1,
		index:    []int{0, 0, 0, 0},
		centers:  []float64{-2.5},
		withinss: nil},
	{
		ID:       "K = 2",
		data:     []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		weights:  nil,
		k:        2,
		index:    []int{0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		centers:  []float64{3, 8},
		withinss: []float64{10, 10}},
	{
		ID:       "single unique value,weighted",
		data:     []float64{-2.5, -2.5, -2.5, -2.5},
		weights:  []float64{1.2, 1.1, 0.9, 0.8},
		k:        1,
		index:    []int{0, 0, 0, 0},
		centers:  []float64{-2.5},
		withinss: []float64{0}},
	{
		ID:       "N=10,K=3",
		data:     []float64{3, 3, 3, 3, 1, 1, 1, 2, 2, 2},
		weights:  nil,
		k:        3,
		index:    []int{2, 2, 2, 2, 0, 0, 0, 1, 1, 1},
		centers:  []float64{1, 2, 3},
		withinss: []float64{0, 0, 0}},
	{
		ID:       "N=14,K=8",
		data:     []float64{-3, 2.2, -6, 7, 9, 11, -6.3, 75, 82.6, 32.3, -9.5, 62.5, 7, 95.2},
		weights:  nil,
		k:        8,
		index:    []int{1, 1, 0, 2, 2, 2, 0, 5, 6, 3, 0, 4, 2, 7},
		centers:  []float64{-7.266666667, -0.4, 8.5, 32.3, 62.5, 75.0, 82.6, 95.2},
		withinss: []float64{7.526666667, 13.52, 11.0, 0.0, 0.0, 0.0, 0.0, 0.0}},
	{
		ID:       "weighted input (1)",
		data:     []float64{-1, 2, 4, 5, 6},
		weights:  []float64{4, 3, 1, 1, 1},
		k:        3,
		index:    []int{0, 1, 2, 2, 2},
		centers:  []float64{-1, 2, 5},
		withinss: []float64{0, 0, 2}},
	{
		ID:       "weighted input (2)",
		data:     []float64{-0.9, 1, 1.1, 1.9, 2, 2.1},
		weights:  []float64{3, 1, 2, 2, 1, 1},
		k:        3,
		index:    []int{0, 1, 1, 2, 2, 2},
		centers:  []float64{-0.9, (1 + 2.2) / 3, (1.9*2 + 2 + 2.1) / 4},
		withinss: []float64{0, 0.00666667, 0.0275}},
	{
		ID:       "estimate K, example set 1 (1)",
		data:     []float64{0.9, 1, 1.1, 1.9, 2, 2.1},
		weights:  nil,
		k:        2,
		index:    []int{0, 0, 0, 1, 1, 1},
		centers:  []float64{1, 2},
		withinss: []float64{0.02, 0.02}},
	{
		ID:       "estimate K, example set 1 (2)",
		data:     []float64{2.1, 2, 1.9, 1.1, 1, 0.9},
		weights:  nil,
		k:        2,
		index:    []int{1, 1, 1, 0, 0, 0},
		centers:  []float64{1, 2},
		withinss: []float64{0.02, 0.02}},
	{
		ID:       "estimate K, example set 2",
		data:     []float64{3.5, 3.6, 3.7, 3.1, 1.1, 0.9, 0.8, 2.2, 1.9, 2.1},
		weights:  nil,
		k:        3,
		index:    []int{2, 2, 2, 2, 0, 0, 0, 1, 1, 1},
		centers:  []float64{0.933333333333, 2.066666666667, 3.475},
		withinss: []float64{0.0466666666667, 0.0466666666667, 0.2075}},
	{
		ID: "estimate K, example set 3",
		data: []float64{
			-0.8390715, -0.9111303, -0.1455000, 0.7539023, 0.9601703, 0.2836622,
			-0.6536436, -0.9899925, -0.4161468, 0.5403023, 1.0000000, 0.5403023,
			-0.4161468, -0.9899925, -0.6536436, 0.2836622, 0.9601703, 0.7539023,
			-0.1455000, -0.9111303, -0.8390715,
		},
		weights:  nil,
		k:        2,
		index:    []int{0, 0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		centers:  []float64{-0.6592474631, 0.6751193405},
		withinss: []float64{1.0564794374, 0.6232976959}},
	{
		ID: "estimate K, example set 4",
		data: []float64{
			0.3678794412, 0.3346952402, 0.2706705665, 0.2052124966, 0.1493612051,
			0.1056908420, 0.0732625556, 0.0499904844, 0.0336897350, 0.0224772429,
			0.0148725131, 0.0097723548, 0.0063831738, 0.0041481328, 0.0026837010,
			0.0017294811, 0.0011106882, 0.0007110924, 0.0004539993,
		},
		weights:  nil,
		k:        3,
		index:    []int{2, 2, 2, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		centers:  []float64{0.01702193495, 0.15342151455, 0.32441508262},
		withinss: []float64{0.006126754998, 0.004977009034, 0.004883305120}},
	{
		ID: "taps",
		data: []float64{
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
		},
		weights: nil,
		k:       8,
		index: []int{
			0, 1, 2, 3, 4, 5, 6, 7,
			0, 1, 2, 3, 4, 5, 6, 7,
			0, 1, 2, 3, 4, 5, 6, 7,
			0, 1, 2, 3, 4, 5, 6, 7,
			0, 1, 2, 3, 4, 5, 6, 7,
			1, 2, 3, 4, 5, 6, 7,
			0, 1, 2, 3, 4, 5, 6, 7,
			0, 1, 2, 3, 4, 5, 6, 7,
			0, 1, 2, 3, 4, 5, 6, 7,
			0, 1, 2, 3, 4, 5, 6, 7,
			0, 1, 2, 3, 4, 5, 6, 7,
		},
		centers:  []float64{4.52369438, 5.0576875, 5.5780842, 6.1004859, 6.6182161, 7.1533345, 7.6857560, 8.2103333},
		withinss: nil},
}

func TestCenters(t *testing.T) {
	for _, p := range tests {
		means := centers(p.data, p.weights, p.k, p.index)

		if p.centers != nil {
			if len(means) != len(p.centers) {
				t.Errorf("Returned invalid centers:\n   expected: %v\n   got:      %v\n", format(p.centers), format(means))
			} else {
				for j := range p.centers {
					if math.IsNaN(means[j]) || math.Abs(means[j]-p.centers[j]) > 0.0000001 {
						t.Errorf("Returned invalid centers:\n   expected: %v\n   got:      %v\n", format(p.centers), format(means))
						break
					}
				}
			}
		}
	}
}

func TestWithinSS(t *testing.T) {
	for _, p := range tests {
		variance := withinss(p.data, p.weights, p.k, p.index)

		if p.withinss != nil {
			if len(variance) != len(p.withinss) {
				t.Errorf("[%s] Returned invalid withinss:\n   expected: %v\n   got:      %v\n", p.ID, format(p.withinss), format(variance))
			} else {
				for j := range p.withinss {
					if math.IsNaN(variance[j]) || math.Abs(variance[j]-p.withinss[j]) > 0.0000001 {
						t.Errorf("[%s] Returned invalid withinss:\n   expected: %v\n   got:      %v\n", p.ID, format(p.withinss), format(variance))
						break
					}
				}
			}
		}
	}
}

func format(array []float64) string {
	var b bytes.Buffer
	for _, v := range array {
		fmt.Fprintf(&b, "%+0.5f ", v)
	}

	return strings.TrimSpace(string(b.Bytes()))
}
