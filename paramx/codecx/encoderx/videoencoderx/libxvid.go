package videoencoderx

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// LIBXVID 9.20 libxvid
type LIBXVID struct {
	Bf typex.UNumber `json:"bf" flag:"-bf"` // (bframes)

	Flags flagx.ELibxvidFlag `json:"flags" flag:"-flags"`

	Gmc        typex.UBool  `json:"gmc" flag:"-gmc"`
	MeQuality  typex.Level  `json:"me_quality" flag:"-me_quality"`
	Mbd        flagx.Mbd    `json:"mbd" flag:"-mbd"`
	LumiAq     typex.UBool  `json:"lumi_aq" flag:"-lumi_aq"`
	VarianceAq typex.UBool  `json:"variance_aq" flag:"-variance_aq"`
	Trellis    typex.Number `json:"trellis" flag:"-trellis"`
	Ssim       Ssim         `json:"ssim" flag:"-ssim"`
	SsimAcc    typex.Level  `json:"ssim_acc" flag:"-ssim_acc"`
}

type Ssim string

const (
	Ssim_off   Ssim = "off"
	Ssim_avg   Ssim = "avg"
	Ssim_frame Ssim = "frame"
)
