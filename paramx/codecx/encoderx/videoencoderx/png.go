package videoencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// PNG 9.24 png
type PNG struct {
	Dpi typex.UI16 `json:"dpi" flag:"-dpi"`
	//Set physical density of pixels, in dots per inch, unset by default

	Dpm typex.UI16 `json:"dpm" flag:"-dpm"`
	//Set physical density of pixels, in dots per meter, unset by default
}
