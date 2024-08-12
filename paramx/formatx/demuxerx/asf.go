package demuxerx

// 20.4 asf  ASF
type ASF struct {
	NoResyncSearch bool `json:"no_resync_search" flag:"-no_resync_search"`
	//Do not try to resynchronize by looking for a certain optional start code.
}
