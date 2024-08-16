package outputdevicex

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// // OPENGL 4.6 opengl
type OPENGL struct {
	Background typex.Color `json:"background" flag:"-background"`
	//Set background color. Black is a default.

	NoWindow string `json:"no_window" flag:"-no_window"`
	//Disables default SDL window when set to non-zero value. Application must provide OpenGL context and both window_size_cb and window_swap_buffers_cb callbacks when set.

	WindowTitle typex.FileName `json:"window_title" flag:"-window_title"`
	//Set the SDL window title, if not specified default to the filename specified for the output device. Ignored when no_window is set.

	//Set preferred window size, can be a string of the form widthxheight or a video size abbreviation. If not specified it defaults to the size of the input video, downscaled according to the aspect ratio. Mostly usable when no_window is not set.
	WindowSize flagx.VideoSize `json:"window_size" flag:"-window_size"`
}

//ffmpeg  -i INPUT -f opengl "window title"
