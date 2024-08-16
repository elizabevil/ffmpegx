package inputdevicex

import "github.com/elizabevil/ffmpegx/paramx/typex"

// OSS 3.16 oss
type OSS struct {
	Channels typex.Channels `json:"channels" flag:"-channels"`

	SampleRate typex.SampleRate `json:"sample_rate" flag:"-sample_rate"`
}
