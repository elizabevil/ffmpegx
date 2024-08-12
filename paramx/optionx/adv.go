package optionx

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

type ExpertInput struct {
	FindStreamInfo bool `json:"find_stream_info" flag:"-find_stream_info"` //OPT_TYPE_BOOL, OPT_INPUT | OPT_EXPERT | OPT_OFFSET,

	//========
	AccurateSeek bool `json:"accurate_seek" flag:"-accurate_seek"` //          OPT_TYPE_BOOL, OPT_OFFSET | OPT_EXPERT | OPT_INPUT,

	Isync typex.Index `json:"isync" flag:"-isync"` //                  OPT_TYPE_INT, OPT_OFFSET | OPT_EXPERT | OPT_INPUT,

	Itsoffset typex.Offset `json:"itsoffset" flag:"-itsoffset"` //              OPT_TYPE_TIME, OPT_OFFSET | OPT_EXPERT | OPT_INPUT,

	Itsscale typex.Flt `json:"itsscale" flag:"-itsscale"` //               OPT_TYPE_DOUBLE, OPT_PERSTREAM | OPT_EXPERT | OPT_INPUT,

	Re bool `json:"re" flag:"-re"` //                     OPT_TYPE_BOOL, OPT_EXPERT | OPT_OFFSET | OPT_INPUT,

	Readrate typex.Flt `json:"readrate" flag:"-readrate"` //               OPT_TYPE_FLOAT, OPT_OFFSET | OPT_EXPERT | OPT_INPUT,

	ReadrateInitialBurst typex.Flt `json:"readrate_initial_burst" flag:"-readrate_initial_burst"` // OPT_TYPE_DOUBLE, OPT_OFFSET | OPT_EXPERT | OPT_INPUT,

	Bitexact bool `json:"bitexact" flag:"-bitexact"` //               OPT_TYPE_BOOL, OPT_EXPERT | OPT_OFFSET | OPT_OUTPUT | OPT_INPUT,

	DumpAttachment typex.String `json:"dump_attachment" flag:"-dump_attachment"` //     OPT_TYPE_STRING, OPT_SPEC | OPT_EXPERT | OPT_INPUT,

	Hwaccel Hwaccel `json:"hwaccel" flag:"-hwaccel"` //                    OPT_TYPE_STRING, OPT_VIDEO | OPT_EXPERT | OPT_PERSTREAM | OPT_INPUT,

	HwaccelDevice typex.String `json:"hwaccel_device" flag:"-hwaccel_device"` //             OPT_TYPE_STRING, OPT_VIDEO | OPT_EXPERT | OPT_PERSTREAM | OPT_INPUT,

	HwaccelOutputFormat typex.String `json:"hwaccel_output_format" flag:"-hwaccel_output_format"` //      OPT_TYPE_STRING, OPT_VIDEO | OPT_EXPERT | OPT_PERSTREAM | OPT_INPUT,

	Autorotate   bool `json:"autorotate" flag:"-autorotate"`     //                 OPT_TYPE_BOOL,   OPT_PERSTREAM | OPT_EXPERT | OPT_INPUT,
	NoAutorotate bool `json:"noautorotate" flag:"-noautorotate"` //                 OPT_TYPE_BOOL,   OPT_PERSTREAM | OPT_EXPERT | OPT_INPUT,

	GuessLayoutMax typex.Number `json:"guess_layout_max" flag:"-guess_layout_max"` // OPT_TYPE_INT,     OPT_AUDIO | OPT_PERSTREAM | OPT_EXPERT | OPT_INPUT,

	FixSubDuration bool `json:"fix_sub_duration" flag:"-fix_sub_duration"` // OPT_TYPE_BOOL, OPT_EXPERT | OPT_SUBTITLE | OPT_PERSTREAM | OPT_INPUT,

	Bsf typex.String `json:"bsf" flag:"-bsf"` // OPT_TYPE_STRING, OPT_PERSTREAM | OPT_EXPERT | OPT_OUTPUT | OPT_INPUT,

	Sseof typex.Position `json:"sseof" flag:"-sseof"` //                  OPT_TYPE_TIME, OPT_OFFSET | OPT_INPUT | OPT_EXPERT,

	SeekTimestamp typex.Number `json:"seek_timestamp" flag:"-seek_timestamp"` //         OPT_TYPE_INT, OPT_OFFSET | OPT_INPUT | OPT_EXPERT,

	Tag typex.String `json:"tag" flag:"-tag"` //                    OPT_TYPE_STRING, OPT_PERSTREAM | OPT_EXPERT | OPT_OUTPUT | OPT_INPUT | OPT_HAS_ALT,

	ReinitFilter typex.Number `json:"reinit_filter" flag:"-reinit_filter"` //          OPT_TYPE_INT, OPT_PERSTREAM | OPT_INPUT | OPT_EXPERT,

	StreamLoop typex.Bool `json:"stream_loop" flag:"-stream_loop"` //         OPT_TYPE_INT, OPT_EXPERT | OPT_INPUT | OPT_OFFSET,

	Discard flagx.Discard `json:"discard" flag:"-discard"` //             OPT_TYPE_STRING, OPT_PERSTREAM | OPT_INPUT | OPT_EXPERT,

	DisplayRotation typex.Flt `json:"display_rotation" flag:"-display_rotation"` //           OPT_TYPE_DOUBLE, OPT_VIDEO | OPT_PERSTREAM | OPT_INPUT | OPT_EXPERT,

	DisplayHflip bool `json:"display_hflip" flag:"-display_hflip"` //              OPT_TYPE_BOOL,   OPT_VIDEO | OPT_PERSTREAM | OPT_INPUT | OPT_EXPERT,

	DisplayVflip bool `json:"display_vflip" flag:"-display_vflip"` //              OPT_TYPE_BOOL,   OPT_VIDEO | OPT_PERSTREAM | OPT_INPUT | OPT_EXPERT,

	//Specify the resolution of the canvas to render subtitles to; usually, this should be frame size of input video. This only applies when -subtitle_type is set to bitmap.
	CanvasSize typex.VideoSize `json:"canvas_size" flag:"-canvas_size"` // OPT_TYPE_STRING, OPT_SUBTITLE | OPT_PERSTREAM | OPT_INPUT | OPT_EXPERT,

}
type ExpertOutput struct {
	Pre                      typex.String   `json:"pre" flag:"-pre"`                                                 //typex.String, OPT_PERSTREAM | OPT_OUTPUT | OPT_EXPERT | OPT_HAS_ALT,
	MapMetadata              typex.String   `json:"map_metadata" flag:"-map_metadata"`                               //typex.String, OPT_SPEC | OPT_OUTPUT | OPT_EXPERT,
	Apad                     typex.String   `json:"apad" flag:"-apad"`                                               //typex.String, OPT_PERSTREAM | OPT_EXPERT | OPT_OUTPUT,
	Copyinkf                 typex.UBool    `json:"copyinkf" flag:"-copyinkf"`                                       //typex.UBool, OPT_EXPERT | OPT_PERSTREAM | OPT_OUTPUT,
	Copypriorss              typex.Number   `json:"copypriorss" flag:"-copypriorss"`                                 //typex.Number, OPT_EXPERT | OPT_PERSTREAM | OPT_OUTPUT,
	Frames                   typex.Number   `json:"frames" flag:"-frames"`                                           //typex.Number64, OPT_PERSTREAM | OPT_OUTPUT | OPT_EXPERT | OPT_HAS_ALT,
	Q                        typex.Flt      `json:"q" flag:"-q"`                                                     //typex.Flt, OPT_EXPERT | OPT_PERSTREAM | OPT_OUTPUT | OPT_HAS_CANON,
	FilterScript             typex.String   `json:"filter_script" flag:"-filter_script"`                             //typex.String, OPT_PERSTREAM | OPT_EXPERT | OPT_OUTPUT,
	BitsPerRawSample         typex.UNumber  `json:"bits_per_raw_sample" flag:"-bits_per_raw_sample"`                 //typex.Number, OPT_EXPERT | OPT_PERSTREAM | OPT_OUTPUT,
	StatsEncPre              typex.String   `json:"stats_enc_pre" flag:"-stats_enc_pre"`                             //typex.String, OPT_PERSTREAM | OPT_EXPERT | OPT_OUTPUT,
	StatsEncPost             typex.String   `json:"stats_enc_post" flag:"-stats_enc_post"`                           //typex.String, OPT_PERSTREAM | OPT_EXPERT | OPT_OUTPUT,
	StatsMuxPre              typex.String   `json:"stats_mux_pre" flag:"-stats_mux_pre"`                             //typex.String, OPT_PERSTREAM | OPT_EXPERT | OPT_OUTPUT,
	StatsEncPreFmt           typex.String   `json:"stats_enc_pre_fmt" flag:"-stats_enc_pre_fmt"`                     //typex.String, OPT_PERSTREAM | OPT_EXPERT | OPT_OUTPUT,
	StatsEncPostFmt          typex.String   `json:"stats_enc_post_fmt" flag:"-stats_enc_post_fmt"`                   //typex.String, OPT_PERSTREAM | OPT_EXPERT | OPT_OUTPUT,
	StatsMuxPreFmt           typex.String   `json:"stats_mux_pre_fmt" flag:"-stats_mux_pre_fmt"`                     //typex.String, OPT_PERSTREAM | OPT_EXPERT | OPT_OUTPUT,
	Autoscale                typex.UBool    `json:"autoscale" flag:"-autoscale"`                                     //typex.UBool,   OPT_PERSTREAM | OPT_EXPERT | OPT_OUTPUT,
	TimeBase                 typex.TimeBase `json:"time_base" flag:"-time_base"`                                     //typex.String, OPT_EXPERT | OPT_PERSTREAM | OPT_OUTPUT,
	EncTimeBase              typex.String   `json:"enc_time_base" flag:"-enc_time_base"`                             //typex.String, OPT_EXPERT | OPT_PERSTREAM | OPT_OUTPUT,
	MaxMuxingQueueSize       typex.Number   `json:"max_muxing_queue_size" flag:"-max_muxing_queue_size"`             //typex.Number, OPT_PERSTREAM | OPT_EXPERT | OPT_OUTPUT,
	MuxingQueueDataThreshold typex.Number   `json:"muxing_queue_data_threshold" flag:"-muxing_queue_data_threshold"` //typex.Number, OPT_PERSTREAM | OPT_EXPERT | OPT_OUTPUT,
	Qscale                   typex.String   `json:"qscale" flag:"-qscale"`                                           //typex.String, OPT_FUNC_ARG | OPT_EXPERT | OPT_PERFILE | OPT_OUTPUT | OPT_HAS_ALT,
	Profile                  typex.String   `json:"profile" flag:"-profile"`                                         //typex.String, OPT_FUNC_ARG | OPT_EXPERT | OPT_PERFILE | OPT_OUTPUT,
	Fpre                     typex.String   `json:"fpre" flag:"-fpre"`                                               //typex.String, OPT_FUNC_ARG | OPT_EXPERT| OPT_PERFILE | OPT_OUTPUT | OPT_HAS_CANON,
	Fpsmax                   typex.String   `json:"fpsmax" flag:"-fpsmax"`                                           //typex.String, OPT_VIDEO | OPT_PERSTREAM | OPT_OUTPUT | OPT_EXPERT,
	RcOverride               typex.String   `json:"rc_override" flag:"-rc_override"`                                 //typex.String, OPT_VIDEO | OPT_EXPERT  | OPT_PERSTREAM | OPT_OUTPUT,
	Passlogfile              typex.String   `json:"passlogfile" flag:"-passlogfile"`                                 //typex.String, OPT_VIDEO | OPT_EXPERT | OPT_PERSTREAM | OPT_OUTPUT,
	IntraMatrix              typex.String   `json:"intra_matrix" flag:"-intra_matrix"`                               //typex.String, OPT_VIDEO | OPT_EXPERT | OPT_PERSTREAM | OPT_OUTPUT,
	InterMatrix              typex.String   `json:"inter_matrix" flag:"-inter_matrix"`                               //typex.String, OPT_VIDEO | OPT_EXPERT | OPT_PERSTREAM | OPT_OUTPUT,
	ChromaIntraMatrix        typex.String   `json:"chroma_intra_matrix" flag:"-chroma_intra_matrix"`                 //typex.String, OPT_VIDEO | OPT_EXPERT | OPT_PERSTREAM | OPT_OUTPUT,
	FpsMode                  FpsMode        `json:"fps_mode" flag:"-fps_mode"`                                       //typex.String, OPT_VIDEO | OPT_EXPERT | OPT_PERSTREAM | OPT_OUTPUT,
	ForceFps                 typex.UBool    `json:"force_fps" flag:"-force_fps"`                                     //typex.UBool,   OPT_VIDEO | OPT_EXPERT  | OPT_PERSTREAM | OPT_OUTPUT,
	ForceKeyFrames           typex.String   `json:"force_key_frames" flag:"-force_key_frames"`                       //typex.String, OPT_VIDEO | OPT_EXPERT | OPT_PERSTREAM | OPT_OUTPUT,
	FixSubDurationHeartbeat  typex.UBool    `json:"fix_sub_duration_heartbeat" flag:"-fix_sub_duration_heartbeat"`   //typex.UBool,   OPT_VIDEO | OPT_EXPERT | OPT_PERSTREAM | OPT_OUTPUT,
	Timecode                 typex.String   `json:"timecode" flag:"-timecode"`                                       //typex.String,   OPT_VIDEO | OPT_FUNC_ARG | OPT_PERFILE | OPT_OUTPUT | OPT_EXPERT,
	Vpre                     typex.String   `json:"vpre" flag:"-vpre"`                                               //typex.String, OPT_VIDEO | OPT_FUNC_ARG | OPT_EXPERT| OPT_PERFILE | OPT_OUTPUT | OPT_HAS_CANON,
	Apre                     typex.String   `json:"apre" flag:"-apre"`                                               //typex.String, OPT_FUNC_ARG | OPT_AUDIO | OPT_EXPERT| OPT_PERFILE | OPT_OUTPUT | OPT_HAS_CANON,
	Stag                     typex.String   `json:"stag" flag:"-stag"`                                               //typex.String, OPT_SUBTITLE | OPT_FUNC_ARG  | OPT_EXPERT  | OPT_PERFILE | OPT_OUTPUT | OPT_HAS_CANON,
	Spre                     typex.String   `json:"spre" flag:"-spre"`                                               //typex.String, OPT_FUNC_ARG | OPT_SUBTITLE | OPT_EXPERT| OPT_PERFILE | OPT_OUTPUT | OPT_HAS_CANON,
	//======
}

