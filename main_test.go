package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestGetAPI(t *testing.T) {
	k := GetAPIkey("APIKEY")
	v := url.Values{}
	api := "http://api.openweathermap.org/data/2.5/weather?q=London&appendid=" + k
	req, err := http.NewRequest(http.MethodGet, api, strings.NewReader(v.Encode()))
	if err != nil {
		t.Fatalf("GetAPI failed to return from call")
	}
	fmt.Println(req)
}

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
