package audioencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// LIBTWOLAME 8.11 libtwolame
type LIBTWOLAME struct {
	B typex.Bitrate `json:"b" flag:"-b"`
	//Set bit rate in bits/s.Setting this automatically activates constant bit rate (CBR) mode.If this option is unspecified it is set to 128kbps.

	Q    typex.Quality `json:"q" flag:"-q"`
	Mode Mode          `json:"mode" flag:"-mode"`
	//Set quality for variable bit rate (UVBR) mode.This option is valid only using the ffmpeg command-line tool.For library interface users, use global_quality.

	Psymodel typex.Bool `json:"psymodel" flag:"-psymodel"` // (--psyc-mode)
	//Set psychoacoustic model to use in encoding. The argument must be an integer between -1 and 4, inclusive. The higher the value, the better the quality. The default value is 3.

	EnergyLevels typex.Bool `json:"energy_levels" flag:"-energy_levels"` // (--energy)
	//Enable energy levels extensions when set to 1. The default value is 0 (disabled).

	ErrorProtection typex.Bool `json:"error_protection" flag:"-error_protection"` // (--protect)
	//Enable CRC error protection when set to 1. The default value is 0 (disabled).

	Copyright typex.Bool `json:"copyright" flag:"-copyright"` // (--copyright)
	//Set MPEG audio copyright flag when set to 1. The default value is 0 (disabled).

	Original typex.Bool `json:"original" flag:"-original"` // (--original)
	//Set MPEG audio original flag when set to 1. The default value is 0 (disabled).

}

type Mode string

const (
	Mode_auto Mode = "auto"
	//Choose mode automatically based on the input. This is the default.

	Mode_stereo Mode = "stereo"
	//Stereo

	Mode_joint_stereo Mode = "joint_stereo"
	//Joint stereo

	Mode_dual_channel Mode = "dual_channel"
	//Dual channel

	Mode_mono Mode = "mono"
	//Mono
)
