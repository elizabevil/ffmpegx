package videoencoderx

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// // LIBX264, 9.16 libx264, libx264rgb
type LIBX264 struct {
	B  typex.UNumber `json:"b" flag:"-b"`   // (encoding,audio,video)
	Bf typex.UNumber `json:"bf" flag:"-bf"` //(bframes)
	//Number of B-frames between I and P-frames

	G typex.Size `json:"g" flag:"-g"` //(keyint)
	//Maximum GOP size

	Qmin typex.Quantizer `json:"qmin" flag:"-qmin"` // (qpmin)
	//Minimum quantizer scale

	Qmax typex.Quantizer `json:"qmax" flag:"-qmax"` //(qpmax)
	//Maximum quantizer scale

	Qdiff typex.Number `json:"qdiff" flag:"-qdiff"` //(qpstep)
	//Maximum difference between quantizer scales

	Qblur typex.UFloat32 `json:"qblur" flag:"-qblur"` // (qblur)
	//Quantizer curve blur

	Qcomp typex.Flt `json:"qcomp" flag:"-qcomp"` //(qcomp)
	//Quantizer curve compression factor

	Refs        typex.Number `json:"refs" flag:"-refs"`                 //(ref)
	ScThreshold typex.Number `json:"sc_threshold" flag:"-sc_threshold"` // (scenecut)
	//Sets the threshold for the scene change detection.

	Trellis typex.Number `json:"trellis" flag:"-trellis"` //(trellis)
	//Performs Trellis quantization to increase efficiency. Enabled by default.

	Nr typex.Number `json:"nr" flag:"-nr"` //(nr)
	//Noise reduction

	MeRange typex.Number `json:"me_range" flag:"-me_range"` //(merange)
	//Maximum range of the motion search in pixels.

	MeMethod  MeMethod     `json:"me_method" flag:"-me_method"`   //(me)
	ForcedIdr typex.Number `json:"forced-idr" flag:"-forced-idr"` //
	//Normally, when forcing a I-frame type, the encoder can select any type of I-frame. This option forces it to choose an IDR-frame.

	Subq typex.Number `json:"subq" flag:"-subq"` //(subme)
	//Sub-pixel motion estimation method.

	BStrategy typex.Bool `json:"b_strategy" flag:"-b_strategy"` // (b-adapt)
	//Adaptive B-frame placement decision algorithm. Use only on first-pass.

	KeyintMin typex.Number `json:"keyint_min" flag:"-keyint_min"` // (min-keyint)
	//Minimum GOP size.

	Coder   VaapiH264Coder `json:"coder" flag:"-coder"`     //ac,vlc
	Cmp     flagx.Cmp      `json:"cmp" flag:"-cmp"`         //chroma ,sad
	Threads typex.UNumber  `json:"threads" flag:"-threads"` //(threads)
	//Number of encoding threads.

	ThreadType flagx.ThreadType `json:"thread_type" flag:"-thread_type"` //
	Flags      int              `json:"flags" flag:"-flags"`             //
	//Set encoding flags. It can be used to disable closed GOP and enable open GOP by setting it to -cgop. The result is similar to the behavior of x264s --open-gop option.

	RcInitOccupancy typex.Number `json:"rc_init_occupancy" flag:"-rc_init_occupancy"` //(vbv-init)
	//Initial VBV buffer occupancy

	Preset flagx.Preset `json:"preset" flag:"-preset"` // (preset)
	//Set the encoding preset.

	Tune Tune `json:"tune" flag:"-tune"` //(tune)
	//Set tuning of the encoding params.

	Profile int `json:"profile" flag:"-profile"` //(profile)
	//Set profile restrictions.

	Fastfirstpass int `json:"fastfirstpass" flag:"-fastfirstpass"` //
	//Enable fast settings when encoding first pass, when set to 1. When set to 0, it has the same effect of x264s --slow-firstpass option.

	Crf typex.UI8 `json:"crf" flag:"-crf"` // (crf)
	//Set the quality for constant quality mode.

	CrfMax int `json:"crf_max" flag:"-crf_max"` //(crf-max)
	//In CRF mode, prevents VBV from lowering quality beyond this point.

	Qp typex.Level `json:"qp" flag:"-qp"` //(qp)
	//Set constant quantization rate control method parameter.

	AqMode     AqMode `json:"aq-mode" flag:"-aq-mode"`         //(aq-mode)
	AqStrength int    `json:"aq-strength" flag:"-aq-strength"` //(aq-strength)
	//Set AQ strength, reduce blocking and blurring in flat and textured areas.

	Psy int `json:"psy" flag:"-psy"` //
	//Use psychovisual optimizations when set to 1. When set to 0, it has the same effect as x264s --no-psy option.

	PsyRd int `json:"psy-rd" flag:"-psy-rd"` // (psy-rd)
	//Set strength of psychovisual optimization, in psy-rd:psy-trellis format.

	RcLookahead int `json:"rc-lookahead" flag:"-rc-lookahead"` //(rc-lookahead)
	//Set number of frames to look ahead for frametype and ratecontrol.

	Weightb int `json:"weightb" flag:"-weightb"` //
	//Enable weighted prediction for B-frames when set to 1. When set to 0, it has the same effect as x264s --no-weightb option.

	Weightp Weightp    `json:"weightp" flag:"-weightp"` //(weightp)
	Ssim    typex.Bool `json:"ssim" flag:"-ssim"`       //(ssim)
	//Enable calculation and printing SSIM stats after the encoding.

	IntraRefresh typex.Bool `json:"intra-refresh" flag:"-intra-refresh"` //(intra-refresh)
	//Enable the use of Periodic Intra Refresh instead of IDR frames when set to 1.

	AvcintraClass int `json:"avcintra-class" flag:"-avcintra-class"` // (class)
	//Configure the encoder to generate AVC-Intra. Valid values are 50, 100 and 200

	BlurayCompat typex.Bool `json:"bluray-compat" flag:"-bluray-compat"` //(bluray-compat)
	//Configure the encoder to be compatible with the bluray standard. It is a shorthand for setting "bluray-compat=1 force-cfr=1".

	BBias typex.Number `json:"b-bias" flag:"-b-bias"` // (b-bias)
	//Set the influence on how often B-frames are used.

	BPyramid BPyramid `json:"b-pyramid" flag:"-b-pyramid"` //(b-pyramid)
	//Set method for keeping of some B-frames as references. Possible values:
	MixedRefs typex.Bool `json:"mixed-refs" flag:"-mixed-refs"`

	//Enable the use of one reference per partition, as opposed to one reference per macroblock when set to 1. When set to 0, it has the same effect as x264s --no-mixed-refs option.

	A8X8Dct typex.Bool `json:"8x8dct" flag:"-8x8dct"` //
	//Enable adaptive spatial transform (high profile 8x8 transform) when set to 1. When set to 0, it has the same effect as x264s --no-8x8dct option.

	FastPskip typex.Bool `json:"fast-pskip" flag:"-fast-pskip"` //
	//Enable early SKIP detection on P-frames when set to 1. When set to 0, it has the same effect as x264s --no-fast-pskip option.

	Aud typex.Bool `json:"aud" flag:"-aud"` //(aud)
	//Enable use of access unit delimiters when set to 1.

	Mbtree typex.Bool `json:"mbtree" flag:"-mbtree"` //
	//Enable use macroblock tree ratecontrol when set to 1. When set to 0, it has the same effect as x264s --no-mbtree option.

	Deblock string `json:"deblock" flag:"-deblock"` //(deblock)
	//Set loop filter parameters, in alpha:beta form.

	Cplxblur typex.Bool `json:"cplxblur" flag:"-cplxblur"` //(cplxblur)
	//Set fluctuations reduction in QP (before curve compression).

	Partitions Partitions `json:"partitions" flag:"-partitions"` // (partitions)
	//Set partitions to consider as a comma-separated list of values. Possible values in the list:
	DirectPred DirectPred `json:"direct-pred" flag:"-direct-pred"` // (direct)
	//Set direct MV prediction mode. Possible values:
	SliceMaxSize typex.Size `json:"slice-max-size" flag:"-slice-max-size"` //(slice-max-size)
	//Set the limit of the size of each slice in bytes. If not specified but RTP payload size (ps) is specified, that is used.

	Stats string `json:"stats" flag:"-stats"` //(stats)
	//Set the file name for multi-pass stats.

	NalHrd NalHrd `json:"nal-hrd" flag:"-nal-hrd"` // (nal-hrd)
	//Set signal HRD information (requires vbv-bufsize to be set). Possible values:
	X264Opts   string `json:"x264opts" flag:"-x264opts"`       //opts
	X264Params string `json:"x264-params" flag:"-x264-params"` // opts
	//Override the x264 configuration using a :-separated list of key=value options.
	A53CcBoolean typex.UBool `json:"a53cc" flag:"-a53cc"` //
	//Import closed captions (which must be ATSC compatible format) into output. Only the mpeg2 and h264 decoders provide these. Default is 1 (on).

	UduSeiBoolean typex.UBool `json:"udu_sei" flag:"-udu_sei"` //
	//Import user data unregistered SEI if available into output. Default is 0 (off).

	MbInfoBoolean typex.UBool `json:"mb_info" flag:"-mb_info"` //
	//Set mb_info data through AVFrameSideData, only useful when used from the API. Default is 0 (off).
}

