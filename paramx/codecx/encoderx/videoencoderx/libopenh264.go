package videoencoderx

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// // LIBOPENH264 9.11 libopenh264
type LIBOPENH264 struct {
	B typex.Bitrate `json:"b" flag:"-b"`
	//Set the bitrate (as a number of bits per second).

	G typex.Size `json:"g" flag:"-g"`
	//Set the GOP size.

	Maxrate typex.BitrateInt `json:"maxrate" flag:"-maxrate"`
	//Set the max bitrate (as a number of bits per second).

	Flags flagx.Flag `json:"flags" flag:"-flags"`
	//Set global header in the bitstream.

	Slices typex.SliceSize `json:"slices" flag:"-slices"`
	//Set the number of slices, used in parallelized encoding. Default value is 0. This is only used when slice_mode is set to ‘fixed’.

	Loopfilter typex.UBool `json:"loopfilter" flag:"-loopfilter"`
	//Enable loop filter, if set to 1 (automatically enabled). To disable set a value of 0.

	Profile LIBOPENH264Profile `json:"profile" flag:"-profile"`
	//Set profile restrictions. If set to the value of ‘main’ enable CABAC (set the SEncParamExt.iEntropyCodingModeFlag flag to 1).

	MaxNalSize typex.Size `json:"max_nal_size" flag:"-max_nal_size"`
	//Set maximum NAL size in bytes.

	AllowSkipFrames typex.UBool `json:"allow_skip_frames" flag:"-allow_skip_frames"`
	//Allow skipping frames to hit the target bitrate if set to 1.
}
type LIBOPENH264Profile string

const (
	LIBOPENH264Profile_constrained_baseline LIBOPENH264Profile = "constrained_baseline"
	LIBOPENH264Profile_main                 LIBOPENH264Profile = "main"
	LIBOPENH264Profile_high                 LIBOPENH264Profile = "high"
)
