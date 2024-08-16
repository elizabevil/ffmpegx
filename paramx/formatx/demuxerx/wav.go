package demuxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// 20.26 wav  WAV
type WAV struct {
	//This demuxer accepts the following options:

	MaxSize typex.UNumber `json:"max_size" flag:"-max_size"`
	//Specify the maximum packet size in bytes for the demuxed packets.By default this is set to 0, which means that a sensible value is chosen based on the input format.
}
