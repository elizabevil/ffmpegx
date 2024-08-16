package inputdevicex

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

//26.5 decklink

type Decklink struct {
	ListDevices typex.BoolStr `json:"list_devices" flag:"-list_devices"` //
	//If set to true, print a list of devices and exit. Defaults to false. This option is deprecated, please use the -sources option of ffmpeg to list the available input devices.

	ListFormats bool `json:"list_formats" flag:"-list_formats"` //
	//If set to true, print a list of supported formats and exit. Defaults to false.

	FormatCode bool `json:"format_code" flag:"-format_code"` //<FourCC>
	//This sets the input video format to the format given by the FourCC. To see the supported values of your device(s) use list_formats. Note that there is a FourCC 'pal ' that can also be used as pal (3 letters). Default behavior is autodetection of the input video format, if the hardware supports it.

	RawFormat RawFormat `json:"raw_format" flag:"-raw_format"` //
	//Set the pixel format of the captured video. Available values are:
	TeletextLines bool `json:"teletext_lines" flag:"-teletext_lines"` //
	//If set to nonzero, an additional teletext stream will be captured from the vertical ancillary data. Both SD PAL (576i) and HD (1080i or 1080p) sources are supported. In case of HD sources, OP47 packets are decoded.

	//This option is a bitmask of the SD PAL VBI lines captured, specifically lines 6 to 22, and lines 318 to 335. Line 6 is the LSB in the mask. Selected lines which do not contain teletext information will be ignored. You can use the special all constant to select all possible lines, or standard to skip lines 6, 318 and 319, which are not compatible with all receivers.

	//For SD sources, ffmpeg needs to be compiled with --enable-libzvbi. For HD sources, on older (pre-4K) DeckLink card models you have to capture in 10 bit mode.

	Channels ChannelNumber `json:"channels" flag:"-channels"` //
	//Defines number of audio channels to capture. Must be ‘2, ‘8 or ‘16. Defaults to ‘2.

	DuplexMode flagx.DuplexMode `json:"duplex_mode" flag:"-duplex_mode"` //
	//Sets the decklink device duplex/profile mode. Must be ‘unset, ‘half, ‘full, ‘one_sub_device_full, ‘one_sub_device_half, ‘two_sub_device_full, ‘four_sub_device_half Defaults to ‘unset.

	//Note: DeckLink SDK 11.0 have replaced the duplex property by a profile property. For the DeckLink Duo 2 and DeckLink Quad 2, a profile is shared between any 2 sub-devices that utilize the same connectors. For the DeckLink 8K Pro, a profile is shared between all 4 sub-devices. So DeckLink 8K Pro support four profiles.

	//Valid profile modes for DeckLink 8K Pro(with DeckLink SDK >= 11.0): ‘one_sub_device_full, ‘one_sub_device_half, ‘two_sub_device_full, ‘four_sub_device_half

	//Valid profile modes for DeckLink Quad 2 and DeckLink Duo 2: ‘half, ‘full

	TimecodeFormat bool `json:"timecode_format" flag:"-timecode_format"` //
	//Timecode type to include in the frame and video stream metadata. Must be ‘none, ‘rp188vitc, ‘rp188vitc2, ‘rp188ltc, ‘rp188hfr, ‘rp188any, ‘vitc, ‘vitc2, or ‘serial. Defaults to ‘none (not included).

	//In order to properly support 50/60 fps timecodes, the ordering of the queried timecode types for ‘rp188any is HFR, VITC1, VITC2 and LTC for >30 fps content. Note that this is slightly different to the ordering used by the DeckLink API, which is HFR, VITC1, LTC, VITC2.

	VideoInput VideoInput `json:"video_input" flag:"-video_input"` //
	//Sets the video input source. Must be ‘unset, ‘sdi, ‘hdmi, ‘optical_sdi, ‘component, ‘composite or ‘s_video. Defaults to ‘unset.

	AudioInput AudioInput `json:"audio_input" flag:"-audio_input"` //
	//Sets the audio input source. Must be ‘unset, ‘embedded, ‘aes_ebu, ‘analog, ‘analog_xlr, ‘analog_rca or ‘microphone. Defaults to ‘unset.

	VideoPts PTS `json:"video_pts" flag:"-video_pts"` //
	//Sets the video packet timestamp source. Must be ‘video, ‘audio, ‘reference, ‘wallclock or ‘abs_wallclock. Defaults to ‘video.

	AudioPts PTS `json:"audio_pts" flag:"-audio_pts"` //
	//Sets the audio packet timestamp source. Must be ‘video, ‘audio, ‘reference, ‘wallclock or ‘abs_wallclock. Defaults to ‘audio.

	DrawBars bool `json:"draw_bars" flag:"-draw_bars"` //
	//If set to ‘true, color bars are drawn in the event of a signal loss. Defaults to ‘true.

	QueueSize typex.Size `json:"queue_size" flag:"-queue_size"` //
	//Sets maximum input buffer size in bytes. If the buffering reaches this value, incoming frames will be dropped. Defaults to ‘1073741824.

	AudioDepth AudioDepth `json:"audio_depth" flag:"-audio_depth"` //
	//Sets the audio sample bit depth. Must be ‘16 or ‘32. Defaults to ‘16.

	DecklinkCopyts bool `json:"decklink_copyts" flag:"-decklink_copyts"` //
	//If set to true, timestamps are forwarded as they are without removing the initial offset. Defaults to false.

	TimestampAlign typex.Bool `json:"timestamp_align" flag:"-timestamp_align"` //
	//Capture start time alignment in seconds. If set to nonzero, input frames are dropped till the system timestamp aligns with configured value. Alignment difference of up to one frame duration is tolerated. This is useful for maintaining input synchronization across N different hardware devices deployed for N-way redundancy. The system time of different hardware devices should be synchronized with protocols such as NTP or PTP, before using this option. Note that this method is not foolproof. In some border cases input synchronization may not happen due to thread scheduling jitters in the OS. Either sync could go wrong by 1 frame or in a rarer case timestamp_align seconds. Defaults to ‘0.

	WaitForTc bool `json:"wait_for_tc" flag:"-wait_for_tc"` // (bool)
	//Drop frames till a frame with timecode is received. Sometimes serial timecode isnt received with the first input frame. If that happens, the stored stream timecode will be inaccurate. If this option is set to true, input frames are dropped till a frame with timecode is received. Option timecode_format must be specified. Defaults to false.

	EnableKlv bool `json:"enable_klv" flag:"-enable_klv"` //(bool)
	//If set to true, extracts KLV data from VANC and outputs KLV packets. KLV VANC packets are joined based on MID and PSC fields and aggregated into one KLV packet. Defaults to false.
}

