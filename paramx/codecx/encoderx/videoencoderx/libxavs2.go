package videoencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

/*
Libxavs2 16.18
xavs2 AVS2-P2/IEEE1857.4 encoder wrapper.
This encoder requires the presence of the libxavs2 headers and library during configuration. You need to explicitly configure the build with --enable-libxavs2.
The following standard libavcodec options are used:
b / bit_rate
g / gop_size
bf / max_b_frames
The encoder also has its own specific options:
*/
type Libxavs2 struct {
	LcuRowThreads typex.Level `json:"lcu_row_threads" flag:"-lcu_row_threads"` //Set the number of parallel threads for rows from 1 to 8 (default 5).

	InitialQp typex.Level `json:"initial_qp" flag:"-initial_qp"` //Set the xavs2 quantization parameter from 1 to 63 (default 34). This is used to set the initial qp for the first frame.

	Qp typex.Level `json:"qp" flag:"-qp"` //Set the xavs2 quantization parameter from 1 to 63 (default 34). This is used to set the qp value under constant-QP mode.

	MaxQp typex.Level `json:"max_qp" flag:"-max_qp"` //Set the max qp for rate control from 1 to 63 (default 55).

	MinQp typex.Level `json:"min_qp" flag:"-min_qp"` //Set the min qp for rate control from 1 to 63 (default 20).

	SpeedLevel typex.Level `json:"speed_level" flag:"-speed_level"` //Set the Speed level from 0 to 9 (default 0). Higher is better but slower.

	LogLevel typex.Level `json:"log_level" flag:"-log_level"` //Set the log level from -1 to 3 (default 0). -1: none, 0: error, 1: warning, 2: info, 3: debug.

	Xavs2Params string `json:"xavs2-params" flag:"-xavs2-params"` //Set xavs2 options using a list of key=value couples separated by ":".
}

//ffmpeg -i input -c:v libxavs2 -xavs2-params RdoqLevel=0 output.avs2
