package videoencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// // GIF 9.3 GIF
type GIF struct {
	Gifflags Gifflags `json:"gifflags" flag:"-gifflags"`
	//Sets the flags used for GIF encoding.

	//Default is enabled.

	Gifimage typex.UBool `json:"gifimage" flag:"-gifimage"`
	//Enables encoding one full GIF image per frame, rather than an animated GIF.

	//Default value is 0.

	GlobalPalette *typex.UBool `json:"global_palette" flag:"-global_palette,1"`
	//Writes a palette to the global GIF header where feasible.

	//If disabled, every frame will always have a palette written, even if there is a global palette supplied.

	//Default value is 1.
}

type Gifflags string

const (
	Gifflags_offsetting Gifflags = "offsetting"

	Gifflags_transdiff Gifflags = "transdiff"
)
