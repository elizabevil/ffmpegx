package codecx

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

type Decoding struct {
	Flags flagx.Flag `json:"flags" flag:"-flags"` // (decoding/encoding,audio,video,subtitles)

	Ar typex.Ar `json:"ar" flag:"-ar"` // (decoding/encoding,audio)

	Ac typex.Channels `json:"ac" flag:"-ac"` // (decoding/encoding,audio)

	Bug flagx.Bug `json:"bug" flag:"-bug"` // (decoding,video)

	Strict typex.Number `json:"strict" flag:"-strict"` // (decoding/encoding,audio,video)

	ErrDetect flagx.ErrDetect `json:"err_detect" flag:"-err_detect"` // (decoding,audio,video)

	Idct     typex.UNumber `json:"idct" flag:"-idct"`      // (decoding/encoding,video)
	IdctFlag flagx.Idct    `json:"idct_flag" flag:"-idct"` // (decoding/encoding,video)

	Ec flagx.Ec `json:"ec" flag:"-ec"` // (decoding,video)

	Debug flagx.Bug `json:"debug" flag:"-debug"` // (decoding/encoding,audio,video,subtitles)

	Flags2 flagx.Flags2 `json:"flags2" flag:"-flags2"` // (decoding/encoding,audio,video,subtitles)

	ExportSideData flagx.ExportSideData `json:"export_side_data" flag:"-export_side_data"` // (decoding/encoding,audio,video,subtitles)

	Threads typex.UNumber `json:"threads" flag:"-threads"` // (decoding/encoding,video)

	SkipTop typex.Number `json:"skip_top" flag:"-skip_top"` // (decoding,video)

	SkipBottom typex.Number `json:"skip_bottom" flag:"-skip_bottom"` // (decoding,video)

	Lowres typex.Size `json:"lowres" flag:"-lowres"` // (decoding,audio,video)

	SkipLoopFilter int `json:"skip_loop_filter" flag:"-skip_loop_filter"` // (decoding,video)

	SkipIdct int `json:"skip_idct" flag:"-skip_idct"` // (decoding,video)

	SkipFrame     typex.Bool      `json:"skip_frame" flag:"-skip_frame"`      // (decoding,video)
	SkipFrameFlag flagx.SkipFrame `json:"skip_frame_flag" flag:"-skip_frame"` // (decoding,video)

	ChannelLayout typex.ChannelLayout `json:"channel_layout" flag:"-channel_layout"` // (decoding/encoding,audio)

	ColorPrimaries     int                  `json:"color_primaries" flag:"-color_primaries"`      // (decoding/encoding,video)
	ColorPrimariesFlag flagx.ColorPrimaries `json:"color_primaries_flag" flag:"-color_primaries"` // (decoding/encoding,video)

	ColorTrc     int            `json:"color_trc" flag:"-color_trc"`      // (decoding/encoding,video)
	ColorTrcFlag flagx.ColorTrc `json:"color_trc_flag" flag:"-color_trc"` // (decoding/encoding,video)

	Colorspace flagx.Colorspace `json:"colorspace" flag:"-colorspace"` // (decoding/encoding, video)

	ColorRange     typex.Bool       `json:"color_range" flag:"-color_range"`      // (decoding/encoding,video)
	ColorRangeFlag flagx.ColorRange `json:"color_range_flag" flag:"-color_range"` // (decoding/encoding,video)

	ChromaSampleLocation     int                            `json:"chroma_sample_location" flag:"-chroma_sample_location"`      // (decoding/encoding,video)
	ChromaSampleLocationFlag flagx.ChromaSampleLocationType `json:"chroma_sample_location_flag" flag:"-chroma_sample_location"` // (decoding/encoding,video)

	ThreadType flagx.ThreadType `json:"thread_type" flag:"-thread_type"` // (decoding/encoding,video)

	RequestSampleFmt typex.SampleFmt `json:"request_sample_fmt" flag:"-request_sample_fmt"` // (decoding,audio)

	SubCharenc typex.Encoding `json:"sub_charenc" flag:"-sub_charenc"` // (decoding,subtitles)

	SkipAlpha typex.UBool `json:"skip_alpha" flag:"-skip_alpha"` // (decoding,video)

	MaxPixels typex.Size `json:"max_pixels" flag:"-max_pixels"` // (decoding/encoding,video)

	ApplyCropping typex.UBool `json:"apply_cropping" flag:"-apply_cropping"` // (decoding,video)

}

type ApplyCropping = typex.String

const (
	ApplyCropping_none      ApplyCropping = "none"
	ApplyCropping_all       ApplyCropping = "all"
	ApplyCropping_codec     ApplyCropping = "codec"
	ApplyCropping_container ApplyCropping = "container"
)
