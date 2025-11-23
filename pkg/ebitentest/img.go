package ebitentest

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var buttonsImg = func() *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFileSystem(embedFS, "assets/buttons.png")
	if err != nil {
		log.Fatal(err)
	}
	return img
}()

var btnBlue = buttonsImg.SubImage(image.Rect(0, 0, 32, 32)).(*ebiten.Image)
var btnBluePressed = buttonsImg.SubImage(image.Rect(32, 0, 64, 32)).(*ebiten.Image)
var btnWhite = buttonsImg.SubImage(image.Rect(64, 0, 96, 32)).(*ebiten.Image)
var btnWhitePressed = buttonsImg.SubImage(image.Rect(96, 0, 128, 32)).(*ebiten.Image)
var btnGray = buttonsImg.SubImage(image.Rect(128, 0, 160, 32)).(*ebiten.Image)
var btnGrayPressed = buttonsImg.SubImage(image.Rect(160, 0, 192, 32)).(*ebiten.Image)
var btnUp = buttonsImg.SubImage(image.Rect(192, 0, 224, 32)).(*ebiten.Image)
var btnUpPressed = buttonsImg.SubImage(image.Rect(224, 0, 256, 32)).(*ebiten.Image)
var imgs = []*ebiten.Image{
	btnBlue,
	btnBluePressed,
	btnGray,
	btnGrayPressed,
	btnWhite,
	btnWhitePressed,
	btnUp,
	btnUpPressed,
}
var buttonBg = func() *ebiten.Image {
	bg := ebiten.NewImage(16*32, 4*32)
	bg.Fill(color.NRGBA{R: 0, G: 0, B: 0, A: 200})
	return bg
}()
