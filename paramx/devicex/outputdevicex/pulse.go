package outputdevicex

import "github.com/elizabevil/ffmpegx/paramx/typex"

// PULSE 4.8 pulse
type PULSE struct {
	Server typex.Server `json:"server" flag:"-server"`

	//Connect to a specific PulseAudio server, specified by an IP address. Default server is used when not provided.

	Name typex.Name `json:"name" flag:"-name"`

	//Specify the application name PulseAudio will use when showing active clients, by default it is the LIBAVFORMAT_IDENT string.

	StreamName typex.StreamName `json:"stream_name" flag:"-stream_name"`

	//Specify the stream name PulseAudio will use when showing active streams, by default it is set to the specified output name.

	Device typex.Device `json:"device" flag:"-device"`
	//Specify the device to use.Default device is used when not provided.List of output devices can be obtained with command pactl list sinks.

	BufferSize     typex.Size      `json:"buffer_size" flag:"-buffer_size"`
	BufferDuration typex.DurationI `json:"buffer_duration" flag:"-buffer_duration"`
	Prebuf         typex.Size      `json:"prebuf" flag:"-prebuf"`
	//Specify pre-buffering size in bytes.The server does not start with playback before at least prebuf bytes are available in the buffer.By default this option is initialized to the same value as buffer_size or buffer_duration (whichever is bigger).

	Minreq typex.Size `json:"minreq" flag:"-minreq"`
	//Specify minimum request size in bytes.The server does not request less than minreq bytes from the client, instead waits until the buffer is free enough to request more bytes at once.It is recommended to not set this option, which will initialize this to a value that is deemed sensible by the server.
	//
}
