package codecx

import "github.com/elizabevil/ffmpegx/paramx/typex"

type Common struct {
	BitsPerCodedSample typex.UNumber  `json:"bits_per_coded_sample" flag:"-bits_per_coded_sample"`
	Delay              typex.Number   `json:"delay" flag:"-delay"`
	TimeBase           typex.TimeBase `json:"time_base" flag:"-time_base"`
	HasBFrames         typex.UNumber  `json:"has_b_frames" flag:"-has_b_frames"`
	BlockAlign         typex.UNumber  `json:"block_align" flag:"-block_align"`
	RcOverrideCount    typex.Number   `json:"rc_override_count" flag:"-rc_override_count"`
	SliceCount         typex.UNumber  `json:"slice_count" flag:"-slice_count"`
	SliceFlags         typex.Number   `json:"slice_flags" flag:"-slice_flags"`
	LogLevelOffset     typex.Number   `json:"log_level_offset" flag:"-log_level_offset"`
	PktTimebase        typex.UNumber  `json:"pkt_timebase" flag:"-pkt_timebase"`
}
