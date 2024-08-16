package typex

const (
	AV_OPT_TYPE_FLAGS = iota + 1
	AV_OPT_TYPE_INT
	AV_OPT_TYPE_INT64
	AV_OPT_TYPE_DOUBLE
	AV_OPT_TYPE_FLOAT
	AV_OPT_TYPE_STRING
	AV_OPT_TYPE_RATIONAL
	AV_OPT_TYPE_BINARY ///< offset must point to a pointer immediately followed by an int for the length
	AV_OPT_TYPE_DICT
	AV_OPT_TYPE_UINT64
	AV_OPT_TYPE_CONST
	AV_OPT_TYPE_IMAGE_SIZE ///< offset must point to two consecutive integers
	AV_OPT_TYPE_PIXEL_FMT
	AV_OPT_TYPE_SAMPLE_FMT
	AV_OPT_TYPE_VIDEO_RATE ///< offset must point to AVRational
	AV_OPT_TYPE_DURATION
	AV_OPT_TYPE_COLOR
	AV_OPT_TYPE_BOOL
	AV_OPT_TYPE_CHLAYOUT

	/**
	 * May be combined with another regular option type to declare an array
	 * option.
	 *
	 * For array options, @ref AVOption.offset should refer to a pointer
	 * corresponding to the option type. The pointer should be immediately
	 * followed by an unsigned int that will store the number of elements in the
	 * array.
	 */

	AV_OPT_TYPE_FLAG_ARRAY = (1 << 16)
)

const (
	/**
	 * A generic parameter which can be set by the user for muxing or encoding.
	 */
	AV_OPT_FLAG_ENCODING_PARAM = (1 << 0)
	/**
	 * A generic parameter which can be set by the user for demuxing or decoding.
	 */
	AV_OPT_FLAG_DECODING_PARAM = (1 << 1)
	AV_OPT_FLAG_AUDIO_PARAM    = (1 << 3)
	AV_OPT_FLAG_VIDEO_PARAM    = (1 << 4)
	AV_OPT_FLAG_SUBTITLE_PARAM = (1 << 5)
	/**
	 * The option is intended for exporting values to the caller.
	 */
	AV_OPT_FLAG_EXPORT = (1 << 6)
	/**
	 * The option may not be set through the AVOptions API, only read.
	 * This flag only makes sense when AV_OPT_FLAG_EXPORT is also set.
	 */
	AV_OPT_FLAG_READONLY = (1 << 7)
	/**
	 * A generic parameter which can be set by the user for bit stream filtering.
	 */
	AV_OPT_FLAG_BSF_PARAM = (1 << 8)

	/**
	 * A generic parameter which can be set by the user at runtime.
	 */
	AV_OPT_FLAG_RUNTIME_PARAM = (1 << 15)
	/**
	 * A generic parameter which can be set by the user for filtering.
	 */
	AV_OPT_FLAG_FILTERING_PARAM = (1 << 16)
	/**
	 * Set if option is deprecated, users should refer to AVOption.help text for
	 * more information.
	 */
	AV_OPT_FLAG_DEPRECATED = (1 << 17)
	/**
	 * Set if option constants can also reside in child objects.
	 */
	AV_OPT_FLAG_CHILD_CONSTS = (1 << 18)
)

//======

const (
	V                        = AV_OPT_FLAG_VIDEO_PARAM
	A                        = AV_OPT_FLAG_AUDIO_PARAM
	S                        = AV_OPT_FLAG_SUBTITLE_PARAM
	E                        = AV_OPT_FLAG_ENCODING_PARAM
	D                        = AV_OPT_FLAG_DECODING_PARAM
	CC                       = AV_OPT_FLAG_CHILD_CONSTS
	AR                       = AV_OPT_TYPE_FLAG_ARRAY
	AV_CODEC_DEFAULT_BITRATE = 200 * 1000
)
