package ckmeans

import (
	"testing"
)

func TestSelectLevelsBIC(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	J := [][]int{{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 5}}
	kmin := 2
	kmax := 2

	kopt := selectLevelsBIC(x, J, kmin, kmax)

	if kopt != 2 {
		t.Errorf("Incorrect kopt - expected:%v, got:%v", 2, kopt)
	}
}
