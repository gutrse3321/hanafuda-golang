package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"hanafuda-demo2/common/constant"
	"hanafuda-demo2/sprites"
)

type Content struct {
}

func (t *Content) Draw(screen *ebiten.Image) {
	//text.Draw(screen, "Welcome Gamers！", common.FontObject.FontHuakang, constant.ScreenWidth/2-(5*12), constant.ScreenHeight/2-12, color.White)

	//TODO test sprite
	i := 0
	x := 0
	y := 0
	for _, tag := range constant.TagFuda {
		if i > 48 {
			break
		}
		if x == 9 {
			x = 0
			y++
		}
		sx := x*constant.Wfuda + 30*x + 20
		sy := y*constant.Hfuda + 30*y + 20
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(sx), float64(sy))
		image := sprites.CardSprites[tag].Texture
		screen.DrawImage(image, op)
		i++
		x++
	}
}

func (t *Content) Update() {
	//input
	if inpututil.IsKeyJustPressed(ebiten.KeyF1) {
		constant.GameMode = constant.ModeGameOver
	}
}
