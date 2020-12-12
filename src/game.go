package src

import (
	"github.com/hajimehoshi/ebiten/v2"
	"hanafuda-demo2/common"
	"hanafuda-demo2/common/constant"
	"hanafuda-demo2/sprites"
	"hanafuda-demo2/src/scenes"
	"image/color"
)

var (
	titleScene   = &scenes.Title{}
	contentScene = &scenes.Content{}
	overScene    = &scenes.Over{}
)

type Game struct {
}

func (g *Game) init() {
	fontModel := &common.Font{}
	spriteModel := &sprites.Sprites{}
	fontModel.Init()
	spriteModel.Init()
}

func (g *Game) Update() error {
	switch constant.GameMode {
	case constant.ModeTitle:
		titleScene.Update()
	case constant.ModeGame:
		contentScene.Update()
	case constant.ModeGameOver:
		overScene.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{R: 139, G: 137, B: 112, A: 0xff})

	switch constant.GameMode {
	case constant.ModeTitle:
		titleScene.Draw(screen)
	case constant.ModeGame:
		contentScene.Draw(screen)
	case constant.ModeGameOver:
		overScene.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return constant.ScreenWidth, constant.ScreenHeight
}

func NewGame() *Game {
	g := &Game{}
	g.init()
	return g
}
