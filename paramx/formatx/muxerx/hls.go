package muxerx

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
	"net/http"
)

// HLS https://ffmpeg.org/ffmpeg-all.html
// 21.46.1 Options
type HLS struct {
	HlsInitTime          typex.DurationI            `json:"hls_init_time" flag:"-hls_init_time"`
	HlsTime              typex.DurationI            `json:"hls_time" flag:"-hls_time"`
	HlsListSize          *typex.Size                `json:"hls_list_size" flag:"-hls_list_size,5"`
	HlsDeleteThreshold   typex.UNumber              `json:"hls_delete_threshold" flag:"-hls_delete_threshold,1"`
	HlsStartNumberSource flagx.HlsStartNumberSource `json:"hls_start_number_source" flag:"-hls_start_number_source"`
	StartNumber          typex.UNumber              `json:"start_number" flag:"-start_number"`
	HlsAllowCache        typex.Bool                 `json:"hls_allow_cache" flag:"-hls_allow_cache"`
	HlsBaseUrl           typex.Url                  `json:"hls_base_url" flag:"-hls_base_url"`
	HlsSegmentFilename   typex.FileName             `json:"hls_segment_filename" flag:"-hls_segment_filename"`
	Strftime             typex.UBool                `json:"strftime" flag:"-strftime"`             //Enable (1) or disable (0)
	StrftimeMkdir        typex.UBool                `json:"strftime_mkdir" flag:"-strftime_mkdir"` //Enable (1) or disable (0)
	HlsSegmentOptions    typex.Dict                 `json:"hls_segment_options" flag:"-hls_segment_options"`
	HlsEnc               typex.UBool                `json:"hls_enc" flag:"-hls_enc"` //Enable (1) or disable (0)
	HlsEncKey            typex.Key                  `json:"hls_enc_key" flag:"-hls_enc_key"`
	HlsEncKeyUrl         typex.Url                  `json:"hls_enc_key_url" flag:"-hls_enc_key_url"`
	HlsEncIv             typex.Key                  `json:"hls_enc_iv" flag:"-hls_enc_iv"`
	HlsSegmentType       HlsSegmentType             `json:"hls_segment_type" flag:"-hls_segment_type"`
	HlsFmp4InitFilename  typex.FileName             `json:"hls_fmp4_init_filename" flag:"-hls_fmp4_init_filename"`
	HlsFmp4InitResend    typex.UBool                `json:"hls_fmp4_init_resend" flag:"-hls_fmp4_init_resend"`
	HlsFlags             flagx.HlsFlags             `json:"hls_flags" flag:"-hls_flags"`
	HlsPlaylistType      flagx.HlsPlaylistType      `json:"hls_playlist_type" flag:"-hls_playlist_type"`
	Method               flagx.HttpMethod           `json:"method" flag:"-method"`
	HttpUserAgent        typex.String               `json:"http_user_agent" flag:"-http_user_agent"`
	VarStreamMap         typex.Param                `json:"var_stream_map" flag:"-var_stream_map"`
	CcStreamMap          typex.Param                `json:"cc_stream_map" flag:"-cc_stream_map"`
	MasterPlName         typex.Name                 `json:"master_pl_name" flag:"-master_pl_name"`
	MasterPlPublishRate  typex.UNumber              `json:"master_pl_publish_rate" flag:"-master_pl_publish_rate"`
	HttpPersistent       typex.UBool                `json:"http_persistent" flag:"-http_persistent"`
	Timeout              typex.Timeout              `json:"timeout" flag:"-timeout"`
	IgnoreIoErrors       typex.UBool                `json:"ignore_io_errors" flag:"-ignore_io_errors"`
	Headers              http.Header                `json:"headers" flag:"-cc_stream_map"`
	HlsKeyInfoFile       typex.FileName             `json:"hls_key_info_file" flag:"hls_key_info_file"` //     "file with key URI and key file path", OFFSET(key_info_file),      AV_OPT_TYPE_STRING, {.str = NULL},            0,       0,         E},
}

//ffmpeg -re -i in.ts -b:v:0 1000k -b:v:1 256k -b:a:0 64k -b:a:1 32k \
//-map 0:v -map 0:a -map 0:v -map 0:a -f hls -var_stream_map "v:0,a:0 v:1,a:1" \
//http://example.com/live/out_%v.m3u8

type HlsSegmentType = typex.String

const (
	HlsSegmentType_mpegts HlsSegmentType = "mpegts"
	HlsSegmentType_fmp4   HlsSegmentType = "fmp4"
)
