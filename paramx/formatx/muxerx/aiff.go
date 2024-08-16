package muxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// 21.9 aiff AIFF
type AIFF struct {
	WriteId3V2 typex.UBool `json:"write_id3v2" flag:"-write_id3v2"`
	//Enable ID3v2 tags writing when set to 1. Default is 0 (disabled).

	Id3V2Version typex.UBool `json:"id3v2_version" flag:"-id3v2_version"`
	//Select ID3v2 version to write. Currently only version 3 and 4 (aka. ID3v2.3 and ID3v2.4) are supported. The default is version 4.
}
