package inputdevicex

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// KMSGRAB 3.11 kmsgrab
type KMSGRAB struct {
	Device string `json:"device" flag:"-device"`
	//DRM device to capture on. Defaults to /dev/dri/card0.

	Format typex.Format `json:"format" flag:"-format"`
	//Pixel format of the framebuffer. This can be autodetected if you are running Linux 5.7 or later, but needs to be provided for earlier versions. Defaults to bgr0, which is the most common format used by the Linux console and Xorg X server.

	FormatModifier int `json:"format_modifier" flag:"-format_modifier"`
	//Format modifier to signal on output frames. This is necessary to import correctly into  APIs. It can be autodetected if you are running Linux 5.7 or later, but will need to be provided explicitly when needed in earlier versions. See the libdrm documentation for possible values.

	CrtcId typex.ID `json:"crtc_id" flag:"-crtc_id"`
	//KMS CRTC ID to define the capture source. The first active plane on the given CRTC will be used.

	PlaneId typex.ID `json:"plane_id" flag:"-plane_id"`
	//KMS plane ID to define the capture source. Defaults to the first active plane found if neither crtc_id nor plane_id are specified.

	Framerate typex.FrameRate `json:"framerate" flag:"-framerate"`
	//Framerate to capture at. This is not synchronised to any page flipping or framebuffer changes - it just defines the interval at which the framebuffer is sampled. Sampling faster than the framebuffer update rate will generate independent frames with the same content. Defaults to 30.
	FramerateFlag flagx.VideoRate `json:"framerate_flag" flag:"-framerate"`
}

/*
Capture from the first active plane, download the result to normal frames and encode. This will only work if the framebuffer is both linear and mappable - if not, the result may be scrambled or fail to download.
  ffmpeg -f kmsgrab -i - -vf 'hwdownload,format=bgr0' output.mp4
Capture from CRTC ID 42 at 60fps, map the result to VAAPI, convert to NV12 and encode as H.264.
  ffmpeg -crtc_id 42 -framerate 60 -f kmsgrab -i - -vf 'hwmap=derive_device=vaapi,scale_vaapi=w=1920:h=1080:format=nv12' -c:v h264_vaapi output.mp4
To capture only part of a plane the output can be cropped - this can be used to capture a single window, as long as it has a known absolute position and size. For example, to capture and encode the middle quarter of a 1920x1080 plane:
  ffmpeg -f kmsgrab -i - -vf 'hwmap=derive_device=vaapi,crop=960:540:480:270,scale_vaapi=960:540:nv12' -c:v h264_vaapi output.mp4
*/
