package videoencoderx

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

/*
Libvvenc
VVenC H.266/VVC encoder wrapper.
This encoder requires the presence of the libvvenc headers and library during configuration. You need to explicitly configure the build with --enable-libvvenc.
The VVenC project website is at https://github.com/fraunhoferhhi/vvenc.
*/
type Libvvenc struct {
	B typex.BitrateInt `json:"b" flag:"-b"`
	//Sets target video bitrate.
	Bitrate typex.Bitrate `json:"bitrate" flag:"-b"`

	G typex.Size `json:"g" flag:"-g"`
	//Set the GOP size. Currently support for g=1 (Intra only) or default.
	Preset flagx.Preset `json:"preset" flag:"-preset"`
	//Set the VVenC preset.

	Levelidc typex.Level `json:"levelidc" flag:"-levelidc"`
	//Set level idc.

	Tier QsvHevcTier `json:"tier" flag:"-tier"`
	//Set vvc tier.

	Qp typex.Level `json:"qp" flag:"-qp"`
	//Set constant quantization parameter.

	Subopt *typex.Bool `json:"subopt" flag:"-subopt,1"`
	//Set subjective (perceptually motivated) optimization. Default is 1 (on).

	Bitdepth8 typex.Bool `json:"bitdepth8" flag:"-bitdepth8"`
	//Set 8bit coding mode instead of using 10bit. Default is 0 (off).

	Period      typex.SecondI `json:"period" flag:"-period"` //	set (intra) refresh period in seconds.
	VvencParams string        `json:"vvenc-params" flag:"-vvenc-params"`
	//Set vvenc options using a list of key=value couples separated by ":". See vvencapp --fullhelp or vvencFFapp --fullhelp for a list of options.
}

/*
For example, the options might be provided as:
	intraperiod=64:decodingrefreshtype=idr:poc0idr=1:internalbitdepth=8
For example the encoding options might be provided with -vvenc-params:
	ffmpeg -i input -c:v libvvenc -b 1M -vvenc-params intraperiod=64:decodingrefreshtype=idr:poc0idr=1:internalbitdepth=8 output.mp4
*/
