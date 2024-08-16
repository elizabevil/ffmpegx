package flagx

type Avioflags = Flag

const (
	Avioflags_direct Avioflags = "direct"
)

type Loglevel = Flag

const (
	Loglevel_prefix_repeat Loglevel = "repeat"
	Loglevel_prefix_level  Loglevel = "level"
	//

	Loglevel_quiet Loglevel = "quiet" //, -8
	//Show nothing at all; be silent.
	Loglevel_panic Loglevel = "panic" //, 0
	//Only show fatal errors which could lead the process to crash, such as an assertion failure. This is not currently used for anything.
	Loglevel_fatal Loglevel = "fatal" //, 8
	//Only show fatal errors. These are errors after which the process absolutely cannot continue.
	Loglevel_error Loglevel = "error" //, 16
	//Show all errors, including ones which can be recovered from.
	Loglevel_warning Loglevel = "warning" //, 24
	//Show all warnings and errors. Any message related to possibly incorrect or unexpected events will be shown.
	Loglevel_info Loglevel = "info" //, 32
	//Show informative messages during processing. This is in addition to warnings and errors. This is the default value.
	Loglevel_verbose Loglevel = "verbose" //, 40
	//Same as info, except more verbose.
	Loglevel_debug Loglevel = "debug" //, 48
	//Show everything, including debugging information.
	Loglevel_trace Loglevel = "trace" //, 56
)

type Preset string

const (
	Preset_none = "none"
	//Do not use a preset.

	Preset_default = "default"
	//Use the encoder default.

	Preset_picture = "picture"
	//Digital picture, like portrait, inner shot

	Preset_photo = "photo"
	//Outdoor photograph, with natural lighting

	Preset_drawing = "drawing"
	//Hand or line drawing, with high-contrast details

	Preset_icon = "icon"
	//Small-sized colorful images

	Preset_text = "text"
	//Text-like
)
