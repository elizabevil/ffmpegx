package demuxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// 20.25 w64  W64
type W64 struct {
	MaxSize typex.UNumber `json:"max_size" flag:"-max_size"`
	//See the same option for the wav demuxer.
}
