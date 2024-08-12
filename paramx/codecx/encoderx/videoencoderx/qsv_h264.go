package videoencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

type QsvH264 struct {
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

	MaxSliceSize typex.SliceSize `json:"max_slice_size" flag:"-max_slice_size"`
	//Maximum encoded slice size in bytes.

	BitrateLimit typex.Bool `json:"bitrate_limit" flag:"-bitrate_limit"`
	//Toggle bitrate limitations. Modifies bitrate to be in the range imposed by the QSV encoder. Setting this flag off may lead to violation of HRD conformance. Mind that specifying bitrate below the QSV encoder range might significantly affect quality. If on this option takes effect in non CQP modes: if bitrate is not in the range imposed by the QSV encoder, it will be changed to be in the range.

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

	Cavlc typex.UBool `json:"cavlc" flag:"-cavlc"`
	//If set, CAVLC is used; if unset, CABAC is used for encoding.

	Vcm typex.UBool `json:"vcm" flag:"-vcm"`
	//Video conferencing mode, please see ratecontrol method.

	IdrInterval typex.Distance `json:"idr_interval" flag:"-idr_interval"`
	//Distance (in I-frames) between IDR frames.

	PicTimingSei typex.UBool `json:"pic_timing_sei" flag:"-pic_timing_sei"`
	//insert picture timing sei with pic_struct_syntax element.

	SingleSeiNalUnit typex.Bool `json:"single_sei_nal_unit" flag:"-single_sei_nal_unit"`
	//Put all the SEI messages into one NALU.

	MaxDecFrameBuffering typex.UNumber `json:"max_dec_frame_buffering" flag:"-max_dec_frame_buffering"`
	//Maximum number of frames buffered in the DPB.

	LookAhead typex.Level `json:"look_ahead" flag:"-look_ahead"`
	//Use UVBR algorithm with look ahead.

	LookAheadDepth typex.Depth `json:"look_ahead_depth" flag:"-look_ahead_depth"`
	//Depth of look ahead in number frames.

	LookAheadDownsampling LookAheadDownsampling `json:"look_ahead_downsampling" flag:"-look_ahead_downsampling"`
	//Downscaling factor for the frames saved for the lookahead analysis.

	IntRefType typex.IntRef `json:"int_ref_type" flag:"-int_ref_type"`
	//Specifies intra refresh type. The major goal of intra refresh is improvement of error resilience without significant impact on encoded bitstream size caused by I frames. The SDK encoder achieves this by encoding part of each frame in refresh cycle using intra MBs. none means no refresh. vertical means vertical refresh, by column of MBs. horizontal means horizontal refresh, by rows of MBs. slice means horizontal refresh by slices without overlapping. In case of slice, in_ref_cycle_size is ignored. To enable intra refresh, B frame should be set to 0.

	IntRefCycleSize typex.IntRef `json:"int_ref_cycle_size" flag:"-int_ref_cycle_size"`
	//Specifies number of pictures within refresh cycle starting from 2. 0 and 1 are invalid values.

	IntRefQpDelta typex.IntRefQpDelta `json:"int_ref_qp_delta" flag:"-int_ref_qp_delta"`
	//Specifies QP difference for inserted intra MBs. This is signed value in [-51, 51] range if target encoding bit-depth for luma samples is 8 and this range is [-63, 63] for 10 bit-depth or [-75, 75] for 12 bit-depth respectively.

	IntRefCycleDist typex.IntRef `json:"int_ref_cycle_dist" flag:"-int_ref_cycle_dist"`
	//Distance between the beginnings of the intra-refresh cycles in frames.

	Profile QsvH264Profile `json:"profile" flag:"-profile"`
	A53Cc   typex.Bool     `json:"a53cc" flag:"-a53cc"`
	//Use A53 Closed Captions (if available).

	Aud typex.Bool `json:"aud" flag:"-aud"`
	//Insert the Access Unit Delimiter NAL.

	Mfmode    typex.BoolMode `json:"mfmode" flag:"-mfmode"`
	RepeatPps typex.UBool    `json:"repeat_pps" flag:"-repeat_pps"`
	//Repeat pps for every frame.

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

	Scenario QsvH264Scenario `json:"scenario" flag:"-scenario"`
	//Provides a hint to encoder about the scenario for the encoding session.
	AvbrAccuracy typex.RateControl `json:"avbr_accuracy" flag:"-avbr_accuracy"`
	//Accuracy of the AVBR ratecontrol (unit of tenth of percent).

	AvbrConvergence typex.RateControl `json:"avbr_convergence" flag:"-avbr_convergence"`
	//Convergence of the AVBR ratecontrol (unit of 100 frames)

	//The parameters avbr_accuracy and avbr_convergence are for the average variable bitrate control (AVBR) algorithm. The algorithm focuses on overall encoding quality while meeting the specified bitrate, target_bitrate, within the accuracy range avbr_accuracy, after a avbr_Convergence period. This method does not follow HRD and the instant bitrate is not capped or padded.

	SkipFrame SkipFrame `json:"skip_frame" flag:"-skip_frame"`
	//Use per-frame metadata "qsv_skip_frame" to skip frame when encoding. This option defines the usage of this metadata.
}