// Expert advanced param
type Expert struct {
	IgnoreUnknown       bool         `json:"ignore_unknown" flag:"-ignore_unknown"`               //         OPT_TYPE_BOOL, OPT_EXPERT,
	CopyUnknown         bool         `json:"copy_unknown" flag:"-copy_unknown"`                   //           OPT_TYPE_BOOL, OPT_EXPERT,
	Stdin               bool         `json:"stdin" flag:"-stdin"`                                 //                  OPT_TYPE_BOOL, OPT_EXPERT,
	NoStdin             bool         `json:"nostdin" flag:"-nostdin"`                             //                  OPT_TYPE_BOOL, OPT_EXPERT,
	Timelimit           typex.Func   `json:"timelimit" flag:"-timelimit"`                         //              OPT_TYPE_FUNC, OPT_FUNC_ARG | OPT_EXPERT,
	FrameDropThreshold  typex.Flt    `json:"frame_drop_threshold" flag:"-frame_drop_threshold"`   //   OPT_TYPE_FLOAT, OPT_EXPERT,
	Copyts              typex.Bool   `json:"copyts" flag:"-copyts"`                               //                 OPT_TYPE_BOOL, OPT_EXPERT,
	StartAtZero         bool         `json:"start_at_zero" flag:"-start_at_zero"`                 //          OPT_TYPE_BOOL, OPT_EXPERT,
	Copytb              typex.UBool  `json:"copytb" flag:"-copytb"`                               //                 OPT_TYPE_INT, OPT_EXPERT,
	DtsDeltaThreshold   typex.Flt    `json:"dts_delta_threshold" flag:"-dts_delta_threshold"`     //    OPT_TYPE_FLOAT, OPT_EXPERT,
	DtsErrorThreshold   typex.Flt    `json:"dts_error_threshold" flag:"-dts_error_threshold"`     //    OPT_TYPE_FLOAT, OPT_EXPERT,
	FilterComplexScript typex.Func   `json:"filter_complex_script" flag:"-filter_complex_script"` // OPT_TYPE_FUNC, OPT_FUNC_ARG | OPT_EXPERT,
	FindStreamInfo      bool         `json:"find_stream_info" flag:"-find_stream_info"`           //    OPT_TYPE_BOOL, OPT_INPUT | OPT_EXPERT | OPT_OFFSET,
	Vstats              typex.Func   `json:"vstats" flag:"-vstats"`                               //                     OPT_TYPE_FUNC,   OPT_VIDEO | OPT_EXPERT,
	VstatsFile          typex.Func   `json:"vstats_file" flag:"-vstats_file"`                     //                OPT_TYPE_FUNC,   OPT_VIDEO | OPT_FUNC_ARG | OPT_EXPERT,
	VstatsVersion       typex.Number `json:"vstats_version" flag:"-vstats_version"`               //             OPT_TYPE_INT,    OPT_VIDEO | OPT_EXPERT,
	Hwaccels            typex.Func   `json:"hwaccels" flag:"-hwaccels"`                           //                   OPT_TYPE_FUNC,   OPT_EXIT | OPT_EXPERT,
	VaapiDevice         typex.Func   `json:"vaapi_device" flag:"-vaapi_device"`                   // OPT_TYPE_FUNC, OPT_FUNC_ARG | OPT_EXPERT,
	QsvDevice           typex.Func   `json:"qsv_device" flag:"-qsv_device"`                       // OPT_TYPE_FUNC, OPT_FUNC_ARG | OPT_EXPERT,
	InitHwDevice        typex.Func   `json:"init_hw_device" flag:"-init_hw_device"`               // OPT_TYPE_FUNC, OPT_FUNC_ARG | OPT_EXPERT,
	FilterHwDevice      typex.Func   `json:"filter_hw_device" flag:"-filter_hw_device"`           // OPT_TYPE_FUNC, OPT_FUNC_ARG | OPT_EXPERT,
	AdriftThreshold     typex.Func   `json:"adrift_threshold" flag:"-adrift_threshold"`           // OPT_TYPE_FUNC, OPT_FUNC_ARG | OPT_EXPERT,
	Qphist              typex.Func   `json:"qphist" flag:"-qphist"`                               // OPT_TYPE_FUNC, OPT_VIDEO | OPT_EXPERT,
}

