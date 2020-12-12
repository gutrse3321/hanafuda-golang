package scenes

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"hanafuda-demo2/common"
	"hanafuda-demo2/common/constant"
	"hanafuda-demo2/sprites"
	"image/color"
)

type Title struct {
}

func (t *Title) Draw(screen *ebiten.Image) {
	//version & other information
	text.Draw(screen, "v0.0.1", common.FontObject.FontFiraCodeM, 50, constant.ScreenHeight-50, color.NRGBA{220, 220, 220, 0xff})
	//start content
	text.Draw(screen, common.TitleMenuString, common.FontObject.FontFiraCode, constant.ScreenWidth/2-(10*12), constant.ScreenHeight-200, color.White)

	//logo
	op := &ebiten.DrawImageOptions{}
	hanafudaTitleW, hanafudaTitleH := sprites.TitleImage.Size()
	op.GeoM.Translate(float64(constant.ScreenWidth/2)-float64(hanafudaTitleW/2), float64(hanafudaTitleH)/2.0)
	screen.DrawImage(sprites.TitleImage, op)
}

func (t *Title) Update() {
	//input
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeyKPEnter) {
		fmt.Println("点了")
		constant.GameMode = constant.ModeGame
	}
}
