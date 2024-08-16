package protocolx

import (
	"github.com/elizabevil/ffmpegx/paramx/parsex"
	"github.com/elizabevil/ffmpegx/paramx/typex"
	"net/url"
)

// 24.43 udp
// User Datagram Protocol.
// The required syntax for an UDP URL is:
// udp://hostname:port[?options]
type UDP struct {
	BufferSize typex.Size `json:"buffer_size" flag:"-buffer_size"` //  =size
	//Set         the UDP maximum socket buffer size in bytes.This is used to set either the receive or send buffer size, depending on what the socket is used for.Default is 32 KB for output, 384 KB for input.See also fifo_size.

	Bitrate typex.Bitrate `json:"bitrate" flag:"-bitrate"` //  =bitrate
	//If set to nonzero, the output will have the specified constant bitrate if the input has enough packets to sustain it.

	BurstBits typex.Size `json:"burst_bits" flag:"-burst_bits"` //  =bits
	//When using bitrate this specifies the maximum number of bits in packet bursts.

	Localport typex.Port `json:"localport" flag:"-localport"` //  =port
	//Override the local UDP port to bind with.

	Localaddr typex.Addr `json:"localaddr" flag:"-localaddr"` //  =addr
	//Local IP address of a network interface used for sending packets or joining multicast groups.

	PktSize typex.Size `json:"pkt_size" flag:"-pkt_size"` //  =size
	//Set the size in bytes of UDP packets.

	Reuse typex.Bytes `json:"reuse" flag:"-reuse"` //  =1|0
	//Explicitly allow or disallow reusing UDP sockets.

	Ttl typex.TTL `json:"ttl" flag:"-ttl"` //  =ttl
	//Set the time to live value (for multicast only).

	Connect typex.Bool `json:"connect" flag:"-connect"` //  =1|0
	//Initialize the UDP socket with connect().In this case, the destination address can’t be changed with ff_udp_set_remote_url later.If the destination address isn’t known at the start, this option can be specified in ff_udp_set_remote_url, too.This allows finding out the source address for the packets with getsockname, and makes writes return with AVERROR(ECONNREFUSED) if "destination unreachable" is received.For receiving, this gives the benefit of only receiving packets from the specified peer address/port.

	Sources string `json:"sources" flag:"-sources"` //  =address[,address]
	//Only receive packets sent from the specified addresses.In case of multicast, also subscribe to multicast traffic coming from these addresses only.

	Block string `json:"block" flag:"-block"` //  =address[,address]
	//Ignore packets sent from the specified addresses.In case of multicast, also exclude the source addresses in the multicast subscription.

	FifoSize typex.Size `json:"fifo_size" flag:"-fifo_size"` //  =units
	//Set the UDP receiving circular buffer size, expressed as a number of packets with size of 188 bytes.If not specified defaults to 7*4096.

	OverrunNonfatal typex.Bool `json:"overrun_nonfatal" flag:"-overrun_nonfatal"` //  =1|0
	//Survive in case of UDP receiving circular buffer overrun.Default value is 0.

	Timeout typex.MicrosecondI `json:"timeout" flag:"-timeout"` //  =microseconds
	//Set raise error timeout, expressed in microseconds.

	//This option is only relevant in read mode: if no data arrived in more than this time interval, raise error.

	Broadcast typex.Bool `json:"broadcast" flag:"-broadcast"` //  =1|0
	//Explicitly allow or disallow UDP broadcasting.

	//Note that broadcasting may not work properly on networks having a broadcast storm protection.
}

/*
Use ffmpeg to stream over UDP to a remote endpoint:
	ffmpeg -i input -f format udp://hostname:port
Use ffmpeg to stream in mpegts format over UDP using 188 sized UDP packets, using a large input buffer:
	ffmpeg -i input -f mpegts udp://hostname:port?pkt_size=188&buffer_size=65535
Use ffmpeg to receive over UDP from a remote endpoint:
	ffmpeg -i udp://[multicast-address]:port ...



*/

func (r UDP) Options() (string, error) {
	return parsex.Options(r)
}
func (r UDP) Url(destination string) (url.URL, error) {
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
