package typex

type CodecType = Codec

// libavutil\utils.c
const (
	AVMEDIA_TYPE_UNKNOWN              = -1 ///< Usually treated as AVMEDIA_TYPE_DATA
	AVMEDIA_TYPE_VIDEO      CodecType = "video"
	AVMEDIA_TYPE_AUDIO      CodecType = "audio"
	AVMEDIA_TYPE_DATA       CodecType = "data" ///< Opaque data information usually continuous
	AVMEDIA_TYPE_SUBTITLE   CodecType = "subtitle"
	AVMEDIA_TYPE_ATTACHMENT CodecType = "attachment" ///< Opaque data information usually sparse
	AVMEDIA_TYPE_NB         CodecType = "nb"
)
