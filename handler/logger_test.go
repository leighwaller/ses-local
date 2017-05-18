package handler

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
)

func TestRequestLoggingAdapter(t *testing.T) {
	var pass bool
	result := RequestLoggingAdapter(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pass = true
	}), "test")

	assert.NotNil(t, result)
	result.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("POST", "/", nil))
	assert.True(t, pass)
}