/*

List input devices:
	ffmpeg -sources decklink
List supported formats:
	ffmpeg -f decklink -list_formats 1 -i 'Intensity Pro'
Capture video clip at 1080i50:
	ffmpeg -format_code Hi50 -f decklink -i 'Intensity Pro' -c:a copy -c:v copy output.avi
Capture video clip at 1080i50 10 bit:
	ffmpeg -raw_format yuv422p10 -format_code Hi50 -f decklink -i 'UltraStudio Mini Recorder' -c:a copy -c:v copy output.avi
Capture video clip at 1080i50 with 16 audio channels:
	ffmpeg -channels 16 -format_code Hi50 -f decklink -i 'UltraStudio Mini Recorder' -c:a copy -c:v copy output.avi

*/

type RawFormat = typex.Format
type TimecodeFormat string
type VideoInput string
type AudioInput string
type ChannelNumber = int8

const (
	ChannelNumber_2  ChannelNumber = iota + 2
	ChannelNumber_8  ChannelNumber = iota << 3
	ChannelNumber_16 ChannelNumber = iota << 3
)
const (
	RawFormat_auto RawFormat = "auto"
	//This is the default which means 8-bit YUV 422 or 8-bit ARGB if format autodetection is used, 8-bit YUV 422 otherwise.

	RawFormat_uyvy422 RawFormat = "uyvy422"
	//8-bit YUV 422.

	RawFormat_yuv422p10 RawFormat = "yuv422p10"
	//10-bit YUV 422.

	RawFormat_argb RawFormat = "argb"
	//8-bit RGB.

	RawFormat_bgra RawFormat = "bgra"
	//8-bit RGB.

	RawFormat_rgb10 RawFormat = "rgb10"
	//10-bit RGB.

)
const (
	TimecodeFormat_none       TimecodeFormat = "none"
	TimecodeFormat_rp188vitc  TimecodeFormat = "rp188vitc"
	TimecodeFormat_rp188vitc2 TimecodeFormat = "rp188vitc2"
	TimecodeFormat_rp188ltc   TimecodeFormat = "rp188ltc"
	TimecodeFormat_rp188hfr   TimecodeFormat = "rp188hfr"
	TimecodeFormat_rp188any   TimecodeFormat = "rp188any"
	TimecodeFormat_vitc       TimecodeFormat = "vitc"
	TimecodeFormat_vitc2      TimecodeFormat = "vitc2"
	TimecodeFormat_serial     TimecodeFormat = "serial"
)
const (
	VideoInput_unset       VideoInput = "unset"
	VideoInput_sdi         VideoInput = "sdi"
	VideoInput_hdmi        VideoInput = "hdmi"
	VideoInput_optical_sdi VideoInput = "optical_sdi"
	VideoInput_component   VideoInput = "component"
	VideoInput_composite   VideoInput = "composite"
	VideoInput_s_video     VideoInput = "s_video"
)
const (
	AudioInput_unset      AudioInput = "unset"
	AudioInput_embedded   AudioInput = "embedded"
	AudioInput_aes_ebu    AudioInput = "aes_ebu"
	AudioInput_analog     AudioInput = "analog"
	AudioInput_analog_xlr AudioInput = "analog_xlr"
	AudioInput_analog_rca AudioInput = "analog_rca"
	AudioInput_microphone AudioInput = "microphone"
)

type PTS string
type AudioDepth string

const (
	AudioDepth_16 AudioDepth = "16"
	AudioDepth_32 AudioDepth = "32"
)
const (
	PTS_video         PTS = "video"
	PTS_audio         PTS = "audio"
	PTS_reference     PTS = "reference"
	PTS_wallclock     PTS = "wallclock"
	PTS_abs_wallclock PTS = "abs_wallclock"
)
