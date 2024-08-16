package muxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// 21.72 webm_chunk
type WebmChunk struct {
	ChunkStartIndex typex.Index `json:"chunk_start_index" flag:"-chunk_start_index"`
	//Index of the first chunk (defaults to 0).

	Header typex.FileName `json:"header" flag:"-header"`
	//FileName of the header where the initialization data will be written.

	AudioChunkDuration typex.MillisecondI `json:"audio_chunk_duration" flag:"-audio_chunk_duration"`
	//Duration of each audio chunk in milliseconds (defaults to 5000).
}

/*
ffmpeg -f v4l2 -i /dev/video0 \
       -f alsa -i hw:0 \
       -map 0:0 \
       -c:v libvpx-vp9 \
       -s 640x360 -keyint_min 30 -g 30 \
       -f webm_chunk \
       -header webm_live_video_360.hdr \
       -chunk_start_index 1 \
       webm_live_video_360_%d.chk \
       -map 1:0 \
       -c:a libvorbis \
       -b:a 128k \
       -f webm_chunk \
       -header webm_live_audio_128.hdr \
       -chunk_start_index 1 \
       -audio_chunk_duration 1000 \
       webm_live_audio_128_%d.chk
*/

// //21.73 webm_dash_manifest
type WebmDashManifest struct {
	AdaptationSets string `json:"adaptation_sets" flag:"-adaptation_sets"`
	//This option has the following syntax: "id=x,streams=a,b,c id=y,streams=d,e" where x and y are the unique identifiers of the adaptation sets and a,b,c,d and e are the indices of the corresponding audio and video streams. Any number of adaptation sets can be added using this option.

	Live typex.Bool `json:"live" flag:"-live"`
	//Set this to 1 to create a live stream DASH Manifest. Default: 0.

	ChunkStartIndex typex.Index `json:"chunk_start_index" flag:"-chunk_start_index"`
	//Start index of the first chunk. This will go in the ‘startNumber’ attribute of the ‘SegmentTemplate’ element in the manifest. Default: 0.

	ChunkDurationMs typex.MillisecondI `json:"chunk_duration_ms" flag:"-chunk_duration_ms"`
	//Duration of each chunk in milliseconds. This will go in the ‘duration’ attribute of the ‘SegmentTemplate’ element in the manifest. Default: 1000.

	UtcTimingUrl typex.Url `json:"utc_timing_url" flag:"-utc_timing_url"`
	//URL of the page that will return the UTC timestamp in ISO format. This will go in the ‘value’ attribute of the ‘UTCTiming’ element in the manifest. Default: None.

	TimeShiftBufferDepth typex.Depth `json:"time_shift_buffer_depth" flag:"-time_shift_buffer_depth"`
	//Smallest time (in seconds) shifting buffer for which any Representation is guaranteed to be available. This will go in the ‘timeShiftBufferDepth’ attribute of the ‘MPD’ element. Default: 60.

	MinimumUpdatePeriod typex.SecondI `json:"minimum_update_period" flag:"-minimum_update_period"`
	//Minimum update period (in seconds) of the manifest. This will go in the ‘minimumUpdatePeriod’ attribute of the ‘MPD’ element. Default: 0.
}

/*
ffmpeg -f webm_dash_manifest -i video1.webm \
       -f webm_dash_manifest -i video2.webm \
       -f webm_dash_manifest -i audio1.webm \
       -f webm_dash_manifest -i audio2.webm \
       -map 0 -map 1 -map 2 -map 3 \
       -c copy \
       -f webm_dash_manifest \
       -adaptation_sets "id=0,streams=0,1 id=1,streams=2,3" \
       manifest.xml
*/
