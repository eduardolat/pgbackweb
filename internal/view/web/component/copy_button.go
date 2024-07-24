package component

import (
	"fmt"
	"strings"

	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/google/uuid"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

// CopyButtonSm is a small copy button.
func CopyButtonSm(textToCopy string) gomponents.Node {
	return copyButton(copyButtonProps{
		TextToCopy: textToCopy,
		Size:       SizeSm,
	})
}

// CopyButtonMd is a medium copy button.
func CopyButtonMd(textToCopy string) gomponents.Node {
	return copyButton(copyButtonProps{
		TextToCopy: textToCopy,
	})
}

// CopyButtonLg is a large copy button.
func CopyButtonLg(textToCopy string) gomponents.Node {
	return copyButton(copyButtonProps{
		TextToCopy: textToCopy,
		Size:       SizeLg,
	})
}

// copyButtonProps is the properties for the CopyButton component.
type copyButtonProps struct {
	// TextToCopy is the text that should be copied to the
	// clipboard when the button is clicked.
	TextToCopy string
	// Size is the size of the button. Can be "sm", "md" (default) or "lg".
	Size size
}

// copyButton is a button that copies text to the clipboard when clicked.
func copyButton(props copyButtonProps) gomponents.Node {
	id := uuid.NewString()
	id = strings.ReplaceAll(id, "-", "")

	sc := copyButtonScript(id, props.TextToCopy)

	return html.Div(
		html.Class("inline-block tooltip tooltip-right"),
		html.Data("tip", "Copy to clipboard"),
		sc.script,
		html.Button(
			components.Classes{
				"btn btn-neutral btn-square btn-ghost": true,
				"btn-sm":                               props.Size == SizeSm,
				"btn-lg":                               props.Size == SizeLg,
			},
			html.ID(id),
			html.Title("Copy to clipboard"),
			sc.copyEvent,
			lucide.Copy(),
		),
	)
}

// copyButtonScript returns a script that copies the given text to the clipboard
// and an event that calls the script when clicked.
func copyButtonScript(
	id string,
	textToCopy string,
) struct {
	script    gomponents.Node
	copyEvent gomponents.Node
} {
	escapedTextToCopy := strings.ReplaceAll(textToCopy, "`", "\\`")

	rawScript := fmt.Sprintf(
		"<script>function copy%s(){ copyToClipboard(`%s`); }</script>",
		id,
		escapedTextToCopy,
	)

	script := gomponents.Raw(rawScript)
	copyEvent := gomponents.Attr("onclick", fmt.Sprintf("copy%s()", id))

	return struct {
		script    gomponents.Node
		copyEvent gomponents.Node
	}{
		script:    script,
		copyEvent: copyEvent,
	}
}