type Hwaccel = typex.String

const (
	Hwaccel_none Hwaccel = "none"
	//Do not use any hardware acceleration (the default).
	Hwaccel_auto Hwaccel = "auto"
	//Automatically select the hardware acceleration method.
	Hwaccel_vdpau Hwaccel = "vdpau"
	//Use VDPAU (Video Decode and Presentation API for Unix) hardware acceleration.
	Hwaccel_dxva2 Hwaccel = "dxva2"
	//Use DXVA2 (DirectX Video Acceleration) hardware acceleration.
	Hwaccel_d3d11va Hwaccel = "d3d11va"
	//Use D3D11VA (DirectX Video Acceleration) hardware acceleration.
	Hwaccel_vaapi Hwaccel = "vaapi"
	//Use VAAPI (Video Acceleration API) hardware acceleration.
	Hwaccel_qsv Hwaccel = "qsv"
	//Use the Intel QuickSync Video acceleration for video transcoding.
)

// -apply_cropping[:stream_specifier] source (input,per-stream)
// -init_hw_device type[=name]@source
// -init_hw_device list

type InitHwDevice = typex.String

const (
	InitHwDevice_cuda    InitHwDevice = "cuda"
	InitHwDevice_dxva2   InitHwDevice = "dxva2"
	InitHwDevice_d3d11va InitHwDevice = "d3d11va"
	InitHwDevice_vaapi   InitHwDevice = "vaapi"
	InitHwDevice_vdpau   InitHwDevice = "vdpau"
	InitHwDevice_qsv     InitHwDevice = "qsv"
	InitHwDevice_opencl  InitHwDevice = "opencl"
	InitHwDevice_vulkan  InitHwDevice = "vulkan"
)
