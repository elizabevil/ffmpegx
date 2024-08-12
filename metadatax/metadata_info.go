package metadatax

type StreamVideoOnly struct {
	Width              int    `json:"width"`
	Height             int    `json:"height"`
	CodedWidth         int    `json:"coded_width"`
	CodedHeight        int    `json:"coded_height"`
	ClosedCaptions     int    `json:"closed_captions"`
	FilmGrain          int    `json:"film_grain"`
	HasBFrames         int    `json:"has_b_frames"`
	SampleAspectRatio  string `json:"sample_aspect_ratio"`
	DisplayAspectRatio string `json:"display_aspect_ratio"`
	PixFmt             string `json:"pix_fmt"`
	Level              int    `json:"level"`

	ColorRange     string `json:"color_range"`
	ColorSpace     string `json:"color_space"`
	ColorTransfer  string `json:"color_transfer"`
	ColorPrimaries string `json:"color_primaries"`
	ChromaLocation string `json:"chroma_location"`
	FieldOrder     string `json:"field_order"`
	Refs           int    `json:"refs"`

	//H264ParseContext ?
	IsAvc            string `json:"is_avc"`
	NalLengthSize    string `json:"nal_length_size"`
	GotFirst         int    `json:"got_first"`
	PictureStructure int    `json:"picture_structure"`
	ParsedExtradata  int    `json:"parsed_extradata"`
	Poc              int    `json:"poc"`
	PocTid0          int    `json:"pocTid0"`

	ColorTrc string `json:"color_trc"`
}

type StreamAudioOnly struct {
	SampleFmt      string `json:"sample_fmt"`
	SampleRate     string `json:"sample_rate"`
	Channels       int    `json:"channels"`
	ChannelLayout  string `json:"channel_layout"`
	BitsPerSample  int    `json:"bits_per_sample"`
	InitialPadding int    `json:"initial_padding"`
}

type BaseStream struct {
	//fmt_ctx->iformat->flags & AVFMT_SHOW_IDS)

	Index          int    `json:"index"`
	CodecName      string `json:"codec_name"`
	CodecLongName  string `json:"codec_long_name"`
	Profile        string `json:"profile"`
	CodecType      string `json:"codec_type"`
	CodecTagString string `json:"codec_tag_string"`
	CodecTag       string `json:"codec_tag"`

	Id           string `json:"id"`
	RFrameRate   string `json:"r_frame_rate"`
	AvgFrameRate string `json:"avg_frame_rate"`
	TimeBase     string `json:"time_base"`
	StartPts     int    `json:"start_pts"`
	StartTime    string `json:"start_time"`
	DurationTs   int    `json:"duration_ts"`
	Duration     string `json:"duration"` // duration_ts *  time_base

	BitRate          string `json:"bit_rate"`
	BitsPerRawSample string `json:"bits_per_raw_sample"`
	NbFrames         string `json:"nb_frames"`
	ExtradataSize    int    `json:"extradata_size"`
	ExtradataHash    string `json:"extradata_hash"`

	MaxBitRate    int    `json:"max_bit_rate"`
	NbReadFrames  string `json:"nb_read_frames"`
	NbReadPackets string `json:"nb_read_packets"`
	Extradata     string `json:"extradata"`
	PtsWrapBits   int    `json:"pts_wrap_bits"` //

	Disposition Tags `json:"disposition,omitempty"`
	Tags        Tags `json:"tags,omitempty"`
}

type StreamSubtitleOnly struct {
}

type Disposition struct {
	Default         int `json:"default"`
	Dub             int `json:"dub"`
	Original        int `json:"original"`
	Comment         int `json:"comment"`
	Lyrics          int `json:"lyrics"`
	Karaoke         int `json:"karaoke"`
	Forced          int `json:"forced"`
	HearingImpaired int `json:"hearing_impaired"`
	VisualImpaired  int `json:"visual_impaired"`
	CleanEffects    int `json:"clean_effects"`
	AttachedPic     int `json:"attached_pic"`
	TimedThumbnails int `json:"timed_thumbnails"`
	NonDiegetic     int `json:"non_diegetic"`
	Captions        int `json:"captions"`
	Descriptions    int `json:"descriptions"`
	Metadata        int `json:"metadata"`
	Dependent       int `json:"dependent"`
	StillImage      int `json:"still_image"`
}

type Format struct {
	Filename       string `json:"filename"`
	NbStreams      int    `json:"nb_streams"`
	NbPrograms     int    `json:"nb_programs"`
	FormatName     string `json:"format_name"`
	FormatLongName string `json:"format_long_name"`
	StartTime      string `json:"start_time"`
	Duration       string `json:"duration"`
	Size           string `json:"size"`
	BitRate        string `json:"bit_rate"`
	ProbeScore     int    `json:"probe_score"`

	Tags Tags `json:"tags,omitempty"`
}

func (s StreamAudioOnly) GetChannelLayout() ChannelLayoutName {
	return ChannelLayoutMap[s.ChannelLayout]
}
