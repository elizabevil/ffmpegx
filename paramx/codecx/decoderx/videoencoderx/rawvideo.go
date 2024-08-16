package videodecoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// RAWVIDEO 4.2 rawvideo
type RAWVIDEO struct {
	Top *typex.Bool `json:"top" flag:"-top,-1"`
}
