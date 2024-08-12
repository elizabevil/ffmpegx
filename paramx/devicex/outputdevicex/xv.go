package outputdevicex

import "github.com/elizabevil/ffmpegx/paramx/typex"

// XV 4.12 xv
type XV struct {
	//can be a string in the format hostname[:number[.screen_number]].
	DisplayName typex.Filename `json:"display_name" flag:"-display_name"`
	WindowId    typex.Index    `json:"window_id" flag:"-window_id"`
	//When set to non-zero value then device doesn’t create new window, but uses existing one with provided window_id. By default this options is set to zero and device creates its own window.

	WindowSize typex.VideoSize `json:"window_size" flag:"-window_size"`
	//Set the created window size, can be a string of the form widthxheight or a video size abbreviation. If not specified it defaults to the size of the input video. Ignored when window_id is set.

	WindowX typex.Position `json:"window_x" flag:"-window_x"`
	WindowY typex.Position `json:"window_y" flag:"-window_y"`
	//Set the X and Y window offsets for the created window. They are both set to 0 by default. The values may be ignored by the window manager. Ignored when window_id is set.

	WindowTitle typex.Filename `json:"window_title" flag:"-window_title"`
	//Set the window title, if not specified default to the filename specified for the output device. Ignored when window_id is set.
	//
}

/*
Decode, display and encode video input with ffmpeg at the same time:
	ffmpeg -i INPUT OUTPUT -f xv display
Decode and display the input video to multiple X11 windows:
	ffmpeg -i INPUT -f xv normal -vf negate -f xv negated
*/
