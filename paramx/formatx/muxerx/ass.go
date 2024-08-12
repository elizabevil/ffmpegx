package muxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// 21.18 ass ASS
type ASS struct {
	IgnoreReadorder typex.UBool `json:"ignore_readorder" flag:"-ignore_readorder"`
	//Write dialogue events immediately, even if they are out-of-order, default is false, otherwise they are cached until the expected time event is found.
}
