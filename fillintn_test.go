package gofillslicewithrand_test

import (
	"testing"

	fillslicewithrand "github.com/yinyin/go-fillslicewithrand"
)

func testFillIntN_BxNx(t *testing.T, bufSize int, nValue int) {
	b := make([]int, bufSize)
	fillslicewithrand.FillIntN(b, nValue)
	nonZero := false
	for _, v := range b {
		if v >= nValue {
			t.Errorf("unexpect value: %v", v)
		} else if v != 0 {
			nonZero = true
		}
	}
	if !nonZero {
		t.Error("all zero")
	}
	t.Log(b)
}

func TestFillIntN_B16N128(t *testing.T) {
	testFillIntN_BxNx(t, 16, 128)
}

func TestFillIntN_B16N13(t *testing.T) {
	testFillIntN_BxNx(t, 16, 13)
}

func TestFillIntN_B32N1000(t *testing.T) {
	testFillIntN_BxNx(t, 32, 1000)
}
