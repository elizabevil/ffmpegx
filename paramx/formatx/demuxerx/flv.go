package demuxerx

// 20.10 flv, live_flv, kux
type FLV struct {
	FlvMetadata bool `json:"flv_metadata" flag:"-flv_metadata"`
	//Allocate the streams according to the onMetaData array content.

	FlvIgnorePrevtag bool `json:"flv_ignore_prevtag" flag:"-flv_ignore_prevtag"`
	//Ignore the size of previous tag value.

	FlvFullMetadata bool `json:"flv_full_metadata" flag:"-flv_full_metadata"`
	//Output all context of the onMetadata.
}

//ffmpeg -f flv -i myfile.flv ...
//ffmpeg -f live_flv -i rtmp://<any.server>/anything/key ....
