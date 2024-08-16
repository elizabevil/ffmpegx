package demuxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// 20.19 mpjpeg  MPJPEG
type MPJPEG struct {
	StrictMimeBoundary typex.UBool `json:"strict_mime_boundary" flag:"-strict_mime_boundary"`
	//Default implementation applies a relaxed standard to multi-part MIME boundary detection, to prevent regression with numerous existing endpoints not generating a proper MIME MJPEG stream. Turning this option on by setting it to 1 will result in a stricter check of the boundary value.

}