type LookAheadDownsampling string

const (
	a__unknown LookAheadDownsampling = "unknown"
	a__auto    LookAheadDownsampling = "auto"
	a__off     LookAheadDownsampling = "off"
	a__2x      LookAheadDownsampling = "2x"
	a__4x      LookAheadDownsampling = "4x"
)

type QsvH264Profile string

const (
	QsvH264Profile_unknown  QsvH264Profile = "unknown"
	QsvH264Profile_baseline QsvH264Profile = "baseline"
	QsvH264Profile_main     QsvH264Profile = "main"
	QsvH264Profile_high     QsvH264Profile = "high"
)

type QsvH264Scenario string

const (
	QsvH264Scenario_unknown           QsvH264Scenario = "unknown"
	QsvH264Scenario_displayremoting   QsvH264Scenario = "displayremoting"
	QsvH264Scenario_videoconference   QsvH264Scenario = "videoconference"
	QsvH264Scenario_archive           QsvH264Scenario = "archive"
	QsvH264Scenario_livestreaming     QsvH264Scenario = "livestreaming"
	QsvH264Scenario_cameracapture     QsvH264Scenario = "cameracapture"
	QsvH264Scenario_videosurveillance QsvH264Scenario = "videosurveillance"
	QsvH264Scenario_gamestreaming     QsvH264Scenario = "gamestreaming"
	QsvH264Scenario_remotegaming      QsvH264Scenario = "remotegaming"
)

type SkipFrame string

const (
	SkipFrame_no_skip SkipFrame = "no_skip"
	//Frame skipping is disabled.

	SkipFrame_insert_dummy SkipFrame = "insert_dummy"
	//Encoder inserts into bitstream frame where all macroblocks are encoded as skipped.

	SkipFrame_insert_nothing SkipFrame = "insert_nothing"
	//Similar to insert_dummy, but encoder inserts nothing into bitstream. The skipped frames are still used in brc. For example, gop still include skipped frames, and the frames after skipped frames will be larger in size.

	SkipFrame_brc_only SkipFrame = "brc_only"
	//skip_frame metadata indicates the number of missed frames before the current frame.
)

/*
ffmpeg -re -i in.ts -b:v 1000k -b:a 64k -a53cc 1 -f hls \
  -cc_stream_map "ccgroup:cc,instreamid:CC1,language:en" \
  -master_pl_name master.m3u8


ffmpeg -re -i in.ts -b:v:0 1000k -b:v:1 256k -b:a:0 64k -b:a:1 32k \
  -a53cc:0 1 -a53cc:1 1\
  -map 0:v -map 0:a -map 0:v -map 0:a -f hls \
  -cc_stream_map "ccgroup:cc,instreamid:CC1,language:en ccgroup:cc,instreamid:CC2,language:sp" \
  -var_stream_map "v:0,a:0,ccgroup:cc v:1,a:1,ccgroup:cc" \
  -master_pl_name master.m3u8

*/
