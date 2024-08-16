package outputdevicex

import "github.com/elizabevil/ffmpegx/paramx/typex"

// SDL 4.9 sdl
type SDL struct {
	WindowBorderless typex.Bool `json:"window_borderless" flag:"-window_borderless"`
	//Set SDL window border off. Default value is 0 (enable window border).

	WindowEnableQuit *typex.Bool `json:"window_enable_quit" flag:"-window_enable_quit"`
	//Enable quit action (using window button or keyboard key) when non-zero value is provided. Default value is 1 (enable quit action).

	WindowFullscreen typex.Bool `json:"window_fullscreen" flag:"-window_fullscreen"`
	//Set fullscreen mode when non-zero value is provided. Default value is zero.

	WindowSize typex.VideoSize `json:"window_size" flag:"-window_size"`
	//Set the SDL window size, can be a string of the form widthxheight or a video size abbreviation. If not specified it defaults to the size of the input video, downscaled according to the aspect ratio.

	WindowTitle typex.FileName `json:"window_title" flag:"-window_title"`
	//Set the SDL window title, if not specified default to the filename specified for the output device.

	WindowX typex.Position `json:"window_x" flag:"-window_x"`
	WindowY typex.Position `json:"window_y" flag:"-window_y"`
	//Set the position of the window on the screen.
	//
}

/*
ffmpeg -i INPUT -c:v rawvideo -pix_fmt yuv420p -window_size qcif -f sdl "SDL output"
*/
