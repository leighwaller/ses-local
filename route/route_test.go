package route

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIndexRoute(t *testing.T) {
	r := CreateRouter()
	index := r.Get("index")

	assert.NotNil(t, index)
	template, _ := index.GetPathTemplate()
	assert.Equal(t, "/", template)
	assert.Equal(t, "index", index.GetName())
	assert.NotNil(t, index.GetHandler())
}
