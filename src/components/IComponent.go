/**
 * @Author: Tomonori
 * @Date: 2019/11/12 11:34
 * @Desc:
 */
package components

import "github.com/hajimehoshi/ebiten/v2"

type IComponent interface {
	New(screen *ebiten.Image, fun ...func())
	update()
	draw(screen *ebiten.Image)
}
