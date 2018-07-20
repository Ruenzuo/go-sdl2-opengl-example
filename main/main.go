package main

import (
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()
	// Game Boy resolution: 160x144
	// 3X for HDPI screens
	const width = int32(160 * 3)
	const height = int32(144 * 3)
	window, err := sdl.CreateWindow("go-sdl2-opengl-example", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_OPENGL)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, 0)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()
	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STATIC, width, height)
	if err != nil {
		panic(err)
	}
	defer texture.Destroy()

	// 8888 pixel format is 4 bytes to represent pixels: RBG+A channels
	// No fucking clue why sdl takes []byte instead of uint32
	pixels := make([]byte, width*height*4)
	for i := range pixels {
		pixels[i] = 255
	}
	running := true
	for running {
		// the number of bytes in a row of pixel data
		pitch := width * int32(unsafe.Sizeof(uint32(0)))
		texture.Update(nil, pixels, int(pitch))
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}
		renderer.Clear()
		renderer.Copy(texture, nil, nil)
		renderer.Present()
	}
}
