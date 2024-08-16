package metadatax

type PacketAndFrame struct {
	Type         string `json:"type"`
	CodecType    string `json:"codec_type"`
	StreamIndex  int    `json:"stream_index"`
	Pts          int    `json:"pts"`
	PtsTime      string `json:"pts_time"`
	Dts          int    `json:"dts"`
	DtsTime      string `json:"dts_time"`
	Duration     int    `json:"duration"`
	DurationTime string `json:"duration_time"`
	Size         string `json:"size"`
	Pos          string `json:"pos"`
	Flags        string `json:"flags"`
	Data         string `json:"data"`
}
type Packet struct {
	CodecType    string `json:"codec_type"`
	StreamIndex  int    `json:"stream_index"`
	Pts          int    `json:"pts"`
	PtsTime      string `json:"pts_time"`
	Dts          int    `json:"dts"`
	DtsTime      string `json:"dts_time"`
	Duration     int    `json:"duration"`
	DurationTime string `json:"duration_time"`
	Size         string `json:"size"`
	Pos          string `json:"pos"`
	Flags        string `json:"flags"`
}
type Frame struct {
	MediaType               string `json:"media_type"`
	StreamIndex             int    `json:"stream_index"`
	KeyFrame                int    `json:"key_frame"`
	Pts                     int    `json:"pts"`
	PtsTime                 string `json:"pts_time"`
	PktDts                  int    `json:"pkt_dts"`
	PktDtsTime              string `json:"pkt_dts_time"`
	BestEffortTimestamp     int    `json:"best_effort_timestamp"`
	BestEffortTimestampTime string `json:"best_effort_timestamp_time"`
	PktDuration             int    `json:"pkt_duration"`
	PktDurationTime         string `json:"pkt_duration_time"`
	PktPos                  string `json:"pkt_pos"`
	PktSize                 string `json:"pkt_size"`
	SampleFmt               string `json:"sample_fmt"`
	NbSamples               int    `json:"nb_samples"`
	Channels                int    `json:"channels"`
	ChannelLayout           string `json:"channel_layout"`
}

type PixelFormat struct {
	Name         string      `json:"name"`
	NbComponents int         `json:"nb_components"`
	Log2ChromaW  int         `json:"log2_chroma_w"`
	Log2ChromaH  int         `json:"log2_chroma_h"`
	BitsPerPixel int         `json:"bits_per_pixel"`
	Flags        Flag        `json:"flags,omitempty"`
	Components   []Component `json:"components,omitempty"`
}

type Component struct {
	Index    int `json:"index"`
	BitDepth int `json:"bit_depth"`
}

type Flag struct {
	BigEndian int `json:"big_endian"`
	Palette   int `json:"palette"`
	Bitstream int `json:"bitstream"`
	Hwaccel   int `json:"hwaccel"`
	Planar    int `json:"planar"`
	Rgb       int `json:"rgb"`
	Alpha     int `json:"alpha"`
}
type LibraryVersion struct {
	Name    string `json:"name"`
	Major   int    `json:"major"`
	Minor   int    `json:"minor"`
	Micro   int    `json:"micro"`
	Version int    `json:"version"`
	Ident   string `json:"ident"`
}
type ProgramVersion struct {
	Version       string `json:"version"`
	Copyright     string `json:"copyright"`
	CompilerIdent string `json:"compiler_ident"`
	Configuration string `json:"configuration"`
}
type Chapter struct {
}

type Program struct {
}
