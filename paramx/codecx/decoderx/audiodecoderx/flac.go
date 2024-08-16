package audiodecoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// FLAC 5.2 flac
type FLAC struct {
	UseBuggyLpc *typex.UBool `json:"use_buggy_lpc" flag:"-use_buggy_lpc,1"`
	//The lavc FLAC encoder used to produce buggy streams with high lpc values (like the default value). This option makes it possible to decode such streams correctly by using lavc’s old buggy lpc logic for decoding.
}
