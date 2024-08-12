package optionx

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

type Scaler struct {
	SwsFlags flagx.SwsFlags `json:"sws_flags" flag:"-sws_flags"`
	Srcw     typex.UNumber  `json:"srcw" flag:"-srcw"` //(API only)
	//Set source width.

	Srch typex.UNumber `json:"srch" flag:"-srch"` //(API only)
	//Set source height.

	Dstw typex.UNumber `json:"dstw" flag:"-dstw"` //(API only)
	//Set destination width.

	Dsth typex.UNumber `json:"dsth" flag:"-dsth"` //(API only)
	//Set destination height.

	SrcFormat typex.UNumber `json:"src_format" flag:"-src_format"` //(API only)
	//Set source pixel format (must be expressed as an integer).

	DstFormat typex.UNumber `json:"dst_format" flag:"-dst_format"` //(API only)
	//Set destination pixel format (must be expressed as an integer).

	SrcRange typex.UBool `json:"src_range" flag:"-src_range"` // (boolean)
	//If value is set to 1, indicates source is full range. Default value is 0, which indicates source is limited range.

	DstRange typex.UBool `json:"dst_range" flag:"-dst_range"` // (boolean)
	//If value is set to 1, enable full range for destination. Default value is 0, which enables limited range.

	Param0 typex.Number `json:"param0" flag:"-param0"`
	Param1 typex.Number `json:"param1" flag:"-param1"`
	//Set scaling algorithm parameters. The specified values are specific of some scaling algorithms and ignored by others. The specified values are floating point number values.
	SwsDither flagx.SwsDither `json:"sws_dither" flag:"-sws_dither"`
	//Set the dithering algorithm. Accepts one of the following values. Default value is ‘auto’.
	Alphablend Alphablend `json:"alphablend" flag:"-alphablend"`
	//Set the alpha blending to use when the input has alpha but the output does not. Default value is ‘none’.
}

type Alphablend = typex.Flags

const (
	Alphablend_uniform_color Alphablend = "uniform_color"
	//Blend onto a uniform background color

	Alphablend_checkerboard Alphablend = "checkerboard"
	//Blend onto a checkerboard

	Alphablend_none Alphablend = "none"
	//No blending
)
