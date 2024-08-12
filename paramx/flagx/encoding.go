package flagx

type ThreadType = Flag

const (
	ThreadType_slice ThreadType = "slice"
	ThreadType_frame ThreadType = "frame"
)

type Debug = Flag

const (
	Debug_slice          Debug = "slice"
	Debug_pict           Debug = "pict"
	Debug_rc             Debug = "rc"
	Debug_bitstream      Debug = "bitstream"
	Debug_mb_type        Debug = "mb_type"
	Debug_qp             Debug = "qp"
	Debug_dct_coeff      Debug = "dct_coeff"
	Debug_green_metadata Debug = "green_metadata"
	Debug_skip           Debug = "skip"
	Debug_startcode      Debug = "startcode"
	Debug_er             Debug = "er"
	Debug_mmco           Debug = "mmco"
	Debug_bugs           Debug = "bugs"
	Debug_buffers        Debug = "buffers"
	Debug_thread_ops     Debug = "thread_ops"
	Debug_nomc           Debug = "nomc"
)

type Flags2 = Flag

const (
	Flags2_fast Flags2 = "fast"

	Flags2_noout Flags2 = "noout"

	Flags2_ignorecrop Flags2 = "ignorecrop"

	Flags2_local_header Flags2 = "local_header"

	Flags2_chunks Flags2 = "chunks"

	Flags2_showall Flags2 = "showall"

	Flags2_export_mvs Flags2 = "export_mvs"

	Flags2_skip_manual Flags2 = "skip_manual"

	Flags2_ass_ro_flush_noop Flags2 = "ass_ro_flush_noop"

	Flags2_icc_profiles Flags2 = "icc_profiles"
)

type ExportSideData = Flag

const (
	ExportSideData_mvs ExportSideData = "mvs"

	ExportSideData_prft ExportSideData = "prft"

	ExportSideData_venc_params ExportSideData = "venc_params"

	ExportSideData_film_grain ExportSideData = "film_grain"
)

type Cmp = Flag

const (
	Cmp_sad Cmp = "sad"
	//sum of absolute differences, fast (default)

	Cmp_sse Cmp = "sse"
	//sum of squared errors

	Cmp_satd Cmp = "satd"
	//sum of absolute Hadamard transformed differences

	Cmp_dct Cmp = "dct"
	//sum of absolute DCT transformed differences

	Cmp_psnr Cmp = "psnr"
	//sum of squared quantization errors (avoid, low quality)

	Cmp_bit Cmp = "bit"
	//number of bits needed for the block

	Cmp_rd Cmp = "rd"
	//rate distortion optimal, slow

	Cmp_zero Cmp = "zero"
	//0

	Cmp_vsad Cmp = "vsad"
	//sum of absolute vertical differences

	Cmp_vsse Cmp = "vsse"
	//sum of squared vertical differences

	Cmp_nsse Cmp = "nsse"
	//noise preserving sum of squared differences

	Cmp_w53 Cmp = "w53"
	//5/3 wavelet, only used in snow

	Cmp_w97 Cmp = "w97"
	//9/7 wavelet, only used in snow

	Cmp_dctmax Cmp = "dctmax"
	Cmp_chroma Cmp = "chroma"
)

type Subcmp = Cmp
type Mbcmp = Cmp
type Ildctcmp = Cmp
type Precmp = Cmp

type Mbd = Flag

const (
	MBD_simple Mbd = "simple"
	//use mbcmp (default)

	MBD_bits Mbd = "bits"
	//use fewest bits

	MBD_rd Mbd = "rd"
	//use best rate distortion

)

type ColorPrimaries = Flag

const (
	ColorPrimaries_bt709 ColorPrimaries = "bt709" //BT.709

	ColorPrimaries_bt470m ColorPrimaries = "bt470m" //BT.470 M

	ColorPrimaries_bt470bg ColorPrimaries = "bt470bg" //BT.470 BG

	ColorPrimaries_smpte170m ColorPrimaries = "smpte170m" //SMPTE 170 M

	ColorPrimaries_smpte240m ColorPrimaries = "smpte240m" //SMPTE 240 M

	ColorPrimaries_film ColorPrimaries = "film" //Film

	ColorPrimaries_bt2020 ColorPrimaries = "bt2020" //BT.2020

	ColorPrimaries_smpte428   ColorPrimaries = "smpte428"
	ColorPrimaries_smpte428_1 ColorPrimaries = "smpte428_1" //SMPTE ST 428-1
	ColorPrimaries_smpte431   ColorPrimaries = "smpte431"   //SMPTE 431-2

	ColorPrimaries_smpte432 ColorPrimaries = "smpte432" //SMPTE 432-1

	ColorPrimaries_jedec_p22 ColorPrimaries = "jedec-p22" //JEDEC P22
)

