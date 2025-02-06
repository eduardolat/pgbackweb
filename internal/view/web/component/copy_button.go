package component

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	nodx "github.com/nodxdev/nodxgo"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

// CopyButtonSm is a small copy button.
func CopyButtonSm(textToCopy string) nodx.Node {
	return copyButton(copyButtonProps{
		TextToCopy: textToCopy,
		Size:       SizeSm,
	})
}

// CopyButtonMd is a medium copy button.
func CopyButtonMd(textToCopy string) nodx.Node {
	return copyButton(copyButtonProps{
		TextToCopy: textToCopy,
	})
}

// CopyButtonLg is a large copy button.
func CopyButtonLg(textToCopy string) nodx.Node {
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
func copyButton(props copyButtonProps) nodx.Node {
	id := uuid.NewString()
	id = strings.ReplaceAll(id, "-", "")

	sc := copyButtonScript(id, props.TextToCopy)

	return nodx.Div(
		nodx.Class("inline-block tooltip tooltip-right"),
		nodx.Data("tip", "Copy to clipboard"),
		sc.script,
		nodx.Button(
			nodx.ClassMap{
				"btn btn-neutral btn-square btn-ghost": true,
				"btn-sm":                               props.Size == SizeSm,
				"btn-lg":                               props.Size == SizeLg,
			},
			nodx.Id(id),
			nodx.TitleAttr("Copy to clipboard"),
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
	script    nodx.Node
	copyEvent nodx.Node
} {
	escapedTextToCopy := strings.ReplaceAll(textToCopy, "`", "\\`")

	rawScript := fmt.Sprintf(
		"<script>function copy%s(){ copyToClipboard(`%s`); }</script>",
		id,
		escapedTextToCopy,
	)

	script := nodx.Raw(rawScript)
	copyEvent := nodx.Attr("onclick", fmt.Sprintf("copy%s()", id))

	return struct {
		script    nodx.Node
		copyEvent nodx.Node
	}{
		script:    script,
		copyEvent: copyEvent,
	}
}
