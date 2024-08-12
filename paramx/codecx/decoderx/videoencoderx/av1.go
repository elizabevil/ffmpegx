package videodecoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// AV1 4.1 av1
type AV1 struct {
	Top typex.Level `json:"top" flag:"-top"`
	//Specify the assumed field type of the input video.

	//-1
	//the video is assumed to be progressive (default)

	//0
	//bottom-field-first is assumed

	//1
	//top-field-first is assumed
}
