package optionx

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

type Global struct {
	Overwrite     bool          `json:"y"                       flag:"-y"`
	ExitWhenExist bool          `json:"n"                       flag:"-n"` //Do not overwrite output files, and exit immediately if a specified output file already exists.
	RecastMedia   bool          `json:"recast_media"            flag:"-recast_media"`
	FilterThreads int           `json:"filter_threads"          flag:"-filter_threads"`
	Stats         bool          `json:"stats"                   flag:"-stats"`
	NoStats       bool          `json:"nostats"                 flag:"-nostats"`
	StatsPeriod   typex.SecondF `json:"stats_period"            flag:"-stats_period"`
	Progress      typex.Url     `json:"progress"                flag:"-progress"`

	DebugTs      bool           `json:"debug_ts"                flag:"-debug_ts"`
	Benchmark    bool           `json:"benchmark"               flag:"-benchmark"`
	BenchmarkAll bool           `json:"benchmark_all"           flag:"-benchmark_all"`
	TimeLimit    typex.Duration `json:"timelimit"               flag:"-timelimit"` //Exit after ffmpeg has been running for duration seconds in CPU user timex.
	Dump         bool           `json:"dump"                    flag:"-dump"`
	Hex          bool           `json:"hex"                     flag:"-hex"`
	//Vsync                   bool              `json:"vsync"                   flag:"-vsync"`
	FilterComplex           typex.Filtergraph `json:"filter_complex"          flag:"-filter_complex"`
	FilterComplexThreads    typex.NbThreads   `json:"filter_complex_threads"  flag:"-filter_complex_threads"` //The default is the number of available CPUs.
	Lavfi                   typex.Filtergraph `json:"lavfi"                   flag:"-lavfi"`
	SdpFile                 bool              `json:"sdp_file"                flag:"-sdp_file"`
	AbortOn                 flagx.AbortOn     `json:"abort_on"                flag:"-abort_on"`
	MaxErrorRate            bool              `json:"max_error_rate"          flag:"-max_error_rate"` // Range is a floatingpoint number between 0 to 1. Default is 2/3.
	Xerror                  bool              `json:"xerror"                  flag:"-xerror"`         //Stop and exit on error
	AutoConversionFilters   bool              `json:"auto_conversion_filters" flag:"-auto_conversion_filters"`
	NoAutoConversionFilters bool              `json:"noauto_conversion_filters" flag:"-noauto_conversion_filters"`
	Cpuflags                string            `json:"cpuflags"                flag:"-cpuflags"`
	Cpucount                typex.UNumber     `json:"cpucount"                flag:"-cpucount"`
}
