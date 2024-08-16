package demuxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// // 20.16 libopenmpt  LIBOPENMPT
type LIBOPENMPT struct {
	Subsong typex.LevelUI16 `json:"subsong" flag:"-subsong"`
	//Set the subsong index. This can be either ’all’, ’auto’, or the index of the subsong. Subsong indexes start at 0. The default is ’auto’.

	//The default value is to let libopenmpt choose.

	Layout typex.Level `json:"layout" flag:"-layout"`
	//Set the channel layout. Valid values are 1, 2, and 4 channel layouts. The default value is STEREO.

	SampleRate typex.SampleRate `json:"sample_rate" flag:"-sample_rate"`
	//Set the sample rate for libopenmpt to output. Range is from 1000 to INT_MAX. The value default is 48000.

}
