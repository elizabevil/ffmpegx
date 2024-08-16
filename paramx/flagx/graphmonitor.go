package flagx

//graphmonitor

type ModeFlags = Flag
type MFlags = ModeFlags

const (
	ModeFlags_full       ModeFlags = "full"
	ModeFlags_compact    ModeFlags = "compact"
	ModeFlags_nozero     ModeFlags = "nozero"
	ModeFlags_noeof      ModeFlags = "noeof"
	ModeFlags_nodisabled ModeFlags = "nodisabled"
)

type GraphmonitorFlag Flag //Set = Flags which enable which stats are shown in video.
const (
	none GraphmonitorFlag = "none"
	//All flags turned off.

	all GraphmonitorFlag = "all"
	//All flags turned on.

	queue GraphmonitorFlag = "queue"
	//Display number of queued frames in each link.

	frame_count_in GraphmonitorFlag = "frame_count_in"
	//Display number of frames taken from filter.

	frame_count_out GraphmonitorFlag = "frame_count_out"
	//Display number of frames given out from filter.

	frame_count_delta GraphmonitorFlag = "frame_count_delta"
	//Display delta number of frames between above two values.

	pts GraphmonitorFlag = "pts"
	//Display current filtered frame pts.

	pts_delta GraphmonitorFlag = "pts_delta"
	//Display pts delta between current and previous frame.

	time GraphmonitorFlag = "time"
	//Display current filtered frame time.

	time_delta GraphmonitorFlag = "time_delta"
	//Display time delta between current and previous frame.

	timebase GraphmonitorFlag = "timebase"
	//Display time base for filter link.

	format GraphmonitorFlag = "format"
	//Display used format for filter link.

	size GraphmonitorFlag = "size"
	//Display video size or number of audio channels in case of audio used by filter link.

	rate GraphmonitorFlag = "rate"
	//Display video frame rate or sample rate in case of audio used by filter link.

	eof GraphmonitorFlag = "eof"
	//Display link output status.

	sample_count_in GraphmonitorFlag = "sample_count_in"
	//Display number of samples taken from filter.

	sample_count_out GraphmonitorFlag = "sample_count_out"
	//Display number of samples given out from filter.

	sample_count_delta GraphmonitorFlag = "sample_count_delta"
	//Display delta number of samples between above two values.

	disabled GraphmonitorFlag = "disabled"
	//Show the timeline filter status.
)
