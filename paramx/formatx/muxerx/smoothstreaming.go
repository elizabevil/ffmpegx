package muxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

type SmoothStreaming struct {
	WindowSize typex.Size `json:"window_size" flag:"-window_size"`
	//Specify the number of fragments kept in the manifest.Default 0 (keep all).

	ExtraWindowSize typex.Size `json:"extra_window_size" flag:"-extra_window_size"`
	//Specify the number of fragments kept outside of the manifest before removing from disk.Default 5.

	LookaheadCount typex.Count `json:"lookahead_count" flag:"-lookahead_count"`
	//Specify the number of lookahead fragments.Default 2.

	MinFragDuration typex.MicrosecondI `json:"min_frag_duration" flag:"-min_frag_duration"`
	//Specify the minimum fragment duration (in microseconds).Default 5000000.

	RemoveAtExit typex.UBool `json:"remove_at_exit" flag:"-remove_at_exit"`
	//Specify whether to remove all fragments when finished.Default 0 (do not remove).
}
