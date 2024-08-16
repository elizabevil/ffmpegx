package flagx

type DuplexMode = Flag

const (
	DuplexMode_unset                DuplexMode = "unset"
	DuplexMode_half                 DuplexMode = "half"
	DuplexMode_full                 DuplexMode = "full"
	DuplexMode_one_sub_device_full  DuplexMode = "one_sub_device_full"
	DuplexMode_one_sub_device_half  DuplexMode = "one_sub_device_half"
	DuplexMode_two_sub_device_full  DuplexMode = "two_sub_device_full"
	DuplexMode_four_sub_device_half DuplexMode = "four_sub_device_half"
)

type PixelFormat = Flag

const (
	PixelFormat_monob        PixelFormat = "monob"
	PixelFormat_rgb555be     PixelFormat = "rgb555be"
	PixelFormat_rgb555le     PixelFormat = "rgb555le"
	PixelFormat_rgb565be     PixelFormat = "rgb565be"
	PixelFormat_rgb565le     PixelFormat = "rgb565le"
	PixelFormat_rgb24        PixelFormat = "rgb24"
	PixelFormat_bgr24        PixelFormat = "bgr24"
	PixelFormat_0rgb         PixelFormat = "0rgb"
	PixelFormat_bgr0         PixelFormat = "bgr0"
	PixelFormat_0bgr         PixelFormat = "0bgr"
	PixelFormat_rgb0         PixelFormat = "rgb0"
	PixelFormat_bgr48be      PixelFormat = "bgr48be"
	PixelFormat_uyvy422      PixelFormat = "uyvy422"
	PixelFormat_yuva444p     PixelFormat = "yuva444p"
	PixelFormat_yuva444p16le PixelFormat = "yuva444p16le"
	PixelFormat_yuv444p      PixelFormat = "yuv444p"
	PixelFormat_yuv422p16    PixelFormat = "yuv422p16"
	PixelFormat_yuv422p10    PixelFormat = "yuv422p10"
	PixelFormat_yuv444p10    PixelFormat = "yuv444p10"
	PixelFormat_yuv420p      PixelFormat = "yuv420p"
	PixelFormat_nv12         PixelFormat = "nv12"
	PixelFormat_yuyv422      PixelFormat = "yuyv422"
	PixelFormat_gray         PixelFormat = "gray"
)
