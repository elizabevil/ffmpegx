package protocolx

import (
	"github.com/elizabevil/ffmpegx/paramx/parsex"
	"github.com/elizabevil/ffmpegx/paramx/typex"
	"net/url"
)

// srt://hostname:port[?options]
// or options srt://hostname:port
type SrtMode string

const (
	SrtMode_caller     SrtMode = "caller"
	SrtMode_listener   SrtMode = "listener"
	SrtMode_rendezvous SrtMode = "rendezvous"
)

type Srt struct {
	ConnectTimeout typex.MillisecondI `json:"connect_timeout"` // =milliseconds
	//Connection timeout; SRT cannot connect for RTT > 1500 msec (2 handshake exchanges) with the default connect timeout of 3 seconds. This option applies to the caller and rendezvous connection modes. The connect timeout is 10 times the value set for the rendezvous mode (which can be used as a workaround for this connection problem with earlier versions).

	Ffs typex.Bytes `json:"ffs"` // =bytes
	//Flight Flag Size (Window Size), in bytes. FFS is actually an internal parameter and you should set it to not less than recv_buffer_size and mss. The default value is relatively large, therefore unless you set a very large receiver buffer, you do not need to change this option. Default value is 25600.

	Inputbw typex.Bytes `json:"inputbw"` // =bytes/seconds
	//Sender nominal input rate, in bytes per seconds. Used along with oheadbw, when maxbw is set to relative (0), to calculate maximum sending rate when recovery packets are sent along with the main media stream: inputbw * (100 + oheadbw) / 100 if inputbw is not set while maxbw is set to relative (0), the actual input rate is evaluated inside the library. Default value is 0.

	Iptos typex.IPTypeOf `json:"iptos"` // =tos
	//IP Type of Service. Applies to sender only. Default value is 0xB8.

	Ipttl typex.IPTypeOf `json:"ipttl"` // =ttl
	//IP Time To Live. Applies to sender only. Default value is 64.

	Latency typex.MicrosecondI `json:"latency"` // =microseconds
	//Timestamp-based Packet Delivery Delay. Used to absorb bursts of missed packet retransmissions. This flag sets both rcvlatency and peerlatency to the same value. Note that prior to version 1.3.0 this is the only flag to set the latency, however this is effectively equivalent to setting peerlatency, when side is sender and rcvlatency when side is receiver, and the bidirectional stream sending is not supported.

	ListenTimeout typex.MicrosecondI `json:"listen_timeout"` // =microseconds
	//Set socket listen timeout.

	Maxbw typex.Bytes `json:"maxbw"` // =bytes/seconds
	//Maximum sending bandwidth, in bytes per seconds. -1 infinite (CSRTCC limit is 30mbps) 0 relative to input rate (see inputbw) >0 absolute limit value Default value is 0 (relative)

	Mode SrtMode `json:"mode"` // =caller|listener|rendezvous
	//Connection mode. caller opens client connection. listener starts server to listen for incoming connections. rendezvous use Rendez-Vous connection mode. Default value is caller.

	Mss typex.Bytes `json:"mss"` // =bytes
	//Maximum Segment Size, in bytes. Used for buffer allocation and rate calculation using a packet counter assuming fully filled packets. The smallest MSS between the peers is used. This is 1500 by default in the overall internet. This is the maximum size of the UDP packet and can be only decreased, unless you have some unusual dedicated network settings. Default value is 1500.

	Nakreport typex.Bool `json:"nakreport"` // =1|0
	//If set to 1, Receiver will send ‘UMSG_LOSSREPORT‘ messages periodically until a lost packet is retransmitted or intentionally dropped. Default value is 1.

	Oheadbw typex.PercentI `json:"oheadbw"` // =percents
	//Recovery bandwidth overhead above input rate, in percents. See inputbw. Default value is 25%.

	Passphrase string `json:"passphrase"` // =string
	//HaiCrypt Encryption/Decryption Passphrase string, length from 10 to 79 characters. The passphrase is the shared secret between the sender and the receiver. It is used to generate the Key Encrypting Key using PBKDF2 (Password-Based Key Derivation Function). It is used only if pbkeylen is non-zero. It is used on the receiver only if the received data is encrypted. The configured passphrase cannot be recovered (write-only).

	EnforcedEncryption typex.Bool `json:"enforced_encryption"` // =1|0
	//If true, both connection parties must have the same password set (including empty, that is, with no encryption). If the password doesn’t match or only one side is unencrypted, the connection is rejected. Default is true.

	Kmrefreshrate typex.Packets `json:"kmrefreshrate"` // =packets
	//The number of packets to be transmitted after which the encryption key is switched to a new key. Default is -1. -1 means auto (0x1000000 in srt library). The range for this option is integers in the 0 - INT_MAX.

	Kmpreannounce typex.Packets `json:"kmpreannounce"` // =packets
	//The interval between when a new encryption key is sent and when switchover occurs. This value also applies to the subsequent interval between when switchover occurs and when the old encryption key is decommissioned. Default is -1. -1 means auto (0x1000 in srt library). The range for this option is integers in the 0 - INT_MAX.

	Snddropdelay typex.MillisecondI `json:"snddropdelay"` // =microseconds
	//The sender’s extra delay before dropping packets. This delay is added to the default drop delay time interval value.
	//Special value -1: Do not drop packets on the sender at all.

	PayloadSize typex.Bytes `json:"payload_size"` //
	//Sets the maximum declared size of a packet transferred during the single call to the sending function in Live mode. Use 0 if this value isn’t used (which is default in file mode). Default is -1 (automatic), which typically means MPEG-TS; if you are going to use SRT to send any different kind of payload, such as, for example, wrapping a live stream in very small frames, then you can use a bigger maximum frame size, though not greater than 1456 bytes.

	PktSize typex.Bytes `json:"pkt_size"` // =bytes
	//Alias for ‘payload_size’.

	Peerlatency typex.MicrosecondI `json:"peerlatency"` // =microseconds
	//The latency value (as described in rcvlatency) that is set by the sender side as a minimum value for the receiver.

	Pbkeylen typex.Bytes `json:"pbkeylen"` // =bytes
	//Sender encryption key length, in bytes. Only can be set to 0, 16, 24 and 32. Enable sender encryption if not 0. Not required on receiver (set to 0), key size obtained from sender in HaiCrypt handshake. Default value is 0.

	Rcvlatency typex.MicrosecondI `json:"rcvlatency"` // =microseconds
	//The time that should elapse since the moment when the packet was sent and the moment when it’s delivered to the receiver application in the receiving function. This time should be a buffer time large enough to cover the time spent for sending, unexpectedly extended RTT time, and the time needed to retransmit the lost UDP packet. The effective latency value will be the maximum of this options’ value and the value of peerlatency set by the peer side. Before version 1.3.0 this option is only available as latency.

	RecvBufferSize typex.Bytes `json:"recv_buffer_size"` // =bytes
	//Set UDP receive buffer size, expressed in bytes.

	SendBufferSize typex.Bytes `json:"send_buffer_size"` // =bytes
	//Set UDP send buffer size, expressed in bytes.

	Timeout typex.MicrosecondI `json:"timeout"` // =microseconds
	//Set raise error timeouts for read, write and connect operations. Note that the SRT library has internal timeouts which can be controlled separately, the value set here is only a cap on those.

	Tlpktdrop typex.Bool `json:"tlpktdrop"` // =1|0
	//Too-late Packet Drop. When enabled on receiver, it skips missing packets that have not been delivered in time and delivers the following packets to the application when their time-to-play has come. It also sends a fake ACK to the sender. When enabled on sender and enabled on the receiving peer, the sender drops the older packets that have no chance of being delivered in time. It was automatically enabled in the sender if the receiver supports it.

	Sndbuf typex.Bytes `json:"sndbuf"` // =bytes
	//Set send buffer size, expressed in bytes.

	Rcvbuf typex.Bytes `json:"rcvbuf"` // =bytes
	//Set receive buffer size, expressed in bytes.

	//Receive buffer must not be greater than ffs.

	Lossmaxttl typex.Packets `json:"lossmaxttl"` // =packets
	//The value up to which the Reorder Tolerance may grow. When Reorder Tolerance is > 0, then packet loss report is delayed until that number of packets come in. Reorder Tolerance increases every time a "belated" packet has come, but it wasn’t due to retransmission (that is, when UDP packets tend to come out of order), with the difference between the latest sequence and this packet’s sequence, and not more than the value of this option. By default it’s 0, which means that this mechanism is turned off, and the loss report is always sent immediately upon experiencing a "gap" in sequences.

	//minversion
	//The minimum SRT version that is required from the peer. A connection to a peer that does not satisfy the minimum version requirement will be rejected.

	//The version format in hex is 0xXXYYZZ for x.y.z in human readable form.

	Streamid string `json:"streamid"` // =string
	//A string limited to 512 characters that can be set on the socket prior to connecting. This stream ID will be able to be retrieved by the listener side from the socket that is returned from srt_accept and was connected by a socket with that set stream ID. SRT does not enforce any special interpretation of the contents of this string. This option doesn’t make sense in Rendezvous connection; the result might be that simply one side will override the value from the other side and it’s the matter of luck which one would win

	SrtStreamid string `json:"srt_streamid"` // =string
	//Alias for ‘streamid’ to avoid conflict with ffmpeg command line option.

	Smoother string `json:"smoother"` // =live|file
	//The type of Smoother used for the transmission for that socket, which is responsible for the transmission and congestion control.
	//The Smoother type must be exactly the same on both connecting parties, otherwise the connection is rejected.

	Messageapi typex.Bool `json:"messageapi"` // =1|0
	//When set, this socket uses the Message API, otherwise it uses Buffer API. Note that in live mode (see transtype) there’s only message API available. In File mode you can chose to use one of two modes:

	//Stream API (default, when this option is false). In this mode you may send as many data as you wish with one sending instruction, or even use dedicated functions that read directly from a file. The internal facility will take care of any speed and congestion control. When receiving, you can also receive as many data as desired, the data not extracted will be waiting for the next call. There is no boundary between data portions in the Stream mode.

	//Message API. In this mode your single sending instruction passes exactly one piece of data that has boundaries (a message). Contrary to Live mode, this message may span across multiple UDP packets and the only size limitation is that it shall fit as a whole in the sending buffer. The receiver shall use as large buffer as necessary to receive the message, otherwise the message will not be given up. When the message is not complete (not all packets received or there was a packet loss) it will not be given up.

	Transtype string `json:"transtype"` // =live|file
	//Sets the transmission type for the socket, in particular, setting this option sets multiple other parameters to their default values as required for a particular transmission type.

	//live: Set options as for live transmission. In this mode, you should send by one sending instruction only so many data that fit in one UDP packet, and limited to the value defined first in payload_size (1316 is default in this mode). There is no speed control in this mode, only the bandwidth control, if configured, in order to not exceed the bandwidth with the overhead transmission (retransmitted and control packets).

	//file: Set options as for non-live transmission. See messageapi for further explanations

	Linger typex.SecondI `json:"linger"` // =seconds
	//The number of seconds that the socket waits for unsent data when closing. Default is -1. -1 means auto (off with 0 seconds in live mode, on with 180 seconds in file mode). The range for this option is integers in the 0 - INT_MAX.

	Tsbpd typex.Bool `json:"tsbpd"` // =1|0
	//When true, use Timestamp-based Packet Delivery mode. The default behavior depends on the transmission type: enabled in live mode, disabled in file mode.
}

func (r Srt) Options() (string, error) {
	return parsex.Options(r)
}
func (r Srt) Url(hostname string) (url.URL, error) {
	//srt://hostname:port[?options]
	options, err := r.Options()
	if err != nil {
		return url.URL{}, err
	}
	u := url.URL{
		Scheme:   "srt",
		Host:     hostname,
		RawQuery: options,
	}
	return u, nil
}
