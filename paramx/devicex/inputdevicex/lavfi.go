package inputdevicex

import "github.com/elizabevil/ffmpegx/paramx/typex"

// LAVFI 3.12 lavfi
type LAVFI struct {
	Graph typex.Graph `json:"graph" flag:"-graph"`
	//Specify the filtergraph to use as input. Each video open output must be labelled by a unique string of the form "outN", where N is a number starting from 0 corresponding to the mapped input stream generated by the device. The first unlabelled output is automatically assigned to the "out0" label, but all the others need to be specified explicitly.

	//The suffix "+subcc" can be appended to the output label to create an extra stream with the closed captions packets attached to that output (experimental; only for EIA-608 / CEA-708 for now). The subcc streams are created after all the normal streams, in the order of the corresponding stream. For example, if there is "out19+subcc", "out7+subcc" and up to "out42", the stream #43 is subcc for stream #7 and stream #44 is subcc for stream #19.

	//If not specified defaults to the filename specified for the input device.

	GraphFile typex.Filename `json:"graph_file" flag:"-graph_file"`
	//Set the filename of the filtergraph to be read and sent to the other filters. Syntax of the filtergraph is the same as the one specified by the option graph.

	Dumpgraph string `json:"dumpgraph" flag:"-dumpgraph"`
	//Dump graph to stderr.

}

/*


ffplay -dumpgraph 1 -f lavfi "
color=s=100x100:c=red  [l];
color=s=100x100:c=blue [r];
nullsrc=s=200x100, zmq [bg];
[bg][l]   overlay     [bg+l];
[bg+l][r] overlay@@my=x=100 "

Create a color video stream and play it back with ffplay:
ffplay -f lavfi -graph "color=c=pink [out0]" dummy
As the previous example, but use filename for specifying the graph description, and omit the "out0" label:
ffplay -f lavfi color=c=pink
Create three different video test filtered sources and play them:
ffplay -f lavfi -graph "testsrc [out0]; testsrc,hflip [out1]; testsrc,negate [out2]" test3
Read an audio stream from a file using the amovie source and play it back with ffplay:
ffplay -f lavfi "amovie=test.wav"
Read an audio stream and a video stream and play it back with ffplay:
ffplay -f lavfi "movie=test.avi[out0];amovie=test.wav[out1]"
Dump decoded frames to images and Closed Captions to an RCWT backup:
ffmpeg -f lavfi -i "movie=test.ts[out0+subcc]" -map v frame%08d.png -map s -c copy -f rcwt subcc.bin

*/
