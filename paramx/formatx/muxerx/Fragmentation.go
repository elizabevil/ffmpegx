package muxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// //FRAGMENTATION 21.4.1 Fragmentation FRAGMENTATION
// //4.4.1 Fragmentation
type FRAGMENTATION struct {
	Brand typex.Brand `json:"brand" flag:"-brand"`
	//Override major brand.

	EmptyHdlrName typex.UBool `json:"empty_hdlr_name" flag:"-empty_hdlr_name"`
	//Enable to skip writing the name inside a hdlr box.Default is false.

	EncryptionKey typex.Key `json:"encryption_key" flag:"-encryption_key"`
	//set the media encryption key in hexadecimal format

	EncryptionKid typex.ID `json:"encryption_kid" flag:"-encryption_kid"`
	//set the media encryption key identifier in hexadecimal format

	EncryptionScheme typex.Scheme `json:"encryption_scheme" flag:"-encryption_scheme"`
	//configure         the encryption scheme, allowed values are ‘none’, and ‘cenc-aes-ctr’

	FragDuration typex.MicrosecondI `json:"frag_duration" flag:"-frag_duration"`
	//Create fragments that are duration microseconds long.

	FragInterleave typex.UNumber `json:"frag_interleave" flag:"-frag_interleave"`
	//Interleave samples within fragments (max number of consecutive samples, lower is tighter interleaving, but with more overhead.It is set to 0 by default.

	FragSize typex.Size `json:"frag_size" flag:"-frag_size"`
	//create fragments that contain up to size bytes of payload data

	IodsAudioProfile typex.Level `json:"iods_audio_profile" flag:"-iods_audio_profile"`
	//specify iods number for the audio profile atom (from -1 to 255), default is -1

	IodsVideoProfile typex.Level `json:"iods_video_profile" flag:"-iods_video_profile"`
	//specify iods number for the video profile atom (from -1 to 255), default is -1

	IsmLookahead typex.UI8 `json:"ism_lookahead" flag:"-ism_lookahead"`
	//specify number of lookahead entries for ISM files (from 0 to 255), default is 0

	MinFragDuration typex.MicrosecondI `json:"min_frag_duration" flag:"-min_frag_duration"`
	//do not create fragments that are shorter than duration microseconds long

	MoovSize typex.Size `json:"moov_size" flag:"-moov_size"`
	//Reserves space for the moov atom at the beginning of the file instead of placing the moov atom at the end.If the space reserved is insufficient, muxing will fail.

	MovGamma typex.UFlt `json:"mov_gamma" flag:"-mov_gamma"`
	//specify gamma value for gama atom (as a decimal number from 0 to 10), default is 0.0, must be set together with + movflags

	Movflags Movflags `json:"movflags" flag:"-movflags"`

	MovieTimescale typex.ScaleI `json:"movie_timescale" flag:"-movie_timescale"`
	//Set the timescale written in the movie header box (mvhd).Range is 1 to INT_MAX.Default is 1000.

	Rtpflags Rtpflags     `json:"rtpflags" flag:"-rtpflags"`
	SkipIods *typex.UBool `json:"skip_iods" flag:"-skip_iods,1"`
	//skip writing iods atom (default value is true)

	UseEditlist typex.Bool `json:"use_editlist" flag:"-use_editlist"`
	//use edit list (default value is auto)

	UseStreamIdsAsTrackIds typex.UBool `json:"use_stream_ids_as_track_ids" flag:"-use_stream_ids_as_track_ids"`
	//use stream ids as track ids (default value is false)

	VideoTrackTimescale typex.ScaleI `json:"video_track_timescale" flag:"-video_track_timescale"`
	//Set the timescale used for video tracks.Range is 0 to INT_MAX.If set to 0, the timescale is automatically set based on the native stream time base.Default is 0.

	WriteBtrt typex.Bool `json:"write_btrt" flag:"-write_btrt"`
	//Force or disable writing bitrate box inside stsd box of a track.The box contains decoding buffer size (in bytes), maximum bitrate and average bitrate for the track.The box will be skipped if none of these values can be computed.Default is -1 or auto, which will write the box only in MP4 mode.

	WritePrft WritePrft `json:"write_prft" flag:"-write_prft,-1"`
	//Write producer time reference box (PRFT) with a specified time source for the NTP field in the PRFT box.Set value as ‘wallclock’ to specify timesource as wallclock time and ‘pts’ to specify timesource as input packets’ PTS values.

	WriteTmcd typex.Bool `json:"write_tmcd" flag:"-write_tmcd"`
	//Specify on to force writing a timecode track, off to disable it and auto to write a timecode track only for mov and mp4 output (default).

	//Setting value to ‘pts’ is applicable only for a live encoding use case, where PTS values are set as as wallclock time at the source.For example, an encoding use case with decklink capture source where video_pts and audio_pts are set to ‘abs_wallclock’.
}

type WritePrft = typex.String
type Rtpflags = typex.String
type Movflags = typex.String

const (
	WritePrft_pts       WritePrft = "pts"
	WritePrft_wallclock WritePrft = "wallclock"
)

