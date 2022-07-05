package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"hanafuda-demo2/common/constant"
	"hanafuda-demo2/sprites"
)

type Content struct {
	cards [50]*sprites.Card
}

// NewContent init cards sprites
func NewContent() *Content {
	return &Content{
		cards: sprites.NewSprites(),
	}
}

func (t *Content) Draw(screen *ebiten.Image) {
	//text.Draw(screen, "Welcome Gamers！", common.FontObject.FontHuakang, constant.ScreenWidth/2-(5*12), constant.ScreenHeight/2-12, color.White)
	//TODO test sprite
	i := 0
	x := float64(0)
	y := float64(0)
	for {
		if i+1 > constant.CountFuda {
			break
		}
		if x == 10 {
			x = 0
			y++
		}
		sx := x*constant.Wfuda + 30*x + 20
		sy := y*constant.Hfuda + 30*y + 20

		card := t.cards[i]
		card.Translate(sx, sy)
		card.Persist(screen)
		i++
		x++
	}
}

func (t *Content) Update() {
	//input
	if inpututil.IsKeyJustPressed(ebiten.KeyF1) {
		constant.GameMode = constant.ModeGameOver
	}

	i := 0
	for {
		if i+1 > constant.CountFuda {
			break
		}
		t.cards[i].Click()
		i++
	}
}
