package audioencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// LIBOPENCORE-AMRNB 8.8 libopencore-amrnb
type LibopencoreAmrnb struct {
	B typex.Bitrate `json:"b" flag:"-b"` //

	/*
		4750
		5150
		5900
		6700
		7400
		7950
		10200
		12200
	*/
	Dtx typex.Bool `json:"dtx" flag:"-dtx"`
	//Allow discontinuous transmission (generate comfort noise) when set to 1. The default value is 0 (disabled).
}
