package videoencoderx

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// PIXEL Format libwebp 9.15 libwebp
//
//	9.15.1 Pixel Format
type PIXEL struct {
	Lossless typex.UBool `json:"lossless" flag:"-lossless"`
	//Enables/Disables use of lossless mode. Default is 0.

	CompressionLevel typex.Level `json:"compression_level" flag:"-compression_level"`
	//For lossy, this is a quality/speed tradeoff. Higher values give better quality for a given size at the cost of increased encoding time. For lossless, this is a size/speed tradeoff. Higher values give smaller size at the cost of increased encoding time. More specifically, it controls the number of extra algorithms and compression tools used, and varies the combination of these tools. This maps to the method option in libwebp. The valid range is 0 to 6. Default is 4.

	Quality typex.Quality `json:"quality" flag:"-quality"`
	//For lossy encoding, this controls image quality. For lossless encoding, this controls the effort and time spent in compression. Range is 0 to 100. Default is 75.

	Preset flagx.Preset `json:"preset" flag:"-preset"`
	//Configuration preset. This does some automatic settings based on the general type of the image.
}
