package EWL2

import (
	"reflect"
	"testing"
)

func TestEWL2DP(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	w := []float64{}
	s := [][]float64{[]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}
	j := [][]int{{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}

	FillDPMatrix(x, w, s, j)

	S := [][]float64{[]float64{0, 0.5, 2, 5, 10, 17.5, 28, 42, 60, 82.5}, []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 20}}
	J := [][]int{{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 5}}

	if !reflect.DeepEqual(s, S) {
		t.Errorf("Invalid S\n   expected:%v\n   got:     %v", S, s)
	}

	if !reflect.DeepEqual(j, J) {
		t.Errorf("Invalid J\n   expected:%v\n   got:     %v", J, j)
	}
}
