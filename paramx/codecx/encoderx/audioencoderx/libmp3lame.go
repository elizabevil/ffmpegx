package audioencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// LIBMP3LAME 8.7 libmp3lame
type LIBMP3LAME struct {
	B typex.Bitrate `json:"b" flag:"-b"` //
	Q typex.Quality `json:"q" flag:"-q"` // (-V)
	//Set constant quality setting for UVBR. This option is valid only using the ffmpeg command-line tool. For library interface users, use global_quality.

	CompressionLevel typex.Level `json:"compression_level" flag:"-compression_level"` //(-q)
	//Set algorithm quality. Valid arguments are integers in the 0-9 range, with 0 meaning highest quality but slowest, and 9 meaning fastest while producing the worst quality.

	Cutoff typex.Cutoff `json:"cutoff" flag:"-cutoff"` //(--lowpass)
	//Set lowpass cutoff frequency. If unspecified, the encoder dynamically adjusts the cutoff.

	Reservoir typex.Bitrate `json:"reservoir" flag:"-reservoir"` //
	//Enable use of bit reservoir when set to 1. Default value is 1. LAME has this enabled by default, but can be overridden by use --nores option.

	JointStereo typex.Bool `json:"joint_stereo" flag:"-joint_stereo"` //(-m j)
	//Enable the encoder to use (on a frame by frame basis) either L/R stereo or mid/side stereo. Default value is 1.

	Abr typex.Bool `json:"abr" flag:"-abr"` //(--abr)
	//Enable the encoder to use ABR when set to 1. The lame --abr sets the target bitrate, while this options only tells FFmpeg to use ABR still relies on b to set bitrate.

	Copyright typex.Bool `json:"copyright" flag:"-copyright"` //(-c)
	//Set MPEG audio copyright flag when set to 1. The default value is 0 (disabled).

	Original typex.Bool `json:"original" flag:"-original"` //(-o)
	//Set MPEG audio original flag when set to 1. The default value is 1 (enabled).

}
