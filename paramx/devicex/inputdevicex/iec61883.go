package inputdevicex

import "github.com/elizabevil/ffmpegx/paramx/typex"

// IEC61883 3.9 iec61883
type IEC61883 struct {
	Dvtype Dvtype `json:"dvtype" flag:"-dvtype"`
	//Override autodetection of DV/HDV. This should only be used if auto detection does not work, or if usage of a different device type should be prohibited. Treating a DV device as HDV (or vice versa) will not work and result in undefined behavior. The values auto, dv and hdv are supported.

	Dvbuffer typex.Size `json:"dvbuffer" flag:"-dvbuffer"`
	//Set maximum size of buffer for incoming data, in frames. For DV, this is an exact value. For HDV, it is not frame exact, since HDV does not have a fixed frame size.

	Dvguid string `json:"dvguid" flag:"-dvguid"`
	//Select the capture device by specifying its GUID. Capturing will only be performed from the specified device and fails if no device with the given GUID is found. This is useful to select the input if multiple devices are connected at the same time. Look at /sys/bus/firewire/devices to find out the GUIDs.

}

/*
ffmpeg -f iec61883 -i auto -dvbuffer 100000 out.mpg
*/

type Dvtype = typex.String

const (
	Dvtype_auto Dvtype = "auto"
	Dvtype_dv   Dvtype = "dv"
	Dvtype_hdv  Dvtype = "hdv"
)
