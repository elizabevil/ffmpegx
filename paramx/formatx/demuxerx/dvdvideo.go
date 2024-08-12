package demuxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// DVD 20.7 dvdvideo
// VD-Video demuxer, powered by libdvdnav and libdvdread.
type DVD struct {
	Title int `json:"title" flag:"-title"`
	//The title number to play. Must be set if pgc and pg are not set. Not applicable to menus. Default is 0 (auto), which currently only selects the first available title (title 1) and notifies the user about the implications.

	ChapterStart int `json:"chapter_start" flag:"-chapter_start,1"`
	//The chapter, or PTT (part-of-title), number to start at. Not applicable to menus. Default is 1.

	ChapterEnd int `json:"chapter_end" flag:"-chapter_end"`
	//The chapter, or PTT (part-of-title), number to end at. Not applicable to menus. Default is 0, which is a special value to signal end at the last possible chapter.

	Angle int `json:"angle" flag:"-angle,1"`
	//The video angle number, referring to what is essentially an additional video stream that is composed from alternate frames interleaved in the VOBs. Not applicable to menus. Default is 1.

	Region int `json:"region" flag:"-region"`
	//The region code to use for playback. Some discs may use this to default playback at a particular angle in different regions. This option will not affect the region code of a real DVD drive, if used as an input. Not applicable to menus. Default is 0, "world".

	Menu bool `json:"menu" flag:"-menu"`
	//Demux menu assets instead of navigating a title. Requires exact coordinates of the menu (menu_lu, menu_vts, pgc, pg). Default is false.

	MenuLu int `json:"menu_lu" flag:"-menu_lu"`
	//The menu language to demux. In DVD, menus are grouped by language. Default is 0, the first language unit.

	MenuVts int `json:"menu_vts" flag:"-menu_vts"`
	//The VTS where the menu lives, or 0 if it is a VMG menu (root-level). Default is 0, VMG menu.

	Pgc int `json:"pgc" flag:"-pgc"`
	//The entry PGC to start playback, in conjunction with pg. Alternative to setting title. Chapter markers are not supported at this time. Must be explicitly set for menus. Default is 0, automatically resolve from value of title.

	Pg int `json:"pg" flag:"-pg"`
	//The entry PG to start playback, in conjunction with pgc. Alternative to setting title. Chapter markers are not supported at this time. Default is 0, automatically resolve from value of title, or start from the beginning (PG 1) of the menu.

	Preindex bool `json:"preindex" flag:"-preindex"`
	//Enable this to have accurate chapter (PTT) markers and duration measurement, which requires a slow second pass read in order to index the chapter marker timestamps from NAV packets. This is non-ideal extra work for real optical drives. It is recommended and faster to use this option with a backup of the DVD structure stored on a hard drive. Not compatible with pgc and pg. Not applicable to menus. Default is 0, false.

	Trim *typex.UBool `json:"trim" flag:"-trim,1"`
	//Skip padding cells (i.e. cells shorter than 1 second) from the beginning. There exist many discs with filler segments at the beginning of the PGC, often with junk data intended for controlling a real DVD player’s buffering speed and with no other material data value. Not applicable to menus. Default is 1, true.
}

/*
Open title 3 from a given DVD structure:
	ffmpeg -f dvdvideo -title 3 -i <path to DVD> ...

Open chapters 3-6 from title 1 from a given DVD structure:
	ffmpeg -f dvdvideo -chapter_start 3 -chapter_end 6 -title 1 -i <path to DVD> ...

Open only chapter 5 from title 1 from a given DVD structure:
	ffmpeg -f dvdvideo -chapter_start 5 -chapter_end 5 -title 1 -i <path to DVD> ...

Demux menu with language 1 from VTS 1, PGC 1, starting at PG 1:
	ffmpeg -f dvdvideo -menu 1 -menu_lu 1 -menu_vts 1 -pgc 1 -pg 1 -i <path to DVD> ...

*/
