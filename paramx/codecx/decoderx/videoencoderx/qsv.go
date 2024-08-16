package videodecoderx

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// QSV 4.7 QSV Decoders
type QSV struct {
	AsyncDepth typex.UNumber `json:"async_depth" flag:"-async_depth"`
	//Internal parallelization depth, the higher the value the higher the latency.

	GpuCopy GpuCopy `json:"gpu_copy" flag:"-gpu_copy"`
	//Extra options for hevc_qsv.
	LoadPlugin  flagx.LoadPlugin `json:"load_plugin" flag:"-load_plugin"`
	LoadPlugins string           `json:"load_plugins" flag:"-load_plugins"`
}

type GpuCopy = typex.String

const (
	GpuCopy_default GpuCopy = "default"
	GpuCopy_on      GpuCopy = "on"
	GpuCopy_off     GpuCopy = "off"
)
