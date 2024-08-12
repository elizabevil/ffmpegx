package videoencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// LIBTHEORA 9.12 libtheora
type LIBTHEORA struct {
	B typex.Bitrate `json:"b" flag:"-b"`

	//Set the video bitrate in bit/s for CBR (Constant Bit Rate) mode. In case UVBR (Variable Bit Rate) mode is enabled this option is ignored.

	Flags typex.Bool `json:"flags" flag:"-flags"`
	//Used to enable constant quality mode (UVBR) encoding through the qscale flag, and to enable the pass1 and pass2 modes.

	G typex.Size `json:"g" flag:"-g"`
	//Set the GOP size.

	GlobalQuality typex.Number `json:"global_quality" flag:"-global_quality"`
	//Set the global quality as an integer in lambda units.

	//Only relevant when UVBR mode is enabled with flags +qscale. The value is converted to QP units by dividing it by FF_QP2LAMBDA, clipped in the [0 - 10] range, and then multiplied by 6.3 to get a value in the native libtheora range [0-63]. A higher value corresponds to a higher quality.

	Q typex.Level `json:"q" flag:"-q"`
	//Enable UVBR mode when set to a non-negative value, and set constant quality value as a double floating point value in QP units.

	//The value is clipped in the [0-10] range, and then multiplied by 6.3 to get a value in the native libtheora range [0-63].

	//This option is valid only using the ffmpeg command-line tool. For library interface users, use global_quality.
}

/*
Set maximum constant quality (UVBR) encoding with ffmpeg:
ffmpeg -i INPUT -codec:v libtheora -q:v 10 OUTPUT.ogg
Use ffmpeg to convert a CBR 1000 kbps Theora video stream:
ffmpeg -i INPUT -codec:v libtheora -b:v 1000k OUTPUT.ogg
*/
