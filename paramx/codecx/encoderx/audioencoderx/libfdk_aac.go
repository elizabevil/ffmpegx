package audioencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// // LIBFDK_AAC 8.5 libfdk_aac
type LIBFDK_AAC struct {
	B typex.Bitrate `json:"b" flag:"-b"`
	//Set bit rate in bits/s.If the bitrate is not explicitly specified, it is automatically set to a suitable value depending on the selected profile.

	//In case UVBR mode is enabled the option is ignored.

	Ar typex.SampleRate `json:"ar" flag:"-ar"`
	//Set audio sampling rate (in Hz).

	Channels typex.Channels `json:"channels" flag:"-channels"`
	//Set the number of audio channels.

	Flags typex.Bool `json:"flags" flag:"-flags"`
	//Enable fixed quality, UVBR (Variable Bit Rate) mode.Note that UVBR is implicitly enabled when the vbr value is positive.

	Cutoff typex.Frequency `json:"cutoff" flag:"-cutoff"`
	//Set cutoff frequency.If not specified (or explicitly set to 0) it will use a value automatically computed by the library.Default value is 0.

	Profile AacProfile `json:"profile" flag:"-profile"`
	//Set audio profile.
	//The following are private options of the libfdk_aac encoder.
	Afterburner *typex.UBool `json:"afterburner" flag:"-afterburner"`
	//Enable afterburner feature if set to 1, disabled if set to 0. This improves the quality but also the required processing power.

	//Default value is 1.

	EldSbr typex.UBool `json:"eld_sbr" flag:"-eld_sbr"`
	//Enable SBR (Spectral Band Replication) for ELD if set to 1, disabled if set to 0.

	//Default value is 0.

	EldV2 typex.UBool `json:"eld_v2" flag:"-eld_v2"`
	//Enable ELDv2 (LD-MPS extension for ELD stereo signals) for ELDv2 if set to 1, disabled if set to 0.
	Signaling Signaling `json:"signaling" flag:"-signaling"`
	//Set SBR/PS signaling style.
	Latm typex.UBool `json:"latm" flag:"-latm"`
	//Output LATM/LOAS encapsulated data if set to 1, disabled if set to 0.

	//Default value is 0.
	HeaderPeriod typex.UI16 `json:"header_period" flag:"-header_period"`
	//Set StreamMuxConfig and PCE repetition period (in frames) for sending in-band configuration buffers within LATM/LOAS transport layer.

	//Must be a 16-bits non-negative integer.

	//Default value is 0.
	Vbr typex.Level `json:"vbr" flag:"-vbr"`
	//Set UVBR mode, from 1 to 5. 1 is lowest quality (though still pretty good) and 5 is highest quality. A value of 0 will disable UVBR, and CBR (Constant Bit Rate) is enabled.

	FrameLength typex.Level `json:"frame_length" flag:"-frame_length"`
	//Set the audio frame length in samples. Default value is the internal default of the library. Refer to the library’s documentation for information about supported values.
}

type AacProfile string
type Signaling string

const (
	AacProfile_aac_low AacProfile = "aac_low"

	AacProfile_aac_he AacProfile = "aac_he"

	AacProfile_aac_he_v2 AacProfile = "aac_he_v2"

	AacProfile_aac_ld  AacProfile = "aac_ld"
	AacProfile_aac_eld AacProfile = "aac_eld"
)

const (
	Signaling_default Signaling = "default"
	//choose signaling implicitly (explicit hierarchical by default, implicit if global header is disabled)

	Signaling_implicit Signaling = "implicit"
	//implicit backwards compatible signaling

	Signaling_explicit_sbr Signaling = "explicit_sbr"
	//explicit SBR, implicit PS signaling

	Signaling_explicit_hierarchical Signaling = "explicit_hierarchical"
	//explicit hierarchical signaling
)

/*

8.5.2 Examples
Use ffmpeg to convert an audio file to UVBR AAC in an M4A (MP4) container:
ffmpeg -i input.wav -codec:a libfdk_aac -vbr 3 output.m4a
Use ffmpeg to convert an audio file to CBR 64k kbps AAC, using the High-Efficiency AAC profile:
ffmpeg -i input.wav -c:a libfdk_aac -profile:a aac_he -b:a 64k output.m4a
*/
