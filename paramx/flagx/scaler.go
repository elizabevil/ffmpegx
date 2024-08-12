package flagx

//Scaler Options

// sws_flags

type SwsFlags = Flag

const (
	SwsFlags_fast_bilinear   SwsFlags = "fast_bilinear"
	SwsFlags_bilinear        SwsFlags = "bilinear"
	SwsFlags_bicubic         SwsFlags = "bicubic"
	SwsFlags_experimental    SwsFlags = "experimental"
	SwsFlags_neighbor        SwsFlags = "neighbor"
	SwsFlags_area            SwsFlags = "area"
	SwsFlags_bicublin        SwsFlags = "bicublin"
	SwsFlags_gauss           SwsFlags = "gauss"
	SwsFlags_sinc            SwsFlags = "sinc"
	SwsFlags_lanczos         SwsFlags = "lanczos"
	SwsFlags_spline          SwsFlags = "spline"
	SwsFlags_print_info      SwsFlags = "print_info"
	SwsFlags_accurate_rnd    SwsFlags = "accurate_rnd"
	SwsFlags_full_chroma_int SwsFlags = "full_chroma_int"
	SwsFlags_full_chroma_inp SwsFlags = "full_chroma_inp"
	SwsFlags_bitexac         SwsFlags = "bitexac"
)

type SwsDither = Flag

const (
	SwsDither_auto     SwsDither = "auto"
	SwsDither_none     SwsDither = "none"
	SwsDither_bayer    SwsDither = "bayer"
	SwsDither_ed       SwsDither = "ed"
	SwsDither_a_dither SwsDither = "a_dither"
	SwsDither_x_dithe  SwsDither = "x_dithe"
)
