package demuxerx

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// https://ffmpeg.org/ffmpeg-all.html#mov_002fmp4_002f3gp
// 20.17 mov/mp4/3gp
// Registered extensions: mov, mp4, m4a, 3gp, 3g2, mj2, psp, m4b, ism, ismv, isma, f4v

type MPEG4 struct {
	EnableDrefs typex.UBool `json:"enable_drefs" flag:"-enable_drefs"`
	//Enable loading of external tracks, disabled by default. Enabling this can theoretically leak information in some use cases.

	UseAbsolutePath typex.UBool `json:"use_absolute_path" flag:"-use_absolute_path"`
	//Allows loading of external tracks via absolute paths, disabled by default. Enabling this poses a security risk. It should only be enabled if the source is known to be non-malicious.

	SeekStreamsIndividually typex.UBool `json:"seek_streams_individually" flag:"-seek_streams_individually"`
	//When seeking, identify the closest point in each stream individually and demux packets in that stream from identified point. This can lead to a different sequence of packets compared to demuxing linearly from the beginning. Default is true.

	IgnoreEditlist typex.UBool `json:"ignore_editlist" flag:"-ignore_editlist"`
	//Ignore any edit list atoms. The demuxer, by default, modifies the stream index to reflect the timeline described by the edit list. Default is false.

	AdvancedEditlist *typex.UBool `json:"advanced_editlist" flag:"-advanced_editlist,1"`
	//Modify the stream index to reflect the timeline described by the edit list. ignore_editlist must be set to false for this option to be effective. If both ignore_editlist and this option are set to false, then only the start of the stream index is modified to reflect initial dwell time or starting timestamp described by the edit list. Default is true.

	IgnoreChapters typex.UBool `json:"ignore_chapters" flag:"-ignore_chapters"`
	//Don’t parse chapters. This includes GoPro ’HiLight’ tags/moments. Note that chapters are only parsed when input is seekable. Default is false.

	UseMfraFor flagx.UseMfraFor `json:"use_mfra_for" flag:"-use_mfra_for"`
	//For seekable fragmented input, set fragment’s starting timestamp from media fragment random access box, if present.

	UseTfdt typex.UBool `json:"use_tfdt" flag:"-use_tfdt"`
	//For fragmented input, set fragment’s starting timestamp to baseMediaDecodeTime from the tfdt box. Default is enabled, which will prefer to use the tfdt box to set DTS. Disable to use the earliest_presentation_time from the sidx box. In either case, the timestamp from the mfra box will be used if it’s available and use_mfra_for is set to pts or dts.

	ExportAll typex.UBool `json:"export_all" flag:"-export_all"`
	//Export unrecognized boxes within the udta box as metadata entries. The first four characters of the box type are set as the key. Default is false.

	ExportXmp typex.UBool `json:"export_xmp" flag:"-export_xmp"`
	//Export entire contents of XMP_ box and uuid box as a string with key xmp. Note that if export_all is set and this option isn’t, the contents of XMP_ box are still exported but with key XMP_. Default is false.

	ActivationBytes typex.Key `json:"activation_bytes" flag:"-activation_bytes"`
	//4-byte key required to decrypt Audible AAX and AAX+ files. See Audible AAX subsection below.

	AudibleFixedKey typex.Key `json:"audible_fixed_key" flag:"-audible_fixed_key"`
	//Fixed key used for handling Audible AAX/AAX+ files. It has been pre-set so should not be necessary to specify.

	DecryptionKey typex.Key `json:"decryption_key" flag:"-decryption_key"`
	//16-byte key, in hex, to decrypt files encrypted using ISO Common Encryption (CENC/AES-128 CTR; ISO/IEC 23001-7).

	MaxSttsDelta typex.UBool `json:"max_stts_delta" flag:"-max_stts_delta"`
	//Very high sample deltas written in a trak’s stts box may occasionally be intended but usually they are written in error or used to store a negative value for dts correction when treated as signed 32-bit integers. This option lets the user set an upper limit, beyond which the delta is clamped to 1. Values greater than the limit if negative when cast to int32 are used to adjust onward dts.

	//Unit is the track time scale. Range is 0 to UINT_MAX. Default is UINT_MAX - 48000*10 which allows up to a 10 second dts correction for 48 kHz audio streams while accommodating 99.9% of uint32 range.

	InterleavedRead typex.UBool `json:"interleaved_read" flag:"-interleaved_read"`
	//Interleave packets from multiple tracks at demuxer level. For badly interleaved files, this prevents playback issues caused by large gaps between packets in different tracks, as MOV/MP4 do not have packet placement requirements. However, this can cause excessive seeking on very badly interleaved files, due to seeking between tracks, so disabling it may prevent I/O issues, at the expense of playback.
}

//Audible AAX files are encrypted M4B files, and they can be decrypted by specifying a 4 byte activation secret.

//ffmpeg -activation_bytes 1CEB00DA -i test.aax -vn -c:a copy output.mp4
