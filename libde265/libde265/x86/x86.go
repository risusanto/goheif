package x86

import "embed"

//go:embed *.cc *.h
var _ embed.FS
