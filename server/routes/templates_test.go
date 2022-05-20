package routes

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplateRenderer(t *testing.T) {
	renderer := GetTemplateRenderer()
	t.Run("test Render with no parameter", func(t *testing.T) {
		output := bytes.Buffer{}
		err := renderer.Render(&output, "ping.gohtml", nil)
		assert.NoError(t, err)
		assert.Equal(t, "pong", output.String())
	})

	t.Run("test Render with parameter", func(t *testing.T) {
		output := bytes.Buffer{}
		err := renderer.Render(&output, "ping_with_ip.gohtml", "127.0.0.1")
		assert.NoError(t, err)
		assert.Equal(t, "pong from ip 127.0.0.1", output.String())
	})

}
