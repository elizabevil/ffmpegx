package videoencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// MICROSOFT 9.22 Microsoft RLE
type MicrosoftRLE struct {
	G typex.Size `json:"g" flag:"-g"`
	//Keyframe interval. A keyframe is inserted at least every -g frames, sometimes sooner.
}
