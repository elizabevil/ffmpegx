package optionx

import "github.com/elizabevil/ffmpegx/paramx/typex"

type Encoding struct {
	Packetsize            typex.UNumber      `json:"packetsize" flagh:"-packetsize"`                           // "set packet size", OFFSET(packet_size), AV_OPT_TYPE_INT, {.i64 = DEFAULT }, 0, INT_MAX, E},
	StartTimeRealtime     typex.Number       `json:"start_time_realtime" flagh:"-start_time_realtime"`         // "wall-clock time when stream begins (PTS==0)", OFFSET(start_time_realtime), AV_OPT_TYPE_INT64, {.i64 = AV_NOPTS_VALUE}, INT64_MIN, INT64_MAX, E},
	AudioPreload          typex.MicrosecondI `json:"audio_preload" flagh:"-audio_preload"`                     // "microseconds by which audio packets should be interleaved earlier", OFFSET(audio_preload), AV_OPT_TYPE_INT, {.i64 = 0}, 0, INT_MAX-1, E},
	ChunkDuration         typex.MicrosecondI `json:"chunk_duration" flagh:"-chunk_duration"`                   // "microseconds for each chunk", OFFSET(max_chunk_duration), AV_OPT_TYPE_INT, {.i64 = 0}, 0, INT_MAX-1, E},
	ChunkSize             typex.Size         `json:"chunk_size" flagh:"-chunk_size"`                           // "size in bytes for each chunk", OFFSET(max_chunk_size), AV_OPT_TYPE_INT, {.i64 = 0}, 0, INT_MAX-1, E},
	FlushPackets          typex.Bool         `json:"flush_packets" flagh:"-flush_packets"`                     // "enable flushing of the I/O context after each packet", OFFSET(flush_packets), AV_OPT_TYPE_INT, {.i64 = -1}, -1, 1, E},
	MetadataHeaderPadding typex.UNumber      `json:"metadata_header_padding" flagh:"-metadata_header_padding"` // "set number of bytes to be written as padding in a metadata header", OFFSET(metadata_header_padding), AV_OPT_TYPE_INT, {.i64 = -1}, -1, INT_MAX, E},
	OutputTsOffset        typex.Number       `json:"output_ts_offset" flagh:"-output_ts_offset"`               // "set output timestamp offset", OFFSET(output_ts_offset), AV_OPT_TYPE_DURATION, {.i64 = 0}, -INT64_MAX, INT64_MAX, E},
	DumpSeparator         typex.String       `json:"dump_separator" flagh:"-dump_separator"`                   // "set information dump field separator", OFFSET(dump_separator), AV_OPT_TYPE_STRING, {.str = ", "}, 0, 0, D|E},
}
type Decoding struct {
	Probesize                   typex.Size         `json:"probesize" flag:"-probesize"`                                             // "set probing size", OFFSET(probesize), AV_OPT_TYPE_INT64, {.i64 = 5000000 }, 32, INT64_MAX, D},
	Formatprobesize             typex.Size         `json:"formatprobesize" flag:"-formatprobesize"`                                 // "number of bytes to probe file format", OFFSET(format_probesize), AV_OPT_TYPE_INT, {.i64 = PROBE_BUF_MAX}, 0, INT_MAX-1, D},
	Seek2Any                    typex.UBool        `json:"seek2any" flag:"-seek2any"`                                               // "allow seeking to non-keyframes on demuxer level when supported", OFFSET(seek2any), AV_OPT_TYPE_BOOL, {.i64 = 0 }, 0, 1, D},
	Analyzeduration             typex.MicrosecondI `json:"analyzeduration" flag:"-analyzeduration"`                                 // "specify how many microseconds are analyzed to probe the input", OFFSET(max_analyze_duration), AV_OPT_TYPE_INT64, {.i64 = 0 }, 0, INT64_MAX, D},
	Cryptokey                   typex.Binary       `json:"cryptokey" flag:"-cryptokey"`                                             // "decryption key", OFFSET(key), AV_OPT_TYPE_BINARY, {.dbl = 0}, 0, 0, D},
	Indexmem                    typex.UNumber      `json:"indexmem" flag:"-indexmem"`                                               // "max memory used for timestamp index (per stream)", OFFSET(max_index_size), AV_OPT_TYPE_INT, {.i64 = 1<<20 }, 0, INT_MAX, D},
	Rtbufsize                   typex.UNumber      `json:"rtbufsize" flag:"-rtbufsize"`                                             // "max memory used for buffering real-time frames", OFFSET(max_picture_buffer), AV_OPT_TYPE_INT, {.i64 = 3041280 }, 0, INT_MAX, D}, /* defaults to 1s of 15fps 352x288 YUYV422 video */
	MaxDelay                    typex.UNumber      `json:"max_delay" flag:"-max_delay"`                                             // "maximum muxing or demuxing delay in microseconds", OFFSET(max_delay), AV_OPT_TYPE_INT, {.i64 = -1 }, -1, INT_MAX, E|D},
	Fpsprobesize                typex.Number       `json:"fpsprobesize" flag:"-fpsprobesize"`                                       // "number of frames used to probe fps", OFFSET(fps_probe_size), AV_OPT_TYPE_INT, {.i64 = -1}, -1, INT_MAX-1, D},
	UseWallclockAsTimestamps    typex.UBool        `json:"use_wallclock_as_timestamps" flag:"-use_wallclock_as_timestamps"`         // "use wallclock as timestamps", OFFSET(use_wallclock_as_timestamps), AV_OPT_TYPE_BOOL, {.i64 = 0}, 0, 1, D},
	SkipInitialBytes            typex.Number       `json:"skip_initial_bytes" flag:"-skip_initial_bytes"`                           // "set number of bytes to skip before reading header and frames", OFFSET(skip_initial_bytes), AV_OPT_TYPE_INT64, {.i64 = 0}, 0, INT64_MAX-1, D},
	CorrectTsOverflow           typex.UBool        `json:"correct_ts_overflow" flag:"-correct_ts_overflow"`                         // "correct single timestamp overflows", OFFSET(correct_ts_overflow), AV_OPT_TYPE_BOOL, {.i64 = 1}, 0, 1, D},
	SkipEstimateDurationFromPts typex.UBool        `json:"skip_estimate_duration_from_pts" flag:"-skip_estimate_duration_from_pts"` // "skip duration calculation in estimate_timings_from_pts", OFFSET(skip_estimate_duration_from_pts), AV_OPT_TYPE_BOOL, {.i64 = 0}, 0, 1, D},
	MaxTsProbe                  typex.UNumber      `json:"max_ts_probe" flag:"-max_ts_probe"`                                       // "maximum number of packets to read while waiting for the first timestamp", OFFSET(max_ts_probe), AV_OPT_TYPE_INT, { .i64 = 50 }, 0, INT_MAX, D },
	CodecWhitelist              typex.String       `json:"codec_whitelist" flag:"-codec_whitelist"`                                 // "List of decoders that are allowed to be used", OFFSET(codec_whitelist), AV_OPT_TYPE_STRING, { .str = NULL },  0, 0, D },
	FormatWhitelist             typex.String       `json:"format_whitelist" flag:"-format_whitelist"`                               // "List of demuxers that are allowed to be used", OFFSET(format_whitelist), AV_OPT_TYPE_STRING, { .str = NULL },  0, 0, D },
	ProtocolWhitelist           typex.String       `json:"protocol_whitelist" flag:"-protocol_whitelist"`                           // "List of protocols that are allowed to be used", OFFSET(protocol_whitelist), AV_OPT_TYPE_STRING, { .str = NULL },  0, 0, D },
	ProtocolBlacklist           typex.String       `json:"protocol_blacklist" flag:"-protocol_blacklist"`                           // "List of protocols that are not allowed to be used", OFFSET(protocol_blacklist), AV_OPT_TYPE_STRING, { .str = NULL },  0, 0, D },
	MaxStreams                  typex.UNumber      `json:"max_streams" flag:"-max_streams"`                                         // "maximum number of streams", OFFSET(max_streams), AV_OPT_TYPE_INT, { .i64 = 1000 }, 0, INT_MAX, D },
	MaxProbePackets             typex.UNumber      `json:"max_probe_packets" flag:"-max_probe_packets"`                             // "Maximum number of packets to probe a codec", OFFSET(max_probe_packets), AV_OPT_TYPE_INT, { .i64 = 2500 }, 0, INT_MAX, D },
}

