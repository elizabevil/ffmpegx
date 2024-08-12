package demuxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// // 20.14 libgme  LIBGME
type LIBGME struct {
	TrackIndex typex.Index `json:"track_index" flag:"-track_index"`
	//Set the index of which track to demux. The demuxer can only export one track. Track indexes start at 0. Default is to pick the first track. Number of tracks is exported as tracks metadata entry.

	SampleRate typex.SampleRate `json:"sample_rate" flag:"-sample_rate"`
	//Set the sampling rate of the exported track. Range is 1000 to 999999. Default is 44100.

	MaxSize typex.UNumber `json:"max_size" flag:"-max_size"` //(bytes)
	//The demuxer buffers the entire file into memory. Adjust this value to set the maximum buffer size, which in turn, acts as a ceiling for the size of files that can be read. Default is 50 MiB.
}
