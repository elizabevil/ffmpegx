package demuxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// 20.8 ea  EA
type EA struct {
	MergeAlpha typex.UBool `json:"merge_alpha" flag:"-merge_alpha"`
	//Normally the VP6 alpha channel (if exists) is returned as a secondary video stream, by setting this option you can make the demuxer return a single video stream which contains the alpha channel in addition to the ordinary video.
}
