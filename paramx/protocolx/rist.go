package protocolx

import "github.com/elizabevil/ffmpegx/paramx/typex"

type RIST struct {
	RistProfile RistProfile `json:"RistProfile" flag:"-RistProfile"`
	//

	BufferSize typex.Size `json:"buffer_size" flag:"-buffer_size"`
	//Set internal RIST buffer size in milliseconds for retransmission of data. Default value is 0 which means the librist default (1 sec). Maximum value is 30 seconds.

	FifoSize typex.UI16 `json:"fifo_size" flag:"-fifo_size"`
	//Size of the librist receiver output fifo in number of packets. This must be a power of 2. Defaults to 8192 (vs the librist default of 1024).

	OverrunNonfatal typex.UBool `json:"overrun_nonfatal" flag:"-overrun_nonfatal"`
	//Survive in case of librist fifo buffer overrun. Default value is 0.

	PktSize typex.Size `json:"pkt_size" flag:"-pkt_size"`
	//Set maximum packet size for sending data. 1316 by default.

	LogLevel typex.UNumber `json:"log_level" flag:"-log_level"`
	//Set loglevel for RIST logging messages. You only need to set this if you explicitly want to enable debug level messages or packet loss simulation, otherwise the regular loglevel is respected.

	Secret string `json:"secret" flag:"-secret"`
	//Set override of encryption secret, by default is unset.

	Encryption typex.UNumber `json:"encryption" flag:"-encryption"`
	//Set encryption type, by default is disabled. Acceptable values are 128 and 256.
}
type RistProfile string

const (
	RistProfile_simple RistProfile = "simple"
	RistProfile_main   RistProfile = "main"
	//This one is default.
	RistProfile_advanced RistProfile = "advanced"
)
