package math

import (
	"fmt"
	"testing"
)

func TestAbs(t *testing.T) {
	tests := []struct {
		n      int64
		expect int64
	}{
		{0, 0},
		{5, 5},
		{-5, 5},
	}
	for _, test := range tests {
		if Abs(test.n) != test.expect {
			t.Errorf("Abs(%d) != %d", test.n, test.expect)
		}
	}
}

func TestSign(t *testing.T) {
	tests := []struct {
		n      float64
		expect float64
	}{
		{0, 0},
		{5, 1},
		{-5, -1},
	}
	for _, test := range tests {
		if Sign(test.n) != test.expect {
			t.Errorf("Sign(%d) != %d", test.n, test.expect)
		}
	}
}

func TestPerfectNumber(t *testing.T) {
	tests := []struct {
		int
		bool
	}{
		{6, true},
		{28, true},
		{496, true},
		{8128, true},
		{1735, false},
	}

	for _, test := range tests {
		if PerfectNumber(test.int) != test.bool {
			t.Errorf("PerfectNumber(%d) != %t", test.int, test.bool)
		}
	}
}

func TestAmicableNumbers(t *testing.T) {
	tests := []struct {
		a int
		b int
	}{
		{220, 284},
		{1184, 1210},
		{2620, 2924},
		{5020, 5564},
		{6232, 6368},
		{0, 0},
		{1, 0},
		{222, 0},
	}
	for _, test := range tests {
		if test.b != AmicableNumbers(test.a) {
			t.Errorf("AmicableNumbers(%d) != %d", test.a, test.b)
		}
	}
}

func TestPrimeFactors(t *testing.T) {
	if fmt.Sprintf("%v", PrimeFactors(23)) != `[23]` {
		t.Error(23)
	}
	if fmt.Sprintf("%v", PrimeFactors(12)) != `[2 2 3]` {
		t.Error(12)
	}
	if fmt.Sprintf("%v", PrimeFactors(360)) != `[2 2 2 3 3 5]` {
		t.Error(360)
	}
	if fmt.Sprintf("%v", PrimeFactors(97)) != `[97]` {
		t.Error(97)
	}
}
