package audioencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

//import "github.com/elizabevil/ffmpegx/paramx/typex"

// // LIBVORBIS 8.13 libvorbis
type LIBVORBIS struct {
	B typex.Bitrate `json:"b" flag:"-b"`
	//Set bit rate in bits/s.Setting this automatically activates constant bit rate (CBR) mode.If this option is unspecified it is set to 128kbps.

	Q typex.Quality `json:"q" flag:"-q"`
	//Set constant quality setting for UVBR. The value should be a float number in the range of -1.0 to 10.0. The higher the value, the better the quality. The default value is ‘3.0’.
	Cutoff typex.Cutoff `json:"cutoff" flag:"-cutoff"` //(--advanced-encode-option lowpass_frequency=N)
	//Set cutoff bandwidth in Hz, a value of 0 disables cutoff. oggenc’s related option is expressed in kHz. The default value is ‘0’ (cutoff disabled).

	Minrate typex.BitrateInt `json:"minrate" flag:"-m"` //(-m)
	//Set minimum bitrate expressed in bits/s. oggenc -m is expressed in kilobits/s.

	Maxrate typex.BitrateInt `json:"maxrate" flag:"-M"` //(-M)
	//Set maximum bitrate expressed in bits/s. oggenc -M is expressed in kilobits/s. This only has effect on ABR mode.

	Iblock typex.Iblock `json:"iblock" flag:"-iblock"` //(--advanced-encode-option impulse_noisetune=N)
	//Set noise floor bias for impulse blocks. The value is a float number from -15.0 to 0.0. A negative bias instructs the encoder to pay special attention to the crispness of transients in the encoded audio. The tradeoff for better transient response is a higher bitrate.
}
