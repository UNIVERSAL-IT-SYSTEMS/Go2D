package sdl

// #cgo LDFLAGS: -lSDL2 -lpthread -lSDL_ttf -lSDL_image
// #cgo CFLAGS: -D_REENTRANT
// #include <SDL2/SDL.h>
import "C"
import "errors"

// This file contains OpenGL-specific functions from SDL.

type GLContext C.SDL_GLContext

func GL_CreateContext(_win *Window) *GLContext {
	context := GLContext(C.SDL_GL_CreateContext(_win.window))
	return &context
}

func GL_SetSwapInterval(_interval int) error {
	if C.SDL_GL_SetSwapInterval(C.int(_interval)) != 0 {
		return errors.New(GetError())
	}
	return nil
}
