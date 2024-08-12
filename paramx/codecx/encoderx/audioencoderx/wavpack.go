package audioencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// WAVPACK 8.15 wavpack
type WAVPACK struct {
	FrameSize typex.FrameSize `json:"frame_size" flag:"-frame_size"` //  (--blocksize)
	//For this encoder, the range for this option is between 128 and 131072. Default is automatically decided based on sample rate and number of channel.
	//For the complete formula of calculating default, see libavcodec/wavpackenc.c.

	//compression_level (-f, -h, -hh, and -x)

	//8.15.1.2 Private options
	JointStereo  typex.BoolMode `json:"joint_stereo" flag:"-j"` // (-j)
	OptimizeMono typex.BoolMode `json:"optimize_mono" flag:"-optimize_mono"`
}
