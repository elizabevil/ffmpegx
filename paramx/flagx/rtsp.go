package flagx

type RtspTransport = Flag

const (
	RtspTransport_udp RtspTransport = "udp"
	RtspTransport_tcp RtspTransport = "tcp"
)

type Transport = Flag

const (
	Transport_udp Transport = "udp"
	//Use UDP as lower transport protocol.

	Transport_tcp Transport = "tcp"
	//Use TCP (interleaving within the RTSP control channel) as lower transport protocol.

	Transport_udp_multicast Transport = "udp_multicast"
	//Use UDP multicast as lower transport protocol.

	Transport_http Transport = "http"
	//Use HTTP tunneling as lower transport protocol, which is useful for passing proxies.

	Transport_https Transport = "https"
	//Use HTTPs tunneling as lower transport protocol, which is useful for passing proxies and widely used for security consideration.

)

type RtspFlags = Flag

const (
	RtspFlags_filter_src RtspFlags = "filter_src"
	//Accept packets only from negotiated peer address and port.

	RtspFlags_listen RtspFlags = "listen"
	//Act as a server, listening for an incoming connection.

	RtspFlags_prefer_tcp RtspFlags = "prefer_tcp"
	//Try TCP for RTP transport first, if TCP is available as RTSP RTP transport.

	RtspFlags_satip_raw RtspFlags = "satip_raw"
	//Export raw MPEG-TS stream instead of demuxing. The flag will simply write out the raw stream, with the original PAT/PMT/PIDs intact.
	RtspFlags_none RtspFlags = "none"
	//Default value is ‘none.
)

type AllowedMediaTypes = Flag

const (
	AllowedMediaTypes_video    AllowedMediaTypes = "video"
	AllowedMediaTypes_audio    AllowedMediaTypes = "audio"
	AllowedMediaTypes_data     AllowedMediaTypes = "data"
	AllowedMediaTypes_subtitle AllowedMediaTypes = "subtitle"
)
