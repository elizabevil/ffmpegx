package demuxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// GIF https://ffmpeg.org/ffmpeg-all.html
// 20.11 gif
type GIF struct {
	MinDelay typex.Delay `json:"min_delay" flag:"-min_delay"`
	//Set the minimum valid delay between frames in hundredths of seconds. Range is 0 to 6000. Default value is 2.

	MaxGifDelay typex.Delay `json:"max_gif_delay" flag:"-max_gif_delay"`
	//Set the maximum valid delay between frames in hundredth of seconds. Range is 0 to 65535. Default value is 65535 (nearly eleven minutes), the maximum value allowed by the specification.

	DefaultDelay typex.Delay `json:"default_delay" flag:"-default_delay"`
	//Set the default delay between frames in hundredths of seconds. Range is 0 to 6000. Default value is 10.

	IgnoreLoop *typex.UBool `json:"ignore_loop" flag:"-ignore_loop,1"` //Default value is 1
	//GIF files can contain information to loop a certain number of times (or infinitely). If ignore_loop is set to 1, then the loop setting from the input will be ignored and looping will not occur. If set to 0, then looping will occur and will cycle the number of times according to the GIF. Default value is 1.

}

//ffmpeg -i input.mp4 -ignore_loop 0 -i input.gif -filter_complex overlay=shortest=1 out.mkv
