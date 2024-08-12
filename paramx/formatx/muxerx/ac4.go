package muxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// AC4 21.6 ac4 AC4
type AC4 struct {
	WriteCrc typex.UBool `json:"write_crc" flag:"-write_crc"`
	//when enabled, write a CRC checksum for each packet to the output, default is false

}
