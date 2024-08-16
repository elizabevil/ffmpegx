package muxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// // 21.7 adts ADTS
type ADTS struct {
	WriteId3V2 typex.UBool `json:"write_id3v2" flag:"-write_id3v2"`
	//Enable to write ID3v2.4 tags at the start of the stream. Default is disabled.

	WriteApetag typex.UBool `json:"write_apetag" flag:"-write_apetag"`
	//Enable to write APE tags at the end of the stream. Default is disabled.

	WriteMpeg2 typex.UBool `json:"write_mpeg2" flag:"-write_mpeg2"`
	//Enable to set MPEG version bit in the ADTS frame header to 1 which indicates MPEG-2. Default is 0, which indicates MPEG-4.
}
