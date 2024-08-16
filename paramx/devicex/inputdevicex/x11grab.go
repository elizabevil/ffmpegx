package inputdevicex

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// X11GRAB 3.21 x11grab
type X11GRAB struct {
	SelectRegion typex.UBool `json:"select_region" flag:"-select_region"`
	//Specify whether to select the grabbing area graphically using the pointer. A value of 1 prompts the user to select the grabbing area graphically by clicking and dragging. A single click with no dragging will select the whole screen. A region with zero width or height will also select the whole screen. This option overwrites the video_size, grab_x, and grab_y options. Default value is 0.

	DrawMouse *typex.UBool `json:"draw_mouse" flag:"-draw_mouse,1"`

	//Specify whether to draw the mouse pointer. A value of 0 specifies not to draw the pointer. Default value is 1.

	FollowMouse typex.UNumber `json:"follow_mouse" flag:"-follow_mouse"`
	//Make the grabbed area follow the mouse. The argument can be centered or a number of pixels PIXELS.

	//When it is specified with "centered", the grabbing region follows the mouse pointer and keeps the pointer at the center of region; otherwise, the region follows only when the mouse pointer reaches within PIXELS (greater than zero) to the edge of region.
	Framerate typex.FrameRate `json:"framerate" flag:"-framerate"`
	//Set the grabbing frame rate. Default value is ntsc, corresponding to a frame rate of 30000/1001.
	FramerateFlag flagx.VideoRate `json:"framerate_flag" flag:"-framerate"`

	ShowRegion typex.UBool `json:"show_region" flag:"-show_region"`
	//Show grabbed region on screen.

	//If show_region is specified with 1, then the grabbing region will be indicated on screen. With this option, it is easy to know what is being grabbed if only a portion of the screen is grabbed.

	RegionBorder typex.Level `json:"region_border" flag:"-region_border"`
	//Set the region border thickness if -show_region 1 is used. Range is 1 to 128 and default is 3 (XCB-based x11grab only).
	WindowId typex.ID `json:"window_id" flag:"-window_id"`
	//Coordinate this window, instead of the whole screen. Default value is 0, which maps to the whole screen (root window).

	//The id of a window can be found using the xwininfo program, possibly with options -tree and -root.

	//If the window is later enlarged, the new area is not recorded. Video ends when the window is closed, unmapped (i.e., iconified) or shrunk beyond the video size (which defaults to the initial window size).

	//This option disables options follow_mouse and select_region.

	VideoSize flagx.VideoSize `json:"video_size" flag:"-video_size"`
	//Set the video frame size. Default is the full desktop or window.

	GrabX typex.Coordinate `json:"grab_x" flag:"-grab_x"`
	GrabY typex.Coordinate `json:"grab_y" flag:"-grab_y"`
	//Set the grabbing region coordinates. They are expressed as offset from the top left corner of the X11 window and correspond to the x_offset and y_offset parameters in the device name. The default value for both options is 0.
}
