package demuxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// 20.5 concat  CONCAT
type CONCAT struct {
	Safe *typex.UBool `json:"safe" flag:"-safe,1"`
	//If set to 1, reject unsafe file paths and directives. A file path is considered safe if it does not contain a protocol specification and is relative and all components only contain characters from the portable character set (letters, digits, period, underscore and hyphen) and have no period at the beginning of a component.
	//If set to 0, any file name is accepted.

	//The default is 1.

	AutoConvert typex.UBool `json:"auto_convert" flag:"-auto_convert"`
	//If set to 1, try to perform automatic conversions on packet data to make the streams concatenable. The default is 1.

	//Currently, the only conversion is adding the h264_mp4toannexb bitstream filter to H.264 streams in MP4 format. This is necessary in particular if there are resolution changes.

	SegmentTimeMetadata typex.UBool `json:"segment_time_metadata" flag:"-segment_time_metadata"`
	//If set to 1, every packet will contain the lavf.concat.start_time and the lavf.concat.duration packet metadata values which are the start_time and the duration of the respective file segments in the concatenated output expressed in microseconds. The duration metadata is only set if it is known based on the concat file. The default is 0.
}
