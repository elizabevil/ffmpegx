package flagx

type FrameTypeFlag = Flag
type FtFlag = FrameTypeFlag

const (
	FrameTypeFlag_if FrameTypeFlag = "if"
	//intra-coded frames (I-frames)

	FrameTypeFlag_pf FrameTypeFlag = "pf"
	//predicted frames (P-frames)

	FrameTypeFlag_bf FrameTypeFlag = "bf"
	//bi-directionally predicted frames (B-frames)

)

type SkipFrame = Flag

const (
	SkipFrame_none    SkipFrame = "none"    //Discard no frame.
	SkipFrame_default SkipFrame = "default" //Discard useless frames like 0-sized frames.
	SkipFrame_noref   SkipFrame = "noref"   //Discard all non-reference frames.
	SkipFrame_bidir   SkipFrame = "bidir"   //Discard all bidirectional frames.
	SkipFrame_nokey   SkipFrame = "nokey"   //Discard all frames excepts keyframes.
	SkipFrame_nointra SkipFrame = "nointra" //Discard all frames except I frames.
	SkipFrame_all     SkipFrame = "all"     //Discard all frames.
)

type Discard = SkipFrame

const (
	Discard_none    Discard = SkipFrame_none    //Discard no frame.
	Discard_default Discard = SkipFrame_default //Default, which discards no frames.
	Discard_noref   Discard = SkipFrame_noref   //Discard all non-reference frames.
	Discard_bidir   Discard = SkipFrame_bidir   //Discard all bidirectional frames.
	Discard_nokey   Discard = SkipFrame_nokey   //Discard all frames excepts keyframes.
	Discard_all     Discard = SkipFrame_all     //Discard all frames.
)
