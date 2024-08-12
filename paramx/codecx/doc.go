package codecx

//https://ffmpeg.org/ffmpeg-all.html#Codec-Options
//
//Possible values:
//
//‘mv4’
//
//‘qpel’
//
//‘loop’
//
//‘qscale’
//
//‘pass1’
//
//‘pass2’
//
//‘gray’
//Only decode/encode grayscale.
//
//‘psnr’
//‘truncated’
//Input bitstream might be randomly truncated.
//
//‘drop_changed’
//Don’t output frames whose parameters differ from first decoded frame in stream. Error AVERROR_INPUT_CHANGED is returned when a frame is dropped.
//
//‘ildct’
//
//‘low_delay’
//Force low delay.
//
//‘global_header’
//Place global headers in extradata instead of every keyframe.
//
//‘bitexact’
//
//‘aic’
//Apply H263 advanced intra coding / mpeg4 ac prediction.
//
//‘ilme’
//Apply interlaced motion estimation.
//
//‘cgop’
//
//‘output_corrupt’
//Output even potentially corrupted frames.
//
//time_base rational number
//
//
//
//
//Each submitted frame except the last must contain exactly frame_size samples per channel. May be 0 when the codec has CODEC_CAP_VARIABLE_FRAME_SIZE set, in that case the frame size is not restricted. It is set by some decoders to indicate constant frame size.
//
//frame_number int
//delay int
//
//
//
//
//
//Must be an int between -1 and 16. 0 means that B-frames are disabled. If a value of -1 is used, it will choose an automatic value depending on the encoder.
//
//Default value is 0.
//
//codec_tag int
//Workaround not auto detected encoder bugs.
//
//Possible values:
//
//‘autodetect’
//‘xvid_ilace’
//
//‘ump4’
//(autodetected if fourcc==UMP4)
//
//‘no_padding’
//
//‘amv’
//‘qpel_chroma’
//‘std_qpel’
//
//‘qpel_chroma2’
//‘direct_blocksize’
//
//‘edge’
//
//‘hpel_chroma’
//‘dc_clip’
//‘ms’
//Workaround various bugs in microsoft broken decoders.
//
//‘trunc’
//trancated frames
//
//Specify how strictly to follow the standards.
//
//Possible values:
//
//‘very’
//strictly conform to an older more strict version of the spec or reference software
//
//‘strict’
//strictly conform to all the things in the spec no matter what consequences
//
//‘normal’
//‘unofficial’
//allow unofficial extensions
//
//‘experimental’
//
//Possible values:
//
//‘crccheck’
//verify embedded CRCs
//
//‘bitstream’
//detect bitstream specification deviations
//
//‘buffer’
//detect improper bitstream length
//
//‘explode’
//abort decoding on minor error detection
//
//‘ignore_err’
//ignore decoding errors, and continue decoding. This is useful if you want to analyze the content of a video and thus want everything to be decoded no matter what. This option will not result in a video that is pleasing to watch in case of errors.
//
//‘careful’
//consider things that violate the spec and have not been seen in the wild as errors
//
//‘compliant’
//consider all spec non compliancies as errors
//
//‘aggressive’
//consider things that a sane encoder should not do as an error
//
//has_b_frames int
//block_align int
//rc_override_count int
//
//
//
//Possible values:
//
//‘auto’
//
//‘fastint’
//fast int
//
//‘int’
//accurate int
//
//‘mmx’
//‘altivec’
//‘faan’
//float32ing point AAN DCT
//
//Compress bright areas stronger than medium ones.
//
//Compress dark areas stronger than medium ones.
//
//Select IDCT implementation.
//
//Possible values:
//
//‘auto’
//‘int’
//‘simple’
//‘simplemmx’
//‘simpleauto’
//Automatically pick a IDCT compatible with the simple one
//
//‘arm’
//‘altivec’
//‘sh4’
//‘simplearm’
//‘simplearmv5te’
//‘simplearmv6’
//‘simpleneon’
//‘xvid’
//‘faani’
//float32ing point AAN IDCT
//
//slice_count int
//Possible values:
//
//‘guess_mvs’
//
//‘deblock’
//use strong deblock filter for damaged MBs
//
//‘favor_inter’
//favor predicting from the previous frame instead of the current
//
//bits_per_coded_sample int
//Print specific debug info.
//
//Possible values:
//
//‘pict’
//picture info
//
//‘rc’
//rate control
//
//‘bitstream’
//‘mb_type’
//
//‘qp’
//
//‘dct_coeff’
//‘green_metadata’
//display complexity metadata for the upcoming frame, GoP or for a given duration.
//
//‘skip’
//‘startcode’
//‘er’
//error recognition
//
//‘mmco’
//
//‘bugs’
//‘buffers’
//picture buffer allocations
//
//‘thread_ops’
//threading operations
//
//‘nomc’
//skip motion compensation
//
//Possible values:
//
//‘sad’
//
//‘sse’
//sum of squared errors
//
//‘satd’
//sum of absolute Hadamard transformed differences
//
//‘dct’
//sum of absolute DCT transformed differences
//
//‘psnr’
//
//‘bit’
//number of bits needed for the block
//
//‘rd’
//rate distortion optimal, slow
//
//‘zero’
//0
//
//‘vsad’
//sum of absolute vertical differences
//
//‘vsse’
//sum of squared vertical differences
//
//‘nsse’
//noise preserving sum of squared differences
//
//‘w53’
//5/3 wavelet, only used in snow
//
//‘w97’
//9/7 wavelet, only used in snow
//
//‘dctmax’
//‘chroma’
//Possible values:
//
//‘sad’
//
//‘sse’
//sum of squared errors
//
//‘satd’
//sum of absolute Hadamard transformed differences
//
//‘dct’
//sum of absolute DCT transformed differences
//
//‘psnr’
//
//‘bit’
//number of bits needed for the block
//
//‘rd’
//rate distortion optimal, slow
//
//‘zero’
//0
//
//‘vsad’
//sum of absolute vertical differences
//
//‘vsse’
//sum of squared vertical differences
//
//‘nsse’
//noise preserving sum of squared differences
//
//‘w53’
//5/3 wavelet, only used in snow
//
//‘w97’
//9/7 wavelet, only used in snow
//
//‘dctmax’
//‘chroma’
//Possible values:
//
//‘sad’
//
//‘sse’
//sum of squared errors
//
//‘satd’
//sum of absolute Hadamard transformed differences
//
//‘dct’
//sum of absolute DCT transformed differences
//
//‘psnr’
//
//‘bit’
//number of bits needed for the block
//
//‘rd’
//rate distortion optimal, slow
//
//‘zero’
//0
//
//‘vsad’
//sum of absolute vertical differences
//
//‘vsse’
//sum of squared vertical differences
//
//‘nsse’
//noise preserving sum of squared differences
//
//‘w53’
//5/3 wavelet, only used in snow
//
//‘w97’
//9/7 wavelet, only used in snow
//
//‘dctmax’
//‘chroma’
//Possible values:
//
//‘sad’
//
//‘sse’
//sum of squared errors
//
//‘satd’
//sum of absolute Hadamard transformed differences
//
//‘dct’
//sum of absolute DCT transformed differences
//
//‘psnr’
//
//‘bit’
//number of bits needed for the block
//
//‘rd’
//rate distortion optimal, slow
//
//‘zero’
//0
//
//‘vsad’
//sum of absolute vertical differences
//
//‘vsse’
//sum of squared vertical differences
//
//‘nsse’
//noise preserving sum of squared differences
//
//‘w53’
//5/3 wavelet, only used in snow
//
//‘w97’
//9/7 wavelet, only used in snow
//
//‘dctmax’
//‘chroma’
//‘(1024, INT_MAX)’
//full motion estimation(slowest)
//
//‘(768, 1024]’
//umh motion estimation
//
//‘(512, 768]’
//hex motion estimation
//
//‘(256, 512]’
//l2s diamond motion estimation
//
//‘[2,256]’
//var diamond motion estimation
//
//‘(-1, 2)’
//small diamond motion estimation
//
//‘-1’
//funny diamond motion estimation
//
//‘(INT_MIN, -1)’
//sab diamond motion estimation
//
//Possible values:
//
//‘sad’
//
//‘sse’
//sum of squared errors
//
//‘satd’
//sum of absolute Hadamard transformed differences
//
//‘dct’
//sum of absolute DCT transformed differences
//
//‘psnr’
//
//‘bit’
//number of bits needed for the block
//
//‘rd’
//rate distortion optimal, slow
//
//‘zero’
//0
//
//‘vsad’
//sum of absolute vertical differences
//
//‘vsse’
//sum of squared vertical differences
//
//‘nsse’
//noise preserving sum of squared differences
//
//‘w53’
//5/3 wavelet, only used in snow
//
//‘w97’
//9/7 wavelet, only used in snow
//
//‘dctmax’
//‘chroma’
//
//slice_flags int
//
//Possible values:
//
//‘simple’
//
//‘bits’
//use fewest bits
//
//‘rd’
//use best rate distortion
//
//Possible values:
//
//‘fast’
//Allow non spec compliant speedup tricks.
//
//‘noout’
//Skip bitstream encoding.
//
//‘ignorecrop’
//Ignore cropping information from sps.
//
//‘local_header’
//Place global headers at every keyframe instead of in extradata.
//
//‘chunks’
//Frame data might be split into multiple chunks.
//
//‘showall’
//Show all frames before the first keyframe.
//
//‘export_mvs’
//
//‘skip_manual’
//Do not skip samples and export skip information as frame side data.
//
//‘ass_ro_flush_noop’
//Do not reset ASS ReadOrder field on flush.
//
//‘icc_profiles’
//Generate/parse embedded ICC profiles from/to colorimetry tags.
//
//Possible values:
//
//‘mvs’
//
//‘prft’
//
//‘venc_params’
//
//‘film_grain’
//
//Possible values:
//
//‘auto, 0’
//automatically select the number of threads to set
//
//Default value is ‘auto’.
//
//Possible values:
//
//‘unknown’
//Decoding at 1= 1/2, 2=1/4, 3=1/8 resolutions.
//
//
//
//Make decoder discard processing depending on the frame type selected by the option value.
//
//skip_loop_filter skips frame loop filtering, skip_idct skips frame IDCT/dequantization, skip_frame skips decoding.
//
//Possible values:
//
//‘none’
//Discard no frame.
//
//‘default’
//Discard useless frames like 0-sized frames.
//
//‘noref’
//Discard all non-reference frames.
//
//‘bidir’
//Discard all bidirectional frames.
//
//‘nokey’
//Discard all frames excepts keyframes.
//
//‘nointra’
//Discard all frames except I frames.
//
//‘all’
//Discard all frames.
//
//Default value is ‘default’.
//
//Refine the two motion vectors used in bidirectional macroblocks.
//
//bits_per_raw_sample int
//
//Possible values:
//
//‘bt709’
//BT.709
//
//‘bt470m’
//BT.470 M
//
//‘bt470bg’
//BT.470 BG
//
//‘smpte170m’
//SMPTE 170 M
//
//‘smpte240m’
//SMPTE 240 M
//
//‘film’
//Film
//
//‘bt2020’
//BT.2020
//
//‘smpte428’
//‘smpte428_1’
//SMPTE ST 428-1
//
//‘smpte431’
//SMPTE 431-2
//
//‘smpte432’
//SMPTE 432-1
//
//‘jedec-p22’
//JEDEC P22
//
//Possible values:
//
//‘bt709’
//BT.709
//
//‘gamma22’
//BT.470 M
//
//‘gamma28’
//BT.470 BG
//
//‘smpte170m’
//SMPTE 170 M
//
//‘smpte240m’
//SMPTE 240 M
//
//‘linear’
//Linear
//
//‘log’
//‘log100’
//Log
//
//‘log_sqrt’
//‘log316’
//Log square root
//
//‘iec61966_2_4’
//‘iec61966-2-4’
//IEC 61966-2-4
//
//‘bt1361’
//‘bt1361e’
//BT.1361
//
//‘iec61966_2_1’
//‘iec61966-2-1’
//IEC 61966-2-1
//
//‘bt2020_10’
//‘bt2020_10bit’
//BT.2020 - 10 bit
//
//‘bt2020_12’
//‘bt2020_12bit’
//BT.2020 - 12 bit
//
//‘smpte2084’
//SMPTE ST 2084
//
//‘smpte428’
//‘smpte428_1’
//SMPTE ST 428-1
//
//‘arib-std-b67’
//ARIB STD-B67
//
//Possible values:
//
//‘rgb’
//RGB
//
//‘bt709’
//BT.709
//
//‘fcc’
//FCC
//
//‘bt470bg’
//BT.470 BG
//
//‘smpte170m’
//SMPTE 170 M
//
//‘smpte240m’
//SMPTE 240 M
//
//‘ycocg’
//YCOCG
//
//‘bt2020nc’
//‘bt2020_ncl’
//BT.2020 NCL
//
//‘bt2020c’
//‘bt2020_cl’
//BT.2020 CL
//
//‘smpte2085’
//SMPTE 2085
//
//‘chroma-derived-nc’
//Chroma-derived NCL
//
//‘chroma-derived-c’
//Chroma-derived CL
//
//‘ictcp’
//ICtCp
//
//If used as input parameter, it serves as a hint to the decoder, which color_range the input has. Possible values:
//
//‘tv’
//‘mpeg’
//‘limited’
//
//‘pc’
//‘jpeg’
//‘full’
//
//Possible values:
//
//‘left’
//‘center’
//‘topleft’
//‘top’
//‘bottomleft’
//‘bottom’
//log_level_offset int
//Number of slices, used in parallelized encoding.
//
//Select which multithreading methods to use.
//
//
//Possible values:
//
//‘slice’
//Decoding more than one part of a single frame at once.
//
//Multithreading using slices works only when the video was encoded with slices.
//
//‘frame’
//Decoding more than one frame at once.
//
//Default value is ‘slice+frame’.
//
//Possible values:
//
//‘ma’
//Main Audio Service
//
//‘ef’
//Effects
//
//‘vi’
//Visually Impaired
//
//‘hi’
//Hearing Impaired
//
//‘di’
//Dialogue
//
//‘co’
//Commentary
//
//‘em’
//Emergency
//
//‘vo’
//Voice Over
//
//‘ka’
//Karaoke
//
//pkt_timebase rational number
//‘progressive’
//Progressive video
//
//‘tt’
//Interlaced video, top field coded and displayed first
//
//‘bb’
//Interlaced video, bottom field coded and displayed first
//
//‘tb’
//Interlaced video, top coded first, bottom displayed first
//
//‘bt’
//Interlaced video, bottom coded first, top displayed first
//
//
//"," separated list of allowed decoders. By default all are allowed.
//
//Separator used to separate the fields printed on the command line about the Stream parameters. For example, to separate the fields with newlines and indentation:
//
//ffprobe -dump_separator "
//"  -i ~/videos/matrixbench_mpeg2.mpg
//Maximum number of pixels per image. This value can be used to avoid out of memory failures due to large images.
//
//Enable cropping if cropping parameters are multiples of the required alignment for the left and top parameters. If the alignment is not met the cropping will be partially applied to maintain alignment. Default is 1 // (enabled). Note: The required alignment depends on if AV_CODEC_FLAG_UNALIGNED is set and the CPU. AV_CODEC_FLAG_UNALIGNED cannot be changed from the command line. Also hardware decoders will not apply left/top Cropping.
