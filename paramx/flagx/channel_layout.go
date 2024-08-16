package flagx

// ChannelLayout A channel layout specifies the spatial disposition of the channels in a multi-channel audio stream. To specify a channel layout, FFmpeg makes use of a special syntax.
// Individual channels are identified by an id, as given by the table below:
type ChannelLayout = Flag

const (
	ChannelLayout_FL ChannelLayout = "FL" //front left

	ChannelLayout_FR ChannelLayout = "FR" //front right

	ChannelLayout_FC ChannelLayout = "FC" //front center

	ChannelLayout_LFE ChannelLayout = "LFE" //low frequency

	ChannelLayout_BL ChannelLayout = "BL" //back left

	ChannelLayout_BR ChannelLayout = "BR" //back right

	ChannelLayout_FLC ChannelLayout = "FLC" //front left-of-center

	ChannelLayout_FRC ChannelLayout = "FRC" //front right-of-center

	ChannelLayout_BC ChannelLayout = "BC" //back center

	ChannelLayout_SL ChannelLayout = "SL" //side left

	ChannelLayout_SR ChannelLayout = "SR" //side right

	ChannelLayout_TC ChannelLayout = "TC" //top center

	ChannelLayout_TFL ChannelLayout = "TFL" //top front left

	ChannelLayout_TFC ChannelLayout = "TFC" //top front center

	ChannelLayout_TFR ChannelLayout = "TFR" //top front right

	ChannelLayout_TBL ChannelLayout = "TBL" //top back left

	ChannelLayout_TBC ChannelLayout = "TBC" //top back center

	ChannelLayout_TBR ChannelLayout = "TBR" //top back right

	ChannelLayout_DL ChannelLayout = "DL" //downmix left

	ChannelLayout_DR ChannelLayout = "DR" //downmix right

	ChannelLayout_WL ChannelLayout = "WL" //wide left

	ChannelLayout_WR ChannelLayout = "WR" //wide right

	ChannelLayout_SDL ChannelLayout = "SDL" //surround direct left

	ChannelLayout_SDR ChannelLayout = "SDR" //surround direct right

	ChannelLayout_LFE2 ChannelLayout = "LFE2" //low frequency 2

	//Standard channel layout compositions can be specified by using the following identifiers:

	ChannelLayout_mono ChannelLayout = "mono" //FC

	ChannelLayout_stereo ChannelLayout = "stereo" //FL+FR
	//===

	ChannelLayout_2_1 ChannelLayout = "2.1" //FL+FR+LFE

	ChannelLayout_3_0 ChannelLayout = "3.0" //FL+FR+FC

	ChannelLayout_3_0_back ChannelLayout = "3.0(back)" //FL+FR+BC

	ChannelLayout_4_0 ChannelLayout = "4.0" //FL+FR+FC+BC

	ChannelLayout_quad ChannelLayout = "quad" //FL+FR+BL+BR

	ChannelLayout_quad_side ChannelLayout = "quad(side)" //FL+FR+SL+SR

	ChannelLayout_3_1 ChannelLayout = "3.1" //FL+FR+FC+LFE

	ChannelLayout_5_0 ChannelLayout = "5.0" //FL+FR+FC+BL+BR

	ChannelLayout_5_0_side ChannelLayout = "5.0(side)" //FL+FR+FC+SL+SR

	ChannelLayout_4_1 ChannelLayout = "4.1" //FL+FR+FC+LFE+BC

	ChannelLayout_5_1 ChannelLayout = "5.1" //FL+FR+FC+LFE+BL+BR

	ChannelLayout_5_1_side ChannelLayout = "5.1(side)" //FL+FR+FC+LFE+SL+SR

	ChannelLayout_6_0 ChannelLayout = "6.0" //FL+FR+FC+BC+SL+SR

	ChannelLayout_6_0_front ChannelLayout = "6.0(front)" //FL+FR+FLC+FRC+SL+SR

	ChannelLayout_3_1_2 ChannelLayout = "3.1.2" //FL+FR+FC+LFE+TFL+TFR

	ChannelLayout_hexagonal ChannelLayout = "hexagonal" //FL+FR+FC+BL+BR+BC

	ChannelLayout_6_1_r ChannelLayout = "6.1" //FL+FR+FC+LFE+BC+SL+SR

	ChannelLayout_6_1_c ChannelLayout = "6.1" //FL+FR+FC+LFE+BL+BR+BC

	ChannelLayout_6_1_front ChannelLayout = "6.1(front)" //FL+FR+LFE+FLC+FRC+SL+SR

	ChannelLayout_7_0 ChannelLayout = "7.0" //FL+FR+FC+BL+BR+SL+SR

	ChannelLayout_7_0_front ChannelLayout = "7.0(front)" //FL+FR+FC+FLC+FRC+SL+SR

	ChannelLayout_7_1 ChannelLayout = "7.1" //FL+FR+FC+LFE+BL+BR+SL+SR

	ChannelLayout_7_1_wide      ChannelLayout = "7.1(wide)"      //FL+FR+FC+LFE+BL+BR+FLC+FRC
	ChannelLayout_7_1_wide_side ChannelLayout = "7.1(wide-side)" //FL+FR+FC+LFE+FLC+FRC+SL+SR
	ChannelLayout_5_1_2         ChannelLayout = "5.1.2"          //FL+FR+FC+LFE+BL+BR+TFL+TFR

	ChannelLayout_octagonal ChannelLayout = "octagonal" //FL+FR+FC+BL+BR+BC+SL+SR

	ChannelLayout_cube ChannelLayout = "cube" //FL+FR+BL+BR+TFL+TFR+TBL+TBR

	ChannelLayout_5_1_4 ChannelLayout = "5.1.4" //FL+FR+FC+LFE+BL+BR+TFL+TFR+TBL+TBR
	ChannelLayout_7_1_2 ChannelLayout = "7.1.2" //FL+FR+FC+LFE+BL+BR+SL+SR+TFL+TFR
	ChannelLayout_7_1_4 ChannelLayout = "7.1.4" //FL+FR+FC+LFE+BL+BR+SL+SR+TFL+TFR+TBL+TBR
	ChannelLayout_7_2_3 ChannelLayout = "7.2.3" //FL+FR+FC+LFE+BL+BR+SL+SR+TFL+TFR+TBC+LFE2
	ChannelLayout_9_1_4 ChannelLayout = "9.1.4" //FL+FR+FC+LFE+BL+BR+FLC+FRC+SL+SR+TFL+TFR+TBL+TBR

	ChannelLayout_hexadecagonal ChannelLayout = "hexadecagonal" //FL+FR+FC+BL+BR+BC+SL+SR+WL+WR+TBL+TBR+TBC+TFC+TFL+TFR

	ChannelLayout_downmix ChannelLayout = "downmix" //DL+DR

	ChannelLayout_22_2 ChannelLayout = "22.2" //FL+FR+FC+LFE+BL+BR+FLC+FRC+BC+SL+SR+TC+TFL+TFC+TFR+TBL+TBC+TBR+LFE2+TSL+TSR+BFC+BFL+BFR

)

