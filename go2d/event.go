package go2d

import (
	"github.com/papplampe/Go2D/sdl"
)

func EventHandler(_event *sdl.SDLEvent) {
	switch _event.Evtype {
	case sdl.KEYDOWN:
		keyboard := _event.Keyboard()
		if g_game.keydownFun != nil {
			g_game.keydownFun(int(keyboard.Keysym().Scancode))
		}

		if g_game.guiManager != nil {
			g_game.guiManager.KeyDown(int(keyboard.Keysym().Scancode))
		}
		
	case sdl.KEYUP:
		keyboard := _event.Keyboard()
		if g_game.keyupFun != nil {
			g_game.keyupFun(int(keyboard.Keysym().Scancode))
		}

		if g_game.guiManager != nil {
			g_game.guiManager.KeyUp(int(keyboard.Keysym().Scancode))
		}
		
	case sdl.TEXTINPUT:
		keyboard := _event.Keyboard()
		keysym := uint8(keyboard.State)
		if keysym != 0 && keysym > 31 {
			if g_game.textinputFun != nil {
				g_game.textinputFun(keysym)
			}
		}

		if g_game.guiManager != nil {
			g_game.guiManager.TextInput(keysym)
		}
		
	case sdl.MOUSEBUTTONUP:
		mouse := _event.MouseButton()
		if g_game.mouseupFun != nil {
			g_game.mouseupFun(int16(mouse.X), int16(mouse.Y))
		}

		if g_game.guiManager != nil {
			g_game.guiManager.MouseUp(int(mouse.X), int(mouse.Y))
		}
		
	case sdl.MOUSEBUTTONDOWN:
		mouse := _event.MouseButton()
		if g_game.mousedownFun != nil {
			g_game.mousedownFun(int16(mouse.X), int16(mouse.Y))
		}

		if g_game.guiManager != nil {
			g_game.guiManager.MouseDown(int(mouse.X), int(mouse.Y))
		}
		
	case sdl.MOUSEMOTION:
		mouse := _event.MouseMotion()
		if g_game.mousemoveFun != nil {
			g_game.mousemoveFun(int16(mouse.X), int16(mouse.Y))
		}

		if g_game.guiManager != nil {
			g_game.guiManager.MouseMove(int(mouse.X), int(mouse.Y))
		}
		
	case sdl.MOUSEWHEEL:
		mouse := _event.MouseWheel()
		direction := 0
		if 0-mouse.Y < 0 {
			direction = sdl.SCROLL_UP
		} else if (0 - mouse.Y) > 0 {
			direction = sdl.SCROLL_DOWN
		}

		if g_game.mousescrollFun != nil {
			g_game.mousescrollFun(direction)
		}

		if g_game.guiManager != nil {
			g_game.guiManager.MouseScroll(direction)
		}
		
	case sdl.QUIT:
		g_game.Exit()
	}
}
