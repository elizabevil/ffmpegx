package videoencoderx

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

type LIBVPX struct {
	B typex.Bitrate `json:"b" flag:"-b"`
	G typex.Size    `json:"g" flag:"-g"`

	KeyintMin typex.Number    `json:"keyint_min" flag:"-keyint_min"` // (kf-min-dist)
	Qmin      typex.Quantizer `json:"qmin" flag:"-qmin"`             // (min-q)
	Qmax      typex.Quantizer `json:"qmax" flag:"-qmax"`             // (max-q)

	Bufsize         typex.Size   `json:"bufsize" flag:"-bufsize"`                     // (buf-sz, buf-optimal-sz)
	RcInitOccupancy typex.Number `json:"rc_init_occupancy" flag:"-rc_init_occupancy"` // (buf-initial-sz)

	UndershootPct typex.UI8 `json:"undershoot-pct" flag:"-undershoot-pct"` // (buf-initial-sz)
	OvershootPct  typex.UI8 `json:"overshoot-pct" flag:"-overshoot-pct"`

	SkipThreshold typex.Number `json:"skip_threshold" flag:"-skip_threshold"` // (drop-frame)
	Qcomp         typex.Flt    `json:"qcomp" flag:"-qcomp"`                   // (bias-pct)

	Minrate typex.BitrateInt `json:"minrate" flag:"-minrate"`
	Maxrate typex.BitrateInt `json:"maxrate" flag:"-maxrate"`

	Crf      typex.UI8     `json:"crf" flag:"-crf"`   // (end-usage = cq, cq-level)
	Tune     Tune          `json:"tune" flag:"-tune"` // (tune)
	Quality  Quality       `json:"quality" flag:"-quality"`
	Deadline Quality       `json:"deadline" flag:"-deadline"` // (deadline)
	Speed    typex.SpeedI8 `json:"speed" flag:"-speed"`
	CpuUsed  typex.Level   `json:"cpu-used" flag:"-cpu-used"` // (cpu-used)

	Nr           typex.Level   `json:"nr" flag:"-nr"` // (noise-sensitivity)
	StaticThresh typex.UNumber `json:"static-thresh" flag:"-static-thresh"`

	Slices typex.Size `json:"slices" flag:"-slices"` // (token-parts)

	MaxIntraRate typex.UNumber `json:"max-intra-rate" flag:"-max-intra-rate"`

	ForceKeyFrames Ssim `json:"force_key_frames" flag:"-force_key_frames"`

	ErrorResilient typex.Number `json:"error-resilient" flag:"-error-resilient"`

	Sharpness typex.Level `json:"sharpness" flag:"-sharpness"`

	TsParameters string `json:"ts-parameters" flag:"-ts-parameters"`

	//VP8-specific options
	ScreenContentMode typex.Level `json:"screen-content-mode" flag:"-screen-content-mode"`

	//Screen content mode, one of: 0 (off), 1 (screen), 2 (screen with more aggressive rate control).
	//VP9-specific options

	Lossless typex.Bool `json:"lossless" flag:"-lossless"`
	//Enable lossless mode.

	TileColumns typex.TileCount `json:"tile-columns" flag:"-tile-columns"`
	//Set number of tile columns to use. Note this is given as log2(tile_columns). For example, 8 tile columns would be requested by setting the tile-columns option to 3.

	TileRows typex.TileCount `json:"tile-rows" flag:"-tile-rows"`
	//Set number of tile rows to use. Note this is given as log2(tile_rows). For example, 4 tile rows would be requested by setting the tile-rows option to 2.

	FrameParallel typex.Bool `json:"frame-parallel" flag:"-frame-parallel"`
	//Enable frame parallel decodability features.

	AqMode AqMode `json:"aq-mode" flag:"-aq-mode"`
	//Set adaptive quantization mode (0: off (default), 1: variance 2: complexity, 3: cyclic refresh, 4: equator360).

	Colorspace flagx.Colorspace `json:"colorspace" flag:"-colorspace"`

	RowMt typex.Bool `json:"row-mt" flag:"-row-mt"`
	//Enable row based multi-threading.

	TuneContent typex.Level `json:"tune-content" flag:"-tune-content"`
	//Set content type: default (0), screen (1), film (2).

	CorpusComplexity typex.UVBR `json:"corpus-complexity" flag:"-corpus-complexity"`
	//Corpus UVBR mode is a variant of standard UVBR where the complexity distribution midpoint is passed in rather than calculated for a specific clip or chunk.

	//The valid range is [0, 10000]. 0 (default) uses standard UVBR.

	EnableTpl typex.Bool `json:"enable-tpl" flag:"-enable-tpl"`
	//Enable temporal dependency model.

	RefFrameConfig string `json:"ref-frame-config" flag:"-ref-frame-config"`
}
type Quality = typex.String

const (
	best Quality = "best"
	//Use best quality deadline. Poorly named and quite slow, this option should be avoided as it may give worse quality output than good.

	good Quality = "good"
	//Use good quality deadline. This is a good trade-off between speed and quality when used with the cpu-used option.

	realtime Quality = "realtime"
	//Use realtime quality deadline.
)

type TsParameters = typex.String

const (
	ts_number_layers TsParameters = "ts_number_layers"
	//Number of temporal coding layers.

	ts_target_bitrate TsParameters = "ts_target_bitrate"
	//Target bitrate for each temporal layer (in kbps). (bitrate should be inclusive of the lower temporal layer).

	ts_rate_decimator TsParameters = "ts_rate_decimator"
	//Frame rate decimation factor for each temporal layer.

	ts_periodicity TsParameters = "ts_periodicity"
	//Length of the sequence defining frame temporal layer membership.

	ts_layer_id TsParameters = "ts_layer_id"
	//Template defining the membership of frames to temporal layers.

	ts_layering_mode TsParameters = "ts_layering_mode"
	//(optional) Selecting the temporal structure from a set of pre-defined temporal layering modes. Currently supports the following options.
)

/*


ffmpeg -i INPUT -c:v libvpx -ts-parameters ts_number_layers=3:\
ts_target_bitrate=250,500,1000:ts_rate_decimator=4,2,1:\
ts_periodicity=4:ts_layer_id=0,2,1,2:ts_layering_mode=3 OUTPUT
*/
