package component

import "github.com/maragudk/gomponents"

// GMap is a convenience function to render a gomponents.Group
// with a map inside.
func GMap[T any](ts []T, cb func(T) gomponents.Node) gomponents.Node {
	return gomponents.Group(gomponents.Map(ts, cb))
}
