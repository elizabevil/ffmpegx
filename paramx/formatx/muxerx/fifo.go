package muxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// // 21.33 fifo FIFO
type FIFO struct {
	AttemptRecovery typex.UBool `json:"attempt_recovery" flag:"-attempt_recovery"`
	//If failure occurs, attempt to recover the output. This is especially useful when used with network output, since it makes it possible to restart streaming transparently. By default this option is set to false.

	DropPktsOnOverflow typex.UBool `json:"drop_pkts_on_overflow" flag:"-drop_pkts_on_overflow"`
	//If set to true, in case the fifo queue fills up, packets will be dropped rather than blocking the encoder. This makes it possible to continue streaming without delaying the input, at the cost of omitting part of the stream. By default this option is set to false, so in such cases the encoder will be blocked until the muxer processes some of the packets and none of them is lost.

	FifoFormat string `json:"fifo_format" flag:"-fifo_format"`
	//Specify the format name. Useful if it cannot be guessed from the output name suffix.

	FormatOpts string `json:"format_opts" flag:"-format_opts"`
	//Specify format options for the underlying muxer. Muxer options can be specified as a list of key=value pairs separated by ’:’.

	MaxRecoveryAttempts typex.Count `json:"max_recovery_attempts" flag:"-max_recovery_attempts"`
	//Set maximum number of successive unsuccessful recovery attempts after which the output fails permanently. By default this option is set to 0 (unlimited).

	QueueSize typex.Size `json:"queue_size" flag:"-queue_size"`
	//Specify size of the queue as a number of packets. Default value is 60.

	RecoverAnyError typex.UBool `json:"recover_any_error" flag:"-recover_any_error"`
	//If set to true, recovery will be attempted regardless of type of the error causing the failure. By default this option is set to false and in case of certain (usually permanent) errors the recovery is not attempted even when the attempt_recovery option is set to true.

	RecoveryWaitStreamtime typex.UBool `json:"recovery_wait_streamtime" flag:"-recovery_wait_streamtime"`
	//If set to false, the real time is used when waiting for the recovery attempt (i.e. the recovery will be attempted after the time specified by the recovery_wait_time option).

	//If set to true, the time of the processed stream is taken into account instead (i.e. the recovery will be attempted after discarding the packets corresponding to the recovery_wait_time option).

	//By default this option is set to false.

	RecoveryWaitTime typex.SecondI `json:"recovery_wait_time" flag:"-recovery_wait_time"`
	//Specify waiting time in seconds before the next recovery attempt after previous unsuccessful recovery attempt. Default value is 5.

	RestartWithKeyframe typex.UBool `json:"restart_with_keyframe" flag:"-restart_with_keyframe"`
	//Specify whether to wait for the keyframe after recovering from queue overflow or failure. This option is set to false by default.

	Timeshift typex.TimeShift `json:"timeshift" flag:"-timeshift"`
	//Buffer the specified amount of packets and delay writing the output. Note that the value of the queue_size option must be big enough to store the packets for timeshift. At the end of the input the fifo buffer is flushed at realtime speed.

}
