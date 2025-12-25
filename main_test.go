package main

import "testing"

func TestMaxInt(t *testing.T) {
	a, b := 2, 7

	res := MaxInt(a, b)

	if res != b {
		t.Errorf("MaxInt failed: %d %d", res, b)
	}
}

func TestMain(m *testing.M) {
	main()
}
