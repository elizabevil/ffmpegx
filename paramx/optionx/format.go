package optionx

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// Format Options
type Format struct {
	MaxTsProbe        typex.UNumber `json:"max_ts_probe" flag:"-max_ts_probe"`             // "maximum number of packets to read while waiting for the first timestamp", OFFSET(max_ts_probe), AV_OPT_TYPE_INT, { .i64 = 50 }, 0, INT_MAX, D },
	ProtocolBlacklist typex.String  `json:"protocol_blacklist" flag:"-protocol_blacklist"` // "List of protocols that are not allowed to be used", OFFSET(protocol_blacklist), AV_OPT_TYPE_STRING, { .str = NULL },  0, 0, D },
	//=========
	Formatprobesize typex.Size `json:"formatprobesize" flag:"-formatprobesize"` // "number of bytes to probe file format", OFFSET(format_probesize), AV_OPT_TYPE_INT, {.i64 = PROBE_BUF_MAX}, 0, INT_MAX-1, D},
	//======
	MaxDelay  typex.UNumber   `json:"max_delay" flag:"-max_delay"` // "maximum muxing or demuxing delay in microseconds", OFFSET(max_delay), AV_OPT_TYPE_INT, {.i64 = -1 }, -1, INT_MAX, E|D},
	Avioflags flagx.Avioflags `json:"avioflags" flag:"-avioflags"` // NULL, OFFSET(avio_flags), AV_OPT_TYPE_FLAGS, {.i64 = DEFAULT }, INT_MIN, INT_MAX, D|E, .unit = "avioflags"},
	Fflags    flagx.Fflags    `json:"fflags" flag:"-fflags"`       // NULL, OFFSET(flags), AV_OPT_TYPE_FLAGS, {.i64 = AVFMT_FLAG_AUTO_BSF }, INT_MIN, INT_MAX, D|E, .unit = "fflags"},
	Fdebug    Fdebug          `json:"fdebug" flag:"-fdebug"`       // "print specific debug info", OFFSET(debug), AV_OPT_TYPE_FLAGS, {.i64 = DEFAULT }, 0, INT_MAX, E|D, .unit = "fdebug"},

	StartTimeRealtime     typex.Number  `json:"start_time_realtime" flag:"-start_time_realtime"`         // "wall-clock time when stream begins (PTS==0)", OFFSET(start_time_realtime), AV_OPT_TYPE_INT64, {.i64 = AV_NOPTS_VALUE}, INT64_MIN, INT64_MAX, E},
	MetadataHeaderPadding typex.UNumber `json:"metadata_header_padding" flag:"-metadata_header_padding"` // "set number of bytes to be written as padding in a metadata header", OFFSET(metadata_header_padding), AV_OPT_TYPE_INT, {.i64 = -1}, -1, INT_MAX, E},
	FStrict               flagx.Strict  `json:"f_strict" flag:"-f_strict"`                               // "how strictly to follow the standards (deprecated; use strict, save via avconv)", OFFSET(strict_std_compliance), AV_OPT_TYPE_INT, {.i64 = DEFAULT }, INT_MIN, INT_MAX, D|E, .unit = "strict"},
	Strict                flagx.Strict  `json:"strict" flag:"-strict"`                                   // "how strictly to follow the standards", OFFSET(strict_std_compliance), AV_OPT_TYPE_INT, {.i64 = DEFAULT }, INT_MIN, INT_MAX, D|E, .unit = "strict"},
}

