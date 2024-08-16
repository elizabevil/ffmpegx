package audioencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// AC3 8.2 ac3 and ac3_fixed
/*The AC-3 metadata options are used to set parameters that describe the audio, but in most cases do not affect the audio encoding itself. Some of the options do directly affect or influence the decoding and playback of the resulting bitstream, while others are just for informational purposes. A few of the options will add bits to the output stream that could otherwise be used for audio data, and will thus affect the quality of the output. Those will be indicated accordingly with a note in the option list below.
 */
type AC3 struct {
	PerFrameMetadata typex.Bool `json:"per_frame_metadata" flag:"-per_frame_metadata"`
	//8.2.1.2 Downmix Levels
	CenterMixlev   typex.FLT0_1 `json:"center_mixlev" flag:"-center_mixlev"`
	SurroundMixlev typex.FLT0_1 `json:"surround_mixlev" flag:"-surround_mixlev"`

	//8.2.1.3 Audio Production Information
	MixingLevel typex.Level `json:"mixing_level" flag:"-mixing_level"`
	/*Mixing Level. Specifies peak sound pressure level (SPL) in the production environment when the mix was mastered. Valid values are 80 to 111, or -1 for unknown or not indicated. The default value is -1, but that value cannot be used if the Audio Production Information is written to the bitstream. Therefore, if the room_type option is not the default value, the mixing_level option must not be -1.*/

	RoomType RoomType `json:"room_type" flag:"-room_type"`
	/*
		0 notindicated
		Not Indicated (default)
		1 large
		Large Room
		2 small
		Small Room
	*/

	//8.2.1.4 Other Metadata Options
	Copyright typex.Level `json:"copyright" flag:"-copyright"`
	Dialnorm  typex.Bool  `json:"dialnorm" flag:"-dialnorm"` //-31, -1,
	DsurMode  DsurMode    `json:"dsur_mode" flag:"-dsur_mode"`
	Original  *typex.Bool `json:"original" flag:"-original"`
}

type DsurMode = typex.String

const (
	DsurMode_notindicated DsurMode = "notindicated"
	DsurMode_large        DsurMode = "on"
	DsurMode_small        DsurMode = "off"
)

type RoomType = typex.String

const (
	RoomType_notindicated RoomType = "notindicated"
	RoomType_large        RoomType = "large"
	RoomType_small        RoomType = "small"
)

type AdConvType = typex.String

const (
	AdConvType_standard AdConvType = "standard"
	AdConvType_hdcd     AdConvType = "hdcd"
)

// AC3ExtendedBitstream 8.2.2 Extended Bitstream Information
type AC3ExtendedBitstream struct {
	DmixMode DmixMode `json:"dmix_mode" flag:"-dmix_mode"`
	//-1.0, 2.0
	LtrtCmixlev   typex.Flt `json:"ltrt_cmixlev" flag:"-ltrt_cmixlev"`
	LtrtSurmixlev typex.Flt `json:"ltrt_surmixlev" flag:"-ltrt_surmixlev"`
	LoroCmixlev   typex.Flt `json:"loro_cmixlev" flag:"-loro_cmixlev"`
	LoroSurmixlev typex.Flt `json:"loro_surmixlev" flag:"-loro_surmixlev"`
	//8.2.2.2 Extended Bitstream Information - Part 2
	DsurexMode     DsurexMode `json:"dsurex_mode" flag:"-dsurex_mode"`
	DheadphoneMode DsurMode   `json:"dheadphone_mode" flag:"-dheadphone_mode"`
	AdConvType     AdConvType `json:"ad_conv_type" flag:"-ad_conv_type"`
}

// AC3Encoding 8.2.3 Other AC-3 Encoding Options
type AC3Encoding struct {
	StereoRematrixing bool         `json:"stereo_rematrixing" flag:"-stereo_rematrixing"`
	Cutoff            typex.Cutoff `json:"cutoff" flag:"-cutoff"`
}

// 8.2.4 Floating-Point-Only AC-3 Encoding Options
type AC3FloatingPointOnly struct {
	ChannelCoupling typex.UBool `json:"channel_coupling" flag:"-channel_coupling"`
	//Coupling Start Band. Sets the channel coupling start band, from 1 to 15. If a value higher than the bandwidth is used, it will be reduced to 1 less than the coupling end band. If auto is used, the start band will be determined by the encoder based on the bit rate, sample rate, and channel layout. This option has no effect if channel coupling is disabled.
	CplStartBand typex.UBool `json:"cpl_start_band" flag:"-cpl_start_band"` // -1~-15
}

const AC3ENC_PARAM = typex.AV_OPT_FLAG_AUDIO_PARAM | typex.AV_OPT_FLAG_ENCODING_PARAM

type Level = typex.String

const (
	Level_0_3 Level = "1.414"
	//Apply +3dB gain

	Level_0_1_5 Level = "1.189"
	//Apply +1.5dB gain

	Level_0 Level = "1.000"
	//Apply 0dB gain

	Level_1_1_5 Level = "0.841"
	//Apply -1.5dB gain

	Level_1_3_0 Level = "0.707"
	//Apply -3.0dB gain

	Level_1_4_5 Level = "0.595"
	//Apply -4.5dB gain (default)

	Level_1_6_0 Level = "0.500"
	//Apply -6.0dB gain

	Level_dB0 Level = "0.000"
	//Silence Surround Channel(s)
)

type DmixMode = typex.String

const (
	DmixMode_notindicated DmixMode = "notindicated"
	DmixMode_ltrt         DmixMode = "ltrt"
	DmixMode_loro         DmixMode = "loro"
	DmixMode_dplii        DmixMode = "dplii"
)

type DsurexMode = typex.String

const (
	DsurexMode_notindicated DsurexMode = "notindicated"
	DsurexMode_on           DsurexMode = "on"
	DsurexMode_off          DsurexMode = "off"
	DsurexMode_dpliiz       DsurexMode = "dpliiz"
)