type ColorTrc = Flag

const (
	ColorTrc_bt709 ColorTrc = "bt709" //BT.709

	ColorTrc_gamma22 ColorTrc = "gamma22" //BT.470 M

	ColorTrc_gamma28 ColorTrc = "gamma28" //BT.470 BG

	ColorTrc_smpte170m ColorTrc = "smpte170m" //	SMPTE 170 M

	ColorTrc_smpte240m ColorTrc = "smpte240m" //SMPTE 240 M

	ColorTrc_linear ColorTrc = "linear" //Linear

	ColorTrc_log    ColorTrc = "log"
	ColorTrc_log100 ColorTrc = "log100" //Log

	ColorTrc_log_sqrt ColorTrc = "log_sqrt"
	ColorTrc_log316   ColorTrc = "log316" //Log square root

	ColorTrc_iec61966_2_4 ColorTrc = "iec61966_2_4"

	ColorTrc_bt1361  ColorTrc = "bt1361"
	ColorTrc_bt1361e ColorTrc = "bt1361e" //BT.1361

	ColorTrc_iec61966_2_1 ColorTrc = "iec61966_2_1"

	ColorTrc_bt2020_10    ColorTrc = "bt2020_10"
	ColorTrc_bt2020_10bit ColorTrc = "bt2020_10bit" //BT.2020 - 10 bit

	ColorTrc_bt2020_12    ColorTrc = "bt2020_12"
	ColorTrc_bt2020_12bit ColorTrc = "bt2020_12bit" //BT.2020 - 12 bit

	ColorTrc_smpte2084 ColorTrc = "smpte2084" //SMPTE ST 2084

	ColorTrc_smpte428   ColorTrc = "smpte428"
	ColorTrc_smpte428_1 ColorTrc = "smpte428_1" //SMPTE ST 428-1

	ColorTrc_arib_std_b_67 ColorTrc = "arib-std-b67" //ARIB STD-B67
)

type Colorspace = Flag

const (
	Colorspace_rgb Colorspace = "rgb" //RGB

	Colorspace_bt_709 Colorspace = "bt709" //BT.709

	Colorspace_fcc Colorspace = "fcc"
	FCC

	Colorspace_bt_470_bg Colorspace = "bt470bg" //BT.470 BG

	Colorspace_smpte_170_m Colorspace = "smpte170m" //SMPTE 170 M

	Colorspace_smpte_240_m Colorspace = "smpte240m" //SMPTE 240 M

	Colorspace_ycocg Colorspace = "ycocg" //YCOCG

	Colorspace_bt_2020_nc Colorspace = "bt2020nc"
	Colorspace_bt2020_ncl Colorspace = "bt2020_ncl" //BT.2020 NCL

	Colorspace_bt_2020_c Colorspace = "bt2020c"
	Colorspace_bt2020_cl Colorspace = "bt2020_cl" //BT.2020 CL

	Colorspace_smpte_2085 Colorspace = "smpte2085" //SMPTE 2085

	Colorspace_chroma_derived_nc Colorspace = "chroma-derived-nc" //Chroma-derived NCL

	Colorspace_chroma_derived_c Colorspace = "chroma-derived-c" //Chroma-derived CL

	Colorspace_ictcp Colorspace = "ictcp" //ICtCp

)

type ColorRange = Flag

const (
	ColorRange_tv      ColorRange = "tv"
	ColorRange_mpeg    ColorRange = "mpeg"
	ColorRange_limited ColorRange = "limited" //MPEG (219*2^(n-8))

	ColorRange_pc   ColorRange = "pc"
	ColorRange_jpeg ColorRange = "jpeg"
	ColorRange_full ColorRange = "full" //JPEG (2^n-1)
)

type ChromaSampleLocationType = Flag

const (
	ChromaSampleLocationType_left       ChromaSampleLocationType = "left"
	ChromaSampleLocationType_center     ChromaSampleLocationType = "center"
	ChromaSampleLocationType_topleft    ChromaSampleLocationType = "topleft"
	ChromaSampleLocationType_top        ChromaSampleLocationType = "top"
	ChromaSampleLocationType_bottomleft ChromaSampleLocationType = "bottomleft"
	ChromaSampleLocationType_bottom     ChromaSampleLocationType = "bottom"
)

type FieldOrder = Flag

const (
	FieldOrder_progressive FieldOrder = "progressive" //Progressive video

	FieldOrder_tt FieldOrder = "tt" //Interlaced video, top field coded and displayed first

	FieldOrder_bb FieldOrder = "bb" //Interlaced video, bottom field coded and displayed first

	FieldOrder_tb FieldOrder = "tb" //Interlaced video, top coded first, bottom displayed first

	FieldOrder_bt FieldOrder = "bt" //Interlaced video, bottom coded first, top displayed first
)
