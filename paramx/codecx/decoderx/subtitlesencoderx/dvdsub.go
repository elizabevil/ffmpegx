package subtitledecoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// DVDSUB 6.4 dvdsub
type DVDSUB struct {
	Palette typex.Palette `json:"palette" flag:"-palette"`
	/*
		Specify the global palette used by the bitmaps. When stored in VobSub, the palette is normally specified in the index file; in Matroska, the palette is stored in the codec extra-data in the same format as in VobSub. In DVDs, the palette is stored in the IFO file, and therefore not available when reading from dumped VOB files.
		The format for this option is a string containing 16 24-bits hexadecimal numbers (without 0x prefix)
	*/

	IfoPalette typex.Palette `json:"ifo_palette" flag:"-ifo_palette"`
	//Specify the IFO file from which the global palette is obtained. (experimental)
	ForcedSubsOnly typex.Bool `json:"forced_subs_only" flag:"-forced_subs_only"`
	//Only decode subtitle entries marked as forced.Some titles have forced and non-forced subtitles in the same track.Setting this flag to 1 will only keep the forced subtitles.Default value is 0.
}
