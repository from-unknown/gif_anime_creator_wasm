// +build js,wasm

package gifanimecreator

import (
	"bytes"
	"image/png"
	"reflect"
	"syscall/js"
	"time"
	"unsafe"
)

func (g *GifAnimeCreator) setupOnImgLoadCb() {
	g.onImgLoadCb = js.NewCallback(func(args []js.Value) {
		reader := bytes.NewReader(g.inBuf)
		var err error
		g.sourceImg, err = png.Decode(reader)
		if err != nil {
			g.log(err.Error())
			return
		}
		g.log("Ready for operations")
		start := time.Now()
		g.updateImage(start)
	})
}

func (g *GifAnimeCreator) setupInitMemCb() {
	// The length of the image array buffer is passed.
	// Then the buf slice is initialized to that length.
	// And a pointer to that slice is passed back to the browser.
	g.initMemCb = js.NewCallback(func(i []js.Value) {
		length := i[0].Int()
		g.console.Call("log", "length:", length)
		g.inBuf = make([]uint8, length)
		hdr := (*reflect.SliceHeader)(unsafe.Pointer(&g.inBuf))
		ptr := uintptr(unsafe.Pointer(hdr.Data))
		js.Global().Call("gotMem", ptr)
	})
}
