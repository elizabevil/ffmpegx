package optionx

import "github.com/elizabevil/ffmpegx/paramx/typex"

type FFplay struct {
	X            typex.Position `json:"x" flag:"-x"`                         // {.func_arg = opt_width }, "force displayed width", "width" },
	Y            typex.Position `json:"y" flag:"-y"`                         // {.func_arg = opt_height }, "force displayed height", "height" },
	Fs           bool           `json:"fs" flag:"-fs"`                       // {&is_full_screen }, "force full screen" },
	An           bool           `json:"an" flag:"-an"`                       // {&audio_disable }, "disable audio" },
	Vn           bool           `json:"vn" flag:"-vn"`                       // {&video_disable }, "disable video" },
	Sn           bool           `json:"sn" flag:"-sn"`                       // {&subtitle_disable }, "disable subtitling" },
	Ss           typex.Position `json:"ss" flag:"-ss"`                       // {&start_time }, "seek to a given position in seconds", "pos" },
	T            typex.Position `json:"t" flag:"-t"`                         // {&duration }, "play  \"duration\" seconds of audio/video", "duration" },
	Bytes        typex.Number   `json:"bytes" flag:"-bytes"`                 // {&seek_by_bytes }, "seek by bytes 0=off 1=on -1=auto", "val" },
	SeekInterval typex.Flt      `json:"seek_interval" flag:"-seek_interval"` // {&seek_interval }, "set seek interval for left/right keys, in seconds", "seconds" },
	Nodisp       bool           `json:"nodisp" flag:"-nodisp"`               // {&display_disable }, "disable graphical display" },
	Noborder     bool           `json:"noborder" flag:"-noborder"`           // {&borderless }, "borderless window" },
	Alwaysontop  bool           `json:"alwaysontop" flag:"-alwaysontop"`     // {&alwaysontop }, "window always on top" },
	Volume       typex.UI8      `json:"volume" flag:"-volume"`               // {&startup_volume}, "set startup volume 0=min 100=max", "volume" },
	F            typex.Format   `json:"f" flag:"-f"`                         // {.func_arg = opt_format }, "force format", "fmt" },
	WindowTitle  typex.FileName `json:"window_title" flag:"-window_title"`   // {&window_title }, "set window title", "window title" },
	Left         typex.Position `json:"left" flag:"-left"`                   // {&screen_left }, "set the x position for the left of the window", "x pos" },

	Top typex.Position `json:"top" flag:"-top"` // {&screen_top }, "set the y position for the top of the window", "y pos" },

	Loop typex.Number      `json:"loop" flag:"-loop"` // {&loop }, "set number of times the playback shall be looped", "loop count" },
	Vf   typex.Filtergraph `json:"vf" flag:"-vf"`     // {.func_arg = opt_add_vfilter }, "set video filters", "filter_graph" },

	Af       typex.Filtergraph `json:"af" flag:"-af"`             // {&afilters }, "set audio filters", "filter_graph" },
	Showmode Showmode          `json:"showmode" flag:"-showmode"` // {.func_arg = opt_show_mode}, "select show mode (0 = video, 1 = waves, 2 = RDFT)", "mode" },
	I        typex.FileName    `json:"i" flag:"-i"`               // {&dummy}, "read specified file", "input_file"},
}

type Showmode = typex.Flags

const (
	Showmode_video Showmode = "video" //0
	//show video

	Showmode_waves Showmode = "waves" //1
	//show audio waves

	Showmode_rdft Showmode = "rdft" //2
)

type FFplayExpert struct {
	Codec typex.Codec `json:"codec" flag:"-codec"` // {.func_arg = opt_codec}, "force decoder", "decoder_name" },

	//======
	Stats bool `json:"stats" flag:"-stats"` // {&show_status }, "show status"

	Fast bool `json:"fast" flag:"-fast"` // {&fast }, "non spec compliant optimizations"

	Sync typex.Type `json:"sync" flag:"-sync"` // {.func_arg = opt_sync }, "set audio-video sync. type (type=audio/video/ext)", "type" },

	Ast typex.StreamSpecifier `json:"ast" flag:"-ast"` // {&wanted_stream_spec[AVMEDIA_TYPE_AUDIO] }, "select desired audio stream", "stream_specifier" },

	Vst typex.StreamSpecifier `json:"vst" flag:"-vst"` // {&wanted_stream_spec[AVMEDIA_TYPE_VIDEO] }, "select desired video stream", "stream_specifier" },

	Sst typex.StreamSpecifier `json:"sst" flag:"-sst"` // {&wanted_stream_spec[AVMEDIA_TYPE_SUBTITLE] }, "select desired subtitle stream", "stream_specifier" },

	Autorotate   bool `json:"autorotate" flag:"-autorotate"`     // {&autorotate }, "automatically rotate video"
	NoAutorotate bool `json:"noautorotate" flag:"-noautorotate"` // {&autorotate }, "automatically rotate video"

	Genpts bool `json:"genpts" flag:"-genpts"` // {&genpts }, "generate pts"

	Drp *typex.Bool `json:"drp" flag:"-drp"` // {&decoder_reorder_pts }, "let decoder reorder pts 0=off 1=on -1=auto"

	Lowres typex.Number `json:"lowres" flag:"-lowres"` // {&lowres }, ""

	Autoexit bool `json:"autoexit" flag:"-autoexit"` // {&autoexit }, "exit at the end"

	Exitonkeydown bool `json:"exitonkeydown" flag:"-exitonkeydown"` // {&exit_on_keydown }, "exit on key down"

	Exitonmousedown bool `json:"exitonmousedown" flag:"-exitonmousedown"` // {&exit_on_mousedown }, "exit on mouse down"

	Framedrop   bool `json:"framedrop" flag:"-framedrop"`     // {&framedrop }, "drop frames when cpu is too slow"
	NoFramedrop bool `json:"noframedrop" flag:"-noframedrop"` // {&framedrop }, "drop frames when cpu is too slow"

	Infbuf   bool `json:"infbuf" flag:"-infbuf"`     // {&infinite_buffer }, "don't limit the input buffer size (useful with realtime streams)"
	NoInfbuf bool `json:"noinfbuf" flag:"-noinfbuf"` // {&infinite_buffer }, "don't limit the input buffer size (useful with realtime streams)"

	Acodec typex.String `json:"acodec" flag:"-acodec"` // {   &audio_codec_name }, "force audio decoder",    "decoder_name" },

	Scodec typex.String `json:"scodec" flag:"-scodec"` // {&subtitle_codec_name }, "force subtitle decoder", "decoder_name" },

	Vcodec typex.String `json:"vcodec" flag:"-vcodec"` // {   &video_codec_name }, "force video decoder",    "decoder_name" },

	FilterThreads typex.NbThreads `json:"filter_threads" flag:"-filter_threads"` // {&filter_nbthreads }, "number of filter threads per graph" },
	EnableVulkan  bool            `json:"enable_vulkan" flag:"-enable_vulkan"`   // {&enable_vulkan }, "enable vulkan renderer" },

	VulkanParams typex.ParamsKv `json:"vulkan_params" flag:"-vulkan_params"` // {&vulkan_params }, "vulkan configuration using a list of key=value pairs separated by ':'" },

	Hwaccel typex.String `json:"hwaccel" flag:"-hwaccel"` // {&hwaccel }, "use HW accelerated decoding" },

}
