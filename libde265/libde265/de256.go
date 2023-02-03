package libde265

import "C"

import (
	"embed"

	_ "github.com/risusanto/goheif/libde265/libde265/arm"
	_ "github.com/risusanto/goheif/libde265/libde265/x86"
)

//go:embed *.cc *.h
var _ embed.FS
