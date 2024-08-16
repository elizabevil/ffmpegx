package videodecoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// V210 4.8 v210
type V210 struct {
	CustomStride typex.UNumber `json:"custom_stride" flag:"-custom_stride"`
	//Set the line size of the v210 data in bytes. The default value is 0 (autodetect). You can use the special -1 value for a strideless v210 as seen in BOXX files.

}
