package muxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// // 21.68 segment, stream_segment, ssegment Segment,
type Segment struct {
	SegmentList typex.FileName `json:"segment_list" flag:"-segment_list"`
	//Generate also a listfile named name. If not specified no listfile is generated.
	IncrementTc typex.UBool `json:"increment_tc" flag:"-increment_tc"`
	//if set to 1, increment timecode between each segment If this is selected, the input need to have a timecode in the first video stream. Default value is 0.

	ReferenceStream string `json:"reference_stream" flag:"-reference_stream"`
	//Set the reference stream, as specified by the string specifier. If specifier is set to auto, the reference is chosen automatically. Otherwise it must be a stream specifier (see the “Stream specifiers” chapter in the ffmpeg manual) which specifies the reference stream. The default value is auto.

	SegmentFormat typex.Format `json:"segment_format" flag:"-segment_format"`
	//Override the inner container format, by default it is guessed by the filename extension.

	SegmentFormatOptions typex.Dict `json:"segment_format_options" flag:"-segment_format_options"`
	//Set output format options using a :-separated list of key=value parameters. Values containing the : special character must be escaped.

	SegmentListFlags SegmentListFlags `json:"segment_list_flags" flag:"-segment_list_flags"`
	SegmentListSize  typex.Size       `json:"segment_list_size" flag:"-segment_list_size"`
	//Update the list file so that it contains at most size segments. If 0 the list file will contain all the segments. Default value is 0.

	SegmentListEntryPrefix string `json:"segment_list_entry_prefix" flag:"-segment_list_entry_prefix"`
	//Prepend prefix to each entry. Useful to generate absolute paths. By default no prefix is applied.

	SegmentListType SegmentListType `json:"segment_list_type" flag:"-segment_list_type"`
	//Select the listing format.
	SegmentTime typex.SecondI `json:"segment_time" flag:"-segment_time"`
	//Set segment duration to time, the value must be a duration specification. Default value is "2". See also the segment_times option.

	//Note that splitting may not be accurate, unless you force the reference stream key-frames at the given time. See the introductory notice and the examples below.

	MinSegDuration typex.SecondI `json:"min_seg_duration" flag:"-min_seg_duration"`
	//Set minimum segment duration to time, the value must be a duration specification. This prevents the muxer ending segments at a duration below this value. Only effective with segment_time. Default value is "0".

	SegmentAtclocktime typex.UBool `json:"segment_atclocktime" flag:"-segment_atclocktime"`
	//If set to "1" split at regular clock time intervals starting from 00:00 o’clock. The time value specified in segment_time is used for setting the length of the splitting interval.

	//For example with segment_time set to "900" this makes it possible to create files at 12:00 o’clock, 12:15, 12:30, etc.

	//Default value is "0".

	SegmentClocktimeOffset typex.UOffset `json:"segment_clocktime_offset" flag:"-segment_clocktime_offset"`
	//Delay the segment splitting times with the specified duration when using segment_atclocktime.

	//For example with segment_time set to "900" and segment_clocktime_offset set to "300" this makes it possible to create files at 12:05, 12:20, 12:35, etc.

	//Default value is "0".

	SegmentClocktimeWrapDuration typex.UTime `json:"segment_clocktime_wrap_duration" flag:"-segment_clocktime_wrap_duration"`
	//Force the segmenter to only start a new segment if a packet reaches the muxer within the specified duration after the segmenting clock time. This way you can make the segmenter more resilient to backward local time jumps, such as leap seconds or transition to standard time from daylight savings time.

	//Default is the maximum possible duration which means starting a new segment regardless of the elapsed time since the last clock time.

	SegmentTimeDelta typex.Delta `json:"segment_time_delta" flag:"-segment_time_delta"`
	//Specify the accuracy time when selecting the start time for a segment, expressed as a duration specification. Default value is "0".
	SegmentTimes string `json:"segment_times" flag:"-segment_times"`
	//Specify a list of split points. times contains a list of comma separated duration specifications, in increasing order. See also the segment_time option.

	SegmentFrames string `json:"segment_frames" flag:"-segment_frames"`
	//Specify a list of split video frame numbers. frames contains a list of comma separated integer numbers, in increasing order.

	//This option specifies to start a new segment whenever a reference stream key frame is found and the sequential number (starting from 0) of the frame is greater or equal to the next value in the list.

	SegmentWrap typex.Index `json:"segment_wrap" flag:"-segment_wrap"`
	//Wrap around segment index once it reaches limit.

	SegmentStartNumber typex.UNumber `json:"segment_start_number" flag:"-segment_start_number"`
	//Set the sequence number of the first segment. Defaults to 0.

	Strftime typex.UBool `json:"strftime" flag:"-strftime"`
	//Use the strftime function to define the name of the new segments to write. If this is selected, the output segment name must contain a strftime function template. Default value is 0.

	BreakNonKeyframes typex.UBool `json:"break_non_keyframes" flag:"-break_non_keyframes"`
	//If enabled, allow segments to start on frames other than keyframes. This improves behavior on some players when the time between keyframes is inconsistent, but may make things worse on others, and can cause some oddities during seeking. Defaults to 0.

	ResetTimestamps typex.UBool `json:"reset_timestamps" flag:"-reset_timestamps"`
	//Reset timestamps at the beginning of each segment, so that each segment will start with near-zero timestamps. It is meant to ease the playback of the generated segments. May not work with some combinations of muxers/codecs. It is set to 0 by default.

	InitialOffset typex.Offset `json:"initial_offset" flag:"-initial_offset"`
	//Specify timestamp offset to apply to the output packet timestamps. The argument must be a time duration specification, and defaults to 0.

	WriteEmptySegments typex.UBool `json:"write_empty_segments" flag:"-write_empty_segments"`
	//If enabled, write an empty segment if there are no packets during the period a segment would usually span. Otherwise, the segment will be filled with the next packet written. Defaults to 0.
}
type SegmentListType = typex.String
type SegmentListFlags = typex.String

