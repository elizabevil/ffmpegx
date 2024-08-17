package protocolx

import (
	"github.com/elizabevil/ffmpegx/paramx/parsex"
	"github.com/elizabevil/ffmpegx/paramx/typex"
	"net/url"
)

// 24.33 rtp
// RTP The required syntax for an RTP URL is: rtp://hostname[:port][?option=val...]
type RTP struct {
	Ttl typex.UI8 `json:"ttl"` // =n
	//Set the TTL (Time-To-Live) value (for multicast only).

	RtcpPort typex.Port `json:"rtcpport"` // =n
	//Set the remote RTCP port to n.

	LocalRtpPort typex.Port `json:"localrtpport"` // =n
	//Set the local RTP port to n.

	LocalRtcpPort typex.Port `json:"localrtcpport"` // =n'
	//Set the local RTCP port to n.

	PktSize typex.Size `json:"pkt_size"` // =n
	//Set max packet size (in bytes) to n.

	BufferSize typex.Size `json:"buffer_size"`
	//Set the maximum UDP socket buffer size in bytes.

	Connect typex.Bool `json:"connect"` // = 0|1
	//Do a connect() on the UDP socket (if set to 1) or not (if set to 0).

	Sources string `json:"sources"` // ip[, ip]
	//List allowed source IP addresses.

	Block string `json:"block"` //= ip[, ip]
	//List disallowed (blocked) source IP addresses.

	WriteToSource typex.UBool `json:"write_to_source"` //= 0|1
	//Send packets to the source address of the latest received packet (if set to 1) or to a default remote address (if set to 0).

	LocalPort typex.Port `json:"localport"` // =n
	//Set the local RTP port to n.

	LocalAddr typex.Addr `json:"localaddr"` //= addr
	//Local IP address of a network interface used for sending packets or joining multicast groups.

	Timeout typex.MillisecondI `json:"timeout"` // =n
	//Set timeout (in microseconds) of socket I/O operations to n.

	//This is a deprecated option.Instead, localrtpport should be used.
}

func (r RTP) Options() (string, error) {
	return parsex.Options(r)
}

func (r RTP) Url(hostname string) (url.URL, error) {
	//rtp://hostname[:port][?option=val...]
	options, err := r.Options()
	if err != nil {
		return url.URL{}, err
	}
	u := url.URL{
		Scheme:   "rtp",
		Host:     hostname,
		RawQuery: options,
	}
	return u, nil
}
