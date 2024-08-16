package muxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// // 21.42 gif GIF
type GIF struct {
	Loop typex.Count `json:"loop" flag:"-loop"`
	//Set the number of times to loop the output. Use -1 for no loop, 0 for looping indefinitely (default).

	FinalDelay typex.CentiSeconds `json:"final_delay" flag:"-final_delay"`
	//Force the delay (expressed in centiseconds) after the last frame. Each frame ends with a delay until the next frame. The default is -1, which is a special value to tell the muxer to re-use the previous delay. In case of a loop, you might want to customize this value to mark a pause for instance.

}
