package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestIndexHandler(t *testing.T) {

	tests := []struct {
		name                string
		method              string
		uri                 string
		expectedContentType string
	}{
		{
			name:                "Get Main Page",
			method:              "GET",
			uri:                 "/js/",
			expectedContentType: "text/html; charset=utf-8",
		},
		{
			name:   "Get JS",
			method: "GET",
			uri:    "/",
		},
	}
	r := newRouter()
	mockServer := httptest.NewServer(r)
	defer mockServer.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, reqErr := http.NewRequest(tt.method, mockServer.URL+tt.uri, nil)
			response, resErr := http.DefaultClient.Do(req)
			if reqErr != nil || resErr != nil {
				t.Errorf("Unexpected Error %s", reqErr)
			}

			if response.StatusCode != http.StatusOK {
				t.Errorf("Expected status 200, got %d", response.StatusCode)
			}

			fmt.Println(response.Header.Get("Content-Type"))

		})
	}
}
