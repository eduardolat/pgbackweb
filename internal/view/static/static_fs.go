package static

import "embed"

//go:embed *
var StaticFs embed.FS