// FormatInput
type FormatInput struct {
	Packetsize                  typex.Size         `json:"packetsize" flag:"-packetsize"`                                           // "set packet size", OFFSET(packet_size), AV_OPT_TYPE_INT, {.i64 = DEFAULT }, 0, INT_MAX, E},
	Probesize                   typex.Size         `json:"probesize" flag:"-probesize"`                                             // "set probing size", OFFSET(probesize), AV_OPT_TYPE_INT64, {.i64 = 5000000 }, 32, INT64_MAX, D},
	MaxProbePackets             typex.Size         `json:"max_probe_packets" flag:"-max_probe_packets"`                             // "Maximum number of packets to probe a codec", OFFSET(max_probe_packets), AV_OPT_TYPE_INT, { .i64 = 2500 }, 0, INT_MAX, D },
	Analyzeduration             typex.MicrosecondI `json:"analyzeduration" flag:"-analyzeduration"`                                 // "specify how many microseconds are analyzed to probe the input", OFFSET(max_analyze_duration), AV_OPT_TYPE_INT64, {.i64 = 0 }, 0, INT64_MAX, D},
	Seek2Any                    typex.UBool        `json:"seek2any" flag:"-seek2any"`                                               // "allow seeking to non-keyframes on demuxer level when supported", OFFSET(seek2any), AV_OPT_TYPE_BOOL, {.i64 = 0 }, 0, 1, D},
	Cryptokey                   typex.Key          `json:"cryptokey" flag:"-cryptokey"`                                             // "decryption key", OFFSET(key), AV_OPT_TYPE_BINARY, {.dbl = 0}, 0, 0, D},
	Indexmem                    typex.Index        `json:"indexmem" flag:"-indexmem"`                                               // "max memory used for timestamp index (per stream)", OFFSET(max_index_size), AV_OPT_TYPE_INT, {.i64 = 1<<20 }, 0, INT_MAX, D},
	Rtbufsize                   typex.Size         `json:"rtbufsize" flag:"-rtbufsize"`                                             // "max memory used for buffering real-time frames", OFFSET(max_picture_buffer), AV_OPT_TYPE_INT, {.i64 = 3041280 }, 0, INT_MAX, D}, /* defaults to 1s of 15fps 352x288 YUYV422 video */
	Fpsprobesize                typex.Size         `json:"fpsprobesize" flag:"-fpsprobesize"`                                       // "number of frames used to probe fps", OFFSET(fps_probe_size), AV_OPT_TYPE_INT, {.i64 = -1}, -1, INT_MAX-1, D},
	FErrDetect                  flagx.ErrDetect    `json:"f_err_detect" flag:"-f_err_detect"`                                       // "set error detection flags (deprecated; use err_detect, save via avconv)", OFFSET(error_recognition), AV_OPT_TYPE_FLAGS, {.i64 = AV_EF_CRCCHECK }, INT_MIN, INT_MAX, D, .unit = "err_detect"},
	ErrDetect                   flagx.ErrDetect    `json:"err_detect" flag:"-err_detect"`                                           // "set error detection flags", OFFSET(error_recognition), AV_OPT_TYPE_FLAGS, {.i64 = AV_EF_CRCCHECK }, INT_MIN, INT_MAX, D, .unit = "err_detect"},
	SkipInitialBytes            typex.UNumber      `json:"skip_initial_bytes" flag:"-skip_initial_bytes"`                           // "set number of bytes to skip before reading header and frames", OFFSET(skip_initial_bytes), AV_OPT_TYPE_INT64, {.i64 = 0}, 0, INT64_MAX-1, D},
	UseWallclockAsTimestamps    typex.UBool        `json:"use_wallclock_as_timestamps" flag:"-use_wallclock_as_timestamps"`         // "use wallclock as timestamps", OFFSET(use_wallclock_as_timestamps), AV_OPT_TYPE_BOOL, {.i64 = 0}, 0, 1, D},
	FormatWhitelist             typex.List         `json:"format_whitelist" flag:"-format_whitelist"`                               // "List of demuxers that are allowed to be used", OFFSET(format_whitelist), AV_OPT_TYPE_STRING, { .str = NULL },  0, 0, D },
	CorrectTsOverflow           typex.UBool        `json:"correct_ts_overflow" flag:"-correct_ts_overflow"`                         // "correct single timestamp overflows", OFFSET(correct_ts_overflow), AV_OPT_TYPE_BOOL, {.i64 = 1}, 0, 1, D},
	DumpSeparator               typex.Separator    `json:"dump_separator" flag:"-dump_separator"`                                   // "set information dump field separator", OFFSET(dump_separator), AV_OPT_TYPE_STRING, {.str = ", "}, 0, 0, D|E},
	MaxStreams                  typex.UNumber      `json:"max_streams" flag:"-max_streams"`                                         // "maximum number of streams", OFFSET(max_streams), AV_OPT_TYPE_INT, { .i64 = 1000 }, 0, INT_MAX, D },
	SkipEstimateDurationFromPts typex.UBool        `json:"skip_estimate_duration_from_pts" flag:"-skip_estimate_duration_from_pts"` // "skip duration calculation in estimate_timings_from_pts", OFFSET(skip_estimate_duration_from_pts), AV_OPT_TYPE_BOOL, {.i64 = 0}, 0, 1, D},
	DurationProbesize           typex.UBool        `json:"duration_probesize" flag:"-duration_probesize"`                           // "skip duration calculation in estimate_timings_from_pts", OFFSET(skip_estimate_duration_from_pts), AV_OPT_TYPE_BOOL, {.i64 = 0}, 0, 1, D},
	CodecWhitelist              typex.List         `json:"codec_whitelist" flag:"-codec_whitelist"`                                 // "List of decoders that are allowed to be used", OFFSET(codec_whitelist), AV_OPT_TYPE_STRING, { .str = NULL },  0, 0, D },
}

// FormatOutput
type FormatOutput struct {
	AudioPreload       typex.MicrosecondI `json:"audio_preload" flag:"-audio_preload"`               // "microseconds by which audio packets should be interleaved earlier", OFFSET(audio_preload), AV_OPT_TYPE_INT, {.i64 = 0}, 0, INT_MAX-1, E},
	ChunkDuration      typex.MicrosecondI `json:"chunk_duration" flag:"-chunk_duration"`             // "microseconds for each chunk", OFFSET(max_chunk_duration), AV_OPT_TYPE_INT, {.i64 = 0}, 0, INT_MAX-1, E},
	ChunkSize          typex.Size         `json:"chunk_size" flag:"-chunk_size"`                     // "size in bytes for each chunk", OFFSET(max_chunk_size), AV_OPT_TYPE_INT, {.i64 = 0}, 0, INT_MAX-1, E},
	MaxInterleaveDelta typex.UNumber      `json:"max_interleave_delta" flag:"-max_interleave_delta"` // "maximum buffering duration for interleaving", OFFSET(max_interleave_delta), AV_OPT_TYPE_INT64, { .i64 = 10000000 }, 0, INT64_MAX, E },
	AvoidNegativeTs    AvoidNegativeTs    `json:"avoid_negative_ts" flag:"-avoid_negative_ts"`       // "shift timestamps so they start at 0", OFFSET(avoid_negative_ts), AV_OPT_TYPE_INT, {.i64 = -1}, -1, 2, E, .unit = "avoid_negative_ts"},
	FlushPackets       typex.Bool         `json:"flush_packets" flag:"-flush_packets"`               // "enable flushing of the I/O context after each packet", OFFSET(flush_packets), AV_OPT_TYPE_INT, {.i64 = -1}, -1, 1, E},
	OutputTsOffset     typex.Number       `json:"output_ts_offset" flag:"-output_ts_offset"`         // "set output timestamp offset", OFFSET(output_ts_offset), AV_OPT_TYPE_DURATION, {.i64 = 0}, -INT64_MAX, INT64_MAX, E},
}
