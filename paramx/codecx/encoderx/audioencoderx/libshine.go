package audioencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// Libshine 8.10 libshine
type Libshine struct {
	B typex.Bitrate `json:"b" flag:"-b"` // (encoding,audio,video)

	//Set bitrate expressed in bits/s for CBR. shineenc -b option is expressed in kilobits/s.

}
