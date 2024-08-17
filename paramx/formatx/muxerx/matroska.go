package muxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// // 21.57 matroska MATROSKA
type MATROSKA struct {
	Title typex.Name `json:"title" flag:"-title"`
	//Set title name provided to a single track. This gets mapped to the FileDescription element for a stream written as attachment.

	Language typex.Dict `json:"language" flag:"-language"`
	//Specify the language of the track in the Matroska languages form.

	//The language can be either the 3 letters bibliographic ISO-639-2 (ISO 639-2/B) form (like "fre" for French), or a language code mixed with a country code for specialities in languages (like "fre-ca" for Canadian French).

	StereoMode StereoMode `json:"stereo_mode" flag:"-stereo_mode"`
	//Set stereo 3D video layout of two views in a single video track.
	//==============

	ReserveIndexSpace typex.Index `json:"reserve_index_space" flag:"-reserve_index_space"`
	//By default, this muxer writes the index for seeking (called cues in Matroska terms) at the end of the file, because it cannot know in advance how much space to leave for the index at the beginning of the file. However for some use cases – e.g. streaming where seeking is possible but slow – it is useful to put the index at the beginning of the file.

	//If this option is set to a non-zero value, the muxer will reserve size bytes of space in the file header and then try to write the cues there when the muxing finishes. If the reserved space does not suffice, no Cues will be written, the file will be finalized and writing the trailer will return an error. A safe size for most use cases should be about 50kB per hour of video.

	//Note that cues are only written if the output is seekable and this option will have no effect if it is not.

	CuesToFront typex.UBool `json:"cues_to_front" flag:"-cues_to_front"`
	//If set, the muxer will write the index at the beginning of the file by shifting the main data if necessary. This can be combined with reserve_index_space in which case the data is only shifted if the initially reserved space turns out to be insufficient.

	//This option is ignored if the output is unseekable.

	ClusterSizeLimit typex.Limit `json:"cluster_size_limit" flag:"-cluster_size_limit"`
	//Store at most the provided amount of bytes in a cluster.

	//If not specified, the limit is set automatically to a sensible hardcoded fixed value.

	ClusterTimeLimit typex.Limit `json:"cluster_time_limit" flag:"-cluster_time_limit"`
	//Store at most the provided number of milliseconds in a cluster.

	//If not specified, the limit is set automatically to a sensible hardcoded fixed value.

	Dash typex.UBool `json:"dash" flag:"-dash"`
	//Create a WebM file conforming to WebM DASH specification. By default it is set to false.

	DashTrackNumber typex.UNumber `json:"dash_track_number" flag:"-dash_track_number"`
	//Track number for the DASH stream. By default it is set to 1.

	Live typex.UBool `json:"live" flag:"-live"`
	//Write files assuming it is a live stream. By default it is set to false.

	AllowRawVfw typex.UBool `json:"allow_raw_vfw" flag:"-allow_raw_vfw"`
	//Allow raw VFW mode. By default it is set to false.

	FlippedRawRgb typex.UBool `json:"flipped_raw_rgb" flag:"-flipped_raw_rgb"`
	//If set to true, store positive height for raw RGB bitmaps, which indicates bitmap is stored bottom-up. Note that this option does not flip the bitmap which has to be done manually beforehand, e.g. by using the ‘vflip’ filter. Default is false and indicates bitmap is stored top down.

	WriteCrc32 typex.UBool `json:"write_crc32" flag:"-write_crc32"`
	//Write a CRC32 element inside every Level 1 element. By default it is set to true. This option is ignored for WebM.

	DefaultMode DefaultMode `json:"default_mode" flag:"-default_mode"`
	//Control how the FlagDefault of the output tracks will be set. It influences which tracks players should play by default. The default mode is ‘passthrough’.
}

type StereoMode = typex.String
type DefaultMode = typex.String

const (
	StereoMode_mono StereoMode = "mono"
	//video is not stereo

	StereoMode_left_right StereoMode = "left_right"
	//Both views are arranged side by side, Left-eye view is on the left

	StereoMode_bottom_top StereoMode = "bottom_top"
	//Both views are arranged in top-bottom orientation, Left-eye view is at bottom

	StereoMode_top_bottom StereoMode = "top_bottom"
	//Both views are arranged in top-bottom orientation, Left-eye view is on top

	StereoMode_checkerboard_rl StereoMode = "checkerboard_rl"
	//Each view is arranged in a checkerboard interleaved pattern, Left-eye view being first

	StereoMode_checkerboard_lr StereoMode = "checkerboard_lr"
	//Each view is arranged in a checkerboard interleaved pattern, Right-eye view being first

	StereoMode_row_interleaved_rl StereoMode = "row_interleaved_rl"
	//Each view is constituted by a row based interleaving, Right-eye view is first row

	StereoMode_row_interleaved_lr StereoMode = "row_interleaved_lr"
	//Each view is constituted by a row based interleaving, Left-eye view is first row

	StereoMode_col_interleaved_rl StereoMode = "col_interleaved_rl"
	//Both views are arranged in a column based interleaving manner, Right-eye view is first column

	StereoMode_col_interleaved_lr StereoMode = "col_interleaved_lr"
	//Both views are arranged in a column based interleaving manner, Left-eye view is first column

	StereoMode_anaglyph_cyan_red StereoMode = "anaglyph_cyan_red"
	//All frames are in anaglyph format viewable through red-cyan filters

	StereoMode_right_left StereoMode = "right_left"
	//Both views are arranged side by side, Right-eye view is on the left

	StereoMode_anaglyph_green_magenta StereoMode = "anaglyph_green_magenta"
	//All frames are in anaglyph format viewable through green-magenta filters

	StereoMode_block_lr StereoMode = "block_lr"
	//Both eyes laced in one Block, Left-eye view is first

	StereoMode_block_rl StereoMode = "block_rl"

	//Both eyes laced in one Block, Right-eye view is first
)

const (
	DefaultMode_infer DefaultMode = "infer"
	//Every track with disposition default will have the FlagDefault set. Additionally, for each type of track (audio, video or subtitle), if no track with disposition default of this type exists, then the first track of this type will be marked as default (if existing). This ensures that the default flag is set in a sensible way even if the input originated from containers that lack the concept of default tracks.

	DefaultMode_infer_no_subs DefaultMode = "infer_no_subs"
	//This mode is the same as infer except that if no subtitle track with disposition default exists, no subtitle track will be marked as default.

	DefaultMode_passthrough DefaultMode = "passthrough"
	//In this mode the FlagDefault is set if and only if the AV_DISPOSITION_DEFAULT flag is set in the disposition of the corresponding stream.
)
