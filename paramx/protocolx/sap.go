package protocolx

import (
	"github.com/elizabevil/ffmpegx/paramx/parsex"
	"net/url"
)

//Session Announcement Protocol (RFC 2974). This is not technically a protocol handler in libavformat, it is a muxer and demuxer. It is used for signalling of RTP streams, by announcing the SDP for the streams regularly on a separate port.
//sap://destination[:port][?options]

type Sap struct {
	AnnounceAddr string `json:"announce_addr"` // address
	//Specify the destination IP address for sending the announcements to. If omitted, the announcements are sent to the commonly used SAP announcement multicast address 224.2.127.254 (sap.mcast.net), or ff0e::2:7ffe if destination is an IPv6 address.

	AnnouncePort int `json:"announce_port"` // port
	//Specify the port to send the announcements on, defaults to 9875 if not specified.

	Ttl int `json:"ttl"` // ttl
	//Specify the time to live value for the announcements and RTP packets, defaults to 255.

	SamePort byte `json:"same_port"` // 0|1
	//If set to 1, send all RTP streams on the same port pair.
	//If zero (the default), all streams are sent on unique ports, with each stream on a port 2 numbers higher than the previous. VLC/Live555 requires this to be set to 1, to be able to receive the stream. The RTP stack in libavformat for receiving requires all streams to be sent on unique ports.
}

func (r Sap) Options() (string, error) {
	return parsex.Options(r)
}
func (r Sap) Url(destination string) (url.URL, error) {
	//sap://destination[:port][?options]
	//ffmpeg -re -i input -f sap sap://224.0.0.255?same_port=1
	options, err := r.Options()
	if err != nil {
		return url.URL{}, err
	}
	u := url.URL{
		Scheme:   "sap",
		Host:     destination,
		RawQuery: options,
	}
	return u, nil
}
