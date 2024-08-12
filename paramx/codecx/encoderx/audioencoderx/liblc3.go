package audioencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// // LIBLC3 8.6 liblc3
type LIBLC3 struct {
	B typex.Bitrate `json:"b" flag:"-b"`
	//Set the bit rate in bits/s. This will determine the fixed size of the encoded frames, for a selected frame duration.

	Ar typex.SampleRate `json:"ar" flag:"-ar"`
	//Set the audio sampling rate (in Hz).

	Channels typex.Channels `json:"channels" flag:"-channels"`
	//Set the number of audio channels.

	FrameDuration typex.MillisecondI `json:"frame_duration" flag:"-frame_duration"`
	//Set the audio frame duration in milliseconds. Default value is 10ms. Allowed frame durations are 2.5ms, 5ms, 7.5ms and 10ms. LC3 (Bluetooth LE Audio), allows 7.5ms and 10ms; and LC3plus 2.5ms, 5ms and 10ms.

	//The 10ms frame duration is available in LC3 and LC3 plus standard. In this mode, the produced bitstream can be referenced either as LC3 or LC3plus.

	HighResolution typex.Bool `json:"high_resolution" flag:"-high_resolution"`
	//Enable the high-resolution mode if set to 1. The high-resolution mode is available with all LC3plus frame durations and for a sampling rate of 48 KHz, and 96 KHz.

	//The encoder automatically turns off this mode at lower sampling rates and activates it at 96 KHz.

	//This mode should be preferred at high bitrates. In this mode, the audio bandwidth is always up to the Nyquist frequency, compared to LC3 at 48 KHz, which limits the bandwidth to 20 KHz.

}
