package optionx

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// IO input/output
type IO = Common

// Common input/output
type Common struct {
	//5.4 Main options
	F              typex.Fmt               `json:"f" flag:"-f"`
	CodecX         []typex.StreamSpecifier `json:"c_specifier" flag:"-c"`         // -c:v -c:a //-c[:stream_specifier]
	C              string                  `json:"c" flag:"-c"`                   // -c copy simple -codec
	CodecSpecifier []typex.StreamSpecifier `json:"codec_specifier" flag:"-codec"` // -codec copy or -codec: copy would copy all the streams without reencoding.
	Codec          string                  `json:"codec" flag:"-codec"`           // -codec:v -codec:a //codec (input/output,per-stream)

	B       []string                `json:"b" flag:"-b"`       // -c copy simple -codec
	Bitrate []typex.StreamSpecifier `json:"bitrate" flag:"-b"` // b:v

	T  typex.Duration `json:"t" flag:"-t"`
	To typex.Position `json:"to" flag:"-to"`
	Ss typex.Position `json:"ss" flag:"-ss"` //When used as an input option (before -i), seeks in this input file to position.
	Dn bool           `json:"dn" flag:"-dn"`

	//5.5 Video Options
	R  typex.FrameRate `json:"r" flag:"-r"` // -r[:stream_specifier]    // (input/output,per-stream)
	S  typex.Size      `json:"s" flag:"-s"` // -s[:stream_specifier]    // (input/output,per-stream)
	Vn bool            `json:"vn" flag:"-vn"`

	//5.6 Advanced Video options
	PixFmt   typex.Format   `json:"pix_fmt" flag:"-pix_fmt"` // -pix_fmt[:stream_specifier]    // (input/output,per-stream)
	SwsFlags flagx.SwsFlags `json:"sws_flags" flag:"-sws_flags"`
	//5.7 Audio Options
	Ar                typex.Freq            `json:"ar" flag:"-ar"`     // -ar[:stream_specifier]    // (input/output,per-stream)
	ArStreamSpecifier typex.StreamSpecifier `json:"ar_map" flag:"-ar"` // -ar[:stream_specifier]    // (input/output,per-stream)

	Ac     typex.Channels `json:"ac" flag:"-ac"` // -ac[:stream_specifier]    // (input/output,per-stream) //Alias for -channel_layout.
	An     bool           `json:"an" flag:"-an"`
	Acodec typex.Codec    `json:"acodec" flag:"-acodec"` //This is an alias for -codec:a.

	//5.8 Advanced Audio options
	ChLayout      typex.StreamSpecifier `json:"ch_layout" flag:"-ch_layout"`           // -ch_layout[:stream_specifier]    // (input/output,per-stream)
	ChannelLayout typex.StreamSpecifier `json:"channel_layout" flag:"-channel_layout"` // -channel_layout[:stream_specifier]    // (input/output,per-stream)
	// -tag[:stream_specifier]    // (input/output,per-stream)
	//5.9 Subtitle options
	Scodec typex.Codec `json:"scodec" flag:"-scodec"` //Set the subtitle codec. This is an alias for -codec:s. (input/output)
	Sn     bool        `json:"sn" flag:"-sn"`
	//5.11 Advanced options
	Bitexact        bool                  `json:"bitexact" flag:"-bitexact"`
	Bsf             typex.StreamSpecifier `json:"bsf" flag:"-bsf"` // -bsf[:stream_specifier]    // (input/output,per-stream)  -bsf:v h264_mp4toannexb
	Tag             typex.StreamSpecifier `json:"tag" flag:"-tag"`
	ThreadQueueSize typex.Size            `json:"thread_queue_size" flag:"-thread_queue_size"`

	//5.5 Video Options
	Vcodec    typex.Codec     `json:"vcodec" flag:"-vcodec"`
	SampleFmt typex.SampleFmt `json:"sample_fmt" flag:"-sample_fmt"`
	Dcodec    typex.Codec     `json:"dcodec" flag:"-dcodec"`
	Top       typex.Codec     `json:"top" flag:"-top"`
}

type Fdebug typex.String

const (
	Fdebug_Ts Fdebug = "ts"
)
