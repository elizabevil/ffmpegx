package inputdevicex

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// FBDEV 3.7 fbdev
/*
Linux framebuffer input device.

The Linux framebuffer is a graphic hardware-independent abstraction layer to show graphics on a computer monitor, typically on the console. It is accessed through a file device node, usually /dev/fb0.

For more detailed information read the file Documentation/fb/framebuffer.txt included in the Linux source tree.
*/
type FBDEV struct {
	Framerate     typex.FrameRate `json:"framerate" flag:"-framerate"` //Set the frame rate. Default is 25.
	FramerateFlag flagx.VideoRate `json:"framerate_flag" flag:"-framerate"`
}

/*
ffmpeg -f fbdev -framerate 1 -i /dev/fb0 -frames:v 1 screenshot.jpeg
*/
