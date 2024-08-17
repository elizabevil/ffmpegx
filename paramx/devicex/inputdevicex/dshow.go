package inputdevicex

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// 26.6 dshow
// Windows DirectShow input device.
// DirectShow support is enabled when FFmpeg is built with the mingw-w64 project. Currently only audio and video devices are supported.
// Multiple devices may be opened as separate inputs, but they may also be opened on the same input, which should improve synchronism between them.
// TYPE=NAME[:TYPE=NAME]
type Dshow struct {
	VideoSize flagx.VideoSize `json:"video_size" flag:"-video_size"`

	Framerate typex.FrameRate `json:"framerate" flag:"-framerate"`
	//Set the frame rate in the captured video.
	FramerateFlag flagx.VideoRate `json:"framerate_flag" flag:"-framerate"`

	SampleRate typex.SampleRate `json:"sample_rate" flag:"-sample_rate"`
	//Set the sample rate (in Hz) of the captured audio.

	SampleSize typex.Size `json:"sample_size" flag:"-sample_size"`
	//Set the sample size (in bits) of the captured audio.

	Channels typex.Channels `json:"channels" flag:"-channels"`
	//Set the number of channels in the captured audio.

	ListDevices typex.BoolStr `json:"list_devices" flag:"-list_devices"`
	//If set to true, print a list of devices and exit.

	ListOptions typex.BoolStr `json:"list_options" flag:"-list_options"`
	//If set to true, print a list of selected device’s options and exit.

	VideoDeviceNumber typex.Number `json:"video_device_number" flag:"-video_device_number"`
	//Set video device number for devices with the same name (starts at 0, defaults to 0).

	AudioDeviceNumber typex.Number `json:"audio_device_number" flag:"-audio_device_number"`
	//Set audio device number for devices with the same name (starts at 0, defaults to 0).

	PixelFormat flagx.PixelFormat `json:"pixel_format" flag:"-pixel_format"`
	//Select pixel format to be used by DirectShow. This may only be set when the video codec is not set or set to rawvideo.

	AudioBufferSize typex.Size `json:"audio_buffer_size" flag:"-audio_buffer_size"`
	//Set audio device buffer size in milliseconds (which can directly impact latency, depending on the device). Defaults to using the audio device’s default buffer size (typically some multiple of 500ms). Setting this value too low can degrade performance. See also http://msdn.microsoft.com/en-us/library/windows/desktop/dd377582(v=vs.85).aspx

	VideoPinName typex.Name `json:"video_pin_name" flag:"-video_pin_name"`
	//Select video capture pin to use by name or alternative name.

	AudioPinName typex.Name `json:"audio_pin_name" flag:"-audio_pin_name"`
	//Select audio capture pin to use by name or alternative name.

	CrossbarVideoInputPinNumber typex.Number `json:"crossbar_video_input_pin_number" flag:"-crossbar_video_input_pin_number"`
	//Select video input pin number for crossbar device. This will be routed to the crossbar device’s Video Decoder output pin. Note that changing this value can affect future invocations (sets a new default) until system reboot occurs.

	CrossbarAudioInputPinNumber typex.Number `json:"crossbar_audio_input_pin_number" flag:"-crossbar_audio_input_pin_number"`
	//Select audio input pin number for crossbar device. This will be routed to the crossbar device’s Audio Decoder output pin. Note that changing this value can affect future invocations (sets a new default) until system reboot occurs.

	ShowVideoDeviceDialog typex.BoolStr `json:"show_video_device_dialog" flag:"-show_video_device_dialog"`
	//If set to true, before capture starts, popup a display dialog to the end user, allowing them to change video filter properties and configurations manually. Note that for crossbar devices, adjusting values in this dialog may be needed at times to toggle between PAL (25 fps) and NTSC (29.97) input frame rates, sizes, interlacing, etc. Changing these values can enable different scan rates/frame rates and avoiding green bars at the bottom, flickering scan lines, etc. Note that with some devices, changing these properties can also affect future invocations (sets new defaults) until system reboot occurs.

	ShowAudioDeviceDialog typex.UBool `json:"show_audio_device_dialog" flag:"-show_audio_device_dialog"`
	//If set to true, before capture starts, popup a display dialog to the end user, allowing them to change audio filter properties and configurations manually.

	ShowVideoCrossbarConnectionDialog typex.UBool `json:"show_video_crossbar_connection_dialog" flag:"-show_video_crossbar_connection_dialog"`
	//If set to true, before capture starts, popup a display dialog to the end user, allowing them to manually modify crossbar pin routings, when it opens a video device.

	ShowAudioCrossbarConnectionDialog typex.UBool `json:"show_audio_crossbar_connection_dialog" flag:"-show_audio_crossbar_connection_dialog"`
	//If set to true, before capture starts, popup a display dialog to the end user, allowing them to manually modify crossbar pin routings, when it opens an audio device.

	ShowAnalogTvTunerDialog typex.UBool `json:"show_analog_tv_tuner_dialog" flag:"-show_analog_tv_tuner_dialog"`
	//If set to true, before capture starts, popup a display dialog to the end user, allowing them to manually modify TV channels and frequencies.

	ShowAnalogTvTunerAudioDialog typex.UBool `json:"show_analog_tv_tuner_audio_dialog" flag:"-show_analog_tv_tuner_audio_dialog"`
	//If set to true, before capture starts, popup a display dialog to the end user, allowing them to manually modify TV audio (like mono vs. stereo, Language A,B or C).

	AudioDeviceLoad typex.UBool `json:"audio_device_load" flag:"-audio_device_load"`
	//Load an audio capture filter device from file instead of searching it by name. It may load additional parameters too, if the filter supports the serialization of its properties to. To use this an audio capture source has to be specified, but it can be anything even fake one.

	AudioDeviceSave typex.UBool `json:"audio_device_save" flag:"-audio_device_save"`
	//Save the currently used audio capture filter device and its parameters (if the filter supports it) to a file. If a file with the same name exists it will be overwritten.

	VideoDeviceLoad typex.UBool `json:"video_device_load" flag:"-video_device_load"`
	//Load a video capture filter device from file instead of searching it by name. It may load additional parameters too, if the filter supports the serialization of its properties to. To use this a video capture source has to be specified, but it can be anything even fake one.

	VideoDeviceSave typex.UBool `json:"video_device_save" flag:"-video_device_save"`
	//Save the currently used video capture filter device and its parameters (if the filter supports it) to a file. If a file with the same name exists it will be overwritten.

	UseVideoDeviceTimestamps typex.BoolStr `json:"use_video_device_timestamps" flag:"-use_video_device_timestamps"`
	//If set to false, the timestamp for video frames will be derived from the wallclock instead of the timestamp provided by the capture device. This allows working around devices that provide unreliable timestamps.
}
