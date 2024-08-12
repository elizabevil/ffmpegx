package muxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// // 21.73 webm_dash_manifest WEBM_DASH_MANIFEST
type WEBM_DASH_MANIFEST struct {
	AdaptationSets string `json:"adaptation_sets" flag:"-adaptation_sets"`
	//This option has the following syntax: "id=x,streams=a,b,c id=y,streams=d,e" where x and y are the unique identifiers of the adaptation sets and a,b,c,d and e are the indices of the corresponding audio and video streams. Any number of adaptation sets can be added using this option.

	Live typex.UBool `json:"live" flag:"-live"`
	//Set this to 1 to create a live stream DASH Manifest. Default: 0.

	ChunkStartIndex typex.Index `json:"chunk_start_index" flag:"-chunk_start_index"`
	//Start index of the first chunk. This will go in the ‘startNumber’ attribute of the ‘SegmentTemplate’ element in the manifest. Default: 0.

	ChunkDurationMs typex.MillisecondI `json:"chunk_duration_ms" flag:"-chunk_duration_ms"`
	//Duration of each chunk in milliseconds. This will go in the ‘duration’ attribute of the ‘SegmentTemplate’ element in the manifest. Default: 1000.

	UtcTimingUrl typex.Url `json:"utc_timing_url" flag:"-utc_timing_url"`
	//URL of the page that will return the UTC timestamp in ISO format. This will go in the ‘value’ attribute of the ‘UTCTiming’ element in the manifest. Default: None.

	TimeShiftBufferDepth typex.DepthF `json:"time_shift_buffer_depth" flag:"-time_shift_buffer_depth"`
	//Smallest time (in seconds) shifting buffer for which any Representation is guaranteed to be available. This will go in the ‘timeShiftBufferDepth’ attribute of the ‘MPD’ element. Default: 60.

	MinimumUpdatePeriod typex.SecondI `json:"minimum_update_period" flag:"-minimum_update_period"`
	//Minimum update period (in seconds) of the manifest. This will go in the ‘minimumUpdatePeriod’ attribute of the ‘MPD’ element. Default: 0.
}
