package sprites

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"hanafuda-demo2/common/constant"
	"hanafuda-demo2/resources/textures"
	"image"
	_ "image/png"
	"log"
)

var (
	CardImage  *ebiten.Image
	TitleImage *ebiten.Image
	Img        *image.Image
)

func Init() {
	img, _, err := image.Decode(bytes.NewReader(textures.Texture_card))
	if err != nil {
		log.Fatal(err)
	}
	Img = &img

	//title image init
	titleImg, _, err := image.Decode(bytes.NewReader(textures.Texture_title))
	if err != nil {
		log.Fatal(err)
	}
	TitleImage = ebiten.NewImageFromImage(titleImg)
}

func NewSprites() [constant.CountFuda]*Card {
	CardImage = ebiten.NewImageFromImage(*Img)
	//map card & set
	x := 0
	y := 0
	count := 0
	var cards [50]*Card
	for {
		if count+1 > constant.CountFuda {
			return cards
		}
		if x == 10 {
			x = 0
			y++
		}

		sx := x*constant.Wfuda + constant.WFudaSpace*x
		sy := y*constant.Hfuda + constant.WFudaSpace*y
		fudaItem := CardImage.SubImage(image.Rect(sx, sy, sx+constant.Wfuda, sy+constant.Hfuda)).(*ebiten.Image)
		cards[count] = &Card{
			texture: fudaItem,
			Tag:     constant.TagFuda[count],
			ID:      uint8(count),
		}

		x++
		count++
	}

	return cards
}

type Card struct {
	texture *ebiten.Image
	X       int
	Y       int
	tx      float64
	ty      float64
	Angle   int
	Tag     string
	ID      uint8
}

func (c *Card) Translate(tx, ty float64) {
	c.tx, c.ty = tx, ty
}

func (c *Card) Persist(screen *ebiten.Image) {
	c.draw(screen)
}

func (c *Card) Click() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		fx := float64(x)
		fy := float64(y)
		if c.tx <= fx && fx < c.tx+float64(constant.Wfuda) && c.ty <= fy && fy < c.ty+float64(constant.Hfuda) {
			fmt.Println(c.Tag, ", ", c.ID)
		}
	}
}

func (c *Card) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(c.tx, c.ty)
	screen.DrawImage(c.texture, op)
}
