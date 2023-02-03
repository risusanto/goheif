package extra

import "C"

import "embed"

//go:embed *.c *.h
var _ embed.FS
