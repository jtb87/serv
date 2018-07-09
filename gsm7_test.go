package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGSM(t *testing.T) {
	req, err := http.NewRequest("GET", "/exp/utftogsm", nil)
	if err != nil {
		t.Fatal(err)
	}
	a := App{}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(a.utftogsm)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
