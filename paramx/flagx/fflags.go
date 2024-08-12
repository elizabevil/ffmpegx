package flagx

type Fflags = Flag

const (
	Fflags_discardcorrupt Fflags = "discardcorrupt"
	//Discard corrupted packets.

	Fflags_fastseek Fflags = "fastseek"
	//Enable fast, but inaccurate seeks for some formats.

	Fflags_genpts Fflags = "genpts"
	//Generate missing PTS if DTS is present.

	Fflags_igndts Fflags = "igndts"
	//Ignore DTS if PTS is also set. In case the PTS is set, the DTS value is set to NOPTS. This is ignored when the nofillin flag is set.

	Fflags_ignidx Fflags = "ignidx"
	//Ignore index.

	Fflags_nobuffer Fflags = "nobuffer"
	//Reduce the latency introduced by buffering during initial input streams analysis.

	Fflags_nofillin Fflags = "nofillin"
	//Do not fill in missing values in packet fields that can be exactly calculated.

	Fflags_noparse Fflags = "noparse"
	//Disable AVParsers, this needs +nofillin too.

	Fflags_sortdts Fflags = "sortdts"
	//Try to interleave output packets by DTS. At present, available only for AVIs with an index.

	Fflags_autobsf Fflags = "autobsf"
	//Automatically apply bitstream filters as required by the output format. Enabled by default.

	Fflags_bitexact Fflags = "bitexact"
	//Only write platform-, build- and time-independent data. This ensures that file and data checksums are reproducible and match between platforms. Its primary use is for regression testing.

	Fflags_flush_packets Fflags = "flush_packets"
	//Write out packets immediately.

	Fflags_shortest Fflags = "shortest"
	//Stop muxing at the end of the shortest stream. It may be needed to increase max_interleave_delta to avoid flushing the longer streams before EOF.
)
