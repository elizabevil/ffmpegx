package muxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// 21.21 avi AVI
type AVI struct {
	FlippedRawRgb typex.UBool `json:"flipped_raw_rgb" flag:"-flipped_raw_rgb"`
	//If set to true, store positive height for raw RGB bitmaps, which indicates bitmap is stored bottom-up. Note that this option does not flip the bitmap which has to be done manually beforehand, e.g. by using the ‘vflip’ filter. Default is false and indicates bitmap is stored top down.

	ReserveIndexSpace typex.Index `json:"reserve_index_space" flag:"-reserve_index_space"`
	//Reserve the specified amount of bytes for the OpenDML master index of each stream within the file header. By default additional master indexes are embedded within the data packets if there is no space left in the first master index and are linked together as a chain of indexes. This index structure can cause problems for some use cases, e.g. third-party software strictly relying on the OpenDML index specification or when file seeking is slow. Reserving enough index space in the file header avoids these problems.

	//The required index space depends on the output file size and should be about 16 bytes per gigabyte. When this option is omitted or set to zero the necessary index space is guessed.

	//Default value is 0.

	WriteChannelMask typex.UBool `json:"write_channel_mask" flag:"-write_channel_mask"`
	//Write the channel layout mask into the audio stream header.

	//This option is enabled by default. Disabling the channel mask can be useful in specific scenarios, e.g. when merging multiple audio streams into one for compatibility with software that only supports a single audio stream in AVI (see (ffmpeg-filters)the "amerge" section in the ffmpeg-filters manual).
}
