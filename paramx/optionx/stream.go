package optionx

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

type InputStream struct {
	Itsscale            typex.Scale   `json:"itsscale"         flag:"-itsscale"`                   // scale (input, per-stream) scale=w=200:h=100
	ReinitFilter        typex.Number  `json:"reinit_filter"    flag:"-reinit_filter"`              // [:stream_specifier] integer (input, per-stream)
	DumpAttachment      typex.String  `json:"dump_attachment"  flag:"-dump_attachment"`            // [:stream_specifier] filename (input, per-stream)
	DisplayRotation     typex.Flt     `json:"display_rotation" flag:"-display_rotation"`           // [:stream_specifier] rotation (input, per-stream)
	DisplayHflip        bool          `json:"display_hflip"    flag:"-display_hflip"`              // [:stream_specifier] (input, per-stream)
	DisplayVflip        bool          `json:"display_vflip"    flag:"-display_vflip"`              // [:stream_specifier] (input, per-stream)
	ApplyCropping       typex.UBool   `json:"apply_cropping"   flag:"-apply_cropping"`             // [:stream_specifier] source (input, per-stream)
	Hwaccel             Hwaccel       `json:"hwaccel"          flag:"-hwaccel"`                    // [:stream_specifier] hwaccel (input, per-stream)
	HwaccelDevice       typex.String  `json:"hwaccel_device"   flag:"-hwaccel_device"`             // [:stream_specifier] hwaccel_device (input, per-stream)
	GuessLayoutMax      typex.Number  `json:"guess_layout_max" flag:"-guess_layout_max"`           // channels (input, per-stream)
	Discard             flagx.Discard `json:"discard" flag:"-discard"`                             //OPT_TYPE_STRING, OPT_PERSTREAM | OPT_INPUT | OPT_EXPERT,
	HwaccelOutputFormat typex.String  `json:"hwaccel_output_format" flag:"-hwaccel_output_format"` //OPT_TYPE_STRING, OPT_VIDEO | OPT_EXPERT | OPT_PERSTREAM | OPT_INPUT,
	Autorotate          bool          `json:"autorotate" flag:"-autorotate"`                       //OPT_TYPE_BOOL,   OPT_PERSTREAM | OPT_EXPERT | OPT_INPUT,
	NoAutorotate        bool          `json:"noautorotate" flag:"-noautorotate"`                   //                 OPT_TYPE_BOOL,   OPT_PERSTREAM | OPT_EXPERT | OPT_INPUT,

	FixSubDuration bool            `json:"fix_sub_duration" flag:"-fix_sub_duration"` //OPT_TYPE_BOOL, OPT_EXPERT | OPT_SUBTITLE | OPT_PERSTREAM | OPT_INPUT,
	CanvasSize     typex.VideoSize `json:"canvas_size" flag:"-canvas_size"`           //OPT_TYPE_STRING, OPT_SUBTITLE | OPT_PERSTREAM | OPT_INPUT | OPT_EXPERT,
}

