package arm

import "embed"

//go:embed *.cc *.h
var _ embed.FS
