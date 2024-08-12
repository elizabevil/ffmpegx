package subtitledecoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// LIBZVBI-TELETEXT 6.5 libzvbi-teletext
type LibzvbiTeletext struct {
	TxtPage typex.Fmt `json:"txt_page" flag:"-txt_page"`
	//List of teletext page numbers to decode. Pages that do not match the specified list are dropped. You may use the special * string to match all pages, or subtitle to match all subtitle pages. Default value is *.

	TxtDefaultRegion typex.Bool `json:"txt_default_region" flag:"-txt_default_region"`
	//Set default character set used for decoding, a value between 0 and 87 (see ETS 300 706, Section 15, Table 32). Default value is -1, which does not override the libzvbi default. This option is needed for some legacy level 1.0 transmissions which cannot signal the proper charset.

	TxtChopTop typex.UBool `json:"txt_chop_top" flag:"-txt_chop_top"`
	//Discards the top teletext line. Default value is 1.

	TxtFormat TxtFormat `json:"txt_left" flag:"-txt_left"`

	TxtLeft typex.U16Offset `json:"txt_format" flag:"-txt_format"`
	//X offset of generated bitmaps, default is 0.

	TxtTop typex.U16Offset `json:"txt_top" flag:"-txt_top"`
	//Y offset of generated bitmaps, default is 0.

	TxtChopSpaces typex.UBool `json:"txt_chop_spaces" flag:"-txt_chop_spaces"`
	//Chops leading and trailing spaces and removes empty lines from the generated text. This option is useful for teletext based subtitles where empty spaces may be present at the start or at the end of the lines or empty lines may be present between the subtitle lines because of double-sized teletext characters. Default value is 1.

	TxtDuration typex.MillisecondI `json:"txt_duration" flag:"-txt_duration"`
	//Sets the display duration of the decoded teletext pages or subtitles in milliseconds. Default value is -1 which means infinity or until the next subtitle event comes.

	TxtTransparent typex.UBool `json:"txt_transparent" flag:"-txt_transparent"`
	//Force transparent background of the generated teletext bitmaps. Default value is 0 which means an opaque background.

	TxtOpacity typex.Opacity `json:"txt_opacity" flag:"-txt_opacity"`
	//Sets the opacity (0-255) of the teletext background. If txt_transparent is not set, it only affects characters between a start box and an end box, typically subtitles. Default value is 0 if txt_transparent is set, 255 otherwise.
}

type TxtFormat string

const (
	TxtFormat_bitmap TxtFormat = "bitmap"
	//The default format, you should use this for teletext pages, because certain graphics and colors cannot be expressed in simple text or even ASS.

	TxtFormat_text TxtFormat = "text"
	//Simple text based output without formatting.

	TxtFormat_ass TxtFormat = "ass"
	//Formatted ASS output, subtitle pages and teletext pages are returned in different styles, subtitle pages are stripped down to text, but an effort is made to keep the text alignment and the formatting.
)
