package flagx

import "github.com/elizabevil/ffmpegx/paramx/typex"

//flags flags (decoding/encoding,audio,video,subtitles)

type Flag = typex.String

type Flags = Flag

const (
	Flags_mv4            Flags = "mv4"
	Flags_qpel           Flags = "qpel"
	Flags_loop           Flags = "loop"
	Flags_qscale         Flags = "qscale"
	Flags_pass1          Flags = "pass1"
	Flags_pass2          Flags = "pass2"
	Flags_gray           Flags = "gray"
	Flags_psnr           Flags = "psnr"
	Flags_truncated      Flags = "truncated"
	Flags_drop_changed   Flags = "drop_changed"
	Flags_ildct          Flags = "ildct"
	Flags_low_delay      Flags = "low_delay"
	Flags_global_header  Flags = "global_header"
	Flags_bitexact       Flags = "bitexact"
	Flags_aic            Flags = "aic"
	Flags_ilme           Flags = "ilme"
	Flags_cgop           Flags = "cgop"
	Flags_output_corrupt Flags = "output_corrupt"
)

type ELibxvidFlag = Flag

const (
	ELibxvidFlag_mv4           ELibxvidFlag = Flags_mv4
	ELibxvidFlag_aic           ELibxvidFlag = Flags_aic
	ELibxvidFlag_gray          ELibxvidFlag = Flags_gray
	ELibxvidFlag_qpel          ELibxvidFlag = Flags_qpel
	ELibxvidFlag_cgop          ELibxvidFlag = Flags_cgop
	ELibxvidFlag_global_header ELibxvidFlag = Flags_global_header
)

type SideDataPkt = Flag

const (
	SideDataPkt_replaygain                 SideDataPkt = "replaygain"
	SideDataPkt_displaymatrix              SideDataPkt = "displaymatrix"
	SideDataPkt_spherical                  SideDataPkt = "spherical"
	SideDataPkt_stereo3d                   SideDataPkt = "stereo3d"
	SideDataPkt_audio_service_type         SideDataPkt = "audio_service_type"
	SideDataPkt_mastering_display_metadata SideDataPkt = "mastering_display_metadata"
	SideDataPkt_content_light_level        SideDataPkt = "content_light_level"
	SideDataPkt_icc_profile                SideDataPkt = "icc_profile"
)
