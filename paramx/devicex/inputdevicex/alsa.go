package inputdevicex

import "github.com/elizabevil/ffmpegx/paramx/typex"

// ALSA 3.1 alsa
type ALSA struct {
	SampleRate typex.SampleRate `json:"sample_rate" flag:"-sample_rate"`
	//Set the sample rate in Hz. Default is 48000.

	Channels typex.Channels `json:"channels" flag:"-channels"`
	//Set the number of channels. Default is 2.
}
