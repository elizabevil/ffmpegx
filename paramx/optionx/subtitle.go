package optionx

import "github.com/elizabevil/ffmpegx/paramx/typex"

type Subtitle struct {
	Scodec typex.Codec `json:"scodec" flag:"-scodec"` // (input/output)
	//Set the subtitle codec.This is an alias for -codec:s.
	Sn bool `json:"sn" flag:"-sn"`
	//As an input option, blocks all subtitle streams of a file from being filtered or being automatically selected or mapped for any output.See -discard option to disable streams individually.
	//As an output option, disables subtitle recording i.e.automatic selection or mapping of any subtitle stream.For full manual control see the -map option.
	//5.10 Advanced Subtitle options
	FixSubDuration bool `json:"fix_sub_duration" flag:"-fix_sub_duration"`
	//Fix subtitles durations.For each subtitle, wait for the next packet in the same stream and adjust the duration of the first to avoid overlap.This is necessary with some subtitles codecs, especially DVB subtitles, because the duration in the original packet is only a rough estimate and the end is actually marked by an empty subtitle frame.Failing to use this option when necessary can result in exaggerated durations or muxing failures due to non-monotonic timestamps.
	//Note that this option will delay the output of all data until the next subtitle packet is decoded: it may increase memory consumption and latency a lot.
	CanvasSize typex.VideoSize `json:"canvas_size" flag:"-canvas_size"`
	//Set the size of the canvas used to render subtitles.
}
