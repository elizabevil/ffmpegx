package inputdevicex

import "github.com/elizabevil/ffmpegx/paramx/typex"

// JACK 3.10 jack
type JACK struct {
	Channels typex.Channels `json:"channels" flag:"-channels"`
	//Set the number of channels. Default is 2.
}
