package outputdevicex

// ALSA 4.1 alsa
type ALSA struct {
	//
}

/*
Play a file on default ALSA device:
ffmpeg -i INPUT -f alsa default
Play a file on soundcard 1, audio device 7:
ffmpeg -i INPUT -f alsa hw:1,7
*/
