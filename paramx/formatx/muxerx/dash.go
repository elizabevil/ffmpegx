package muxerx

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// // 21.29 dash DASH
type DASH struct {
	AdaptationSets  string          `json:"adaptation_sets" flag:"-adaptation_sets "`
	DashSegmentType DashSegmentType `json:"dash_segment_type" flag:"-dash_segment_type  "`

	ExtraWindowSize typex.Size `json:"extra_window_size" flag:"-extra_window_size"`
	//Set the maximum number of segments kept outside of the manifest before removing from disk.

	FormatOptions string `json:"format_options" flag:"-format_options"`
	//Set container format (mp4/webm) options using a:-separated list of key = value parameters.Values containing: special characters must be escaped.

	FragDuration typex.SecondI `json:"frag_duration" flag:"-frag_duration"`
	//Set the length in seconds of fragments within segments, fractional value can also be set.

	FragType FragType `json:"frag_type" flag:"-frag_type"`

	GlobalSidx typex.UBool `json:"global_sidx" flag:"-global_sidx"`
	//Write global SIDX atom.Applicable only for single file, mp4 output, non-streaming mode.

	HlsMasterName typex.Name `json:"hls_master_name" flag:"-hls_master_name"`
	//HLS master playlist name.Default is master.m3u8.

	HlsPlaylist typex.UBool `json:"hls_playlist" flag:"-hls_playlist"`
	//Generate HLS playlist files.The master playlist is generated with filename specified by the hls_master_name option.One media playlist file is generated for each stream with filenames media_0.m3u8, media_1.m3u8, etc.

	HttpOpts string `json:"http_opts" flag:"-http_opts"`
	//Specify a list of:-separated key = value options to pass to the underlying HTTP protocol.Applicable only for HTTP output.

	HttpPersistent typex.UBool `json:"http_persistent" flag:"-http_persistent"`
	//Use persistent HTTP connections.Applicable only for HTTP output.

	HttpUserAgent string `json:"http_user_agent" flag:"-http_user_agent"`
	//Override User-Agent field in HTTP header.Applicable only for HTTP output.

	IgnoreIoErrors typex.UBool `json:"ignore_io_errors" flag:"-ignore_io_errors"`
	//Ignore IO errors during open and write.Useful for long-duration runs with network output.This is disabled by default.

	IndexCorrection typex.UBool `json:"index_correction" flag:"-index_correction"`
	InitSegName     typex.Name  `json:"init_seg_name" flag:"-init_seg_name"`
	//DASH-templated name to use for the initialization segment.Default is init-stream$RepresentationID$.$ext$.$ext$ is replaced with the file name extension specific for the segment format.

	Ldash typex.UBool `json:"ldash" flag:"-ldash"`
	//Enable Low-latency Dash by constraining the presence and values of some elements.This is disabled by default.

	Lhls typex.UBool `json:"lhls" flag:"-lhls"`
	//Enable Low-latency HLS (LHLS).Add #EXT-X-PREFETCH tag with current segment’s URI.hls.js player folks are trying to standardize an open LHLS spec.The draft spec is available at https: //github.com/video-dev/hlsjs-rfcs/blob/lhls-spec/proposals/0001-lhls.md.

	//This option tries to comply with the above open spec.It enables streaming and hls_playlist options automatically.This is an experimental feature.

	//Note: This is not Apple’s version LHLS.See https: //datatracker.ietf.org/doc/html/draft-pantos-hls-rfc8216bis

	MasterM3U8PublishRate typex.UNumber `json:"master_m3u8_publish_rate" flag:"-master_m3u8_publish_rate"`
	//Publish master playlist repeatedly every after specified number of segment intervals.

	MaxPlaybackRate typex.FLT0_1 `json:"max_playback_rate" flag:"-max_playback_rate"`
	//Set the maximum playback rate indicated as appropriate for the purposes of automatically adjusting playback latency and buffer occupancy during normal playback by clients.

	MediaSegName typex.Name `json:"media_seg_name" flag:"-media_seg_name"`
	//DASH-templated name to use for the media segments.Default is chunk-stream$RepresentationID$-$Number%05d$.$ext$.$ext$ is replaced with the file name extension specific for the segment format.

	Method flagx.HttpMethod `json:"method" flag:"-method"`
	//Use the given HTTP method to create output files.Generally set to PUT or POST.

	MinPlaybackRate typex.FLT0_1 `json:"min_playback_rate" flag:"-min_playback_rate"`
	//Set the minimum playback rate indicated as appropriate for the purposes of automatically adjusting playback latency and buffer occupancy during normal playback by clients.

	MpdProfile   MpdProfile  `json:"mpd_profile" flag:"-mpd_profile"`
	RemoveAtExit typex.UBool `json:"remove_at_exit" flag:"-remove_at_exit"`
	//Enable or disable removal of all segments when finished.This is disabled by default.

	SegDuration typex.SecondI `json:"seg_duration" flag:"-seg_duration"`
	//Set the segment length in seconds (fractional value can be set).The value is treated as average segment duration when the use_template option is enabled and the use_timeline option is disabled and as minimum segment duration for all the other use cases.

	SingleFile typex.Name `json:"single_file" flag:"-single_file"`
	//Enable or disable storing all segments in one file, accessed using byte ranges.This is disabled by default.

	//The name of the single file can be specified with the single_file_name option, if not specified assume the basename of the manifest file with the output format extension.

	SingleFileName typex.Name `json:"single_file_name" flag:"-single_file_name"`
	//DASH-templated name to use for the manifest baseURL element.Imply that the single_file option is set to true.In the template, $ext$ is replaced with the file name extension specific for the segment format.

	Streaming typex.UBool `json:"streaming" flag:"-streaming"`
	//Enable or disable chunk streaming mode of output.In chunk streaming mode, each frame will be a moof fragment which forms a chunk.This is disabled by default.

	TargetLatency typex.SecondI `json:"target_latency" flag:"-target_latency"`
	//Set an intended target latency in seconds for serving (fractional value can be set).Applicable only when the streaming and write_prft options are enabled.This is an informative fields clients can use to measure the latency of the service.

	Timeout typex.SecondI `json:"timeout" flag:"-timeout"`
	//Set timeout for socket I/O operations expressed in seconds (fractional value can be set).Applicable only for HTTP output.

	UpdatePeriod typex.SecondI `json:"update_period" flag:"-update_period"`
	//Set the MPD update period, for dynamic content.The unit is second.If set to 0, the period is automatically computed.

	//Default value is 0.

	UseTemplate typex.UBool `json:"use_template" flag:"-use_template"`
	//Enable or disable use of SegmentTemplate instead of SegmentList in the manifest.This is enabled by default.

	UseTimeline typex.UBool `json:"use_timeline" flag:"-use_timeline"`
	//Enable or disable use of SegmentTimeline within the SegmentTemplate manifest section.This is enabled by default.

	UtcTimingUrl typex.Url `json:"utc_timing_url" flag:"-utc_timing_url"`
	//URL of the page that will return the UTC timestamp in ISO format, for example https: //time.akamai.com/?iso

	WindowSize typex.Size `json:"window_size" flag:"-window_size"`
	//Set the maximum number of segments kept in the manifest, discard the oldest one.This is useful for live streaming.

	//If the value is 0, all segments are kept in the manifest.Default value is 0.

	WritePrft typex.Bool `json:"write_prft" flag:"-write_prft"`
	//Write Producer Reference Time elements on supported streams.This also enables writing prft boxes in the underlying muxer.Applicable only when the utc_url option is enabled.It is set to auto by default, in which case the muxer will attempt to enable it only in modes that require it.
}

type MpdProfile = typex.String
type DashSegmentType = typex.String
type FragType = typex.String

const (
	MpdProfile_dash MpdProfile = "dash"
	//The dash segment files format will be selected based on the stream codec. This is the default mode.

	MpdProfile_dvb_dash MpdProfile = "dvb_dash"
	//the dash segment files will be in ISOBMFF/MP4 format

)
const (
	DashSegmentType_auto DashSegmentType = "auto"
	//The dash segment files format will be selected based on the stream codec. This is the default mode.

	DashSegmentType_mp4 DashSegmentType = "mp4"
	//the dash segment files will be in ISOBMFF/MP4 format

	DashSegmentType_webm DashSegmentType = "webm"
	//the dash segment files will be in WebM format
)

const (
	FragType_none FragType = "none"
	//The dash segment files format will be selected based on the stream codec. This is the default mode.

	FragType_every_frame FragType = "every_frame"
	//the dash segment files will be in ISOBMFF/MP4 format

	FragType_duration FragType = "duration"
	FragType_pframes  FragType = "pframes"
	//the dash segment files will be in WebM format
)
