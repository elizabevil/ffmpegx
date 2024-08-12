package demuxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// 20.23 tedcaptions  TEDCAPTIONS
type TEDCAPTIONS struct {
	StartTime typex.MillisecondI `json:"start_time" flag:"-start_time"`
	//Set the start time of the TED talk, in milliseconds. The default is 15000 (15s). It is used to sync the captions with the downloadable videos, because they include a 15s intro.
}
