package demuxerx

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// 20.20 rawvideo  RAWVIDEO
type RAWVIDEO struct {
	Framerate typex.FrameRate `json:"framerate" flag:"-framerate"`
	//Set input video frame rate. Default value is 25.

	PixelFormat flagx.PixelFormat `json:"pixel_format" flag:"-pixel_format"`
	//Set the input video pixel format. Default value is yuv420p.

	VideoSize typex.VideoSize `json:"video_size" flag:"-video_size"`
	//Set the input video size. This value must be specified explicitly.

}
