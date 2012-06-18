package sdl

// #cgo LDFLAGS: -lSDL2 -lpthread -lSDL_ttf -lSDL_image
// #cgo CFLAGS: -D_REENTRANT
// #include <SDL2/SDL.h>
import "C"

// This file contains OpenGL-specific functions from SDL.

type GLContext C.SDL_GLContext

func GL_CreateContext(_win *Window) *GLContext {
	context := GLContext(C.SDL_GL_CreateContext(_win.window))
	return &context
}
