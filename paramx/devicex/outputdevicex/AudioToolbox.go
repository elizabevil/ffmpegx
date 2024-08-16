package outputdevicex

import "github.com/elizabevil/ffmpegx/paramx/typex"

// AUDIOTOOLBOX 4.2 AudioToolbox
type AUDIOTOOLBOX struct {
	AudioDeviceIndex typex.Index `json:"audio_device_index" flag:"-audio_device_index"`
	//Specify the audio device by its index. Overrides anything given in the output filename.

}

/*

Print the list of supported devices and output a sine wave to the default device:
$ ffmpeg -f lavfi -i sine=r=44100 -f audiotoolbox -list_devices true -
Output a sine wave to the device with the index 2, overriding any output filename:
$ ffmpeg -f lavfi -i sine=r=44100 -f audiotoolbox -audio_device_index 2 -
*/
