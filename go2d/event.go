package go2d

import (
	"github.com/genbattle/Go2D/sdl"
)

func EventHandler(_event *sdl.SDLEvent) {
	switch _event.Evtype {
	case sdl.WINDOWEVENT:
		HandleWindowEvent(_event.Window())

	case sdl.KEYDOWN, sdl.TEXTINPUT:
		HandleKeyboardEvent(_event.Keyboard())

	case sdl.MOUSEBUTTONUP, sdl.MOUSEBUTTONDOWN:
		HandleMouseButtonEvent(_event.MouseButton())

	case sdl.MOUSEMOTION:
		HandleMouseMotionEvent(_event.MouseMotion())

	case sdl.MOUSEWHEEL:
		HandleMouseWheelEvent(_event.MouseWheel())
	}
}

func HandleWindowEvent(_event *sdl.WindowEvent) {
	switch _event.Event {
	case sdl.WINDOWEVENT_CLOSE:
		g_running = false
	}
}

func HandleKeyboardEvent(_event *sdl.KeyboardEvent) {
	switch _event.Evtype {
	case sdl.KEYDOWN:
		if g_game.keydownFun != nil {
			g_game.keydownFun(int(_event.Keysym().Scancode))
		}

		if g_game.guiManager != nil {
			g_game.guiManager.KeyDown(int(_event.Keysym().Scancode))
		}

	case sdl.KEYUP:
		if g_game.keyupFun != nil {
			g_game.keyupFun(int(_event.Keysym().Scancode))
		}

		if g_game.guiManager != nil {
			g_game.guiManager.KeyUp(int(_event.Keysym().Scancode))
		}

	case sdl.TEXTINPUT:
		keysym := uint8(_event.State)
		if keysym != 0 && keysym > 31 {
			if g_game.textinputFun != nil {
				g_game.textinputFun(keysym)
			}
		}

		if g_game.guiManager != nil {
			g_game.guiManager.TextInput(keysym)
		}
	}
}

func HandleMouseButtonEvent(_event *sdl.MouseButtonEvent) {
	switch _event.Evtype {
	case sdl.MOUSEBUTTONUP:
		if g_game.mouseupFun != nil {
			g_game.mouseupFun(int16(_event.X), int16(_event.Y))
		}

		if g_game.guiManager != nil {
			g_game.guiManager.MouseUp(int(_event.X), int(_event.Y))
		}

	case sdl.MOUSEBUTTONDOWN:
		if g_game.mousedownFun != nil {
			g_game.mousedownFun(int16(_event.X), int16(_event.Y))
		}

		if g_game.guiManager != nil {
			g_game.guiManager.MouseDown(int(_event.X), int(_event.Y))
		}
	}
}

func HandleMouseMotionEvent(_event *sdl.MouseMotionEvent) {
	if g_game.mousemoveFun != nil {
		g_game.mousemoveFun(int16(_event.X), int16(_event.Y))
	}

	if g_game.guiManager != nil {
		g_game.guiManager.MouseMove(int(_event.X), int(_event.Y))
	}
}

func HandleMouseWheelEvent(_event *sdl.MouseWheelEvent) {
	direction := 0
	if 0-_event.Y < 0 {
		direction = sdl.SCROLL_UP
	} else if (0 - _event.Y) > 0 {
		direction = sdl.SCROLL_DOWN
	}

	if g_game.mousescrollFun != nil {
		g_game.mousescrollFun(direction)
	}

	if g_game.guiManager != nil {
		g_game.guiManager.MouseScroll(direction)
	}
}
