package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"hanafuda-demo2/common/constant"
	"hanafuda-demo2/src"
	"log"
)

func main() {
	ebiten.SetWindowSize(constant.ScreenWidth, constant.ScreenHeight)
	ebiten.SetWindowTitle("Hanafuda (花札 Demo)")
	if err := ebiten.RunGame(src.NewGame()); err != nil {
		log.Println(err)
	}
}
