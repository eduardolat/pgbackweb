package htmx

import (
	"github.com/maragudk/gomponents"
)

// HxGet returns a gomponents node with the hx-get
// attribute set to the given path.
//
// https://htmx.org/attributes/hx-get/
func HxGet(path string) gomponents.Node {
	return gomponents.Attr("hx-get", path)
}

// HxPost returns a gomponents node with the hx-post
// attribute set to the given path.
//
// https://htmx.org/attributes/hx-post/
func HxPost(path string) gomponents.Node {
	return gomponents.Attr("hx-post", path)
}

// HxPut returns a gomponents node with the hx-put
// attribute set to the given path.
//
// https://htmx.org/attributes/hx-put/
func HxPut(path string) gomponents.Node {
	return gomponents.Attr("hx-put", path)
}

// HxPatch returns a gomponents node with the hx-patch
// attribute set to the given path.
//
// https://htmx.org/attributes/hx-patch/
func HxPatch(path string) gomponents.Node {
	return gomponents.Attr("hx-patch", path)
}

// HxDelete returns a gomponents node with the hx-delete
// attribute set to the given path.
//
// https://htmx.org/attributes/hx-delete/
func HxDelete(path string) gomponents.Node {
	return gomponents.Attr("hx-delete", path)
}

// HxTrigger returns a gomponents node with the hx-trigger
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-trigger/
func HxTrigger(value string) gomponents.Node {
	return gomponents.Attr("hx-trigger", value)
}

// HxTarget returns a gomponents node with the hx-target
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-target/
func HxTarget(value string) gomponents.Node {
	return gomponents.Attr("hx-target", value)
}

// HxSwap returns a gomponents node with the hx-swap
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-swap/
func HxSwap(value string) gomponents.Node {
	return gomponents.Attr("hx-swap", value)
}

// HxIndicator returns a gomponents node with the hx-indicator
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-indicator/
func HxIndicator(value string) gomponents.Node {
	return gomponents.Attr("hx-indicator", value)
}

// HxConfirm returns a gomponents node with the hx-confirm
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-confirm/
func HxConfirm(value string) gomponents.Node {
	return gomponents.Attr("hx-confirm", value)
}

// HxBoost returns a gomponents node with the hx-boost
// attribute set to the given value.
//
// See https://htmx.org/attributes/hx-boost/
func HxBoost(value string) gomponents.Node {
	return gomponents.Attr("hx-boost", value)
}

// HxOn returns a gomponents node with the hx-on:name="value"
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-on/
func HxOn(name string, value string) gomponents.Node {
	return gomponents.Attr("hx-on:"+name, value)
}

// HxPushURL returns a gomponents node with the hx-push-url
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-push-url/
func HxPushURL(value string) gomponents.Node {
	return gomponents.Attr("hx-push-url", value)
}

// HxSelect returns a gomponents node with the hx-select
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-select/
func HxSelect(value string) gomponents.Node {
	return gomponents.Attr("hx-select", value)
}

// HxSelectOOB returns a gomponents node with the hx-select-oob
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-select-oob/
func HxSelectOOB(value string) gomponents.Node {
	return gomponents.Attr("hx-select-oob", value)
}

// HxSwapOOB returns a gomponents node with the hx-swap-oob
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-swap-oob/
func HxSwapOOB(value string) gomponents.Node {
	return gomponents.Attr("hx-swap-oob", value)
}

// HxVals returns a gomponents node with the hx-vals
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-vals/
func HxVals(value string) gomponents.Node {
	return gomponents.Attr("hx-vals", value)
}

// HxDisable returns a gomponents node with the hx-disable
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-disable/
func HxDisable(value string) gomponents.Node {
	return gomponents.Attr("hx-disable", value)
}

// HxDisabledELT returns a gomponents node with the hx-disabled-elt
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-disabled-elt/
func HxDisabledELT(value string) gomponents.Node {
	return gomponents.Attr("hx-disabled-elt", value)
}

// HxDisinherit returns a gomponents node with the hx-disinherit
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-disinherit/
func HxDisinherit(value string) gomponents.Node {
	return gomponents.Attr("hx-disinherit", value)
}

// HxEncoding returns a gomponents node with the hx-encoding
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-encoding/
func HxEncoding(value string) gomponents.Node {
	return gomponents.Attr("hx-encoding", value)
}

// HxExt returns a gomponents node with the hx-ext
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-ext/
func HxExt(value string) gomponents.Node {
	return gomponents.Attr("hx-ext", value)
}

// HxHeaders returns a gomponents node with the hx-headers
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-headers/
func HxHeaders(value string) gomponents.Node {
	return gomponents.Attr("hx-headers", value)
}

// HxHistory returns a gomponents node with the hx-history
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-history/
func HxHistory(value string) gomponents.Node {
	return gomponents.Attr("hx-history", value)
}

// HxHistoryElt returns a gomponents node with the hx-history-elt
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-history-elt/
func HxHistoryElt(value string) gomponents.Node {
	return gomponents.Attr("hx-history-elt", value)
}

// HxInclude returns a gomponents node with the hx-include
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-include/
func HxInclude(value string) gomponents.Node {
	return gomponents.Attr("hx-include", value)
}

// HxParams returns a gomponents node with the hx-params
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-params/
func HxParams(value string) gomponents.Node {
	return gomponents.Attr("hx-params", value)
}

// HxPreserve returns a gomponents node with the hx-preserve
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-preserve/
func HxPreserve(value string) gomponents.Node {
	return gomponents.Attr("hx-preserve", value)
}

// HxPrompt returns a gomponents node with the hx-prompt
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-prompt/
func HxPrompt(value string) gomponents.Node {
	return gomponents.Attr("hx-prompt", value)
}

// HxReplaceURL returns a gomponents node with the hx-replace-url
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-replace-url/
func HxReplaceURL(value string) gomponents.Node {
	return gomponents.Attr("hx-replace-url", value)
}

// HxRequest returns a gomponents node with the hx-request
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-request/
func HxRequest(value string) gomponents.Node {
	return gomponents.Attr("hx-request", value)
}

// HxSync returns a gomponents node with the hx-sync
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-sync/
func HxSync(value string) gomponents.Node {
	return gomponents.Attr("hx-sync", value)
}

// HxValidate returns a gomponents node with the hx-validate
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-validate/
func HxValidate(value string) gomponents.Node {
	return gomponents.Attr("hx-validate", value)
}
