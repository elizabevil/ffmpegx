package audioencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// LIBOPUS 8.9 libopus
/*
Most libopus options are modelled after the opusenc utility from opus-tools. The following is an option mapping chart describing options supported by the libopus wrapper, and their opusenc-equivalent in parentheses.
*/
type LIBOPUS struct {
	B typex.Bitrate `json:"b" flag:"-b"` //
	//Set the bit rate in bits/s. FFmpeg’s b option is expressed in bits/s, while opusenc’s bitrate in kilobits/s.

	Vbr              Vbr         `json:"vbr" flag:"-vbr"`                             //(vbr, hard-cbr, and cvbr)
	CompressionLevel typex.Level `json:"compression_level" flag:"-compression_level"` //(comp)
	//Set encoding algorithm complexity. Valid options are integers in the 0-10 range. 0 gives the fastest encodes but lower quality, while 10 gives the highest quality but slowest encoding. The default is 10.

	FrameDuration typex.Duration `json:"frame_duration" flag:"-frame_duration"` //(framesize)
	//Set maximum frame size, or duration of a frame in milliseconds. The argument must be exactly the following: 2.5, 5, 10, 20, 40, 60. Smaller frame sizes achieve lower latency but less quality at a given bitrate. Sizes greater than 20ms are only interesting at fairly low bitrates. The default is 20ms.

	PacketLoss typex.Level `json:"packet_loss" flag:"-packet_loss"` //(expect-loss)
	//Set expected packet loss percentage. The default is 0.

	Fec typex.Bool `json:"fec" flag:"-fec"` //(n/a)
	//Enable inband forward error correction. packet_loss must be non-zero to take advantage - frequency of FEC ’side-data’ is proportional to expected packet loss. Default is disabled.
	Application Application  `json:"application" flag:"-application"` //(N.A.)
	Cutoff      typex.Cutoff `json:"cutoff" flag:"-cutoff"`           //(N.A.)
	//Set cutoff bandwidth in Hz. The argument must be exactly one of the following: 4000, 6000, 8000, 12000, or 20000, corresponding to narrowband, mediumband, wideband, super wideband, and fullband respectively. The default is 0 (cutoff disabled). Note that libopus forces a wideband cutoff for bitrates < 15 kbps, unless CELT-only (application set to ‘lowdelay’) mode is used.

	MappingFamily typex.UI8 `json:"mapping_family" flag:"-mapping_family"` //(mapping_family)
	//Set channel mapping family to be used by the encoder. The default value of -1 uses mapping family 0 for mono and stereo inputs, and mapping family 1 otherwise. The default also disables the surround masking and LFE bandwidth optimzations in libopus, and requires that the input contains 8 channels or fewer.

	//Other values include 0 for mono and stereo, 1 for surround sound with masking and LFE bandwidth optimizations, and 255 for independent streams with an unspecified channel layout.

	ApplyPhaseInv typex.UBool `json:"apply_phase_inv" flag:"-apply_phase_inv"` //(N.A.) (requires libopus >= 1.2)
	//If set to 0, disables the use of phase inversion for intensity stereo, improving the quality of mono downmixes, but slightly reducing normal stereo quality. The default is 1 (phase inversion enabled).

}

type Vbr string
type Application string

const (
	Vbr_off Vbr = "off" // (hard-cbr)
	//Use constant bit rate encoding.

	Vbr_on Vbr = "on" // (vbr)
	//Use variable bit rate encoding (the default).

	Vbr_constrained Vbr = "constrained" // (cvbr)
	//Use constrained variable bit rate encoding.
)

const (
	Application_voip Application = "voip"
	//Favor improved speech intelligibility.

	Application_audio Application = "audio"
	//Favor faithfulness to the input (the default).

	Application_lowdelay Application = "lowdelay"
	//Restrict to only the lowest delay modes by disabling voice-optimized modes.

)
