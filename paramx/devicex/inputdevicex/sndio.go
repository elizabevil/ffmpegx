package inputdevicex

import "github.com/elizabevil/ffmpegx/paramx/typex"

// SNDIO 3.18 sndio
type SNDIO struct {
	Channels typex.Channels `json:"channels" flag:"-channels"`
	//Set the number of channels. Default is 2.
	SampleRate typex.SampleRate `json:"sample_rate" flag:"-sample_rate"`
	//	Set the sample rate in Hz. Default is 48000.
}
