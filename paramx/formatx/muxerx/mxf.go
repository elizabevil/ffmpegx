package muxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// 21.63 mxf, mxf_d10, mxf_opatom
type MXF struct {
	StoreUserComments typex.UBool `json:"store_user_comments" flag:"-store_user_comments"`
	//Set if user comments should be stored if available or never. IRT D-10 does not allow user comments. The default is thus to write them for mxf and mxf_opatom but not for mxf_d10
}

// NUT 21.65 nut
// ffmpeg -i INPUT -f_strict experimental -syncpoints none - | processor
