package inputdevicex

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// AVFOUNDATION 3.3 avfoundation
type AVFOUNDATION struct {
	ListDevices typex.BoolStr `json:"list_devices" flag:"-list_devices"`
	//If set to true, a list of all available input devices is given showing all device names and indices.

	VideoDeviceIndex typex.Index `json:"video_device_index" flag:"-video_device_index"`
	//Specify the video device by its index. Overrides anything given in the input filename.

	AudioDeviceIndex typex.Index `json:"audio_device_index" flag:"-audio_device_index"`
	//Specify the audio device by its index. Overrides anything given in the input filename.

	PixelFormat flagx.PixelFormat `json:"pixel_format" flag:"-pixel_format"`

	Framerate typex.FrameRate `json:"framerate" flag:"-framerate"`
	//Set the grabbing frame rate.Default is ntsc, corresponding to a frame rate of 30000/1001.
	FramerateFlag flagx.VideoRate `json:"framerate_flag" flag:"-framerate"`

	VideoSize flagx.VideoSize `json:"video_size" flag:"-video_size"`

	//Set the video frame size.

	CaptureCursor typex.UBool `json:"capture_cursor" flag:"-capture_cursor"`
	//Capture the Mouse pointer.Default is 0.

	CaptureMouseClicks typex.UBool `json:"capture_mouse_clicks" flag:"-capture_mouse_clicks"`
	//Capture the screen mouse clicks.Default is 0.

	CaptureRawData typex.UBool `json:"capture_raw_data" flag:"-capture_raw_data"`
	//Capture the raw device data.Default is 0. Using this option may result in receiving the underlying data delivered to the AVFoundation framework.E.g.for muxed devices that sends raw DV data to the framework (like tape-based camcorders), setting this option to false results in extracted video frames captured in the designated pixel format only.Setting this option to true results in receiving the raw DV stream untouched.
}

/*
Print the list of AVFoundation supported devices and exit:
$ ffmpeg -f avfoundation -list_devices true -i ""
Record video from video device 0 and audio from audio device 0 into out.avi:
$ ffmpeg -f avfoundation -i "0:0" out.avi
Record video from video device 2 and audio from audio device 1 into out.avi:
$ ffmpeg -f avfoundation -video_device_index 2 -i ":1" out.avi
Record video from the system default video device using the pixel format bgr0 and do not record any audio into out.avi:
$ ffmpeg -f avfoundation -pixel_format bgr0 -i "default:none" out.avi
Record raw DV data from a suitable input device and write the output into out.dv:
$ ffmpeg -f avfoundation -capture_raw_data true -i "zr100:none" out.dv
*/
