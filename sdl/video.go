package sdl

// #include <SDL2/SDL.h>
import "C"
import "errors"

type GLattr C.SDL_GLattr

func GL_SetAttribute(attribute GLattr, value int) (err error) {
	if C.SDL_GL_SetAttribute(C.SDL_GLattr(attribute), C.int(value)) != 0 {
		return errors.New(C.GoString(C.SDL_GetError()))
	}
	return nil
}
