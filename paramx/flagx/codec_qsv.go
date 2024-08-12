package flagx

type LoadPlugin string

const (
	LoadPlugin_none    LoadPlugin = "none"
	LoadPlugin_hevc_sw LoadPlugin = "hevc_sw"
	LoadPlugin_hevc_hw LoadPlugin = "hevc_hw"
)
