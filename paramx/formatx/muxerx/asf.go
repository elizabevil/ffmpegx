package muxerx

import (
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// 21.17 asf, asf_stream ASF,
type ASF struct {
	PacketSize typex.Size `json:"packet_size" flag:"-packet_size"`
	//Set the muxer packet size as a number of bytes. By tuning this setting you may reduce data fragmentation or muxer overhead depending on your source. Default value is 3200, minimum is 100, maximum is 64Ki.
}
