package demuxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// 20.24 vapoursynth  VAPOURSYNTH
type VAPOURSYNTH struct {
	MaxScriptSize typex.UNumber `json:"max_script_size" flag:"-max_script_size"`
	//The demuxer buffers the entire script into memory. Adjust this value to set the maximum buffer size, which in turn, acts as a ceiling for the size of scripts that can be read. Default is 1 MiB.
}
