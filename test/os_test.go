package test

import (
	"github.com/4thel00z/installd/pkg/v1/libinstall"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTemplatesFromPath(t *testing.T) {
	templates, err := libinstall.TemplatesFromPath("resources")
	assert.Nil(t, err, "some shite crashed yo %e", err)
	assert.True(t, len(templates) > 0, "could not loud the templates %+v", templates)
}
