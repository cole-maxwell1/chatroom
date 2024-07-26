package web

import "embed"

//go:embed "javascript"
//go:embed "css"
var Files embed.FS
