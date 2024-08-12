package optionx

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// https://ffmpeg.org/ffmpeg-all.html
// OPT_INPUT
type Input struct {
	Input          string           `json:"i"                flag:"-i"`
	Inputs         []string         `json:"i_s"              flag:"-i"`
	StreamLoop     typex.Bool       `json:"stream_loop"      flag:"-stream_loop"` // -1,0,
	Sseof          typex.Position   `json:"sseof"            flag:"-sseof"`       //0 is at EOF
	Isync          typex.InputIndex `json:"isync"            flag:"-isync"`       // Assign an input as a sync source.
	Itsoffset      typex.U16Offset  `json:"itsoffset"        flag:"-itsoffset"`   //
	Readrate       typex.SpeedRate  `json:"readrate"         flag:"-readrate"`    //Limit input read speed.
	Re             bool             `json:"re"               flag:"-re"`          //Read input at native frame rate. This is equivalent to setting -readrate 1.
	AccurateSeek   bool             `json:"accurate_seek"    flag:"-accurate_seek"`
	NoAccurateSeek bool             `json:"noaccurate_seek"    flag:"-noaccurate_seek"`
	SeekTimestamp  typex.Number     `json:"seek_timestamp"   flag:"-seek_timestamp"`
	Discard        flagx.Discard    `json:"discard"          flag:"-discard"`
}
