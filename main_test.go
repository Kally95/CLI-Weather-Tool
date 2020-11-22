package main

import (
	"testing"
)

func TestKtoC(t *testing.T) {
	kelvin := 284.15
	got := KtoC(kelvin)
	want := "11.00"
	if got != want {
		t.Errorf("KtoC(%v) = %v; want = %v", kelvin, got, want)
	}

}
