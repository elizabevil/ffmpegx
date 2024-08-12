package codecx

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// DE decoding | encoding
type DE struct {
	Flags flagx.Flag `json:"flags" flag:"-flags"` // (decoding/encoding,audio,video,subtitles)

	Ar typex.Ar `json:"ar" flag:"-ar"` // (decoding/encoding,audio)

	Ac typex.Channels `json:"ac" flag:"-ac"` // (decoding/encoding,audio)

	Strict typex.Number `json:"strict" flag:"-strict"` // (decoding/encoding,audio,video)

	Idct int `json:"idct" flag:"-idct"` // (decoding/encoding,video)

	Debug flagx.Bug `json:"debug" flag:"-debug"` // (decoding/encoding,audio,video,subtitles)

	Flags2 flagx.Flags2 `json:"flags2" flag:"-flags2"` // (decoding/encoding,audio,video,subtitles)

	ExportSideData flagx.ExportSideData `json:"export_side_data" flag:"-export_side_data"` // (decoding/encoding,audio,video,subtitles)

	Threads typex.UNumber `json:"threads" flag:"-threads"` // (decoding/encoding,video)

	ChannelLayout typex.ChannelLayout `json:"channel_layout" flag:"-channel_layout"` // (decoding/encoding,audio)

	ColorPrimaries     typex.ColorUI8       `json:"color_primaries" flag:"-color_primaries"`      // (decoding/encoding,video)
	ColorPrimariesFlag flagx.ColorPrimaries `json:"color_primaries_flag" flag:"-color_primaries"` // (decoding/encoding,video)

	ColorTrc     typex.ColorUI8 `json:"color_trc" flag:"-color_trc"`      // (decoding/encoding,video)
	ColorTrcFlag flagx.ColorTrc `json:"color_trc_flag" flag:"-color_trc"` // (decoding/encoding,video)

	Colorspace flagx.Colorspace `json:"colorspace" flag:"-colorspace"` // (decoding/encoding, video)

	ColorRange     typex.Bool       `json:"color_range" flag:"-color_range"`      // (decoding/encoding,video)
	ColorRangeFlag flagx.ColorRange `json:"color_range_flag" flag:"-color_range"` // (decoding/encoding,video)

	ChromaSampleLocation     int                            `json:"chroma_sample_location" flag:"-chroma_sample_location"`      // (decoding/encoding,video)
	ChromaSampleLocationFlag flagx.ChromaSampleLocationType `json:"chroma_sample_location_flag" flag:"-chroma_sample_location"` // (decoding/encoding,video)

	ThreadType flagx.ThreadType `json:"thread_type" flag:"-thread_type"` // (decoding/encoding,video)

	MaxPixels typex.Size `json:"max_pixels" flag:"-max_pixels"` // (decoding/encoding,video)

}
