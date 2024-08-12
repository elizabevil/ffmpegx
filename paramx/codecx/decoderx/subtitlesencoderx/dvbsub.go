package subtitledecoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// DVBSUB 6.3 dvbsub
type DVBSUB struct {
	ComputeClut typex.Bool `json:"compute_clut" flag:"-compute_clut"`
	/*
		-2
		Compute clut once if no matching CLUT is in the stream.
		-1
		Compute clut if no matching CLUT is in the stream.
		0
		Never compute CLUT
		1
		Always compute CLUT and override the one provided in the stream.
	*/

	DvbSubstream *typex.Level `json:"dvb_substream" flag:"-dvb_substream,-1"`
	//Selects the dvb substream, or all substreams if -1 which is default.
}
