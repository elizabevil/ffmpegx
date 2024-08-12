package muxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// 21.62 mpegts MPEGTS
type MPEG struct {
	Muxrate typex.RateI `json:"muxrate" flag:"-muxrate"`
	//Set user-defined mux rate expressed as a number of bits/s. If not specied the automatically computed mux rate is employed. Default value is 0.

	Preload typex.MicrosecondI `json:"preload" flag:"-preload"`
	//Set initial demux-decode delay in microseconds. Default value is 500000.

}
