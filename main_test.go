package main

import (
	"fmt"
	"testing"
)

func TestGetAPI(t *testing.T) {
	k := GetAPIkey("APIKEY")
	l := "London"
	resp, err := GetAPI(l, k)
	if err != nil {
		t.Fatalf("GetAPI(%v, %v) = %v", l, k, resp)
	}
	fmt.Println(resp.Status)

	// Currently returning a 401 (not authorised) *must fix*
}

func TestKtoC(t *testing.T) {
	kelvin := 284.15
	got := KtoC(kelvin)
	want := "11.00"
	if got != want {
		t.Errorf("KtoC(%v) = %v; want = %v", kelvin, got, want)
	}

}
