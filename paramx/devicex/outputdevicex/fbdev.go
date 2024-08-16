package outputdevicex

import "github.com/elizabevil/ffmpegx/paramx/typex"

// FBDEV 4.5 fbdev
type FBDEV struct {
	Xoffset typex.Offset `json:"xoffset" flag:"-xoffset"`
	Yoffset typex.Offset `json:"yoffset" flag:"-yoffset"`
	//Set x/y coordinate of top left corner. Default is 0.

}
