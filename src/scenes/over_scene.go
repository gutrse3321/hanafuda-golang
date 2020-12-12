package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"hanafuda-demo2/common"
	"hanafuda-demo2/common/constant"
	"hanafuda-demo2/src/components"
	"image/color"
)

type Over struct {
	btnBox *ebiten.Image
	btnBg  *ebiten.Image

	restartButton *components.Button
}

func (o *Over) Draw(screen *ebiten.Image) {
	text.Draw(screen, "Game Over", common.FontObject.FontFiraCodeBig, constant.ScreenWidth/2-(11*12), constant.ScreenHeight/2-24, color.White)

	//restart button
	o.restartButton = components.NewButton(
		"Restart",
		"large",
		"black",
	)
	o.restartButton.Translate(float64(constant.ScreenWidth/2-o.restartButton.Width/2), float64(constant.ScreenHeight-200))
	o.restartButton.Persist(screen)
}

func (o *Over) Update() {
	//event
	o.restartButton.Click(func() {
		constant.GameMode = constant.ModeGame
	})
}
