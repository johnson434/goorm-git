package util_test

import (
	"testing"

	"github.com/johnson434/goorm-git/util"
)

func TestCalculator(t *testing.T) {
	if util.Add(1, 5) != 6 {
		t.Error("Wrong result")
	}
}

func TestCalculator2(t *testing.T) {
	if util.Add(1, 9) != 10 {
		t.Error("Wrong result")
	}
}
