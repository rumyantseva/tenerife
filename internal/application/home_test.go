package application

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest(
		"GET", "http://example.com/foo", nil,
		)
	w := httptest.NewRecorder()
	HomeHandler(w, req)

	if http.StatusOK != w.Code {
		t.Errorf("have %+v, want %+v", w.Code, http.StatusOK)
	}
}