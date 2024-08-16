package videoencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// // VAAPI 9.28 VAAPI encoders
type VAAPI struct {
	LowPower typex.Bool `json:"low_power" flag:"-low_power"`
	//Some drivers/platforms offer a second encoder for some codecs intended to use less power than the default encoder; setting this option will attempt to use that encoder. Note that it may support a reduced feature set, so some other options may not be available in this mode.

	IdrInterval typex.UNumber `json:"idr_interval" flag:"-idr_interval"`
	//Set the number of normal intra frames between full-refresh (IDR) frames in open-GOP mode. The intra frames are still IRAPs, but will not include global headers and may have non-decodable leading pictures.

	BDepth typex.Depth `json:"b_depth" flag:"-b_depth"`
	//Set the B-frame reference depth. When set to one (the default), all B-frames will refer only to P- or I-frames. When set to greater values multiple layers of B-frames will be present, frames in each layer only referring to frames in higher layers.

	AsyncDepth typex.Depth `json:"async_depth" flag:"-async_depth"`
	//Maximum processing parallelism. Increase this to improve single channel performance. This option doesn’t work if driver doesn’t implement vaSyncBuffer function. Please make sure there are enough hw_frames allocated if a large number of async_depth is used.

	MaxFrameSize typex.Size `json:"max_frame_size" flag:"-max_frame_size"`
	//Set the allowed max size in bytes for each frame. If the frame size exceeds the limitation, encoder will adjust the QP value to control the frame size. Invalid in CQP rate control mode.

	RcMode RcMode `json:"rc_mode" flag:"-rc_mode"`
	//Set the rate control mode to use. A given driver may only support a subset of modes
	Blbrc typex.Bool `json:"blbrc" flag:"-blbrc"`
	//Enable block level rate control, which assigns different bitrate block by block. Invalid for CQP mode.

	// ---Each encoder also has its own specific options:

	//av1_vaapi
	//profile sets the value of seq_profile. tier sets the value of seq_tier. level sets the value of seq_level_idx.

	Tiles typex.TileCount `json:"tiles" flag:"-tiles"`
	//Set the number of tiles to encode the input video with, as columns x rows. (default is auto, which means use minimal tile column/row number).

	TileGroups typex.TileCount `json:"tile_groups" flag:"-tile_groups"`
	//Set tile groups number. All the tiles will be distributed as evenly as possible to each tile group. (default is 1).

	//h264_vaapi
	//profile sets the value of profile_idc and the constraint_set*_flags. level sets the value of level_idc.
	Coder VaapiH264Coder `json:"coder" flag:"-coder"`
	Aud   typex.Bool     `json:"aud" flag:"-aud"`
	//Include access unit delimiters in the stream (not included by default).

	Sei VaapiH264Sei `json:"sei" flag:"-sei"`
	//Set SEI message types to include. Some combination of the following values:

	//hevc_vaapi
	//profile and level set the values of general_profile_idc and general_level_idc respectively.

	//Include access unit delimiters in the stream (not included by default).

	Tier QsvHevcTier `json:"tier" flag:"-tier"`
	//Set general_tier_flag. This may affect the level chosen for the stream if it is not explicitly specified.

	//mjpeg_vaapi
	//Only baseline DCT encoding is supported. The encoder always uses the standard quantisation and huffman tables - global_quality scales the standard quantisation table (range 1-100).

	Jfif typex.Size `json:"jfif" flag:"-jfif"`
	//Include JFIF header in each frame (not included by default).

	Huffman typex.Bool `json:"huffman" flag:"-huffman"`
	//Include standard huffman tables (on by default). Turning this off will save a few hundred bytes in each output frame, but may lose compatibility with some JPEG decoders which don’t fully handle MJPEG.

	//mpeg2_vaapi
	//profile and level set the value of profile_and_level_indication.

	//vp8_vaapi
	//B-frames are not supported.
	LoopFilterLevel     typex.Level `json:"loop_filter_level" flag:"-loop_filter_level"`
	LoopFilterSharpness typex.Level `json:"loop_filter_sharpness" flag:"-loop_filter_sharpness"`
	//vp9_vaapi
	//global_quality sets the q_idx used for P-frames (range 0-255).

}

type RcMode = typex.String

const (
	auto RcMode = "auto"
	//Choose the mode automatically based on driver support and the other options. This is the default.

	CQP RcMode = "CQP"
	//Constant-quality.

	CBR RcMode = "CBR"
	//Constant-bitrate.

	VBR RcMode = "UVBR"
	//Variable-bitrate.

	ICQ RcMode = "ICQ"
	//Intelligent constant-quality.

	QVBR RcMode = "QVBR"
	//Quality-defined variable-bitrate.

	AVBR RcMode = "AVBR"
	//Average variable bitrate.
)

type VaapiH264Coder = typex.String

const (
	ac    VaapiH264Coder = "ac"
	cabac VaapiH264Coder = "cabac"
	vlc   VaapiH264Coder = "vlc"
	cavlc VaapiH264Coder = "cavlc"
)

type VaapiH264Sei = typex.String

const (
	identifier VaapiH264Sei = "identifier"
	//Include a user_data_unregistered message containing information about the encoder.

	timing VaapiH264Sei = "timing"
	//Include picture timing parameters (buffering_period and pic_timing messages).

	recovery_point VaapiH264Sei = "recovery_point"
	//Include recovery points where appropriate (recovery_point messages).
)
