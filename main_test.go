package main

import (
	"testing"
)

// func TestGetAPI(t *testing.T) {
// 	k := GetAPIkey("APIKEY")
// 	l := "London"
// 	resp, err := GetAPI(l, k)
// 	if err != nil {
// 		t.Fatalf("GetAPI(%v, %v) = %v", l, k, resp)
// 	}
// 	fmt.Println(resp.Status)
// 	// Currently returning a 401 (not authorised) *must fix*
// }

func TestKtoC(t *testing.T) {
	testCases := []struct {
		arg  float64
		want string
	}{
		{284.15, "11.00"},
		{295.20, "22.05"},
		{289.63, "16.48"},
	}
	for _, tc := range testCases {
		t.Logf("Testing: %v", tc.arg)
		got := KtoC(tc.arg)
		if got != tc.want {
			t.Errorf("KtoC(%v) = %v; want %v", tc.arg, got, tc.want)
		}
	}
}
