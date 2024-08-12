package videoencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

const (
	HAP_COMP_NONE    = 0xA0
	HAP_COMP_SNAPPY  = 0xB0
	HAP_COMP_COMPLEX = 0xC0

	HAP_FMT_RGBDXT1   = 0x0B
	HAP_FMT_RGBADXT5  = 0x0E
	HAP_FMT_YCOCGDXT5 = 0x0F
	HAP_FMT_RGTC1     = 0x01
)

// HAP 9.4 Hap
type HAP struct {
	Format     Format      `json:"format" flag:"-format"`
	Chunks     typex.Level `json:"chunks" flag:"-chunks"`
	Compressor Compressor  `json:"compressor" flag:"-compressor"`
}

type Format string

const (
	hap       Format = "hap"
	hap_alpha Format = "hap_alpha"
	hap_q     Format = "hap_q"
)

type Compressor string

const (
	none   Compressor = "none"
	snappy Compressor = "snappy"
)
