package videodecoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// LIBXEVD 4.6 libxevd
type LIBXEVD struct {
	Threads typex.UNumber `json:"threads" flag:"-threads"`
	//Force to use a specific number of threads
}
