package videoencoderx

// VBN 9.29 vbn
type VBN struct {
	Format VbnFmt `json:"format" flag:"-format"`
	//Sets the texture compression used by the VBN file. Can be dxt1, dxt5 or raw. Default is dxt5.
}
type VbnFmt = typex.String

const (
	VbnFmt_dxt1 VbnFmt = "dxt1"
	VbnFmt_dxt5 VbnFmt = "dxt5"
	VbnFmt_raw  VbnFmt = "raw"
)
