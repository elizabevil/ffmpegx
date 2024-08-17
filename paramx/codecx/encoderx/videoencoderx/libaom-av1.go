package videoencoderx

import (
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// // LIBAOM-AV1 9.7 libaom-av1
type LibaomAv1 struct {
	B typex.BitrateInt `json:"b" flag:"-b"` // int `json:"b" flag:"-g"` //(encoding, video)
	G typex.Size       `json:"g" flag:"-g"` // int `json:"b" flag:"-g"` //(encoding, video)

	Qmin typex.Quantizer `json:"qmin" flag:"-qmin"`
	Qmax typex.Quantizer `json:"qmax" flag:"-qmax"`
	//Set minimum/maximum quantisation values. Valid range is from 0 to 63 int //(warning: this does not match the quantiser values actually used by AV1 - divide by four to map real quantiser values to this range). Defaults to min/max int (no constraint).

	Minrate         typex.BitrateInt `json:"minrate" flag:"-minrate"`
	Maxrate         typex.BitrateInt `json:"maxrate" flag:"-maxrate"`
	Bufsize         typex.Size       `json:"bufsize" flag:"-bufsize"`
	RcInitOccupancy typex.Number     `json:"rc_init_occupancy" flag:"-rc_init_occupancy"`
	//Set rate control buffering parameters. Not used if not set - defaults to unconstrained variable bitrate.

	Threads typex.UNumber `json:"threads" flag:"-threads"`
	//Set the number of threads to use while encoding. This may require the tiles or row-mt options to also be set to actually use the specified number of threads fully. Defaults to the number of hardware threads supported by the host machine.

	Profile typex.FileName `json:"profile" flag:"-profile"`
	//Set the encoding profile. Defaults to using the profile which matches the bit depth and chroma subsampling of the input.

	//The wrapper also has some specific options:

	CpuUsed typex.Level `json:"cpu-used" flag:"-cpu-used"`
	//Set the quality/encoding speed tradeoff. Valid range is from 0 to 8, higher numbers indicating greater speed and lower quality. The default value is 1, which will be slow and high quality.

	AutoAltRef typex.Bool `json:"auto-alt-ref" flag:"-auto-alt-ref"`
	//Enable use of alternate reference frames. Defaults to the internal default of the library.

	ArnrMaxFrames typex.Number `json:"arnr-max-frames" flag:"-arnr-max-frames"` //(frames)
	//Set altref noise reduction max frame count. Default is -1.

	ArnrStrength typex.Level `json:"arnr-strength" flag:"-arnr-strength"` //(strength)
	//Set altref noise reduction filter strength. Range is -1 to 6. Default is -1.

	AqMode AqMode `json:"aq-mode" flag:"-aq-mode"` //(aq-mode)
	Tune   Tune   `json:"tune" flag:"-tune"`       //(tune)

	LagInFrames typex.UNumber `json:"lag-in-frames" flag:"-lag-in-frames"`
	//Set the maximum number of frames which the encoder may keep in flight at any one time for lookahead purposes. Defaults to the internal default of the library.

	ErrorResilience typex.Number `json:"error-resilience" flag:"-error-resilience"`
	Crf             typex.UI8    `json:"crf" flag:"-crf"`
	//Set the quality/size tradeoff for constant-quality int(no bitrate target) and constrained-quality int(with maximum bitrate target) modes. Valid range is 0 to 63, higher numbers indicating lower quality and smaller output size. Only used if set; by default only the bitrate target is used.

	StaticThresh typex.UNumber `json:"static-thresh" flag:"-static-thresh"`
	//Set a change threshold on blocks below which they will be skipped by the encoder. Defined in arbitrary units as a nonnegative integer, defaulting to zero int (no blocks are skipped).

	DropThreshold typex.Frames `json:"drop-threshold" flag:"-drop-threshold"`
	//Set a threshold for dropping frames when close to rate control bounds. Defined as a percentage of the target buffer - when the rate control buffer falls below this percentage, frames will be dropped until it has refilled above the threshold. Defaults to zero int (no frames are dropped).

	DenoiseNoiseLevel typex.UNumber `json:"denoise-noise-level" flag:"-denoise-noise-level"` //(level)
	//Amount of noise to be removed for grain synthesis. Grain synthesis is disabled if this option is not set or set to 0.

	DenoiseBlockSize typex.Size `json:"denoise-block-size" flag:"-denoise-block-size"` //(pixels)
	//Block size used for denoising for grain synthesis. If not set, AV1 codec uses the default value of 32.

	UndershootPct typex.PercentI `json:"undershoot-pct" flag:"-undershoot-pct"` //(pct)
	///Set datarate undershoot (min) percentage of the target bitrate. Range is -1 to 100. Default is -1.
	OvershootPct typex.UI16 `json:"overshoot-pct" flag:"-overshoot-pct"` //(pct)
	//Set datarate overshoot (max) percentage of the target bitrate. Range is -1 to 1000. Default is -1.

	MinsectionPct typex.PercentI `json:"minsection-pct" flag:"-minsection-pct"` //(pct)
	//Minimum percentage variation of the GOP bitrate from the target bitrate. If minsection-pct is not set, the libaomenc wrapper computes it as follows: (minrate * 100 / bitrate). Range is -1 to 100. Default is -1 (unset).
	MaxsectionPct typex.UI16 `json:"maxsection-pct" flag:"-maxsection-pct"` //(pct)
	//Maximum percentage variation of the GOP bitrate from the target bitrate. If maxsection-pct is not set, the libaomenc wrapper computes it as follows: (maxrate * 100 / bitrate). Range is -1 to 5000. Default is -1 (unset).

	FrameParallel typex.Bool `json:"frame-parallel" flag:"-frame-parallel"` //(boolean)
	//Enable frame parallel decodability features. Default is true.

	Tiles typex.TileCount `json:"tiles" flag:"-tiles"`
	//Set the number of tiles to encode the input video with, as columns x rows. Larger numbers allow greater parallelism in both encoding and decoding, but may decrease coding efficiency. Defaults to the minimum number of tiles required by the size of the input video int (this is 1x1 int (that is, a single tile) for sizes up to and including 4K).

	TileColumns typex.TileCount `json:"tile-columns" flag:"-tile-columns"`
	TileRows    typex.TileCount `json:"tile-rows" flag:"-tile-rows"`
	//Set the number of tiles as log2 of the number of tile rows and columns. Provided for compatibility with libvpx/VP9.

	///========
	RowMt typex.Bool `json:"row-mt" flag:"-row-mt"` //(Requires libaom >= 1.0.0-759-g90a15f4f2)
	//Enable row based multi-threading. Disabled by default.

	EnableCdef typex.UBool `json:"enable-cdef" flag:"-enable-cdef"` //(boolean)
	//Enable Constrained Directional Enhancement Filter. The libaom-av1 encoder enables CDEF by default.

	EnableRestoration typex.UBool `json:"enable-restoration" flag:"-enable-restoration"` //(boolean)
	//Enable Loop Restoration Filter. Default is true for libaom-av1.

	EnableGlobalMotion typex.UBool `json:"enable-global-motion" flag:"-enable-global-motion"` //(boolean)
	//Enable the use of global motion for block prediction. Default is true.

	EnableIntrabc typex.UBool `json:"enable-intrabc" flag:"-enable-intrabc"` //(boolean)
	//Enable block copy mode for intra block prediction. This mode is useful for screen content. Default is true.

	EnableRectPartitions typex.UBool `json:"enable-rect-partitions" flag:"-enable-rect-partitions"` //(boolean) int (Requires libaom >= v2.0.0)
	//Enable rectangular partitions. Default is true.

	Enable1To4Partitions typex.UBool `json:"enable-1to4-partitions" flag:"-enable-1to4-partitions"` //(boolean) int (Requires libaom >= v2.0.0)
	//Enable 1:4/4:1 partitions. Default is true.

	EnableAbPartitions typex.UBool `json:"enable-ab-partitions" flag:"-enable-ab-partitions"` //(boolean) int (Requires libaom >= v2.0.0)
	//Enable AB shape partitions. Default is true.

	EnableAngleDelta typex.UBool `json:"enable-angle-delta" flag:"-enable-angle-delta"` //(boolean) int (Requires libaom >= v2.0.0)
	//Enable angle delta intra prediction. Default is true.

	EnableCflIntra typex.UBool `json:"enable-cfl-intra" flag:"-enable-cfl-intra"` //(boolean) int (Requires libaom >= v2.0.0)
	//Enable chroma predicted from luma intra prediction. Default is true.

	EnableFilterIntra typex.UBool `json:"enable-filter-intra" flag:"-enable-filter-intra"` //(boolean) int (Requires libaom >= v2.0.0)
	//Enable filter intra predictor. Default is true.

	EnableIntraEdgeFilter typex.UBool `json:"enable-intra-edge-filter" flag:"-enable-intra-edge-filter"` //(boolean) int (Requires libaom >= v2.0.0)
	//Enable intra edge filter. Default is true.

	EnableSmoothIntra typex.UBool `json:"enable-smooth-intra" flag:"-enable-smooth-intra"` //(boolean) int (Requires libaom >= v2.0.0)
	//Enable smooth intra prediction mode. Default is true.

	EnablePaethIntra typex.UBool `json:"enable-paeth-intra" flag:"-enable-paeth-intra"` //(boolean) int (Requires libaom >= v2.0.0)
	//Enable paeth predictor in intra prediction. Default is true.

	EnablePalette typex.UBool `json:"enable-palette" flag:"-enable-palette"` //(boolean) int (Requires libaom >= v2.0.0)
	//Enable palette prediction mode. Default is true.

	EnableFlipIdtx typex.UBool `json:"enable-flip-idtx" flag:"-enable-flip-idtx"` //(boolean) int (Requires libaom >= v2.0.0)
	//Enable extended transform type, including FLIPADST_DCT, DCT_FLIPADST, FLIPADST_FLIPADST, ADST_FLIPADST, FLIPADST_ADST, IDTX, V_DCT, H_DCT, V_ADST, H_ADST, V_FLIPADST, H_FLIPADST. Default is true.

	EnableTx64 typex.UBool `json:"enable-tx64" flag:"-enable-tx64"` //(boolean) int (Requires libaom >= v2.0.0)
	//Enable 64-pt transform. Default is true.

	ReducedTxTypeSet typex.Bool `json:"reduced-tx-type-set" flag:"-reduced-tx-type-set"` //(boolean) int (Requires libaom >= v2.0.0)
	//Use reduced set of transform types. Default is false.

	UseIntraDctOnly typex.Bool `json:"use-intra-dct-only" flag:"-use-intra-dct-only"` //(boolean) int (Requires libaom >= v2.0.0)
	//Use DCT only for INTRA modes. Default is false.

	UseInterDctOnly typex.Bool `json:"use-inter-dct-only" flag:"-g"` //(boolean) int (Requires libaom >= v2.0.0)
	//Use DCT only for INTER modes. Default is false.

	UseIntraDefaultTxOnly typex.Bool `json:"use-intra-default-tx-only" flag:"-g"` //(boolean) int (Requires libaom >= v2.0.0)
	//Use Default-transform only for INTRA modes. Default is false.

	EnableRefFrameMvs typex.UBool `json:"enable-ref-frame-mvs" flag:"-g"` //(boolean) int (Requires libaom >= v2.0.0)
	//Enable temporal mv prediction. Default is true.

	EnableReducedReferenceSet typex.UBool `json:"enable-reduced-reference-set" flag:"-g"` //(boolean) int (Requires libaom >= v2.0.0)
	//Use reduced set of single and compound references. Default is false.

	EnableObmc typex.UBool `json:"enable-obmc" flag:"-g"` //(boolean) int (Requires libaom >= v2.0.0)
	//Enable obmc. Default is true.

	EnableDualFilter typex.UBool `json:"enable-dual-filter" flag:"-enable-dual-filter"` //(boolean) int (Requires libaom >= v2.0.0)
	//Enable dual filter. Default is true.

	EnableDiffWtdComp typex.UBool `json:"enable-diff-wtd-comp" flag:"-enable-diff-wtd-comp"` //(boolean) int (Requires libaom >= v2.0.0)
	//Enable difference-weighted compound. Default is true.

	EnableDistWtdComp typex.UBool `json:"enable-dist-wtd-comp" flag:"-enable-dist-wtd-comp"` //(boolean) int (Requires libaom >= v2.0.0)
	//Enable distance-weighted compound. Default is true.

	EnableOnesidedComp typex.UBool `json:"enable-onesided-comp" flag:"-enable-onesided-comp"` //(boolean) int (Requires libaom >= v2.0.0)
	//Enable one sided compound. Default is true.

	EnableInterinterWedge typex.UBool `json:"enable-interinter-wedge" flag:"-enable-interinter-wedge"` //(boolean) int (Requires libaom >= v2.0.0)
	//Enable interinter wedge compound. Default is true.

	EnableInterintraWedge typex.UBool `json:"enable-interintra-wedge" flag:"-enable-interintra-wedge"` //(boolean) int (Requires libaom >= v2.0.0)
	//Enable interintra wedge compound. Default is true.

	EnableMaskedComp typex.UBool `json:"enable-masked-comp" flag:"-enable-masked-comp"` //(boolean) int (Requires libaom >= v2.0.0)
	//Enable masked compound. Default is true.

	EnableInterintraComp typex.UBool `json:"enable-interintra-comp" flag:"-enable-interintra-comp"` //(boolean) int (Requires libaom >= v2.0.0)
	//Enable interintra compound. Default is true.

	EnableSmoothInterintra typex.UBool `json:"enable-smooth-interintra" flag:"-enable-smooth-interintra"` //(boolean) int (Requires libaom >= v2.0.0)
	//Enable smooth interintra mode. Default is true.

	AomParams string `json:"aom-params" flag:"-aom-params"`
}

type AqMode = typex.String

const (
	AqMode_none AqMode = "none" // (0)’
	//Disabled.

	AqMode_variance AqMode = "variance" // (1)’
	//Variance-based.

	AqMode_complexity AqMode = "complexity" // (2)’
	//Complexity-based.

	AqMode_cyclic AqMode = "cyclic" // (3)’
	//Cyclic refresh.
)

type Tune = typex.String

const (
	Tune_psnr        Tune = "psnr"        // (0)
	Tune_ssim        Tune = "ssim"        // (1)
	Tune_none        Tune = "none"        // (1)
	Tune_zerolatency Tune = "zerolatency" // (1)
)

/*

For example to specify libaom encoding options with -aom-params:

ffmpeg -i input -c:v libaom-av1 -b:v 500K -aom-params tune=psnr:enable-tpl-model=1 output.mp4

*/
