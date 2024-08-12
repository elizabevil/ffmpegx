package demuxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// Mpegts 20.18 mpegts
// MPEG-2 transport stream demuxer.
type Mpegts struct {
	ResyncSize typex.Size `json:"resync_size" flag:"-resync_size"`
	//Set size limit for looking up a new synchronization.Default value is 65536.

	SkipUnknownPmt typex.Bool `json:"skip_unknown_pmt" flag:"-skip_unknown_pmt"`
	//Skip PMTs for programs not defined in the PAT.Default value is 0.

	FixTeletextPts *typex.Bool `json:"fix_teletext_pts" flag:"-fix_teletext_pts,1"` //1
	//Override teletext packet PTS and DTS values with the timestamps calculated from the PCR of the first program which the teletext stream is part of and is not discarded.Default value is 1, set this option to 0 if you want your teletext packet PTS and DTS values untouched.

	TsPacketsize typex.Size `json:"ts_packetsize" flag:"-ts_packetsize"`
	//Output option carrying the raw packet size in bytes.Show the detected raw packet size, cannot be set by the user.

	ScanAllPmts *typex.Bool `json:"scan_all_pmts" flag:"-scan_all_pmts,-1"`
	//Scan and combine all PMTs.The value is an integer with value from -1 to 1 (-1 means automatic setting, 1 means enabled, 0 means disabled).Default value is -1.

	MergePmtVersions typex.Level `json:"merge_pmt_versions" flag:"-merge_pmt_versions"`
	//Re-use existing streams when a PMT’s version is updated and elementary streams move to different PIDs.Default value is 0.

	MaxPacketSize typex.Size `json:"max_packet_size" flag:"-max_packet_size"`
	//Set maximum size, in bytes, of packet emitted by the demuxer.Payloads above this size are split across multiple packets.Range is 1 to INT_MAX/2. Default is 204800 bytes.
}
