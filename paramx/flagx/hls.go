package flagx

type HlsStartNumberSource = Flag

const (
	HlsStartNumberSource_Generic  HlsStartNumberSource = "generic"
	HlsStartNumberSource_Epoch    HlsStartNumberSource = "epoch"
	HlsStartNumberSource_EpochUs  HlsStartNumberSource = "epoch_us"
	HlsStartNumberSource_Datetime HlsStartNumberSource = "datetime"
)

type HlsFlags = Flag

const (
	HlsFlags_single_file HlsFlags = "single_file"

	HlsFlags_delete_segments HlsFlags = "delete_segments"

	HlsFlags_append_list HlsFlags = "append_list"

	HlsFlags_round_durations HlsFlags = "round_durations"

	HlsFlags_discont_start HlsFlags = "discont_start"

	HlsFlags_omit_endlist HlsFlags = "omit_endlist"

	HlsFlags_periodic_rekey HlsFlags = "periodic_rekey"

	HlsFlags_independent_segments HlsFlags = "independent_segments"

	HlsFlags_iframes_only HlsFlags = "iframes_only"

	HlsFlags_split_by_time HlsFlags = "split_by_time"

	HlsFlags_program_date_time HlsFlags = "program_date_time"

	HlsFlags_second_level_segment_index HlsFlags = "second_level_segment_index"

	HlsFlags_second_level_segment_size HlsFlags = "second_level_segment_size"

	HlsFlags_second_level_segment_duration HlsFlags = "second_level_segment_duration"

	HlsFlags_temp_file HlsFlags = "temp_file"
)

type HlsPlaylistType = Flag

const (
	//If type is ‘event’, emit #EXT-X-PLAYLIST-TYPE:EVENT in the m3u8 header. This forces hls_list_size to 0; the playlist can only be appended to.
	HlsPlaylistType_event HlsPlaylistType = "event"
	//If type is ‘vod’, emit #EXT-X-PLAYLIST-TYPE:VOD in the m3u8 header. This forces hls_list_size to 0; the playlist must not change.
	HlsPlaylistType_vod HlsPlaylistType = "vod"
)
