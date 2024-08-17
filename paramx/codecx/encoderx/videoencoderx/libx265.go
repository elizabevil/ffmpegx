package videoencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

type Libx265 struct {
	B       typex.BitrateInt `json:"b" flag:"-b"`
	Bitrate typex.Bitrate    `json:"bitrate" flag:"-b"`
	//Sets target video bitrate.

	Bf typex.UNumber `json:"bf" flag:"-bf"`
	G  typex.Size    `json:"g" flag:"-g"`
	//Set the GOP size.

	KeyintMin typex.Number `json:"keyint_min" flag:"-keyint_min"`
	//Minimum GOP size.

	Refs typex.Level `json:"refs" flag:"-refs"`
	//Number of reference frames each P-frame can use. The range is from 1-16.

	Preset typex.String `json:"preset" flag:"-preset"`
	//Set the x265 preset.

	Tune typex.String `json:"tune" flag:"-tune"`
	//Set the x265 tune parameter.

	Profile typex.Profile `json:"profile" flag:"-profile"`
	//Set profile restrictions.

	Crf typex.UI8 `json:"crf" flag:"-crf"`
	//Set the quality for constant quality mode.

	Qp typex.Level `json:"qp" flag:"-qp"`
	//Set constant quantization rate control method parameter.

	Qmin typex.Quantizer `json:"qmin" flag:"-qmin"`
	//Minimum quantizer scale.

	Qmax typex.Quantizer `json:"qmax" flag:"-qmax"`
	//Maximum quantizer scale.

	Qdiff typex.Number `json:"qdiff" flag:"-qdiff"`
	//Maximum difference between quantizer scales.

	Qblur typex.UFlt `json:"qblur" flag:"-qblur"` //	Quantizer curve blur

	Qcomp typex.Flt `json:"qcomp" flag:"-qcomp"` //	Quantizer curve compression factor

	IQfactor  typex.Flt    `json:"i_qfactor" flag:"-i_qfactor"`
	BQfactor  typex.Flt    `json:"b_qfactor" flag:"-b_qfactor"`
	ForcedIdr typex.Number `json:"forced-idr" flag:"-forced-idr"`
	//Normally, when forcing a I-frame type, the encoder can select any type of I-frame. This option forces it to choose an IDR-frame.

	UduSei typex.UBool `json:"udu_sei" flag:"-udu_sei"`
	//Import user data unregistered SEI if available into output. Default is 0 (off).

	X265Params typex.Dict `json:"x265-params" flag:"-x265-params"`
	//Set x265 options using a list of key=value couples separated by ":". See x265 --help for a list of options.

}

//For example to specify libx265 encoding options with -x265-params:
//ffmpeg -i input -c:v libx265 -x265-params crf=26:psy-rd=1 output.mp4
