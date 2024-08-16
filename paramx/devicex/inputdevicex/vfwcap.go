package inputdevicex

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// VFWCAP 3.20 vfwcap
type VFWCAP struct {
	VideoSize flagx.VideoSize `json:"video_size" flag:"-video_size"`

	//Set the video frame size.

	Framerate typex.FrameRate `json:"framerate" flag:"-framerate"`
	//Set the grabbing frame rate.Default is ntsc, corresponding to a frame rate of 30000/1001.
	FramerateFlag flagx.VideoRate `json:"framerate_flag" flag:"-framerate"`
	//Set the grabbing frame rate. Default value is ntsc, corresponding to a frame rate of 30000/1001.

}
