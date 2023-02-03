package x86

import "C"

import "embed"

//go:embed *.cc *.h
var _ embed.FS
