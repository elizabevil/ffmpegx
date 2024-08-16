package videoencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// MPEG2 MPEG-2 video encoder.
type MPEG2 struct {
	Profile Profile `json:"profile" flag:"-profile"`

	Level Level `json:"level" flag:"-level"`

	SeqDispExt *typex.Bool `json:"seq_disp_ext" flag:"-seq_disp_ext,-1"`
	//Specifies if the encoder should write a sequence_display_extension to the output.
	//-1	auto	 Decide automatically to write it or not (this is the default) by checking if the data to be written is different from the default or unspecified values.
	//0     never    Never write it.
	//1	    always	 Always write it.

	VideoFormat VideoFormat `json:"video_format" flag:"-video_format"`
	//Specifies the video_format written into the sequence display extension indicating the source of the video pictures. The default is ‘unspecified’, can be ‘component’, ‘pal’, ‘ntsc’, ‘secam’ or ‘mac’. For maximum compatibility, use ‘component’.

	A53Cc *typex.Bool `json:"a53cc" flag:"-a53cc,1"`
	//Import closed captions (which must be ATSC compatible format) into output. Default is 1 (on).
}
type Level = typex.String
type VideoFormat = typex.String

type Profile = typex.String

const (
	Level_high     Level = "high"
	Level_high1440 Level = "high1440"
	Level_main     Level = "main"
	Level_low      Level = "low"

	Profile_422  Profile = "422"
	Profile_high Profile = "high"
	Profile_ss   Profile = "ss" //Spatially Scalable

	Profile_snr    Profile = "snr" //SNR Scalable
	Profile_main   Profile = "main"
	Profile_simple Profile = "simple"
)
const (
	VideoFormat_component VideoFormat = "component"
	VideoFormat_pal       VideoFormat = "pal"
	VideoFormat_ntsc      VideoFormat = "ntsc"
	VideoFormat_secam     VideoFormat = "secam"
	VideoFormat_mac       VideoFormat = "mac"
)
