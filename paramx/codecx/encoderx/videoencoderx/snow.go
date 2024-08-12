package videoencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// SNOW 9.27 snow
type SNOW struct {
	IterativeDiaSize typex.UNumber `json:"iterative_dia_size" flag:"-iterative_dia_size"`

	//dia size for the iterative motion estimation
}
