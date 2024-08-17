package outputdevicex

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// // CACA 4.3 caca
type CACA struct {
	WindowTitle typex.FileName `json:"window_title" flag:"-window_title"`
	//Set the CACA window title, if not specified default to the filename specified for the output device.

	WindowSize flagx.VideoSize `json:"window_size" flag:"-window_size"`
	//Set the CACA window size, can be a string of the form widthxheight or a video size abbreviation. If not specified it defaults to the size of the input video.

	Driver typex.Driver `json:"driver" flag:"-driver"`
	//Set display driver.

	Algorithm typex.Algorithm `json:"algorithm" flag:"-algorithm"`
	//Set dithering algorithm. Dithering is necessary because the picture being rendered has usually far more colours than the available palette. The accepted values are listed with -list_dither algorithms.

	Antialias typex.Antialias `json:"antialias" flag:"-antialias"`
	//Set antialias method. Antialiasing smoothens the rendered image and avoids the commonly seen staircase effect. The accepted values are listed with -list_dither antialiases.

	Charset typex.Charset `json:"charset" flag:"-charset"`
	//Set which characters are going to be used when rendering text. The accepted values are listed with -list_dither charsets.

	Color typex.Color `json:"color" flag:"-color"`
	//Set color to be used when rendering text. The accepted values are listed with -list_dither colors.

	ListDrivers typex.BoolStr `json:"list_drivers" flag:"-list_drivers"`
	//If set to true, print a list of available drivers and exit.

	ListDither ListDither `json:"list_dither" flag:"-list_dither"`
	//List available dither options related to the argument. The argument must be one of algorithms, antialiases, charsets, colors.
}

type ListDither = typex.String

const (
	ListDither_algorithms  ListDither = "algorithms"
	ListDither_antialiases ListDither = "antialiases"
	ListDither_charsets    ListDither = "charsets"
	ListDither_colors      ListDither = "colors"
)

/*


The following command shows the ffmpeg output is an CACA window, forcing its size to 80x25:
	ffmpeg -i INPUT -c:v rawvideo -pix_fmt rgb24 -window_size 80x25 -f caca -
Show the list of available drivers and exit:
	ffmpeg -i INPUT -pix_fmt rgb24 -f caca -list_drivers true -
Show the list of available dither colors and exit:
	ffmpeg -i INPUT -pix_fmt rgb24 -f caca -list_dither colors -
*/
