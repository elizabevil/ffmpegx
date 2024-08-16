package muxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// Ogg 21.66 ogg
// ogg container muxer.
type Ogg struct {
	PageDuration *typex.MicrosecondI `json:"page_duration" flag:"-page_duration,1"`
	//Preferred page duration, in microseconds. The muxer will attempt to create pages that are approximately duration microseconds long. This allows the user to compromise between seek granularity and container overhead. The default is 1 second. A value of 0 will fill all segments, making pages as large as possible. A value of 1 will effectively use 1 packet-per-page in most situations, giving a small seek granularity at the cost of additional container overhead.

	SerialOffset typex.UOffset `json:"serial_offset" flag:"-serial_offset"`
	//Serial value from which to set the streams serial number. Setting it to different and sufficiently large values ensures that the produced ogg files can be safely chained.
}
