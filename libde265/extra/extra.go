package extra

import "embed"

//go:embed *.c *.h
var _ embed.FS
