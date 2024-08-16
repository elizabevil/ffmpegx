package flagx

type UseMfraFor = Flag

const (
	UseMfraFor_auto UseMfraFor = "auto"
	//Auto-detect whether to set mfra timestamps as PTS or DTS (default)

	UseMfraFor_dts UseMfraFor = "dts"
	//Set mfra timestamps as DTS

	UseMfraFor_pts UseMfraFor = "pts"
	//Set mfra timestamps as PTS

	UseMfraFor_0 UseMfraFor = "0"
	//Don’t use mfra box to set timestamps
)
