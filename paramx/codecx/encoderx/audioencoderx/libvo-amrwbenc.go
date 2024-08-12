package audioencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// LIBVO-AMRWBENC 8.12 libvo-amrwbenc
type LibvoAmrwbenc struct {
	B typex.Bitrate `json:"b" flag:"-b"` //

	/*
		‘6600’
		‘8850’
		‘12650’
		‘14250’
		‘15850’
		‘18250’
		‘19850’
		‘23050’
		‘23850’
	*/
	Dtx typex.Bool `json:"dtx" flag:"-dtx"`
	//Allow discontinuous transmission (generate comfort noise) when set to 1. The default value is 0 (disabled).
}
