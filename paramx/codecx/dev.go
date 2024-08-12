package codecx

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// DEV decoding encoding & video
type DEV struct {
	Idct int `json:"idct" flag:"-idct"` // (decoding/encoding, video)

	IdctFlag flagx.Idct `json:"idct_flag" flag:"-idct"` // (decoding/encoding,video)

	Threads typex.UNumber `json:"threads" flag:"-threads"` // (decoding/encoding, video)

	ColorPrimaries typex.ColorUI8 `json:"color_primaries" flag:"-color_primaries"` // (decoding/encoding, video)

	ColorPrimariesFlag flagx.ColorPrimaries `json:"color_primaries_flag" flag:"-color_primaries"` // (decoding/encoding, video)

	ColorTrc typex.ColorUI8 `json:"color_trc" flag:"-color_trc"` // (decoding/encoding, video)

	ColorTrcFlag flagx.ColorTrc `json:"color_trc_flag" flag:"-color_trc"` // (decoding/encoding, video)

	Colorspace flagx.Colorspace `json:"colorspace" flag:"-colorspace"` // (decoding/encoding, video)

	ColorRange typex.Bool `json:"color_range" flag:"-color_range"` // (decoding/encoding, video)

	ColorRangeFlag flagx.ColorRange `json:"color_range_flag" flag:"-color_range"` // (decoding/encoding,video)

	ChromaSampleLocation int `json:"chroma_sample_location" flag:"-chroma_sample_location"` // (decoding/encoding, video)

	ChromaSampleLocationFlag flagx.ChromaSampleLocationType `json:"chroma_sample_location_flag" flag:"-chroma_sample_location"` // (decoding/encoding,video)

	ThreadType flagx.ThreadType `json:"thread_type" flag:"-thread_type"` // (decoding/encoding, video)

	MaxPixels typex.Size `json:"max_pixels" flag:"-max_pixels"` // (decoding/encoding, video)

}
