package chip8

// typedef unsigned char Uint8;
// void AudioSpecCallback(void *userdata, Uint8 *stream, int len);
import "C"
import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"math"
	"reflect"
	"unsafe"
)

type Buzzer struct {
	chip *CHIP8
	spec *sdl.AudioSpec
}

const (
	toneHz   = 440
	sampleHz = 720
	dPhase   = 2 * math.Pi * toneHz / sampleHz
)

func NewBuzzer() *Buzzer {
	_spec := &sdl.AudioSpec{
		Freq:     sampleHz,
		Format:   sdl.AUDIO_U8,
		Channels: 2,
		Samples:  sampleHz * 0.3,
		Callback: sdl.AudioCallback(C.AudioSpecCallback),
	}
	return &Buzzer{
		spec: _spec,
	}
}

//export AudioSpecCallback
func AudioSpecCallback(userdata unsafe.Pointer, stream *C.Uint8, length C.int) {
	n := int(length)
	hdr := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(stream)), Len: n, Cap: n}
	buf := *(*[]C.Uint8)(unsafe.Pointer(&hdr))

	var phase float64
	for i := 0; i < n; i += 2 {
		phase += dPhase
		sample := C.Uint8((math.Sin(phase) + 0.999999) * 128)
		buf[i] = sample
		buf[i+1] = sample
	}
}

func (buzzer *Buzzer) beep() {
	if err := sdl.OpenAudio(buzzer.spec, nil); err != nil {
		log.Println(err)
		return
	}
	sdl.PauseAudio(false)
	sdl.Delay(300)
	sdl.CloseAudio()
}
