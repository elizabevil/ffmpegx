package muxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// 21.22 avif AVIF
type AVIF struct {
	Loop typex.Count `json:"loop" flag:"-loop"`
	//number of times to loop an animated AVIF, 0 specify an infinite loop, default is 0

	MovieTimescale typex.ScaleI `json:"movie_timescale" flag:"-movie_timescale"`
	//Set the timescale written in the movie header box (mvhd). Range is 1 to INT_MAX. Default is 1000.
}
