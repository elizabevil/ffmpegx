package muxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// 21.19 ast AST
type AST struct {
	Loopstart typex.MillisecondI `json:"loopstart" flag:"-loopstart"`
	//Specify loop start position expressesd in milliseconds, from -1 to INT_MAX, in case -1 is set then no loop is specified (default -1) and the loopend value is ignored.

	Loopend typex.MillisecondI `json:"loopend" flag:"-loopend"`
	//Specify loop end position expressed in milliseconds, from 0 to INT_MAX, default is 0, in case 0 is set it assumes the total stream duration.
}
