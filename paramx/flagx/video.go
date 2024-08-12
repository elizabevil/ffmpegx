package flagx

import "github.com/elizabevil/ffmpegx/paramx/typex"

type VideoRate = Flag

// Specify the size of the sourced video, it may be a Flag of the form widthxheight, or the name of a size abbreviation.
const (
	VideoRate_ntsc VideoRate = "ntsc" //30000/1001

	VideoRate_pal VideoRate = "pal" //25/1

	VideoRate_qntsc VideoRate = "qntsc" //30000/1001

	VideoRate_qpal VideoRate = "qpal" //25/1

	VideoRate_sntsc VideoRate = "sntsc" //30000/1001

	VideoRate_spal VideoRate = "spal" //25/1

	VideoRate_film VideoRate = "film" //24/1

	VideoRate_ntsc_film VideoRate = "ntsc-film" //24000/1001

)

type VideoSize = typex.VideoSize

const (
	VideoSize_ntsc VideoSize = "ntsc" // 720x480

	VideoSize_pal VideoSize = "pal" // 720x576

	VideoSize_qntsc VideoSize = "qntsc" // 352x240

	VideoSize_qpal VideoSize = "qpal" // 352x288

	VideoSize_sntsc VideoSize = "sntsc" // 640x480

	VideoSize_spal VideoSize = "spal" // 768x576

	VideoSize_film VideoSize = "film" // 352x240

	VideoSize_ntsc_film VideoSize = "ntsc-film" // 352x240

	VideoSize_sqcif VideoSize = "sqcif" // 128x96

	VideoSize_qcif VideoSize = "qcif" // 176x144

	VideoSize_cif VideoSize = "cif" // 352x288

	VideoSize_4cif VideoSize = "4cif" // 704x576

	VideoSize_16cif VideoSize = "16cif" // 1408x1152

	VideoSize_qqvga VideoSize = "qqvga" // 160x120

	VideoSize_qvga VideoSize = "qvga" // 320x240

	VideoSize_vga VideoSize = "vga" // 640x480

	VideoSize_svga VideoSize = "svga" // 800x600

	VideoSize_xga VideoSize = "xga" // 1024x768

	VideoSize_uxga VideoSize = "uxga" // 1600x1200

	VideoSize_qxga VideoSize = "qxga" // 2048x1536

	VideoSize_sxga VideoSize = "sxga" // 1280x1024

	VideoSize_qsxga VideoSize = "qsxga" // 2560x2048

	VideoSize_hsxga VideoSize = "hsxga" // 5120x4096

	VideoSize_wvga VideoSize = "wvga" // 852x480

	VideoSize_wxga VideoSize = "wxga" // 1366x768

	VideoSize_wsxga VideoSize = "wsxga" // 1600x1024

	VideoSize_wuxga VideoSize = "wuxga" // 1920x1200

	VideoSize_woxga VideoSize = "woxga" // 2560x1600

	VideoSize_wqsxga VideoSize = "wqsxga" // 3200x2048

	VideoSize_wquxga VideoSize = "wquxga" // 3840x2400

	VideoSize_whsxga VideoSize = "whsxga" // 6400x4096

	VideoSize_whuxga VideoSize = "whuxga" // 7680x4800

	VideoSize_cga VideoSize = "cga" // 320x200

	VideoSize_ega VideoSize = "ega" // 640x350

	VideoSize_hd480 VideoSize = "hd480" // 852x480

	VideoSize_hd720 VideoSize = "hd720" // 1280x720

	VideoSize_hd1080 VideoSize = "hd1080" // 1920x1080

	VideoSize_2k VideoSize = "2k" // 2048x1080

	VideoSize_2kflat VideoSize = "2kflat" // 1998x1080

	VideoSize_2kscope VideoSize = "2kscope" // 2048x858

	VideoSize_4k VideoSize = "4k" // 4096x2160

	VideoSize_4kflat VideoSize = "4kflat" // 3996x2160

	VideoSize_4kscope VideoSize = "4kscope" // 4096x1716

	VideoSize_nhd VideoSize = "nhd" // 640x360

	VideoSize_hqvga VideoSize = "hqvga" // 240x160

	VideoSize_wqvga VideoSize = "wqvga" // 400x240

	VideoSize_fwqvga VideoSize = "fwqvga" // 432x240

	VideoSize_hvga VideoSize = "hvga" // 480x320

	VideoSize_qhd VideoSize = "qhd" // 960x540

	VideoSize_2kdci VideoSize = "2kdci" // 2048x1080

	VideoSize_4kdci VideoSize = "4kdci" // 4096x2160

	VideoSize_uhd2160 VideoSize = "uhd2160" // 3840x2160

	VideoSize_uhd4320 VideoSize = "uhd4320" // 7680x4320

)
