package audioencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// // AAC 8.1 aac
type AAC struct {
	B typex.Bitrate `json:"b" flag:"-b"`
	//Set bit rate in bits/s.Setting this automatically activates constant bit rate (CBR) mode.If this option is unspecified it is set to 128kbps.

	Q typex.Quality `json:"q" flag:"-q"`
	//Set quality for variable bit rate (UVBR) mode.This option is valid only using the ffmpeg command-line tool.For library interface users, use global_quality.

	Cutoff typex.Cutoff `json:"cutoff" flag:"-cutoff"`
	//Set cutoff frequency.If unspecified will allow the encoder to dynamically adjust the cutoff to improve clarity on low bitrates.

	AccCoder AccCoder `json:"aac_coder" flag:"-aac_coder"`
	//Set AAC encoder coding method.

	AacMs typex.Bool `json:"aac_ms" flag:"-aac_ms"`
	//Sets mid/side coding mode. The default value of "auto" will automatically use M/S with bands which will benefit from such coding. Can be forced for all bands using the value "enable", which is mainly useful for debugging or disabled using "disable".

	AacIs typex.Bool `json:"aac_is" flag:"-aac_is"`
	//Sets intensity stereo coding tool usage. By default, it’s enabled and will automatically toggle IS for similar pairs of stereo bands if it’s beneficial. Can be disabled for debugging by setting the value to "disable".

	AacPns typex.Bool `json:"aac_pns" flag:"-aac_pns"`
	//Uses perceptual noise substitution to replace low entropy high frequency bands with imperceptible white noise during the decoding process. By default, it’s enabled, but can be disabled for debugging purposes by using "disable".

	AacTns typex.Bool `json:"aac_tns" flag:"-aac_tns"`
	//Enables the use of a multitap FIR filter which spans through the high frequency bands to hide quantization noise during the encoding process and is reverted by the decoder. As well as decreasing unpleasant artifacts in the high range this also reduces the entropy in the high bands and allows for more bits to be used by the mid-low bands. By default it’s enabled but can be disabled for debugging by setting the option to "disable".

	AacLtp typex.Bool `json:"aac_ltp" flag:"-aac_ltp"`
	//Enables the use of the long term prediction extension which increases coding efficiency in very low bandwidth situations such as encoding of voice or solo piano music by extending constant harmonic peaks in bands throughout frames. This option is implied by Profile:a aac_low and is incompatible with aac_pred. Use in conjunction with -ar to decrease the samplerate.

	AacPred typex.Bool `json:"aac_pred" flag:"-aac_pred"`
	//Enables the use of a more traditional style of prediction where the spectral coefficients transmitted are replaced by the difference of the current coefficients minus the previous "predicted" coefficients. In theory and sometimes in practice this can improve quality for low to mid bitrate audio. This option implies the aac_main Profile and is incompatible with aac_ltp.

	Profile Profile `json:"profile" flag:"-profile"`
}

type Profile = typex.String
type AccCoder = typex.String

const (
	Profile_aac_low Profile = "aac_low"
	//The default, AAC "Low-complexity" Profile. Is the most compatible and produces decent quality.

	Profile_mpeg2_aac_low Profile = "mpeg2_aac_low"
	//Equivalent to -Profile:a aac_low -aac_pns 0. PNS was introduced with the MPEG4 specifications.

	Profile_aac_ltp Profile = "aac_ltp"
	//Long term prediction Profile, is enabled by and will enable the aac_ltp option. Introduced in MPEG4.

	Profile_aac_main Profile = "aac_main"
	//Main-type prediction Profile, is enabled by and will enable the aac_pred option. Introduced in MPEG2.
)

const (
	AccCoder_twoloop AccCoder = "twoloop"
	//Two loop searching (TLS) method. This is the default method.

	//This method first sets quantizers depending on band thresholds and then tries to find an optimal combination by adding or subtracting a specific value from all quantizers and adjusting some individual quantizer a little. Will tune itself based on whether aac_is, aac_ms and aac_pns are enabled.

	AccCoder_anmr AccCoder = "anmr"
	//Average noise to mask ratio (ANMR) trellis-based solution.

	//This is an experimental coder which currently produces a lower quality, is more unstable and is slower than the default twoloop coder but has potential. Currently has no support for the aac_is or aac_pns options. Not currently recommended.

	AccCoder_fast AccCoder = "fast"
	//Constant quantizer method.

	//Uses a cheaper version of twoloop algorithm that doesn’t try to do as many clever adjustments. Worse with low bitrates (less than 64kbps), but is better and much faster at higher bitrates.
)
