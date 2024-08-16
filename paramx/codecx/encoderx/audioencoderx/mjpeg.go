package audioencoderx

// MJPEG 8.14 mjpeg
type MJPEG struct {
	Huffman Huffman `json:"huffman" flag:"-huffman"`
	//Set the huffman encoding strategy. Possible values:

}
type Huffman = typex.String

const (
	Huffman_default Huffman = "default"
	//Use the default huffman tables. This is the default strategy.

	Huffman_optimal Huffman = "optimal"
	//Compute and use optimal huffman tables.
)
