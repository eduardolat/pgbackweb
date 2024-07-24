package component

import (
	"bytes"
	"testing"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
	"github.com/stretchr/testify/assert"
)

func TestRenderableGroupRenderer(t *testing.T) {
	t.Run("renders a group of string nodes without a parent element", func(t *testing.T) {
		gotRenderer := RenderableGroup([]gomponents.Node{
			gomponents.Text("foo"),
			gomponents.Text("bar"),
		})

		got := bytes.Buffer{}
		err := gotRenderer.Render(&got)
		assert.NoError(t, err)

		expected := "foobar"

		assert.Equal(t, expected, got.String())
	})

	t.Run("renders a group of tag nodes without a parent element", func(t *testing.T) {
		gotRenderer := RenderableGroup([]gomponents.Node{
			html.Span(
				gomponents.Text("foo"),
			),
			html.P(
				gomponents.Text("bar"),
			),
		})

		got := bytes.Buffer{}
		err := gotRenderer.Render(&got)
		assert.NoError(t, err)

		expected := `<span>foo</span><p>bar</p>`

		assert.Equal(t, expected, got.String())
	})
}
