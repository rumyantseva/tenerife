package application

import (
	"github.com/sirupsen/logrus/hooks/test"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	logger, _ := test.NewNullLogger()

	req := httptest.NewRequest(
		"GET", "http://example.com/foo", nil,
		)
	w := httptest.NewRecorder()
	handler := HomeHandler(logger)
	handler(w, req)

	if http.StatusOK != w.Code {
		t.Errorf("have %+v, want %+v", w.Code, http.StatusOK)
	}
}