package optionx

import (
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

//ffprobe [options] input_url

type FFprobe struct {
	F                   typex.Format     `json:"f" flag:"-f"`                                          // force format int"format
	Unit                bool             `json:"unit" flag:"-unit"`                                    // show unit of the displayed values
	Prefix              bool             `json:"prefix" flag:"-prefix"`                                // use SI prefixes for the displayed values
	ByteBinaryPrefix    bool             `json:"byte_binary_prefix" flag:"-byte_binary_prefix"`        // use binary prefixes for byte units
	Sexagesimal         bool             `json:"sexagesimal" flag:"-sexagesimal"`                      // use sexagesimal format HOURS:MM:SS.MICROSECONDS for time units
	Pretty              typex.Format     `json:"pretty" flag:"-pretty"`                                // prettify the format of displayed values, make it more human readable
	OutputFormat        OutputFormat     `json:"output_format" flag:"-output_format"`                  // set the output printing format (available formats are: default, compact, csv, flat, ini, json, xml) int"format
	PrintFormat         OutputFormat     `json:"print_format" flag:"-print_format"`                    // alias for -output_format (deprecated
	Of                  OutputFormat     `json:"of" flag:"-of"`                                        // alias for -output_format int"format
	SelectStreams       typex.StreamName `json:"select_streams" flag:"-select_streams"`                // select the specified streams int"stream_specifier
	Sections            typex.String     `json:"sections" flag:"-sections"`                            // print sections structure and section information, and exit
	ShowData            bool             `json:"show_data" flag:"-show_data"`                          // show packets data
	ShowDataHash        typex.String     `json:"show_data_hash" flag:"-show_data_hash"`                // show packets data hash
	ShowError           bool             `json:"show_error" flag:"-show_error"`                        // "show probing error
	ShowFormat          bool             `json:"show_format" flag:"-show_format"`                      // show format/container info
	ShowFrames          bool             `json:"show_frames" flag:"-show_frames"`                      // show frames info
	ShowEntries         typex.Entries    `json:"show_entries" flag:"-show_entries"`                    // show a set of specified entries int"entry_list
	ShowLog             typex.LevelI8    `json:"show_log" flag:"-show_log"`                            // show log
	ShowPackets         bool             `json:"show_packets" flag:"-show_packets"`                    // show packets info
	ShowPrograms        bool             `json:"show_programs" flag:"-show_programs"`                  // show programs info
	ShowStreamGroups    bool             `json:"show_stream_groups" flag:"-show_stream_groups"`        // show stream groups info
	ShowStreams         bool             `json:"show_streams" flag:"-show_streams"`                    // show streams info
	ShowChapters        bool             `json:"show_chapters" flag:"-show_chapters"`                  // show chapters info
	CountFrames         bool             `json:"count_frames" flag:"-count_frames"`                    // count the number of frames per stream
	CountPackets        bool             `json:"count_packets" flag:"-count_packets"`                  // count the number of packets per stream
	ShowProgramVersion  bool             `json:"show_program_version" flag:"-show_program_version"`    // "show ffprobe version
	ShowLibraryVersions bool             `json:"show_library_versions" flag:"-show_library_versions"`  // show library versions
	ShowVersions        bool             `json:"show_versions" flag:"-show_versions"`                  // show program and library versions
	ShowPixelFormats    bool             `json:"show_pixel_formats" flag:"-show_pixel_formats"`        // show pixel format descriptions
	ShowOptionalFields  *typex.Bool      `json:"show_optional_fields" flag:"-show_optional_fields,-1"` // show optional fields
	ShowPrivateData     bool             `json:"show_private_data" flag:"-show_private_data"`          // show private data
	Private             bool             `json:"private" flag:"-private"`                              // same as show_private_data
	Bitexact            bool             `json:"bitexact" flag:"-bitexact"`                            // force bitexact output
	ReadIntervals       typex.String     `json:"read_intervals" flag:"-read_intervals"`                // set read intervals int"read_intervals
	I                   typex.Filename   `json:"i" flag:"-i"`                                          // read specified file int"input_file"},
	O                   typex.Filename   `json:"o" flag:"-o"`                                          // write to specified output int"output_file"},
}

type OutputFormat = typex.Flags

const (
	OutputFormat_default OutputFormat = "default"
	OutputFormat_compact OutputFormat = "compact"
	OutputFormat_csv     OutputFormat = "csv"
	OutputFormat_flat    OutputFormat = "flat"
	OutputFormat_ini     OutputFormat = "ini"
	OutputFormat_json    OutputFormat = "json"
	OutputFormat_xml     OutputFormat = "xml"
)

type WriterOption struct {
	StringValidation            int `json:"string_validation" flag:"-string_validation"`                         //0, WRITER_STRING_VALIDATION_NB-1, .unit = "sv//
	Sv                          Sv  `json:"sv" flag:"-sv"`                                                       //0, WRITER_STRING_VALIDATION_NB-1, .unit = "sv//
	StringValidationReplacement int `json:"string_validation_replacement" flag:"-string_validation_replacement"` // "set string validation replacement string", OFFSET(string_validation_replacement), AV_OPT_TYPE_STRING, {.str=""}},
	Svr                         int `json:"svr" flag:"-svr"`                                                     // "set string validation replacement string", OFFSET(string_validation_replacement), AV_OPT_TYPE_STRING, {.str="\xEF\xBF\xBD"}},
}

type Sv = typex.Flags

const (
	Sv_ignore  = "ignore"
	Sv_replace = "replace"
	Sv_fail    = "fail"
)