const (
	cmaf Movflags = "cmaf"
	//write CMAF (Common Media Application Format) compatible fragmented MP4 output

	dash Movflags = "dash"
	//write DASH (Dynamic Adaptive Streaming over HTTP) compatible fragmented MP4 output

	default_base_moof Movflags = "default_base_moof"
	//Similarly to the ‘omit_tfhd_offset’ flag, this flag avoids writing the absolute base_data_offset field in tfhd atoms, but does so by using the new default-base-is-moof flag instead. This flag is new from 14496-12:2012. This may make the fragments easier to parse in certain circumstances (avoiding basing track fragment location calculations on the implicit end of the previous track fragment).

	delay_moov Movflags = "delay_moov"
	//delay writing the initial moov until the first fragment is cut, or until the first fragment flush

	disable_chpl Movflags = "disable_chpl"
	//Disable Nero chapter markers (chpl atom). Normally, both Nero chapters and a QuickTime chapter track are written to the file. With this option set, only the QuickTime chapter track will be written. Nero chapters can cause failures when the file is reprocessed with certain tagging programs, like mp3Tag 2.61a and iTunes 11.3, most likely other versions are affected as well.

	faststart Movflags = "faststart"
	//Run a second pass moving the index (moov atom) to the beginning of the file. This operation can take a while, and will not work in various situations such as fragmented output, thus it is not enabled by default.

	frag_custom Movflags = "frag_custom"
	//Allow the caller to manually choose when to cut fragments, by calling av_write_frame(ctx, NULL) to write a fragment with the packets written so far. (This is only useful with other applications integrating libavformat, not from ffmpeg.)

	frag_discont Movflags = "frag_discont"
	//signal that the next fragment is discontinuous from earlier ones

	frag_every_frame Movflags = "frag_every_frame"
	//fragment at every frame

	frag_keyframe Movflags = "frag_keyframe"
	//start a new fragment at each video keyframe

	global_sidx Movflags = "global_sidx"
	//write a global sidx index at the start of the file

	isml Movflags = "isml"
	//create a live smooth streaming feed (for pushing to a publishing point)

	negative_cts_offsets Movflags = "negative_cts_offsets"
	//Enables utilization of version 1 of the CTTS box, in which the CTS offsets can be negative. This enables the initial sample to have DTS/CTS of zero, and reduces the need for edit lists for some cases such as video tracks with B-frames. Additionally, eases conformance with the DASH-IF interoperability guidelines.

	//This option is implicitly set when writing ‘ismv’ (Smooth Streaming) files.

	omit_tfhd_offset Movflags = "omit_tfhd_offset"
	//Do not write any absolute base_data_offset in tfhd atoms. This avoids tying fragments to absolute byte positions in the file/streams.

	prefer_icc Movflags = "prefer_icc"
	//If writing colr atom prioritise usage of ICC profile if it exists in stream packet side data.

	rtphint Movflags = "rtphint"
	//add RTP hinting tracks to the output file

	separate_moof Movflags = "separate_moof"
	//Write a separate moof (movie fragment) atom for each track. Normally, packets for all tracks are written in a moof atom (which is slightly more efficient), but with this option set, the muxer writes one moof/mdat pair for each track, making it easier to separate tracks.

	skip_sidx Movflags = "skip_sidx"
	//Skip writing of sidx atom. When bitrate overhead due to sidx atom is high, this option could be used for cases where sidx atom is not mandatory. When the ‘global_sidx’ flag is enabled, this option is ignored.

	skip_trailer Movflags = "skip_trailer"
	//skip writing the mfra/tfra/mfro trailer for fragmented files

	use_metadata_tags Movflags = "use_metadata_tags"
	//use mdta atom for metadata

	write_colr Movflags = "write_colr"
	//write colr atom even if the color info is unspecified. This flag is experimental, may be renamed or changed, do not use from scripts.

	write_gama Movflags = "write_gama"
	//write deprecated gama atom

	hybrid_fragmented Movflags = "hybrid_fragmented"
	//For recoverability - write the output file as a fragmented file. This allows the intermediate file to be read while being written (in particular, if the writing process is aborted uncleanly). When writing is finished, the file is converted to a regular, non-fragmented file, which is more compatible and allows easier and quicker seeking.

	//If writing is aborted, the intermediate file can manually be remuxed to get a regular, non-fragmented file of what had been written into the unfinished file.
)

const (
	Rtpflags_h264_mode0 Rtpflags = "h264_mode0"
	//use mode 0 for H.264 in RTP

	Rtpflags_latm Rtpflags = "latm"
	//use MP4A-LATM packetization instead of MPEG4-GENERIC for AAC

	Rtpflags_rfc2190 Rtpflags = "rfc2190"
	//use RFC 2190 packetization instead of RFC 4629 for H.263

	Rtpflags_send_bye Rtpflags = "send_bye"
	//send RTCP BYE packets when finishing

	Rtpflags_skip_rtcp Rtpflags = "skip_rtcp"
	//do not send RTCP sender reports

)
