package demuxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

//20.12 hls
//HLS demuxer
//Apple HTTP Live Streaming demuxer.
//This demuxer presents all AVStreams from all variant streams. The id field is set to the bitrate variant index number. By setting the discard flags on AVStreams (by pressing ’a’ or ’v’ in ffplay), the caller can decide which variant streams to actually receive. The total bitrate of the variant that the stream belongs to is available in a metadata key named "variant_bitrate".

type HLS struct {
	LiveStartIndex typex.Index `json:"live_start_index" flag:"-live_start_index"`
	//segment index to start live streams at (negative values are from the end).

	PreferXStart typex.UBool `json:"prefer_x_start" flag:"-prefer_x_start"`
	//prefer to use #EXT-X-START if it’s in playlist instead of live_start_index.

	AllowedExtensions typex.List `json:"allowed_extensions" flag:"-allowed_extensions"`
	//’,’ separated list of file extensions that hls is allowed to access.

	MaxReload typex.Size `json:"max_reload" flag:"-max_reload"`
	//Maximum number of times a insufficient list is attempted to be reloaded. Default value is 1000.

	M3U8HoldCounters typex.Size `json:"m3u8_hold_counters" flag:"-m3u8_hold_counters"`
	//The maximum number of times to load m3u8 when it refreshes without new segments. Default value is 1000.

	HttpPersistent typex.UBool `json:"http_persistent" flag:"-http_persistent"`
	//Use persistent HTTP connections. Applicable only for HTTP streams. Enabled by default.

	HttpMultiple typex.Bool `json:"http_multiple" flag:"-http_multiple"`
	//Use multiple HTTP connections for downloading HTTP segments. Enabled by default for HTTP/1.1 servers.

	HttpSeekable *typex.Bool `json:"http_seekable" flag:"-http_seekable,-1"`
	//Use HTTP partial requests for downloading HTTP segments. 0 = disable, 1 = enable, -1 = auto, Default is auto.

	SegFormatOptions typex.Dict `json:"seg_format_options" flag:"-seg_format_options"`
	//Set options for the demuxer of media segments using a list of key=value pairs separated by :.

	SegMaxRetry typex.Size `json:"seg_max_retry" flag:"-seg_max_retry"`
	//Maximum number of times to reload a segment on error, useful when segment skip on network error is not desired. Default value is 0.
}