//A custom channel layout can be specified as a sequence of terms, separated by ’+’. Each term can be:
//the name of a single channel (e.g. ‘FL’, ‘FR’, ‘FC’, ‘LFE’, etc.), each optionally containing a custom name after a ’@’, (e.g. ‘FL@Left’, ‘FR@Right’, ‘FC@Center’, ‘LFE@Low_Frequency’, etc.)

/*

A standard channel layout can be specified by the following:
	the name of a single channel (e.g. ‘FL’, ‘FR’, ‘FC’, ‘LFE’, etc.)
	the name of a standard channel layout (e.g. ‘mono’, ‘stereo’, ‘4.0’, ‘quad’, ‘5.0’, etc.)
	a number of channels, in decimal, followed by ’c’, yielding the default channel layout for that number of channels (see the function av_channel_layout_default). Note that not all channel counts have a default layout.
	a number of channels, in decimal, followed by ’C’, yielding an unknown channel layout with the specified number of channels. Note that not all channel layout specification Flagss support unknown channel layouts.
	a channel layout mask, in hexadecimal starting with "0x" (see the AV_CH_* macros in libavutil/channel_layout.h.
*/
/*
Before libavutil version 53 the trailing character "c" to specify a number of channels was optional, but now it is required, while a channel layout mask can also be specified as a decimal number (if and only if not followed by "c" or "C").
*/
