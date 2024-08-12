package protocolx

import (
	"github.com/elizabevil/ffmpegx/paramx/parsex"
	"github.com/elizabevil/ffmpegx/paramx/typex"
	"net/url"
)

// 24.41 tcp
// Transmission Control Protocol.
// tcp://hostname:port[?options]
type TCP struct {
	Listen typex.Bool `json:"listen" flag:"-listen"` //    =2|1|0
	//Listen for an incoming connection. 0 disables listen, 1 enables listen in single client mode, 2 enables listen in multi-client mode. Default value is 0.

	LocalAddr typex.Addr `json:"local_addr" flag:"-local_addr"` //    =addr
	//Local IP address of a network interface used for tcp socket connect.

	LocalPort typex.Port `json:"local_port" flag:"-local_port"` //    =port
	//Local port used for tcp socket connect.

	Timeout typex.MicrosecondI `json:"timeout" flag:"-timeout"` //    =microseconds
	//Set raise error timeout, expressed in microseconds.
	//This option is only relevant in read mode: if no data arrived in more than this time interval, raise error.

	ListenTimeout typex.MillisecondI `json:"listen_timeout" flag:"-listen_timeout"` //    =milliseconds
	//Set listen timeout, expressed in milliseconds.

	RecvBufferSize typex.Size `json:"recv_buffer_size" flag:"-recv_buffer_size"` //    =bytes
	//Set receive buffer size, expressed bytes.

	SendBufferSize typex.Size `json:"send_buffer_size" flag:"-send_buffer_size"` //    =bytes
	//Set send buffer size, expressed bytes.

	TcpNodelay typex.Bool `json:"tcp_nodelay" flag:"-tcp_nodelay"` //    =1|0
	//Set TCP_NODELAY to disable Nagle’s algorithm. Default value is 0.

	//Remark: Writing to the socket is currently not optimized to minimize system calls and reduces the efficiency / effect of TCP_NODELAY.

	TcpMss typex.Size `json:"tcp_mss" flag:"-tcp_mss"` //    =bytes
	//Set maximum segment size for outgoing TCP packets, expressed in bytes.
}

func (r TCP) Options() (string, error) {
	return parsex.Options(r)
}
func (r TCP) Url(destination string) (url.URL, error) {
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
