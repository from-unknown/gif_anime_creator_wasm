// +build js,wasm
package gifanimecreator

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"math"
	"reflect"
	"syscall/js"
	"time"
	"unsafe"

	"github.com/soniakeys/quant/median"
)

type GifAnimeCreator struct {
	inBuf                  []uint8
	outBuf                 bytes.Buffer
	onImgLoadCb, initMemCb js.Callback
	sourceImg              image.Image

	console js.Value
	done    chan struct{}
}

func New() *GifAnimeCreator {
	return &GifAnimeCreator{
		console: js.Global().Get("console"),
		done:    make(chan struct{}),
	}
}

func (g *GifAnimeCreator) Start() {
	// Setup callbacks
	g.setupInitMemCb()
	js.Global().Set("initMem", g.initMemCb)

	g.setupOnImgLoadCb()
	js.Global().Set("loadImage", g.onImgLoadCb)

	<-g.done
	g.log("Shutting down app")
	g.onImgLoadCb.Release()
}

func (g *GifAnimeCreator) ConvertImage(argStartFlag string, argEndFlag string, argLoopFlag string) {
	var gifImage []*image.Paletted
	var delays []int
	var disposal []byte

	rect := g.sourceImg.Bounds()
	for i := 0; i < 10; i++ {
		var drawXBound int
		var drawYBound int
		drawXBound = 0
		drawYBound = 0
		if i < 5 { // for first 5 images
			if argLoopFlag == "true" { // if true, start from center
				switch argStartFlag {
				case "right":
					drawXBound = -rect.Dx() + int(math.Floor(float64(rect.Dx())*0.2*float64(i+5)))
				case "left":
					drawXBound = rect.Dx() - int(math.Floor(float64(rect.Dx())*0.2*float64(i+5)))
				case "top":
					drawYBound = rect.Dy() - int(math.Floor(float64(rect.Dy())*0.2*float64(i+5)))
				case "bottom":
					drawYBound = -rect.Dy() + int(math.Floor(float64(rect.Dy())*0.2*float64(i+5)))
				default:
				}
			} else {
				switch argStartFlag {
				case "right":
					drawXBound = -rect.Dx() + int(math.Floor(float64(rect.Dx())*0.2*float64(i)))
				case "left":
					drawXBound = rect.Dx() - int(math.Floor(float64(rect.Dx())*0.2*float64(i)))
				case "top":
					drawYBound = rect.Dy() - int(math.Floor(float64(rect.Dy())*0.2*float64(i)))
				case "bottom":
					drawYBound = -rect.Dy() + int(math.Floor(float64(rect.Dy())*0.2*float64(i)))
				default:
				}
			}
		} else { // for last 5 images
			if argLoopFlag == "true" {
				switch argEndFlag {
				case "right":
					drawXBound = -int(math.Floor(float64(rect.Dx()) * 0.2 * float64(i-10)))
				case "left":
					drawXBound = int(math.Floor(float64(rect.Dx()) * 0.2 * float64(i-10)))
				case "top":
					drawYBound = int(math.Floor(float64(rect.Dy()) * 0.2 * float64(i-10)))
				case "bottom":
					drawYBound = -int(math.Floor(float64(rect.Dy()) * 0.2 * float64(i-10)))
				default:
				}
			} else {
				switch argEndFlag {
				case "right":
					drawXBound = -int(math.Floor(float64(rect.Dx()) * 0.2 * float64(i-5)))
				case "left":
					drawXBound = int(math.Floor(float64(rect.Dx()) * 0.2 * float64(i-5)))
				case "top":
					drawYBound = int(math.Floor(float64(rect.Dy()) * 0.2 * float64(i-5)))
				case "bottom":
					drawYBound = -int(math.Floor(float64(rect.Dy()) * 0.2 * float64(i-5)))
				default:
				}
			}
		}
		drawPoint := image.Pt(drawXBound, drawYBound)

		// gif can have 256 colors and one for Transparent
		const n = 255

		// use Quantizer to choose colors for gif
		var q draw.Quantizer = median.Quantizer(n)
		quantizePallete := q.Quantize(make(color.Palette, 0, n), g.sourceImg)

		newPalette := color.Palette{
			image.Transparent,
		}
		for _, color := range quantizePallete {
			newPalette = append(newPalette, color)
		}
		tmpPalette := image.NewPaletted(rect, newPalette)
		draw.Draw(tmpPalette, rect, g.sourceImg, drawPoint, draw.Src)
		// if use FloydSteinberg, some images output become wired color
		//draw.FloydSteinberg.Draw(tmpPalette, rect, decodedImage, drawPoint)

		gifImage = append(gifImage, tmpPalette)
		delays = append(delays, 0)
		disposal = append(disposal, 2)
	}

	var buffer bytes.Buffer
	// Set size to resized size
	gif.EncodeAll(&buffer, &gif.GIF{
		Image:     gifImage,
		Delay:     delays,
		Disposal:  disposal,
		LoopCount: 0,
	})

	g.outBuf = buffer
}

// updateImage writes the image to a byte buffer and then converts it to base64.
// Then it sets the value to the src attribute of the target image.
func (g *GifAnimeCreator) updateImage(start time.Time) {
	g.console.Call("log", "updateImage:", start.String())
	g.ConvertImage("left", "right", "false")
	out := g.outBuf.Bytes()
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&out))
	ptr := uintptr(unsafe.Pointer(hdr.Data))
	js.Global().Call("displayImage", ptr, len(out))
	g.console.Call("log", "time taken:", time.Now().Sub(start).String())
	g.outBuf.Reset()
}

// utility function to log a msg to the UI from inside a callback
func (s *GifAnimeCreator) log(msg string) {
	js.Global().Get("document").
		Call("getElementById", "status").
		Set("innerText", msg)
}