//ffmpeg -i foo.mpg -c:v libx264 -x264opts keyint=123:min-keyint=20 -an out.mkv

type NalHrd = typex.String

const (
	NalHrd_none     NalHrd = "none"
	NalHrd_spatial  NalHrd = "vbr"
	NalHrd_temporal NalHrd = "cbr"
)

type MotionEst = typex.String

const (
	MotionEst_dia  MotionEst = "dia"
	MotionEst_hex  MotionEst = "hex"
	MotionEst_umh  MotionEst = "umh"
	MotionEst_esa  MotionEst = "esa"
	MotionEst_tesa MotionEst = "tesa"
)

type DirectPred = typex.String

const (
	DirectPred_none     DirectPred = "none"
	DirectPred_spatial  DirectPred = "spatial"
	DirectPred_temporal DirectPred = "temporal"
	DirectPred_auto     DirectPred = "auto"
)

type Partitions = typex.String

const (
	Partitions_p8x8 Partitions = "p8x8"
	//8x8 P-frame partition.

	Partitions_p4x4 Partitions = "p4x4"
	//4x4 P-frame partition.

	Partitions_b8x8 Partitions = "b8x8"
	//4x4 B-frame partition.

	Partitions_i8x8 Partitions = "i8x8"
	//8x8 I-frame partition.

	Partitions_i4x4 Partitions = "i4x4"
	//4x4 I-frame partition. (Enabling p4x4 requires p8x8 to be enabled. Enabling i8x8 requires adaptive spatial transform (8x8dct option) to be enabled.)

	Partitions_none Partitions = "none" // (none)
	//Do not consider any partitions.

	Partitions_all Partitions = "all" //(all)
	//Consider every partition.
)

type MeMethod = typex.String

const (
	MeMethod_dia MeMethod = "dia" // (dia)
	//
	MeMethod_epzs MeMethod = "epzs" // (dia)
	//Diamond search with radius 1 (fastest). epzs is an alias for dia.

	MeMethod_hex MeMethod = "hex" // (hex)
	//Hexagonal search with radius 2.

	MeMethod_umh MeMethod = "umh" // (umh)
	//Uneven multi-hexagon search.

	MeMethod_esa MeMethod = "esa" // (esa)
	//Exhaustive search.

	MeMethod_tesa MeMethod = "tesa" // (tesa)
	//Hadamard exhaustive search (slowest).
)

type Weightp = typex.String

const (
	Weightp_none   Weightp = "none"   // (dia)
	Weightp_simple Weightp = "simple" // (dia)
	Weightp_smart  Weightp = "smart"  // (dia)
	//

)

type BPyramid = typex.String

const (
	BPyramid_none   BPyramid = "none"   // (dia)
	BPyramid_strict BPyramid = "strict" // (dia)
	BPyramid_normal BPyramid = "normal" // (dia)
	//

)
