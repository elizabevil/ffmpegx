package inputdevicex

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// GDIGRAB 3.8 gdigrab
type GDIGRAB struct {
	DrawMouse *typex.UBool `json:"draw_mouse" flag:"-draw_mouse,1"`
	//Specify whether to draw the mouse pointer. Use the value 0 to not draw the pointer. Default value is 1.
	Framerate typex.FrameRate `json:"framerate" flag:"-framerate"`

	FramerateFlag flagx.VideoRate `json:"framerate_flag" flag:"-framerate"`
	//Set the grabbing frame rate. Default value is ntsc, corresponding to a frame rate of 30000/1001.

	ShowRegion typex.UBool `json:"show_region" flag:"-show_region"`
	//Show grabbed region on screen.
	VideoSize flagx.VideoSize `json:"video_size" flag:"-video_size"`

	//Set the video frame size. The default is to capture the full screen if desktop is selected, or the full window size if title=window_title is selected.

	OffsetX typex.Offset `json:"offset_x" flag:"-offset_x"`
	//When capturing a region with video_size, set the distance from the left edge of the screen or desktop.

	//Note that the offset calculation is from the top left corner of the primary monitor on Windows. If you have a monitor positioned to the left of your primary monitor, you will need to use a negative offset_x value to move the region to that monitor.

	OffsetY typex.Offset `json:"offset_y" flag:"-offset_y"`
	//When capturing a region with video_size, set the distance from the top edge of the screen or desktop.

}

/*

ffmpeg -f gdigrab -show_region 1 -framerate 6 -video_size cif -offset_x 10 -offset_y 20 -i desktop out.mpg
*/
