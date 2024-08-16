package audiodecoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// LIBILBC 5.6 libilbc
type LIBILBC struct {
	Enhance typex.UBool `json:"enhance" flag:"-enhance"`
	//Enable the enhancement of the decoded audio when set to 1. The default value is 0 (disabled).
}
