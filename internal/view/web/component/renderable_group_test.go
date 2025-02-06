package component

import (
	"bytes"
	"testing"

	nodx "github.com/nodxdev/nodxgo"
	"github.com/stretchr/testify/assert"
)

func TestRenderableGroupRenderer(t *testing.T) {
	t.Run("renders a group of string nodes without a parent element", func(t *testing.T) {
		gotRenderer := RenderableGroup([]nodx.Node{
			nodx.Text("foo"),
			nodx.Text("bar"),
		})

		got := bytes.Buffer{}
		err := gotRenderer.Render(&got)
		assert.NoError(t, err)

		expected := "foobar"

		assert.Equal(t, expected, got.String())
	})

	t.Run("renders a group of tag nodes without a parent element", func(t *testing.T) {
		gotRenderer := RenderableGroup([]nodx.Node{
			nodx.SpanEl(
				nodx.Text("foo"),
			),
			nodx.P(
				nodx.Text("bar"),
			),
		})

		got := bytes.Buffer{}
		err := gotRenderer.Render(&got)
		assert.NoError(t, err)

		expected := `<span>foo</span><p>bar</p>`

		assert.Equal(t, expected, got.String())
	})
}
