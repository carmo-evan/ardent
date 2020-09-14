package ebiten

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/split-cube-studios/ardent/internal/common"
)

var commonToEbitenKey = map[int]ebiten.Key{
	common.Key0:            ebiten.Key0,
	common.Key1:            ebiten.Key1,
	common.Key2:            ebiten.Key2,
	common.Key3:            ebiten.Key3,
	common.Key4:            ebiten.Key4,
	common.Key5:            ebiten.Key5,
	common.Key6:            ebiten.Key6,
	common.Key7:            ebiten.Key7,
	common.Key8:            ebiten.Key8,
	common.Key9:            ebiten.Key9,
	common.KeyA:            ebiten.KeyA,
	common.KeyB:            ebiten.KeyB,
	common.KeyC:            ebiten.KeyC,
	common.KeyD:            ebiten.KeyD,
	common.KeyE:            ebiten.KeyE,
	common.KeyF:            ebiten.KeyF,
	common.KeyG:            ebiten.KeyG,
	common.KeyH:            ebiten.KeyH,
	common.KeyI:            ebiten.KeyI,
	common.KeyJ:            ebiten.KeyJ,
	common.KeyK:            ebiten.KeyK,
	common.KeyL:            ebiten.KeyL,
	common.KeyM:            ebiten.KeyM,
	common.KeyN:            ebiten.KeyN,
	common.KeyO:            ebiten.KeyO,
	common.KeyP:            ebiten.KeyP,
	common.KeyQ:            ebiten.KeyQ,
	common.KeyR:            ebiten.KeyR,
	common.KeyS:            ebiten.KeyS,
	common.KeyT:            ebiten.KeyT,
	common.KeyU:            ebiten.KeyU,
	common.KeyV:            ebiten.KeyV,
	common.KeyW:            ebiten.KeyW,
	common.KeyX:            ebiten.KeyX,
	common.KeyY:            ebiten.KeyY,
	common.KeyZ:            ebiten.KeyZ,
	common.KeyApostrophe:   ebiten.KeyApostrophe,
	common.KeyBackslash:    ebiten.KeyBackslash,
	common.KeyBackspace:    ebiten.KeyBackspace,
	common.KeyCapsLock:     ebiten.KeyCapsLock,
	common.KeyComma:        ebiten.KeyComma,
	common.KeyDelete:       ebiten.KeyDelete,
	common.KeyDown:         ebiten.KeyDown,
	common.KeyEnd:          ebiten.KeyEnd,
	common.KeyEnter:        ebiten.KeyEnter,
	common.KeyEqual:        ebiten.KeyEqual,
	common.KeyEscape:       ebiten.KeyEscape,
	common.KeyF1:           ebiten.KeyF1,
	common.KeyF2:           ebiten.KeyF2,
	common.KeyF3:           ebiten.KeyF3,
	common.KeyF4:           ebiten.KeyF4,
	common.KeyF5:           ebiten.KeyF5,
	common.KeyF6:           ebiten.KeyF6,
	common.KeyF7:           ebiten.KeyF7,
	common.KeyF8:           ebiten.KeyF8,
	common.KeyF9:           ebiten.KeyF9,
	common.KeyF10:          ebiten.KeyF10,
	common.KeyF11:          ebiten.KeyF11,
	common.KeyF12:          ebiten.KeyF12,
	common.KeyGraveAccent:  ebiten.KeyGraveAccent,
	common.KeyHome:         ebiten.KeyHome,
	common.KeyInsert:       ebiten.KeyInsert,
	common.KeyKP0:          ebiten.KeyKP0,
	common.KeyKP1:          ebiten.KeyKP1,
	common.KeyKP2:          ebiten.KeyKP2,
	common.KeyKP3:          ebiten.KeyKP3,
	common.KeyKP4:          ebiten.KeyKP4,
	common.KeyKP5:          ebiten.KeyKP5,
	common.KeyKP6:          ebiten.KeyKP6,
	common.KeyKP7:          ebiten.KeyKP7,
	common.KeyKP8:          ebiten.KeyKP8,
	common.KeyKP9:          ebiten.KeyKP9,
	common.KeyKPAdd:        ebiten.KeyKPAdd,
	common.KeyKPDecimal:    ebiten.KeyKPDecimal,
	common.KeyKPDivide:     ebiten.KeyKPDivide,
	common.KeyKPEnter:      ebiten.KeyKPEnter,
	common.KeyKPEqual:      ebiten.KeyKPEqual,
	common.KeyKPMultiply:   ebiten.KeyKPMultiply,
	common.KeyKPSubtract:   ebiten.KeyKPSubtract,
	common.KeyLeft:         ebiten.KeyLeft,
	common.KeyLeftBracket:  ebiten.KeyLeftBracket,
	common.KeyMenu:         ebiten.KeyMenu,
	common.KeyMinus:        ebiten.KeyMinus,
	common.KeyNumLock:      ebiten.KeyNumLock,
	common.KeyPageDown:     ebiten.KeyPageDown,
	common.KeyPageUp:       ebiten.KeyPageUp,
	common.KeyPause:        ebiten.KeyPause,
	common.KeyPeriod:       ebiten.KeyPeriod,
	common.KeyPrintScreen:  ebiten.KeyPrintScreen,
	common.KeyRight:        ebiten.KeyRight,
	common.KeyRightBracket: ebiten.KeyRightBracket,
	common.KeyScrollLock:   ebiten.KeyScrollLock,
	common.KeySemicolon:    ebiten.KeySemicolon,
	common.KeySlash:        ebiten.KeySlash,
	common.KeySpace:        ebiten.KeySpace,
	common.KeyTab:          ebiten.KeyTab,
	common.KeyUp:           ebiten.KeyUp,
	common.KeyAlt:          ebiten.KeyAlt,
	common.KeyControl:      ebiten.KeyControl,
	common.KeyShift:        ebiten.KeyShift,
}

var commonToEbitenMouseButton = map[int]ebiten.MouseButton{
	common.MouseButtonLeft:   ebiten.MouseButtonLeft,
	common.MouseButtonRight:  ebiten.MouseButtonRight,
	common.MouseButtonMiddle: ebiten.MouseButtonMiddle,
}