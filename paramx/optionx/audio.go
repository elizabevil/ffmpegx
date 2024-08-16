package optionx

import "github.com/elizabevil/ffmpegx/paramx/typex"

type Audio struct {
	Aframes int `json:"aframes" flag:"-aframes"` // aframes  number (output)
	//Set the number of audio frames to output. This is an obsolete alias for -frames:a, which you should use instead.

	Ar typex.StreamSpecifier `json:"ar" flag:"-ar"` // ar [:stream_specifier] freq (input/output,per-stream)
	//Set the audio sampling frequency. For output streams it is set by default to the frequency of the corresponding input stream. For input streams this option only makes sense for audio grabbing devices and raw demuxers and is mapped to the corresponding demuxer options.

	Aq string `json:"aq" flag:"-aq"` // aq  q (output)
	//Set the audio quality (codec-specific, UVBR). This is an alias for -q:a.

	Ac typex.Channels `json:"ac" flag:"-ac"` // ac [:stream_specifier] channels (input/output,per-stream)
	//Set the number of audio channels. For output streams it is set by default to the number of input audio channels. For input streams this option only makes sense for audio grabbing devices and raw demuxers and is mapped to the corresponding demuxer options.

	An bool `json:"an" flag:"-an"` // an  (input/output)
	//As an input option, blocks all audio streams of a file from being filtered or being automatically selected or mapped for any output. See -discard option to disable streams individually.
	//As an output option, disables audio recording i.e. automatic selection or mapping of any audio stream. For full manual control see the -map option.

	Acodec typex.Codec `json:"acodec" flag:"-acodec"` // acodec  codec (input/output)
	//Set the audio codec. This is an alias for -codec:a.

	//-sample_fmt[:stream_specifier] sample_fmt (output,per-stream)
	//Set the audio sample format. Use -sample_fmts to get a list of supported sample formats.

	Af typex.Filtergraph `json:"af" flag:"-af"` // af  filtergraph (output)
	//Create the filtergraph specified by filtergraph and use it to filter the stream.
	//This is an alias for -filter:a, see the -filter option.

	Atag typex.Tag `json:"atag" flag:"-atag"` //fourcc/tag (output)
	//Force audio tag/fourcc. This is an alias for -tag:a.

	ChLayout typex.ChannelLayout `json:"ch_layout" flag:"-ch_layout"` //stream_specifier]  layout (input/output, per-stream)
	//Alias for -channel_layout.

	ChannelLayout typex.ChannelLayout `json:"channel_layout" flag:"-channel_layout"` // layout (input/output, per-stream)
	//Set the audio channel layout. For output streams it is set by default to the input channel layout. For input streams it overrides the channel layout of the input. Not all decoders respect the overridden channel layout. This option also sets the channel layout for audio grabbing devices and raw demuxers and is mapped to the corresponding demuxer option.

	GuessLayoutMax typex.Number `json:"guess_layout_max" flag:"-guess_layout_max"` // channels (input, per-stream)
	//If some input channel layout is not known, try to guess only if it corresponds to at most the specified number of channels. For example, 2 tells to ffmpeg to recognize 1 channel as mono and 2 channels as stereo but not 6 channels as 5.1. The default is to always try to guess. Use 0 to disable all guessing. Using the -channel_layout option to explicitly specify an input layout also disables guessing.
}
