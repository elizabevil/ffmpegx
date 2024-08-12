# # FFmpex

> ###### go control ffmpeg program

Encapsulates commonly used methods and parameters

> All From =>FFmpegDoc https://www.ffmpeg.org/documentation.html

---

### Usage notes

> go get github.com/elizabevil/ffmpegx@version

- args

```go
//ffmpeg [global_options] {[input_file_options] -i input_url} ... {[output_file_options] output_url} ...

```

- Make args

```go
//make cmd args
//https://www.ffmpeg.org/documentation.html
//ffmpeg [global_options] {[input_file_options] -i input_url} ... {[output_file_options] output_url} ...


hls := muxerx.HLS{
HlsInitTime:        1,
HlsTime:            10,
HlsDeleteThreshold: 2,
HlsListSize:        &typex.ZeroUN, //default 5 --> force 0
HlsSegmentFilename: "%Y%m%d%H%M%S.ts",
}
fmt.Println(hls.Args())// show args
fmt.Println(parsex.DefaultParser.ParamParse(hls))// or this

//[-hls_segment_filename %Y%m%d%H%M%S.ts -hls_delete_threshold 2 -hls_list_size 0 -hls_time 10 -hls_init_time 1]


param := paramx.Param{}
param.GlobalHandle(func(input *optionx.Global) {
input.Overwrite = true
}).InputHandle(func(input *optionx.Input) {
input.Inputs = []string{InputVedio}
}).OutputHandle(func(output *optionx.Output) {
output.Outputs = []string{OutputVideo}
}).CommonHandle(paramx.PositionInput, optionx.Common{
An: true,
}).CommonHandle(paramx.PositionOutput, optionx.Common{
F: "hls",
CodecSpecifier: []typex.StreamSpecifier{
{"v", "h264x"},
},
}).CommonHandle(paramx.PositionOutput, optionx.Common{ //overwrite
F: "flv",
T: 4,
})
fmt.Println(param.Args())
//[-y -an -i ./file/input.mp4 -t 4 -f flv ./file/output.mp4]

//===============
input := optionx.Input{Input: InputVedio}
args := paramx.AnyArgs{"-c": "copy"}
output := optionx.Output{Output: OutputVideoHls}
?? := impl interfacex.IArg
argInterface := paramx.BuildIArgInterface(input, args, ??, output) // Concat args
fmt.Println(argInterface.Args())
//[-i ./file/input.mp4 -c copy ?? output.mp4]

```

- YourParse Args

```go
// you can impl Arg interface
type Arg interface {
    Args() typex.Args
}

type YourParse struct {// your paesr method
}

func (y YourParse) ParamParse(input any) typex.Args {
	return []string{"-c", "copy"}
}

func (y YourParse) ParamItemType(of reflect.Value) (string, bool) {
	return "c", false
}
//================
type YourHls muxerx.HLS // your type
func (r YourHls) Args() typex.Args {
	return YourParse{}.ParamParse(r)
}
fmt.Println(YourHls{}.Args()) //see args
//paramx.BuildIArgInterface(input, args, hls, output)
paramx.BuildIArgInterface(input, args, YourHls{}, output) // complete args
```



- Run

```go
//==or== 
transcoderx.Cmd(bin string, args interfacex.IArg)   //*exec.Cmd 

// makeTranscoder
transcode := transcoderx.NewTranscoder(transcoderx.NewConfig())// with default config
transcode.Args = makeParams() // set ffmpeg/ffplay args
transcode.Cmd()   //*exec.Cmd 

```



- Get Metadata

```go
metadata, err := transcoderx.Metadata(ffprobeBinPath, InputVedio) // use ffprobeBin to get InputVedio's Metadata 
fmt.Println(metadata.Format)  // println Format
fmt.Println(metadata.Streams) // println Streams

transcoderx.Metadata()	//default
transcoderx.MetadataWithArgs() // with args
//args
generic := optionx.Generic{HideBanner: true}
fprobe := optionx.FFprobe{
I:                InputVedio,
PrintFormat:      optionx.OutputFormat_json,
ShowError:        true,
ShowFormat:       true,
ShowStreams:      true,
}
transcoderx.Debug = true
metadata, err := transcoderx.MetadataWithArgs(FfprobeBin, paramx.BuildIArgInterface(generic, fprobe), transcoderx.JsonUnmarshal, func(out []byte, err []byte) {
//bytes
})

if err != nil {
fmt.Println("MetadataX", err)
return
}
fmt.Println(metadata.Format)
fmt.Println(metadata.Streams) // println Streams
//OutputFormat_json: JsonUnmarshal{}
//OutputFormat_xml : XmlUnmarshal{}

```



- Audio Channel_layout

```go
streams := metadata.StreamWithType(AVMEDIA_TYPE_AUDIO) // if this stream `s codec_type is audio
layout := streams[0].GetChannelLayout() //{stereo {1 2 3}}
fmt.Println(metadatax.ChannelNames[layout.Layout.Mask]) //{LFE low frequency}
```



- Use Pipeline  Run

```go

transcoderx.Debug = true // show args
duration := 1000 * time.Millisecond
killDuration := duration * 100

background, cc := context.WithTimeout(context.Background(), duration) // ctx
defer cc()

//Ph:= impl interfacex.ProgressHandle // set your self method to parse progress
runFunc, err := transcoderx.Pipeline(FfmpegBin, Ph, makeParamsToHls(), func(cmd *exec.Cmd) {
cmd.Dir = "/tmp"
})
if err != nil {
return
}
err = runFunc(background, func(process *os.Process) {
go func() {
time.Sleep(killDuration) // after killDuration 
process.Kill()
}()
}, func(progress metadatax.Progress) {
//progress:: {1 0.0 -1.0 N/A 00:00:00.03 N/A  N/A      }
//progress:: {16 0.0 -1.0 N/A 00:00:00.53 N/A  1.06x      }
fmt.Println("progress::", progress)
})

```

- some types fun

```go
fmt.Println(typex.TimeZero) //0000-01-01 00:00:00 +0000 UTC
duration := 3*time.Second + 500*time.Millisecond //3.5s
fmt.Println(typex.TimeDuration(duration, time.Millisecond)) //3500 Millisecond
fmt.Println(typex.TimeDurationSecond(duration))//3.5s
fmt.Println(typex.TimeDurationSecondI(duration))//3 s
fmt.Println(typex.TimeDurationParseSecondF("200ms"))//0.2s
fmt.Println(typex.TimeDurationParseSecondF("200000us"))//0.2s
fmt.Println(typex.TimeDurationParseSecondF("00:01:01"))//61s
fmt.Println(typex.TimeDurationParseSecondF("61s"))//61s


fmt.Println(typex.NewVideoSize(1024, 720))//720x1024
fmt.Println(typex.NewRatio(1024, 720))//720:1024
fmt.Println(typex.NewRate(25, 1))  //25/1
```



- More examples ==>see examples dir 