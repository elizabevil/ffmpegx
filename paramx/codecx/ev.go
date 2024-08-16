package codecx

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// EV encoding & video: video and encoding  are almost all together
type EV struct {
	Bt typex.BitrateInt `json:"bt" flag:"-bt"` // (encoding, video)

	G typex.Size `json:"g" flag:"-g"` // (encoding, video)

	Qcomp typex.Flt `json:"qcomp" flag:"-qcomp"` // (encoding, video)

	Qblur typex.UFloat32 `json:"qblur" flag:"-qblur"` // (encoding, video)

	Qmin typex.Quantizer `json:"qmin" flag:"-qmin"` // (encoding, video)

	Qmax typex.Quantizer `json:"qmax" flag:"-qmax"` // (encoding, video)

	Qdiff typex.Number `json:"qdiff" flag:"-qdiff"` // (encoding, video)

	Bf typex.UNumber `json:"bf" flag:"-bf"` // (encoding, video)

	BQfactor typex.Flt `json:"b_qfactor" flag:"-b_qfactor"` // (encoding, video)

	BQoffset typex.Flt `json:"b_qoffset" flag:"-b_qoffset"` // (encoding, video)

	IQfactor typex.Flt `json:"i_qfactor" flag:"-i_qfactor"` // (encoding, video)

	IQoffset typex.Flt `json:"i_qoffset" flag:"-i_qoffset"` // (encoding, video)

	Dct typex.UNumber `json:"dct" flag:"-dct"` // (encoding, video)

	LumiMask typex.Flt `json:"lumi_mask" flag:"-lumi_mask"` // (encoding, video)

	TcplxMask typex.Flt `json:"tcplx_mask" flag:"-tcplx_mask"` // (encoding, video)

	ScplxMask typex.Flt `json:"scplx_mask" flag:"-scplx_mask"` // (encoding, video)

	PMask typex.Flt `json:"p_mask" flag:"-p_mask"` // (encoding, video)

	DarkMask typex.Flt `json:"dark_mask" flag:"-dark_mask"` // (encoding, video)

	Aspect typex.Level `json:"aspect" flag:"-aspect"` //10 (encoding, video)

	Sar typex.Level `json:"sar" flag:"-sar"` //10 (encoding, video)

	Cmp     int       `json:"cmp" flag:"-cmp"`      // (encoding, video)
	CmpFlag flagx.Cmp `json:"cmp_flag" flag:"-cmp"` // (encoding, video)

	Subcmp     int          `json:"subcmp" flag:"-subcmp"`      // (encoding, video)
	SubcmpFlag flagx.Subcmp `json:"subcmp_flag" flag:"-subcmp"` // (encoding, video)

	Mbcmp     int         `json:"mbcmp" flag:"-mbcmp"`      // (encoding, video)
	MbcmpFlag flagx.Mbcmp `json:"mbcmp_flag" flag:"-mbcmp"` // (encoding, video)

	Ildctcmp     int            `json:"ildctcmp" flag:"-ildctcmp"`      // (encoding, video)
	IldctcmpFlag flagx.Ildctcmp `json:"ildctcmp_flag" flag:"-ildctcmp"` // (encoding, video)

	DiaSize typex.DiaSize `json:"dia_size" flag:"-dia_size"` // (encoding, video)

	LastPred typex.LastPred `json:"last_pred" flag:"-last_pred"` // (encoding, video)

	Precmp     int          `json:"precmp" flag:"-precmp"`      // (encoding, video)
	PrecmpFlag flagx.Precmp `json:"precmp_flag" flag:"-precmp"` // (encoding, video)

	PreDiaSize typex.Size `json:"pre_dia_size" flag:"-pre_dia_size"` // (encoding, video)

	Subq typex.Number `json:"subq" flag:"-subq"` // (encoding, video)

	MeRange typex.Number `json:"me_range" flag:"-me_range"` // (encoding, video)

	Mbd int `json:"mbd" flag:"-mbd"` // (encoding, video)

	RcInitOccupancy typex.Number `json:"rc_init_occupancy" flag:"-rc_init_occupancy"` // (encoding, video)

	Dc typex.LevelI8 `json:"dc" flag:"-dc"` // (encoding, video)

	Nssew typex.Number `json:"nssew" flag:"-nssew"` // (encoding, video)

	Mblmin typex.UVBR `json:"mblmin" flag:"-mblmin"` // (encoding, video)

	Mblmax typex.UVBR `json:"mblmax" flag:"-mblmax"` // (encoding, video)

	BidirRefine typex.BidirRefine `json:"bidir_refine" flag:"-bidir_refine"` // (encoding, video)

	KeyintMin typex.Number `json:"keyint_min" flag:"-keyint_min"` // (encoding, video)

	Refs typex.Number `json:"refs" flag:"-refs"` // (encoding, video)

	Mv0Threshold typex.Size `json:"mv0_threshold" flag:"-mv0_threshold"` // (encoding, video)

	RcMaxVbvUse typex.UFloat32 `json:"rc_max_vbv_use" flag:"-rc_max_vbv_use"` // (encoding, video)

	RcMinVbvUse typex.UFloat32 `json:"rc_min_vbv_use" flag:"-rc_min_vbv_use"` // (encoding, video)

	Slices typex.Size `json:"slices" flag:"-slices"` // (encoding, video)

	FieldOrder flagx.FieldOrder `json:"field_order" flag:"-field_order"` // (video) 	//Set/override the field order of the video. Possible values:
}
