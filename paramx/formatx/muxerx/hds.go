package muxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// 21.45 hds
type HDS struct {
	ExtraWindowSize typex.Size `json:"extra_window_size" flag:"-extra_window_size"`
	//number of fragments kept outside of the manifest before removing from disk
	MinFragDuration typex.MicrosecondI `json:"min_frag_duration" flag:"-min_frag_duration,10000000"`
	//minimum fragment duration (in microseconds), default value is 1 second (10000000)

	RemoveAtExit typex.UBool `json:"remove_at_exit" flag:"-remove_at_exit"`
	//remove all fragments when finished when set to true

	WindowSize typex.Size `json:"window_size" flag:"-window_size"`
	//number of fragments kept in the manifest, if set to a value different from 0. By default all segments are kept in the output directory.
}

//ffmpeg -re -i INPUT -f hds -b:v 200k output.hds
