package videoencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// QSV 9.26 QSV Encoders The family of Intel QuickSync Video encoders (MPEG-2, H.264, HEVC, JPEG/MJPEG, VP9, AV1)
type QSV struct {
	AsyncDepth typex.Depth `json:"async_depth" flag:"-async_depth"`
	//Specifies how many asynchronous operations an application performs before the application explicitly synchronizes the result. If zero, the value is not specified.

	Preset QsvPreset `json:"preset" flag:"-preset"`
	//This option itemizes a range of choices from veryfast (best speed) to veryslow (best quality)
	ForcedIdr typex.UBool `json:"forced_idr" flag:"-forced_idr"`
	//Forcing I frames as IDR frames.

	LowPower typex.Bool `json:"low_power" flag:"-low_power"`
	//For encoders set this flag to ON to reduce power consumption and GPU usage.
}

/*
9.26.2 Global Options -> MSDK Options
Additional libavcodec global options are mapped to MSDK options as follows:
	g/gop_size -> GopPicSize
	bf/max_b_frames+1 -> GopRefDist
	rc_init_occupancy/rc_initial_buffer_occupancy -> InitialDelayInKB
	slices -> NumSlice
	refs -> NumRefFrame
	b_strategy/b_frame_strategy -> BRefType
	cgop/CLOSED_GOP codec flag -> GopOptFlag
For the CQP mode, the i_qfactor/i_qoffset and b_qfactor/b_qoffset set the difference between QPP and QPI, and QPP and QPB respectively.
//Setting the coder option to the value vlc will make the H.264 encoder use CAVLC instead of CABAC.
*/

type QsvPreset string
type QsvVP9Profile string
type QsvAv1Profile = QsvHevcTier

const (
	QsvPreset_veryfast QsvPreset = "veryfast"
	QsvPreset_faster   QsvPreset = "faster"
	QsvPreset_fast     QsvPreset = "fast"
	QsvPreset_medium   QsvPreset = "medium"
	QsvPreset_slow     QsvPreset = "slow"
	QsvPreset_slower   QsvPreset = "slower"
	QsvPreset_veryslow QsvPreset = "veryslow"
)

const (
	QsvVP9Profile_unknown  QsvVP9Profile = "unknown"
	QsvVP9Profile_profile0 QsvVP9Profile = "profile0"
	QsvVP9Profile_profile1 QsvVP9Profile = "profile1"
	QsvVP9Profile_profile2 QsvVP9Profile = "profile2"
	QsvVP9Profile_profile3 QsvVP9Profile = "profile3"
)

type QsvMpeg2 struct {
	Profile QsvH264Profile `json:"profile" flag:"-profile"`
}
type QsvVp9 struct {
	Profile  QsvVP9Profile   `json:"profile" flag:"-profile"`
	TileCols typex.TileCount `json:"tile_cols" flag:"-tile_cols"`
	//Number of columns for tiled encoding (requires libmfx >= 1.29).

	TileRows typex.TileCount `json:"tile_rows" flag:"-tile_rows"`
	//Number of rows for tiled encoding (requires libmfx >= 1.29).
}
type QsvAv1 struct {
	Profile  QsvAv1Profile   `json:"profile" flag:"-profile"`
	TileCols typex.TileCount `json:"tile_cols" flag:"-tile_cols"`
	//Number of columns for tiled encoding.

	TileRows typex.TileCount `json:"tile_rows" flag:"-tile_rows"`
	//Number of rows for tiled encoding.

	AdaptiveI typex.Bool `json:"adaptive_i" flag:"-adaptive_i"`
	//This flag controls insertion of I frames by the QSV encoder. Turn ON this flag to allow changing of frame type from P and B to I.

	AdaptiveB typex.Bool `json:"adaptive_b" flag:"-adaptive_b"`
	//This flag controls changing of frame type from B to P.

	BStrategy typex.Bool `json:"b_strategy" flag:"-b_strategy"`
	//This option controls usage of B frames as reference.

	Extbrc int `json:"extbrc" flag:"-extbrc"`
	//Extended bitrate control.

	LookAheadDepth typex.Depth `json:"look_ahead_depth" flag:"-look_ahead_depth"`
	//Depth of look ahead in number frames, available when extbrc option is enabled.

	LowDelayBrc typex.Bool `json:"low_delay_brc" flag:"-low_delay_brc"`
	//Setting this flag turns on or off LowDelayBRC feautre in qsv plugin, which provides more accurate bitrate control to minimize the variance of bitstream size frame by frame. Value: -1-default 0-off 1-on

	MaxFrameSize typex.Size `json:"max_frame_size" flag:"-max_frame_size"`
	//Set the allowed max size in bytes for each frame. If the frame size exceeds the limitation, encoder will adjust the QP value to control the frame size. Invalid in CQP rate control mode.
}
