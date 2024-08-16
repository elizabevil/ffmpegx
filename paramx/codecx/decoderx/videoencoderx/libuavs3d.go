package videodecoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// LIBUAVS3D 4.5 libuavs3d
type LIBUAVS3D struct {
	FrameThreads typex.UNumber `json:"frame_threads" flag:"-frame_threads"`
	//Set           amount of frame threads to use during decoding.The default value is 0 (autodetect).
}
