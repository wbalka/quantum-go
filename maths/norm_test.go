package maths

import (
	"fmt"
	"testing"
)

func TestNorm100(t *testing.T) {
	oldNorm := 0.0018900383817771435
	newNorm := GetNormalizationConstant(3, 1, 1)
	fmt.Printf("Normalization constant old: %0.17g new: %0.17g ", oldNorm, newNorm)
	if oldNorm != newNorm {
		t.Errorf("Bad normalization")
	}
}
