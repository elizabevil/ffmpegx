package protocolx

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// 24.34 rtsp
// rtsp://hostname[:port]/path
// 24.34.1 Muxer
type RstpMuxer struct {
	RtspTransport flagx.RtspTransport `json:"rtsp_transport" flag:"-rtsp_transport"`
	RtspFlags     flagx.RtspFlags     `json:"rtsp_flags" flag:"-rtsp_flags"`
	MinPort       typex.Port          `json:"min_port" flag:"-min_port"`
	//Set minimum local UDP port. Default value is 5000.

	MaxPort typex.Port `json:"max_port" flag:"-max_port"`
	//Set maximum local UDP port. Default value is 65000.

	BufferSize typex.Size `json:"buffer_size" flag:"-buffer_size"`
	//Set the maximum socket buffer size in bytes.

	PktSize typex.Size `json:"pkt_size" flag:"-pkt_size"`
	//Set max send packet size (in bytes). Default value is 1472.
}

// 24.34.2 Demuxer
type RstpDemuxer struct {
	InitialPause typex.Bool `json:"initial_pause" flag:"-initial_pause"`
	//Do not start playing the stream immediately if set to 1. Default value is 0.

	RtspTransport     flagx.Transport         `json:"rtsp_transport" flag:"-rtsp_transport"`
	RtspFlags         flagx.RtspFlags         `json:"rtsp_flags" flag:"-rtsp_flags"`
	AllowedMediaTypes flagx.AllowedMediaTypes `json:"allowed_media_types" flag:"-allowed_media_types"`
	MinPort           typex.Port              `json:"min_port" flag:"-min_port"`
	//Set minimum local UDP port. Default value is 5000.

	MaxPort typex.Port `json:"max_port" flag:"-max_port"`
	//Set maximum local UDP port. Default value is 65000.

	ListenTimeout typex.Size `json:"listen_timeout" flag:"-listen_timeout"`
	//Set maximum timeout (in seconds) to establish an initial connection. Setting listen_timeout > 0 sets rtsp_flags to ‘listen. Default is -1 which means an infinite timeout when ‘listen mode is set.

	ReorderQueueSize typex.Size `json:"reorder_queue_size" flag:"-reorder_queue_size"`
	//Set number of packets to buffer for handling of reordered packets.

	Timeout typex.MicrosecondI `json:"timeout" flag:"-timeout"`
	//Set socket TCP I/O timeout in microseconds.

	UserAgent string `json:"user_agent" flag:"-user_agent"`
	//Override User-Agent header. If not specified, it defaults to the libavformat identifier string.

	BufferSize typex.Size `json:"buffer_size" flag:"-buffer_size"`
	//Set the maximum socket buffer size in bytes.
}