type OutputStream struct {
	Disposition              typex.Disposition       `json:"disposition"                 flag:"-disposition"`                 // [:stream_specifier] value (output, per-stream)     ffmpeg -i in.mkv -c copy -disposition:s:0 0 -disposition:s:1 default out.mkv
	Frames                   typex.Number            `json:"frames"                      flag:"-frames"`                      // [:stream_specifier] framecount (output, per-stream)    ffmpeg -f fbdev -framerate 1 -i /dev/fb0 -frames:v 1 screenshot.jpeg
	Q                        []typex.StreamSpecifier `json:"q"                           flag:"-q"`                           // [:stream_specifier] q (output, per-stream)
	Qscale                   string                  `json:"qscale"                      flag:"-qscale"`                      // [:stream_specifier] q (output, per-stream)
	Filter                   typex.FilterGraph       `json:"filter"                      flag:"-filter"`                      // [:stream_specifier] filtergraph (output, per-stream)
	Pre                      string                  `json:"pre"                         flag:"-pre"`                         // [:stream_specifier] preset_name (output, per-stream)
	Fpsmax                   string                  `json:"fpsmax"                      flag:"-fpsmax"`                      // [:stream_specifier] fps (output, per-stream)
	Aspect                   typex.Aspect            `json:"aspect"                      flag:"-aspect"`                      // [:stream_specifier] aspect (output, per-stream)
	Pass                     typex.Level             `json:"pass"                        flag:"-pass"`                        // [:stream_specifier] n (output, per-stream)
	Passlogfile              string                  `json:"passlogfile"                 flag:"-passlogfile"`                 // [:stream_specifier] prefix (output, per-stream)
	RcOverride               typex.String            `json:"rc_override"                 flag:"-rc_override"`                 // [:stream_specifier] override (output, per-stream)
	ForceKeyFrames           typex.Key               `json:"force_key_frames"            flag:"-force_key_frames"`            // [:stream_specifier] time[, time...] (output, per-stream)
	Copyinkf                 bool                    `json:"copyinkf"                    flag:"-copyinkf"`                    // [:stream_specifier] (output, per-stream)
	SampleFmt                string                  `json:"sample_fmt"                  flag:"-sample_fmt"`                  // [:stream_specifier] sample_fmt (output, per-stream)
	FpsMode                  FpsMode                 `json:"fps_mode"                    flag:"-fps_mode"`                    // [:stream_specifier] parameter (output, per-stream)
	Apad                     string                  `json:"apad"                        flag:"-apad"`                        // parameters (output, per-stream)
	EncTimeBase              string                  `json:"enc_time_base"               flag:"-enc_time_base"`               // [:stream_specifier] timebase (output, per-stream)
	MaxMuxingQueueSize       typex.Size              `json:"max_muxing_queue_size"       flag:"-max_muxing_queue_size"`       // packets (output, per-stream)
	MuxingQueueDataThreshold typex.Size              `json:"muxing_queue_data_threshold" flag:"-muxing_queue_data_threshold"` // bytes (output, per-stream)
	BitsPerRawSample         typex.UNumber           `json:"bits_per_raw_sample"         flag:"-bits_per_raw_sample"`         // [:stream_specifier] value (output, per-stream)
	StatsEncPre              string                  `json:"stats_enc_pre"               flag:"-stats_enc_pre"`               // [:stream_specifier] path (output, per-stream)
	StatsEncPost             string                  `json:"stats_enc_post"              flag:"-stats_enc_post"`              // [:stream_specifier] path (output, per-stream)
	StatsMuxPre              string                  `json:"stats_mux_pre"               flag:"-stats_mux_pre"`               // [:stream_specifier] path (output, per-stream)
	StatsEncPreFmt           string                  `json:"stats_enc_pre_fmt"           flag:"-stats_enc_pre_fmt"`           // [:stream_specifier] format_spec (output, per-stream)
	StatsEncPostFmt          string                  `json:"stats_enc_post_fmt"          flag:"-stats_enc_post_fmt"`          // [:stream_specifier] format_spec (output, per-stream)
	StatsMuxPreFmt           string                  `json:"stats_mux_pre_fmt"           flag:"-stats_mux_pre_fmt"`           // [:stream_specifier] format_spec (output, per-stream)
	Copypriorss              typex.Number            `json:"copypriorss" flag:"-copypriorss"`                                 //OPT_TYPE_INT, OPT_EXPERT | OPT_PERSTREAM | OPT_OUTPUT,
	FilterScript             typex.String            `json:"filter_script" flag:"-filter_script"`                             //OPT_TYPE_STRING, OPT_PERSTREAM | OPT_EXPERT | OPT_OUTPUT,
	IntraMatrix              typex.String            `json:"intra_matrix" flag:"-intra_matrix"`                               //OPT_TYPE_STRING, OPT_VIDEO | OPT_EXPERT | OPT_PERSTREAM | OPT_OUTPUT,
	InterMatrix              typex.String            `json:"inter_matrix" flag:"-inter_matrix"`                               //OPT_TYPE_STRING, OPT_VIDEO | OPT_EXPERT | OPT_PERSTREAM | OPT_OUTPUT,
	ChromaIntraMatrix        typex.String            `json:"chroma_intra_matrix" flag:"-chroma_intra_matrix"`                 //OPT_TYPE_STRING, OPT_VIDEO | OPT_EXPERT | OPT_PERSTREAM | OPT_OUTPUT,
	ForceFps                 bool                    `json:"force_fps" flag:"-force_fps"`                                     //OPT_TYPE_BOOL,   OPT_VIDEO | OPT_EXPERT  | OPT_PERSTREAM | OPT_OUTPUT,
	Autoscale                bool                    `json:"autoscale" flag:"-autoscale"`                                     //OPT_TYPE_BOOL,   OPT_PERSTREAM | OPT_EXPERT | OPT_OUTPUT,
	Noautoscale              bool                    `json:"noautoscale" flag:"-noautoscale"`                                 //OPT_TYPE_BOOL,   OPT_PERSTREAM | OPT_EXPERT | OPT_OUTPUT,
	FixSubDurationHeartbeat  bool                    `json:"fix_sub_duration_heartbeat" flag:"-fix_sub_duration_heartbeat"`   //OPT_TYPE_BOOL,   OPT_VIDEO | OPT_EXPERT | OPT_PERSTREAM | OPT_OUTPUT,
	TimeBase                 typex.TimeBase          `json:"time_base" flag:"-time_base"`                                     //OPT_TYPE_STRING, OPT_EXPERT | OPT_PERSTREAM | OPT_OUTPUT,
}

type FpsMode = typex.String

const (
	FpsMode_passthrough FpsMode = "passthrough" // (0)
	//Each frame is passed with its timestamp from the demuxer to the muxer.

	FpsMode_cfr FpsMode = "cfr" // (1)
	//Frames will be duplicated and dropped to achieve exactly the requested constant frame rate.

	FpsMode_vfr FpsMode = "vfr" // (2)
	//Frames are passed through with their timestamp or dropped so as to prevent 2 frames from having the same timestamp.

	FpsMode_auto FpsMode = "auto" // (-1)
	//Chooses between cfr and vfr depending on muxer capabilities. This is the default method.
)

type EncTimeBase = typex.String

const (
	EncTimeBase_0 EncTimeBase = "0"
	//Assign a default value according to the media type.
	//For video - use 1/framerate, for audio - use 1/samplerate.

	EncTimeBase_demux EncTimeBase = "demux"
	//Use the timebase from the demuxer.

	EncTimeBase_filter EncTimeBase = "filter"
	//Use the timebase from the filtergraph.

	//a positive number
	//Use the provided number as the timebase.
)
