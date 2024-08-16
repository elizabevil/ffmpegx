package inputdevicex

import "github.com/elizabevil/ffmpegx/paramx/typex"

// OPENAL 3.15 openal
type OPENAL struct {
	Channels typex.Channels `json:"channels" flag:"-channels"`

	//Set the number of channels in the captured audio. Only the values 1 (monaural) and 2 (stereo) are currently supported. Defaults to 2.
	SampleSize typex.Size `json:"sample_size" flag:"-sample_size"`

	//Set the sample size (in bits) of the captured audio. Only the values 8 and 16 are currently supported. Defaults to 16.

	SampleRate typex.SampleRate `json:"sample_rate" flag:"-sample_rate"`

	//Set the sample rate (in Hz) of the captured audio. Defaults to 44.1k.

	ListDevices typex.BoolStr `json:"list_devices" flag:"-list_devices"`
	//If set to true, print a list of devices and exit. Defaults to false.
}

/*

Print the list of OpenAL supported devices and exit:

$ ffmpeg -list_devices true -f openal -i dummy out.ogg
Capture from the OpenAL device DR-BT101 via PulseAudio:

$ ffmpeg -f openal -i 'DR-BT101 via PulseAudio' out.ogg
Capture from the default device (note the empty string ” as filename):

$ ffmpeg -f openal -i '' out.ogg
Capture from two devices simultaneously, writing to two different files, within the same ffmpeg command:

$ ffmpeg -f openal -i 'DR-BT101 via PulseAudio' out1.ogg -f openal -i 'ALSA Default' out2.ogg
Note: not all OpenAL implementations support multiple simultaneous capture - try the latest OpenAL Soft if the above does not work.
*/
