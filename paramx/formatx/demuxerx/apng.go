package demuxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// Apng https://ffmpeg.org/ffmpeg-all.html#Demuxers
// 20.3 apng
// Animated Portable Network Graphics demuxer.
// This demurer is used to demux APNG files. All headers, but the PNG signature, up to (but not including) the first fcTL chunk are transmitted as extradata. Frames are then split as being all the chunks between two fcTL ones, or between the last fcTL and IEND chunks.
type Apng struct {
	IgnoreLoop *typex.UBool `json:"ignore_loop" flag:"-ignore_loop"`
	//Ignore the loop variable in the file if set. Default is enabled.

	MaxFps typex.FrameRate `json:"max_fps" flag:"-max_fps"`
	//Maximum framerate in frames per second. Default of 0 imposes no limit.

	DefaultFps typex.Level `json:"default_fps" flag:"-default_fps"`
	//Default framerate in frames per second when none is specified in the file (0 meaning as fast as possible). Default is 15.

}

// Asf 20.4 asf
// Advanced Systems Format demuxer.
// This demuxer is used to demux ASF files and MMS network streams.
type Asf struct {
	NoResyncSearch bool `json:"no_resync_search" flag:"-no_resync_search"`
	//Do not try to resynchronize by looking for a certain optional start code.
}
