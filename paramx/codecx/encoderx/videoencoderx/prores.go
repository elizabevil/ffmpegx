package videoencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// PRORES 9.25 ProRes
/*
Apple ProRes encoder.

FFmpeg contains 2 ProRes encoders, the prores-aw and prores-ks encoder. The used encoder can be chosen with the -vcodec option.
*/
type PRORES struct {
	Profile ProresKsProfile `json:"profile" flag:"-profile"`
	//Select the ProRes profile to encode
	QuantMat QuantMat `json:"quant_mat" flag:"-quant_mat"`
	//Select quantization matrix.
	BitsPerMb typex.UI16 `json:"bits_per_mb" flag:"-bits_per_mb"`
	//How many bits to allot for coding one macroblock. Different profiles use between 200 and 2400 bits per macroblock, the maximum is 8000.

	MbsPerSlice typex.Level `json:"mbs_per_slice" flag:"-mbs_per_slice"`
	//Number of macroblocks in each slice (1-8); the default value (8) should be good in almost all situations.

	Vendor string `json:"vendor" flag:"-vendor"`
	//Override the 4-byte vendor ID. A custom vendor ID like apl0 would claim the stream was produced by the Apple encoder.

	AlphaBits typex.Level `json:"alpha_bits" flag:"-alpha_bits"`
	//Specify number of bits for alpha component. Possible values are 0, 8 and 16. Use 0 to disable alpha plane coding.
}
type QuantMat = typex.String
type ProresKsProfile = typex.String

const (
	ProresKsProfile_proxy    ProresKsProfile = "proxy"
	ProresKsProfile_lt       ProresKsProfile = "lt"
	ProresKsProfile_standard ProresKsProfile = "standard"
	ProresKsProfile_hq       ProresKsProfile = "hq"
	ProresKsProfile_4444     ProresKsProfile = "4444"
	ProresKsProfile_4444xq   ProresKsProfile = "4444xq"
)

const (
	QuantMat_auto     QuantMat = "auto"
	QuantMat_default  QuantMat = "default"
	QuantMat_proxy    QuantMat = "proxy"
	QuantMat_lt       QuantMat = "lt"
	QuantMat_standard QuantMat = "standard"
	QuantMat_hq       QuantMat = "hq"
)
