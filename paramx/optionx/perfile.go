package optionx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// AdvPerFile
type AdvPerFile struct {
	Map typex.Func `json:"map" flag:"-map"` //                    OPT_TYPE_FUNC, OPT_FUNC_ARG | OPT_EXPERT | OPT_PERFILE | OPT_OUTPUT,

	Timestamp typex.Func `json:"timestamp" flag:"-timestamp"` //              OPT_TYPE_FUNC,   OPT_FUNC_ARG | OPT_PERFILE | OPT_EXPERT | OPT_OUTPUT,

	Dframes typex.Func `json:"dframes" flag:"-dframes"` //                OPT_TYPE_FUNC, OPT_FUNC_ARG | OPT_PERFILE | OPT_EXPERT | OPT_OUTPUT | OPT_HAS_CANON,

	Target typex.Func `json:"target" flag:"-target"` //                 OPT_TYPE_FUNC, OPT_FUNC_ARG | OPT_PERFILE | OPT_EXPERT | OPT_OUTPUT,

	Qscale typex.Func `json:"qscale" flag:"-qscale"` //                 OPT_TYPE_FUNC, OPT_FUNC_ARG | OPT_EXPERT | OPT_PERFILE | OPT_OUTPUT | OPT_HAS_ALT,

	Profile typex.Func `json:"profile" flag:"-profile"` //                OPT_TYPE_FUNC, OPT_FUNC_ARG | OPT_EXPERT | OPT_PERFILE | OPT_OUTPUT,

	Attach typex.Func `json:"attach" flag:"-attach"` //              OPT_TYPE_FUNC, OPT_FUNC_ARG | OPT_PERFILE | OPT_EXPERT | OPT_OUTPUT,

	Vframes typex.Func `json:"vframes" flag:"-vframes"` //                    OPT_TYPE_FUNC,   OPT_VIDEO | OPT_FUNC_ARG | OPT_PERFILE | OPT_OUTPUT | OPT_EXPERT | OPT_HAS_CANON,

	Timecode typex.Func `json:"timecode" flag:"-timecode"` //                   OPT_TYPE_FUNC,   OPT_VIDEO | OPT_FUNC_ARG | OPT_PERFILE | OPT_OUTPUT | OPT_EXPERT,

	Streamid typex.Func `json:"streamid" flag:"-streamid"` //                   OPT_TYPE_FUNC,   OPT_VIDEO | OPT_FUNC_ARG | OPT_EXPERT | OPT_PERFILE | OPT_OUTPUT,

	/* audio options */
	Aframes typex.Func `json:"aframes" flag:"-aframes"` //          OPT_TYPE_FUNC,    OPT_AUDIO | OPT_FUNC_ARG | OPT_PERFILE | OPT_OUTPUT | OPT_EXPERT | OPT_HAS_CANON,

	Atag typex.Tag `json:"atag" flag:"-atag"` //             OPT_TYPE_FUNC,    OPT_AUDIO | OPT_FUNC_ARG  | OPT_EXPERT | OPT_PERFILE | OPT_OUTPUT | OPT_HAS_CANON,

	Stag typex.Func `json:"stag" flag:"-stag"` //   OPT_TYPE_FUNC, OPT_SUBTITLE | OPT_FUNC_ARG  | OPT_EXPERT  | OPT_PERFILE | OPT_OUTPUT | OPT_HAS_CANON,

	Apre typex.Func `json:"apre" flag:"-apre"` // OPT_TYPE_FUNC, OPT_FUNC_ARG | OPT_AUDIO | OPT_EXPERT| OPT_PERFILE | OPT_OUTPUT | OPT_HAS_CANON,

	Vpre typex.Func `json:"vpre" flag:"-vpre"` // OPT_TYPE_FUNC, OPT_VIDEO | OPT_FUNC_ARG | OPT_EXPERT| OPT_PERFILE | OPT_OUTPUT | OPT_HAS_CANON,

	Spre typex.Func `json:"spre" flag:"-spre"` // OPT_TYPE_FUNC, OPT_FUNC_ARG | OPT_SUBTITLE | OPT_EXPERT| OPT_PERFILE | OPT_OUTPUT | OPT_HAS_CANON,

	Fpre typex.Func `json:"fpre" flag:"-fpre"` // OPT_TYPE_FUNC, OPT_FUNC_ARG | OPT_EXPERT| OPT_PERFILE | OPT_OUTPUT | OPT_HAS_CANON,

}
