package constant

type Mode int

const (
	ModeTitle Mode = iota
	ModeGame
	ModeGameOver
)

var (
	ScreenWidth  = 1920
	ScreenHeight = 1080
	GameMode     = ModeTitle
)
