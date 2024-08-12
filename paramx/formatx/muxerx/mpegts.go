package muxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

const M2TS_PMT_PID = 0x0100
const M2TS_PCR_PID = 0x1001
const M2TS_VIDEO_PID = 0x1011
const M2TS_AUDIO_START_PID = 0x1100
const M2TS_PGSSUB_START_PID = 0x1200
const M2TS_TEXTSUB_PID = 0x1800
const M2TS_SECONDARY_AUDIO_START_PID = 0x1A00
const M2TS_SECONDARY_VIDEO_START_PID = 0x1B00
const FIRST_OTHER_PID = 0x0020
const LAST_OTHER_PID = 0x1FFA

// MPEGTS 21.62 mpegts MPEGTS
// /*This muxer implements ISO 13818-1 and part of ETSI EN 300 468.
// The recognized metadata settings in mpegts muxer are service_provider and service_name. If they are not set the default for service_provider is ‘FFmpeg’ and the default for service_name is ‘Service01’.*/
type MPEGTS struct {
	MpegtsTransportStreamId typex.UI16 `json:"mpegts_transport_stream_id" flag:"-mpegts_transport_stream_id"`
	//Set the ‘transport_stream_id’. This identifies a transponder in DVB. Default is 0x0001.

	MpegtsOriginalNetworkId typex.UI16 `json:"mpegts_original_network_id" flag:"-mpegts_original_network_id"`
	//Set the ‘original_network_id’. This is unique identifier of a network in DVB. Its main use is in the unique identification of a service through the path ‘Original_Network_ID, Transport_Stream_ID’. Default is 0x0001.

	MpegtsServiceId typex.UI16 `json:"mpegts_service_id" flag:"-mpegts_service_id"`
	//Set the ‘service_id’, also known as program in DVB. Default is 0x0001.

	MpegtsServiceType typex.UI8  `json:"mpegts_service_type" flag:"-mpegts_service_type"`
	MpegtsPmtStartPid typex.UI16 `json:"mpegts_pmt_start_pid" flag:"-mpegts_pmt_start_pid"`
	//Set the first PID for PMTs. Default is 0x1000, minimum is 0x0020, maximum is 0x1ffa. This option has no effect in m2ts mode where the PMT PID is fixed 0x0100.

	MpegtsStartPid typex.UI16 `json:"mpegts_start_pid" flag:"-mpegts_start_pid"`
	//Set the first PID for elementary streams. Default is 0x0100, minimum is 0x0020, maximum is 0x1ffa. This option has no effect in m2ts mode where the elementary stream PIDs are fixed.

	MpegtsM2TsMode typex.UBool `json:"mpegts_m2ts_mode" flag:"-mpegts_m2ts_mode"`
	//Enable m2ts mode if set to 1. Default value is -1 which disables m2ts mode.

	Muxrate typex.RateI `json:"muxrate" flag:"-muxrate"`
	//Set a constant muxrate. Default is VBR.

	PesPayloadSize typex.Size `json:"pes_payload_size" flag:"-pes_payload_size"`
	//Set minimum PES packet payload in bytes. Default is 2930.

	MpegtsFlags  MpegtsFlags `json:"mpegts_flags" flag:"-mpegts_flags"`
	MpegtsCopyts typex.Bool  `json:"mpegts_copyts" flag:"-mpegts_copyts"`
	//Preserve original timestamps, if value is set to 1. Default value is -1, which results in shifting timestamps so that they start from 0.

	OmitVideoPesLength typex.UBool `json:"omit_video_pes_length" flag:"-omit_video_pes_length"`
	//Omit the PES packet length for video packets. Default is 1 (true).

	PcrPeriod typex.MillisecondI `json:"pcr_period" flag:"-pcr_period"`
	//Override the default PCR retransmission time in milliseconds. Default is -1 which means that the PCR interval will be determined automatically: 20 ms is used for CBR streams, the highest multiple of the frame duration which is less than 100 ms is used for VBR streams.

	PatPeriod typex.SecondF `json:"pat_period" flag:"-pat_period"`
	//Maximum time in seconds between PAT/PMT tables. Default is 0.1.

	SdtPeriod typex.SecondF `json:"sdt_period" flag:"-sdt_period"`
	//Maximum time in seconds between SDT tables. Default is 0.5.

	NitPeriod typex.SecondF `json:"nit_period" flag:"-nit_period"`
	//Maximum time in seconds between NIT tables. Default is 0.5.

	TablesVersion typex.Level `json:"tables_version" flag:"-tables_version"`
}

type MpegtsFlags = typex.String

const (
	MpegtsFlags_resend_headers        MpegtsFlags = "resend_headers"
	MpegtsFlags_latm                  MpegtsFlags = "latm"
	MpegtsFlags_pat_pmt_at_frames     MpegtsFlags = "pat_pmt_at_frames"
	MpegtsFlags_system_b              MpegtsFlags = "system_b"
	MpegtsFlags_initial_discontinuity MpegtsFlags = "initial_discontinuity"
	MpegtsFlags_nit                   MpegtsFlags = "nit"
	MpegtsFlags_omit_rai              MpegtsFlags = "omit_rai"
)
