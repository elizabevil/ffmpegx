package inputdevicex

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// LIBDC1394 3.14 libdc1394
type LIBDC1394 struct {
	Framerate     typex.Frames    `json:"framerate" flag:"-framerate"`
	FramerateFlag flagx.VideoRate `json:"framerate_flag" flag:"-framerate"`
	//Set the frame rate. Default is ntsc, corresponding to a frame rate of 30000/1001.

	PixelFormat flagx.PixelFormat `json:"pixel_format" flag:"-pixel_format"`
	//Select the pixel format. Default is uyvy422.

	VideoSize flagx.VideoSize `json:"video_size" flag:"-video_size"`
	//Set the video size given as a string such as 640x480 or hd720. Default is qvga.
}
