package sdl

// #cgo LDFLAGS: -lSDL2 -lpthread -lSDL2_ttf -lSDL2_image
// #cgo CFLAGS: -D_REENTRANT
// #include <SDL2/SDL.h>
import "C"
import "unsafe"
import "errors"

type cast unsafe.Pointer

func Delay(_ticks uint32) {
	C.SDL_Delay(C.Uint32(_ticks))
}

func GetTicks() uint32 {
	return uint32(C.SDL_GetTicks())
}

func Quit() {
	C.SDL_Quit()
}

func GetError() (ret string) {
	ret = C.GoString(C.SDL_GetError())
	C.SDL_ClearError()
	return
}

func KeyDown(_key int) bool {
	zero := C.int(0)
	var state = uintptr(unsafe.Pointer(C.SDL_GetKeyboardState(&zero))) + uintptr(_key)
	down := (*uint8)(cast(state))
	if *down == 1 {
		return true
	}
	return false
}

func Init(flags uint32) (err error) {
	if C.SDL_Init(C.Uint32(flags)) != 0 {
		return errors.New(C.GoString(C.SDL_GetError()))
	}
	return nil
}

type Window struct {
	window *C.SDL_Window
}

func CreateWindow(_title string, _x int, _y int, _width int, _height int, _flags uint32) (*Window, error) {
	ctitle := C.CString(_title)
	var window *C.SDL_Window = C.SDL_CreateWindow(ctitle,
		C.int(_x),
		C.int(_y),
		C.int(_width),
		C.int(_height),
		C.Uint32(_flags))
	C.free(unsafe.Pointer(ctitle))
	if window == nil {
		return nil, errors.New(GetError())
	}
	return &Window{window}, nil
}

func DestroyWindow(_window *Window) {
	C.SDL_DestroyWindow(_window.window)
}

func PollEvent() (*SDLEvent, bool) {
	var ev *SDLEvent = &SDLEvent{}
	if C.SDL_PollEvent((*C.SDL_Event)(cast(ev))) != 0 {
		return ev, true
	}
	return nil, false
}
