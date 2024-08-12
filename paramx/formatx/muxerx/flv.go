package muxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// 21.38 flv FLV
type FLV struct {
	Flvflags Flvflags `json:"flvflags" flag:"-flvflags"`
}

type Flvflags = typex.String

const (
	Flvflags_aac_seq_header_detect Flvflags = "aac_seq_header_detect"
	//Place AAC sequence header based on audio stream data.

	Flvflags_no_sequence_end Flvflags = "no_sequence_end"
	//Disable sequence end tag.

	Flvflags_no_metadata Flvflags = "no_metadata"
	//Disable metadata tag.

	Flvflags_no_duration_filesize Flvflags = "no_duration_filesize"
	//Disable duration and filesize in metadata when they are equal to zero at the end of stream. (Be used to non-seekable living stream).

	Flvflags_add_keyframe_index Flvflags = "add_keyframe_index"
	//Used to facilitate seeking; particularly for HTTP pseudo streaming.

)
