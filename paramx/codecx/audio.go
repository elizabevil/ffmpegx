package codecx

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

type Audio struct {
	B typex.BitrateInt `json:"b" flag:"-b"` // (encoding,audio,video)

	Ab typex.Bitrate `json:"ab" flag:"-ab"` // (encoding,audio)

	Flags flagx.Flag `json:"flags" flag:"-flags"` // (decoding/encoding,audio,video,subtitles)

	Ar typex.Ar `json:"ar" flag:"-ar"` // (decoding/encoding,audio)

	Ac typex.Channels `json:"ac" flag:"-ac"` // (decoding/encoding,audio)

	Cutoff typex.Cutoff `json:"cutoff" flag:"-cutoff"` // (encoding,audio)

	FrameSize typex.FrameSize `json:"frame_size" flag:"-frame_size"` // (encoding,audio)

	Strict typex.Number `json:"strict" flag:"-strict"` // (decoding/encoding,audio,video)

	ErrDetect flagx.ErrDetect `json:"err_detect" flag:"-err_detect"` // (decoding,audio,video)

	Maxrate typex.BitrateInt `json:"maxrate" flag:"-maxrate"` // (encoding,audio,video)

	Minrate typex.BitrateInt `json:"minrate" flag:"-minrate"` // (encoding,audio,video)

	Bufsize typex.Size `json:"bufsize" flag:"-bufsize"` // (encoding,audio,video)

	Debug flagx.Bug `json:"debug" flag:"-debug"` // (decoding/encoding,audio,video,subtitles)

	GlobalQuality typex.Number `json:"global_quality" flag:"-global_quality"` // (encoding,audio,video)

	Flags2 flagx.Flags2 `json:"flags2" flag:"-flags2"` // (decoding/encoding,audio,video,subtitles)

	ExportSideData flagx.ExportSideData `json:"export_side_data" flag:"-export_side_data"` // (decoding/encoding,audio,video,subtitles)

	Profile int `json:"profile" flag:"-profile"` // (encoding,audio,video)

	Level int `json:"level" flag:"-level"` // (encoding,audio,video)

	Lowres typex.Size `json:"lowres" flag:"-lowres"` // (decoding,audio,video)

	Trellis typex.Number `json:"trellis" flag:"-trellis"` // (encoding,audio,video)

	CompressionLevel typex.Level `json:"compression_level" flag:"-compression_level"` // (encoding,audio,video)

	ChannelLayout typex.ChannelLayout `json:"channel_layout" flag:"-channel_layout"` // (decoding/encoding,audio)

	AudioServiceType flagx.AudioServiceType `json:"audio_service_type" flag:"-audio_service_type"` // (encoding,audio)

	RequestSampleFmt typex.SampleRate `json:"request_sample_fmt" flag:"-request_sample_fmt"` // (decoding,audio)

}
