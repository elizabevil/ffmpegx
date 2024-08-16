package audioencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// // FLAC 8.3 flac
type FLAC struct {
	CompressionLevel typex.Level `json:"compression_level" flag:"-compression_level"`
	//Sets the compression level, which chooses defaults for many other options if they are not set explicitly. Valid values are from 0 to 12, 5 is the default.

	FrameSize typex.FrameSize `json:"frame_size" flag:"-frame_size"`
	//Sets the size of the frames in samples per channel.

	LpcCoeffPrecision typex.Level `json:"lpc_coeff_precision" flag:"-lpc_coeff_precision"`
	//Sets the LPC coefficient precision, valid values are from 1 to 15, 15 is the default.

	LpcType LpcType `json:"lpc_type" flag:"-lpc_type"`

	LpcPasses typex.UNumber `json:"lpc_passes" flag:"-lpc_passes"`
	//Number of passes to use for Cholesky factorization during LPC analysis

	MinPartitionOrder typex.Bool `json:"min_partition_order" flag:"-min_partition_order"`
	//The minimum partition order

	MaxPartitionOrder typex.Bool `json:"max_partition_order" flag:"-max_partition_order"`
	//The maximum partition order

	PredictionOrderMethod PredictionOrderMethod `json:"prediction_order_method" flag:"-prediction_order_method"`
	ChMode                ChMode                `json:"ch_mode" flag:"-ch_mode"`

	ExactRiceParameters typex.UBool `json:"exact_rice_parameters" flag:"-exact_rice_parameters"`
	//Chooses if rice parameters are calculated exactly or approximately. if set to 1 then they are chosen exactly, which slows the code down slightly and improves compression slightly.

	MultiDimQuant typex.UBool `json:"multi_dim_quant" flag:"-multi_dim_quant"`
	//Multi Dimensional Quantization. If set to 1 then a 2nd stage LPC algorithm is applied after the first stage to finetune the coefficients. This is quite slow and slightly improves compression.
}

type LpcType = typex.String

const (
	LpcType_none LpcType = "none"
	//LPC is not used

	LpcType_fixed LpcType = "fixed"
	//fixed LPC coefficients

	LpcType_levinson LpcType = "levinson"
	LpcType_cholesky LpcType = "cholesky"
)

type PredictionOrderMethod = typex.String

const (
	estimation PredictionOrderMethod = "estimation"
	level2     PredictionOrderMethod = "2level"
	level4     PredictionOrderMethod = "4level"
	level8     PredictionOrderMethod = "8level"
	search     PredictionOrderMethod = "search"
	log        PredictionOrderMethod = "log"
)

type ChMode = typex.String

const (
	auto ChMode = "auto"

	indep ChMode = "indep"

	left_side  ChMode = "left_side"
	right_side ChMode = "right_side"
	mid_side   ChMode = "mid_side"
)
