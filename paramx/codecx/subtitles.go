package codecx

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// Subtitle
type Subtitle struct {
	Flags flagx.Flag `json:"flags" flag:"-flags"` // (decoding/encoding,audio,video,subtitles)

	Debug flagx.Bug `json:"debug" flag:"-debug"` // (decoding/encoding,audio,video,subtitles)

	Flags2 flagx.Flags2 `json:"flags2" flag:"-flags2"` // (decoding/encoding,audio,video,subtitles)

	ExportSideData flagx.ExportSideData `json:"export_side_data" flag:"-export_side_data"` // (decoding/encoding,audio,video,subtitles)

	SubCharenc typex.Encoding `json:"sub_charenc" flag:"-sub_charenc"` // (decoding,subtitles)

	//Set the input subtitles character encoding.

}
