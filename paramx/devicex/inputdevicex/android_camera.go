package inputdevicex

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

//26.2 android_camera

type AndroidCamera struct {
	VideoSize flagx.VideoSize `json:"video_size" flag:"-video_size"`

	//Set the video size given as a string such as 640x480 or hd720. Falls back to the first available configuration reported by Android if requested video size is not available or by default.

	Framerate typex.FrameRate `json:"framerate" flag:"-framerate"`
	//Set the video framerate. Falls back to the first available configuration reported by Android if requested framerate is not available or by default (-1).
	FramerateFlag flagx.VideoRate `json:"framerate_flag" flag:"-framerate"`

	CameraIndex typex.Index `json:"camera_index" flag:"-camera_index"`
	//Set the index of the camera to use. Default is 0.

	InputQueueSize typex.Size `json:"input_queue_size" flag:"-input_queue_size"`
	//Set the maximum number of frames to buffer. Default is 5.
}