type Coding struct {
	Fdebug   int `UNumber:"fdebug" flag:"-fdebug"`    // "print specific debug info", OFFSET(debug), AV_OPT_TYPE_FLAGS, {.i64 = DEFAULT }, 0, INT_MAX, E|D, .unit = "fdebug"},
	Ts       int `json:"ts" flag:"-ts"`               // NULL, 0, AV_OPT_TYPE_CONST, {.i64 = FF_FDEBUG_TS }, INT_MIN, INT_MAX, E|D, .unit = "fdebug"},
	MaxDelay int `json:"max_delay" flag:"-max_delay"` // "maximum muxing or demuxing delay in microseconds", OFFSET(max_delay), AV_OPT_TYPE_INT, {.i64 = -1 }, -1, INT_MAX, E|D},
}

type Codec struct {
	Fdebug   int `UNumber:"fdebug" flag:"-fdebug"`    // "print specific debug info", OFFSET(debug), AV_OPT_TYPE_FLAGS, {.i64 = DEFAULT }, 0, INT_MAX, E|D, .unit = "fdebug"},
	Ts       int `json:"ts" flag:"-ts"`               // NULL, 0, AV_OPT_TYPE_CONST, {.i64 = FF_FDEBUG_TS }, INT_MIN, INT_MAX, E|D, .unit = "fdebug"},
	MaxDelay int `json:"max_delay" flag:"-max_delay"` // "maximum muxing or demuxing delay in microseconds", OFFSET(max_delay), AV_OPT_TYPE_INT, {.i64 = -1 }, -1, INT_MAX, E|D},
}
