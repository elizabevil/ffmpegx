package audioencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// OPUS 8.4 opus
type OPUS struct {
	B typex.Bitrate `json:"b" flag:"-b"`

	//Set bit rate in bits/s. If unspecified it uses the number of channels and the layout to make a good guess.
	OpusDelay typex.MillisecondF `json:"opus_delay" flag:"-opus_delay"`

	//Sets the maximum delay in milliseconds. Lower delays than 20ms will very quickly decrease quality.
}
