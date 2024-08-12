package inputdevicex

import "github.com/elizabevil/ffmpegx/paramx/typex"

// PULSE 3.17 pulse
type PULSE struct {
	Server typex.Server `json:"server" flag:"-server"`
	//Connect to a specific PulseAudio server, specified by an IP address. Default server is used when not provided.

	Name typex.Name `json:"name" flag:"-name"`
	//Specify the application name PulseAudio will use when showing active clients, by default it is the LIBAVFORMAT_IDENT string.

	StreamName typex.StreamName `json:"stream_name" flag:"-stream_name"`
	//Specify the stream name PulseAudio will use when showing active streams, by default it is "record".

	SampleRate typex.SampleRate `json:"sample_rate" flag:"-sample_rate"`
	//Specify the samplerate in Hz, by default 48kHz is used.

	Channels typex.Channels `json:"channels" flag:"-channels"`
	//Specify the channels in use, by default 2 (stereo) is set.

	//FrameSize int `json:"frame_size" flag:"-frame_size"`
	//This option does nothing and is deprecated.

	FragmentSize typex.Size `json:"fragment_size" flag:"-fragment_size"`
	//Specify the size in bytes of the minimal buffering fragment in PulseAudio, it will affect the audio latency. By default it is set to 50 ms amount of data.

	Wallclock typex.Bool `json:"wallclock" flag:"-wallclock"`
	//Set the initial PTS using the current time. Default is 1.

}
