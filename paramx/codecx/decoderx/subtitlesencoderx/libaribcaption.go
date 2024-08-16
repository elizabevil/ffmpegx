package subtitledecoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// LIBARIBCAPTION 6.2 libaribcaption
type LIBARIBCAPTION struct {
	SubType          Subtype         `json:"sub_type" flag:"-sub_type"`
	CaptionEncoding  CaptionEncoding `json:"caption_encoding" flag:"-caption_encoding"`
	Font             typex.Font      `json:"font " flag:"-font "`
	AssSingleRect    typex.Bool      `json:"ass_single_rect " flag:"-ass_single_rect "`
	ForceOutlineText typex.Bool      `json:"force_outline_text" flag:"-force_outline_text"`
	OutlineWidth     typex.Bool      `json:"outline_width" flag:"-outline_width"`

	IgnoreBackground typex.Bool `json:"ignore_background" flag:"-ignore_background"`
	//Specify whether to ignore background color rendering.

	//The default is false.

	IgnoreRuby typex.Bool `json:"ignore_ruby" flag:"-ignore_ruby"`
	//Specify whether to ignore rendering for ruby-like (furigana) characters.

	//The default is false.

	ReplaceDrcs typex.Bool `json:"replace_drcs" flag:"-replace_drcs"`
	//Specify whether to render replaced DRCS characters as Unicode characters.

	//The default is true.

	ReplaceMszAscii typex.Bool `json:"replace_msz_ascii" flag:"-replace_msz_ascii"`
	//Specify whether to replace MSZ (Middle Size; half width) fullwidth alphanumerics with halfwidth alphanumerics.

	//The default is true.

	ReplaceMszJapanese typex.Bool `json:"replace_msz_japanese" flag:"-replace_msz_japanese"`
	//Specify whether to replace some MSZ (Middle Size; half width) fullwidth japanese special characters with halfwidth ones.

	//The default is true.

	ReplaceMszGlyph typex.Bool      `json:"replace_msz_glyph" flag:"-replace_msz_glyph"`
	CanvasSize      typex.VideoSize `json:"canvas_size" flag:"-canvas_size"`
}

/*

ffplay -sub_type bitmap MPEG.TS
Display MPEG-TS file with input frame size 1920x1080 by ffplay tool:

ffplay -sub_type bitmap -canvas_size 1920x1080 MPEG.TS
Embed ARIB subtitle in transcoded video:

ffmpeg -sub_type bitmap -i src.m2t -filter_complex "[0:v][0:s]overlay" -vcodec h264 dest.mp4
*/

type Subtype = typex.String
type CaptionEncoding = typex.String

const (
	Subtype_bitmap Subtype = "bitmap"
	//Graphical image.

	Subtype_ass Subtype = "ass"
	//ASS formatted text.

	Subtype_text Subtype = "text"
	//Simple text based output without formatting.
)

const (
	CaptionEncoding_auto CaptionEncoding = "auto"
	//Automatically detect text encoding (default).

	CaptionEncoding_jis CaptionEncoding = "jis"
	//8bit-char JIS encoding defined in ARIB STD B24. This encoding used in Japan for ISDB captions.

	CaptionEncoding_utf8 CaptionEncoding = "utf8"
	//UTF-8 encoding defined in ARIB STD B24. This encoding is used in Philippines for ISDB-T captions.

	CaptionEncoding_latin CaptionEncoding = "latin"
	//Latin character encoding defined in ABNT NBR 15606-1. This encoding is used in South America for SBTVD / ISDB-Tb captions.
)
