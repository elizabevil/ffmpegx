package videoencoderx

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// QsvHevc 9.26.6 HEVC Options
type QsvHevc struct {
	Extbrc typex.Bool `json:"extbrc" flag:"-extbrc"`
	//Extended bitrate control.

	RecoveryPointSei typex.Bool `json:"recovery_point_sei" flag:"-recovery_point_sei"`
	//Set this flag to insert the recovery point SEI message at the beginning of every intra refresh cycle.

	Rdo typex.Bool `json:"rdo" flag:"-rdo"`
	//Enable rate distortion optimization.

	MaxFrameSize typex.FrameSize `json:"max_frame_size" flag:"-max_frame_size"`
	//Maximum encoded frame size in bytes.

	MaxFrameSizeI typex.FrameSize `json:"max_frame_size_i" flag:"-max_frame_size_i"`
	//Maximum encoded frame size for I frames in bytes. If this value is set as larger than zero, then for I frames the value set by max_frame_size is ignored.

	MaxFrameSizeP typex.FrameSize `json:"max_frame_size_p" flag:"-max_frame_size_p"`
	//Maximum encoded frame size for P frames in bytes. If this value is set as larger than zero, then for P frames the value set by max_frame_size is ignored.

	MaxSliceSize typex.FrameSize `json:"max_slice_size" flag:"-max_slice_size"`
	//Maximum encoded slice size in bytes.

	Mbbrc typex.Bool `json:"mbbrc" flag:"-mbbrc"`
	//Setting this flag enables macroblock level bitrate control that generally improves subjective visual quality. Enabling this flag may have negative impact on performance and objective visual quality metric.

	LowDelayBrc typex.Bool `json:"low_delay_brc" flag:"-low_delay_brc"`
	//Setting this flag turns on or off LowDelayBRC feautre in qsv plugin, which provides more accurate bitrate control to minimize the variance of bitstream size frame by frame. Value: -1-default 0-off 1-on

	AdaptiveI typex.Bool `json:"adaptive_i" flag:"-adaptive_i"`
	//This flag controls insertion of I frames by the QSV encoder. Turn ON this flag to allow changing of frame type from P and B to I.

	AdaptiveB typex.Bool `json:"adaptive_b" flag:"-adaptive_b"`
	//This flag controls changing of frame type from B to P.

	PStrategy typex.Bool `json:"p_strategy" flag:"-p_strategy"`
	//Enable P-pyramid: 0-default 1-simple 2-pyramid(bf need to be set to 0).

	BStrategy typex.Bool `json:"b_strategy" flag:"-b_strategy"`
	//This option controls usage of B frames as reference.

	DblkIdc typex.Level `json:"dblk_idc" flag:"-dblk_idc"`
	//This option disable deblocking. It has value in range 0~2.

	IdrInterval typex.Distance `json:"idr_interval" flag:"-idr_interval"`
	//Distance (in I-frames) between IDR frames.

	LoadPlugin flagx.LoadPlugin `json:"load_plugin" flag:"-load_plugin"`
	//A user plugin to load in an internal session.
	LoadPlugins string `json:"load_plugins" flag:"-load_plugins"`
	//A :-separate list of hexadecimal plugin UIDs to load in an internal session.

	LookAheadDepth typex.Depth `json:"look_ahead_depth" flag:"-look_ahead_depth"`
	//Depth of look ahead in number frames, available when extbrc option is enabled.

	Profile QsvHevcProfile `json:"profile" flag:"-profile"`
	//Set the encoding profile (scc requires libmfx >= 1.32).
	Tier QsvHevcTier `json:"tier" flag:"-tier"`
	//Set the encoding tier (only level >= 4 can support high tier). This option only takes effect when the level option is specified.
	Gpb      typex.Bool      `json:"gpb" flag:"-gpb"`
	TileCols typex.TileCount `json:"tile_cols" flag:"-tile_cols"`
	//Number of columns for tiled encoding.

	TileRows typex.TileCount `json:"tile_rows" flag:"-tile_rows"`
	//Number of rows for tiled encoding.

	Aud typex.Bool `json:"aud" flag:"-aud"`
	//Insert the Access Unit Delimiter NAL.

	PicTimingSei typex.Bool `json:"pic_timing_sei" flag:"-pic_timing_sei"`
	//Insert picture timing SEI with pic_struct_syntax element.

	TransformSkip typex.Bool `json:"transform_skip" flag:"-transform_skip"`
	//Turn this option ON to enable transformskip. It is supported on platform equal or newer than ICL.

	IntRefType typex.IntRef `json:"int_ref_type" flag:"-int_ref_type"`
	//Specifies intra refresh type. The major goal of intra refresh is improvement of error resilience without significant impact on encoded bitstream size caused by I frames. The SDK encoder achieves this by encoding part of each frame in refresh cycle using intra MBs. none means no refresh. vertical means vertical refresh, by column of MBs. horizontal means horizontal refresh, by rows of MBs. slice means horizontal refresh by slices without overlapping. In case of slice, in_ref_cycle_size is ignored. To enable intra refresh, B frame should be set to 0.

	IntRefCycleSize typex.IntRef `json:"int_ref_cycle_size" flag:"-int_ref_cycle_size"`
	//Specifies number of pictures within refresh cycle starting from 2. 0 and 1 are invalid values.

	IntRefQpDelta typex.IntRefQpDelta `json:"int_ref_qp_delta" flag:"-int_ref_qp_delta"`
	//Specifies QP difference for inserted intra MBs. This is signed value in [-51, 51] range if target encoding bit-depth for luma samples is 8 and this range is [-63, 63] for 10 bit-depth or [-75, 75] for 12 bit-depth respectively.

	IntRefCycleDist typex.IntRef `json:"int_ref_cycle_dist" flag:"-int_ref_cycle_dist"`
	//Distance between the beginnings of the intra-refresh cycles in frames.

	MaxQpI typex.Quantizer `json:"max_qp_i" flag:"-max_qp_i"`
	//Maximum video quantizer scale for I frame.

	MinQpI typex.Quantizer `json:"min_qp_i" flag:"-min_qp_i"`
	//Minimum video quantizer scale for I frame.

	MaxQpP typex.Quantizer `json:"max_qp_p" flag:"-max_qp_p"`
	//Maximum video quantizer scale for P frame.

	MinQpP typex.Quantizer `json:"min_qp_p" flag:"-min_qp_p"`
	//Minimum video quantizer scale for P frame.

	MaxQpB typex.Quantizer `json:"max_qp_b" flag:"-max_qp_b"`
	//Maximum video quantizer scale for B frame.

	MinQpB typex.Quantizer `json:"min_qp_b" flag:"-min_qp_b"`
	//Minimum video quantizer scale for B frame.

	Scenario QsvHevcScenario `json:"scenario" flag:"-scenario"`
	//Provides a hint to encoder about the scenario for the encoding session.
	AvbrAccuracy typex.RateControl `json:"avbr_accuracy" flag:"-avbr_accuracy"`
	//Accuracy of the AVBR ratecontrol (unit of tenth of percent).

	AvbrConvergence typex.RateControl `json:"avbr_convergence" flag:"-avbr_convergence"`
	//Convergence of the AVBR ratecontrol (unit of 100 frames)

	//The parameters avbr_accuracy and avbr_convergence are for the average variable bitrate control (AVBR) algorithm. The algorithm focuses on overall encoding quality while meeting the specified bitrate, target_bitrate, within the accuracy range avbr_accuracy, after a avbr_Convergence period. This method does not follow HRD and the instant bitrate is not capped or padded.

	SkipFrame SkipFrame `json:"skip_frame" flag:"-skip_frame"`
	//Use per-frame metadata "qsv_skip_frame" to skip frame when encoding. This option defines the usage of this metadata.
}

type QsvHevcProfile = typex.String

const (
	QsvHevcProfile_unknown QsvHevcProfile = "unknown"
	QsvHevcProfile_main    QsvHevcProfile = "main"
	QsvHevcProfile_main10  QsvHevcProfile = "main10"
	QsvHevcProfile_mainsp  QsvHevcProfile = "mainsp"
	QsvHevcProfile_rext    QsvHevcProfile = "rext"
	QsvHevcProfile_scc     QsvHevcProfile = "scc"
)

type QsvHevcTier = typex.String

const (
	QsvHevcTier_main QsvHevcTier = "main"
	QsvHevcTier_high QsvHevcTier = "high"
)

type QsvHevcScenario = QsvH264Scenario
