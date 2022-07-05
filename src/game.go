package src

import (
	"github.com/hajimehoshi/ebiten/v2"
	"hanafuda-demo2/common"
	"hanafuda-demo2/common/constant"
	"hanafuda-demo2/sprites"
	"hanafuda-demo2/src/scenes"
	"image/color"
)

type Game struct {
	titleScene   *scenes.Title
	contentScene *scenes.Content
	overScene    *scenes.Over
}

func (g *Game) init() {
	fontModel := &common.Font{}
	fontModel.Init()
	sprites.Init()

	g.titleScene = &scenes.Title{}
	g.contentScene = scenes.NewContent()
	g.overScene = &scenes.Over{}
}

func (g *Game) Update() error {
	switch constant.GameMode {
	case constant.ModeTitle:
		g.titleScene.Update()
	case constant.ModeGame:
		g.contentScene.Update()
	case constant.ModeGameOver:
		g.overScene.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{R: 139, G: 137, B: 112, A: 0xff})

	switch constant.GameMode {
	case constant.ModeTitle:
		g.titleScene.Draw(screen)
	case constant.ModeGame:
		g.contentScene.Draw(screen)
	case constant.ModeGameOver:
		g.overScene.Draw(screen)
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