const (
	SegmentListType_flat SegmentListType = "flat"
	//Generate a flat list for the created segments, one segment per line.

	SegmentListType_csv SegmentListType = "csv"
	//Generate a list for the created segments, one segment per line, each line matching the format (comma-separated values):
	SegmentListType_ffconcat SegmentListType = "ffconcat"
	//Generate an ffconcat file for the created segments. The resulting file can be read using the FFmpeg concat demuxer.

	//A list file with the suffix ".ffcat" or ".ffconcat" will auto-select this format.

	SegmentListType_m3u8 SegmentListType = "m3u8"
	//Generate an extended M3U8 file, version 3, compliant with http://tools.ietf.org/id/draft-pantos-http-live-streaming.
)

const (
	SegmentListFlags_cache SegmentListFlags = "cache"
	SegmentListFlags_live  SegmentListFlags = "live"
)

/*
4.68.2 Examples
Remux the content of file in.mkv to a list of segments out-000.nut, out-001.nut, etc., and write the list of generated segments to out.list:
	ffmpeg -i in.mkv -codec hevc -flags +cgop -g 60 -map 0 -f segment -segment_list out.list out%03d.nut
Segment input and set output format options for the output segments:
	ffmpeg -i in.mkv -f segment -segment_time 10 -segment_format_options movflags=+faststart out%03d.mp4
Segment the input file according to the split points specified by the segment_times option:
	ffmpeg -i in.mkv -codec copy -map 0 -f segment -segment_list out.csv -segment_times 1,2,3,5,8,13,21 out%03d.nut
Use the ffmpeg force_key_frames option to force key frames in the input at the specified location, together with the segment option segment_time_delta to account for possible roundings operated when setting key frame times.
	ffmpeg -i in.mkv -force_key_frames 1,2,3,5,8,13,21 -codec:v mpeg4 -codec:a pcm_s16le -map 0 \
	-f segment -segment_list out.csv -segment_times 1,2,3,5,8,13,21 -segment_time_delta 0.05 out%03d.nut
In order to force key frames on the input file, transcoding is required.

Segment the input file by splitting the input file according to the frame numbers sequence specified with the segment_frames option:
	ffmpeg -i in.mkv -codec copy -map 0 -f segment -segment_list out.csv -segment_frames 100,200,300,500,800 out%03d.nut
Convert the in.mkv to TS segments using the libx264 and aac encoders:
	ffmpeg -i in.mkv -map 0 -codec:v libx264 -codec:a aac -f ssegment -segment_list out.list out%03d.ts
Segment the input file, and create an M3U8 live playlist (can be used as live HLS source):
	ffmpeg -re -i in.mkv -codec copy -map 0 -f segment -segment_list playlist.m3u8 \
-segment_list_flags +live -segment_time 10 out%03d.mkv

*/
