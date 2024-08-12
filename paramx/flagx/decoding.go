package flagx

type ErrDetect = Flag

const (
	ErrDetect_crccheck   ErrDetect = "crccheck"
	ErrDetect_bitstream  ErrDetect = "bitstream"
	ErrDetect_buffer     ErrDetect = "buffer"
	ErrDetect_explode    ErrDetect = "explode"
	ErrDetect_ignore_err ErrDetect = "ignore_err"
	ErrDetect_careful    ErrDetect = "careful"
	ErrDetect_compliant  ErrDetect = "compliant"
	ErrDetect_aggressive ErrDetect = "aggressive"
)

type MovFlag = Flag

const (
	Movflag_cmaf                 MovFlag = "cmaf"
	Movflag_dash                 MovFlag = "dash"
	Movflag_default_base_moof    MovFlag = "default_base_moof"
	Movflag_delay_moov           MovFlag = "delay_moov"
	Movflag_disable_chpl         MovFlag = "disable_chpl"
	Movflag_faststart            MovFlag = "faststart"
	Movflag_frag_custom          MovFlag = "frag_custom"
	Movflag_frag_discont         MovFlag = "frag_discont"
	Movflag_frag_every_frame     MovFlag = "frag_every_frame"
	Movflag_frag_keyframe        MovFlag = "frag_keyframe"
	Movflag_global_sidx          MovFlag = "global_sidx"
	Movflag_isml                 MovFlag = "isml"
	Movflag_negative_cts_offsets MovFlag = "negative_cts_offsets"
	Movflag_omit_tfhd_offset     MovFlag = "omit_tfhd_offset"
	Movflag_prefer_icc           MovFlag = "prefer_icc"
	Movflag_rtphint              MovFlag = "rtphint"
	Movflag_separate_moof        MovFlag = "separate_moof"
	Movflag_skip_sidx            MovFlag = "skip_sidx"
	Movflag_skip_trailer         MovFlag = "skip_trailer"
	Movflag_use_metadata_tags    MovFlag = "use_metadata_tags"
	Movflag_write_colr           MovFlag = "write_colr"
	Movflag_write_gama           MovFlag = "write_gama"
	Movflag_hybrid_fragmented    MovFlag = "hybrid_fragmented"
)

type AbortOn = Flag

const (
	AbortOnFlag_empty_output        AbortOn = "empty_output"
	AbortOnFlag_empty_output_stream AbortOn = "empty_output_stream"
)

type Bug = Flag

const (
	BugFlag_autodetect       Bug = "autodetect"
	BugFlag_xvid_ilace       Bug = "xvid_ilace"
	BugFlag_ump4             Bug = "ump4"
	BugFlag_no_padding       Bug = "no_padding"
	BugFlag_amv              Bug = "amv"
	BugFlag_qpel_chroma      Bug = "qpel_chroma"
	BugFlag_std_qpel         Bug = "std_qpel"
	BugFlag_qpel_chroma2     Bug = "qpel_chroma2"
	BugFlag_direct_blocksize Bug = "direct_blocksize"
	BugFlag_edge             Bug = "edge"
	BugFlag_hpel_chroma      Bug = "hpel_chroma"
	BugFlag_dc_clip          Bug = "dc_clip"
	BugFlag_ms               Bug = "ms"
	BugFlag_trunc            Bug = "trunc"
)

type AudioServiceType = Flag

const (
	AudioServiceType_ma AudioServiceType = "ma" //Main Audio Service

	AudioServiceType_ef AudioServiceType = "ef" //Effects

	AudioServiceType_vi AudioServiceType = "vi" //Visually Impaired

	AudioServiceType_hi AudioServiceType = "hi" //Hearing Impaired

	AudioServiceType_di AudioServiceType = "di" //Dialogue

	AudioServiceType_co AudioServiceType = "co" //Commentary

	AudioServiceType_em AudioServiceType = "em" //Emergency

	AudioServiceType_vo AudioServiceType = "vo" //Voice Over

	AudioServiceType_ka AudioServiceType = "ka" //Karaoke

)

type Idct = Flag

const (
	Idct_auto       Idct = "auto"
	Idct_int        Idct = "int"
	Idct_simple     Idct = "simple"
	Idct_simplemmx  Idct = "simplemmx"
	Idct_simpleauto Idct = "simpleauto"
	//Automatically pick a IDCT compatible with the simple one

	Idct_arm           Idct = "arm"
	Idct_altivec       Idct = "altivec"
	Idct_sh4           Idct = "sh4"
	Idct_simplearm     Idct = "simplearm"
	Idct_simplearmv5te Idct = "simplearmv5te"
	Idct_simplearmv6   Idct = "simplearmv6"
	Idct_simpleneon    Idct = "simpleneon"
	Idct_xvid          Idct = "xvid"
	Idct_faani         Idct = "faani"
	//floating point AAN IDCT

)

type Ec = Flag

const (
	Ec_guess_mvs Ec = "guess_mvs"
	//iterative motion vector (MV) search (slow)

	Ec_deblock Ec = "deblock"
	//use strong deblock filter for damaged MBs

	Ec_favor_inter Ec = "favor_inter"
	//favor predicting from the previous frame instead of the current
)
