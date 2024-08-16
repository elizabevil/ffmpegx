package metadatax

type Tags = map[string]any

type Stream struct {
	BaseStream
	StreamAudioOnly
	StreamVideoOnly
	StreamSubtitleOnly
}
type StreamAudio struct {
	BaseStream
	StreamAudioOnly
}
type StreamVideo struct {
	BaseStream
	StreamVideoOnly
}
type StreamSubtitle struct {
	BaseStream
	StreamSubtitleOnly
}
