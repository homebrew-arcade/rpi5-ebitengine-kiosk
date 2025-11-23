package ebitentest

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	Esc     = ebiten.KeyEscape
	P1Coin  = ebiten.Key5
	P1Start = ebiten.Key1
	P1Right = ebiten.KeyArrowRight
	P1Left  = ebiten.KeyArrowLeft
	P1Up    = ebiten.KeyArrowUp
	P1Down  = ebiten.KeyArrowDown
	P1B1    = ebiten.KeyControlLeft
	P1B2    = ebiten.KeyAltLeft
	P1B3    = ebiten.KeySpace
	P1B4    = ebiten.KeyShiftLeft
	P1B5    = ebiten.KeyZ
	P1B6    = ebiten.KeyX
	P1B7    = ebiten.KeyC
	P1B8    = ebiten.KeyV

	P2Coin  = ebiten.Key6
	P2Start = ebiten.Key2
	P2Right = ebiten.KeyG
	P2Left  = ebiten.KeyD
	P2Up    = ebiten.KeyR
	P2Down  = ebiten.KeyF
	P2B1    = ebiten.KeyA
	P2B2    = ebiten.KeyS
	P2B3    = ebiten.KeyQ
	P2B4    = ebiten.KeyW
	P2B5    = ebiten.KeyI
	P2B6    = ebiten.KeyK
	P2B7    = ebiten.KeyJ
	P2B8    = ebiten.KeyL
)

type Button struct {
	TX         int
	TY         int
	Key        ebiten.Key
	Img        *ebiten.Image
	ImgPressed *ebiten.Image
	IsPressed  bool
}

func initButtons(g *Game) {
	p2TOffset := 7
	g.btns = []Button{
		{
			Key:        P1Left,
			Img:        btnUp,
			ImgPressed: btnUpPressed,
			TX:         0,
			TY:         1,
		},
		{
			Key:        P1Up,
			Img:        btnUp,
			ImgPressed: btnUpPressed,
			TX:         1,
			TY:         0,
		},
		{
			Key:        P1Down,
			Img:        btnUp,
			ImgPressed: btnUpPressed,
			TX:         1,
			TY:         2,
		},
		{
			Key:        P1Right,
			Img:        btnUp,
			ImgPressed: btnUpPressed,
			TX:         2,
			TY:         1,
		},
		{
			Key:        P1Start,
			Img:        btnWhite,
			ImgPressed: btnWhitePressed,
			TX:         3,
			TY:         0,
		},
		{
			Key:        P1Coin,
			Img:        btnWhite,
			ImgPressed: btnWhitePressed,
			TX:         4,
			TY:         0,
		},
		{
			Key:        P1B7,
			Img:        btnGray,
			ImgPressed: btnGrayPressed,
			TX:         5,
			TY:         0,
		},
		{
			Key:        P1B8,
			Img:        btnGray,
			ImgPressed: btnGrayPressed,
			TX:         6,
			TY:         0,
		},
		{
			Key:        P1B1,
			Img:        btnBlue,
			ImgPressed: btnBluePressed,
			TX:         3,
			TY:         1,
		},
		{
			Key:        P1B2,
			Img:        btnBlue,
			ImgPressed: btnBluePressed,
			TX:         4,
			TY:         1,
		},
		{
			Key:        P1B3,
			Img:        btnBlue,
			ImgPressed: btnBluePressed,
			TX:         5,
			TY:         1,
		},
		{
			Key:        P1B4,
			Img:        btnBlue,
			ImgPressed: btnBluePressed,
			TX:         3,
			TY:         2,
		},
		{
			Key:        P1B5,
			Img:        btnBlue,
			ImgPressed: btnBluePressed,
			TX:         4,
			TY:         2,
		},
		{
			Key:        P1B6,
			Img:        btnBlue,
			ImgPressed: btnBluePressed,
			TX:         5,
			TY:         2,
		},
		{
			Key:        P2Left,
			Img:        btnUp,
			ImgPressed: btnUpPressed,
			TX:         0 + p2TOffset,
			TY:         1,
		},
		{
			Key:        P2Up,
			Img:        btnUp,
			ImgPressed: btnUpPressed,
			TX:         1 + p2TOffset,
			TY:         0,
		},
		{
			Key:        P2Down,
			Img:        btnUp,
			ImgPressed: btnUpPressed,
			TX:         1 + p2TOffset,
			TY:         2,
		},
		{
			Key:        P2Right,
			Img:        btnUp,
			ImgPressed: btnUpPressed,
			TX:         2 + p2TOffset,
			TY:         1,
		},
		{
			Key:        P2Start,
			Img:        btnWhite,
			ImgPressed: btnWhitePressed,
			TX:         3 + p2TOffset,
			TY:         0,
		},
		{
			Key:        P2Coin,
			Img:        btnWhite,
			ImgPressed: btnWhitePressed,
			TX:         4 + p2TOffset,
			TY:         0,
		},
		{
			Key:        P2B7,
			Img:        btnGray,
			ImgPressed: btnGrayPressed,
			TX:         5 + p2TOffset,
			TY:         0,
		},
		{
			Key:        P2B8,
			Img:        btnGray,
			ImgPressed: btnGrayPressed,
			TX:         6 + p2TOffset,
			TY:         0,
		},
		{
			Key:        P2B1,
			Img:        btnBlue,
			ImgPressed: btnBluePressed,
			TX:         3 + p2TOffset,
			TY:         1,
		},
		{
			Key:        P2B2,
			Img:        btnBlue,
			ImgPressed: btnBluePressed,
			TX:         4 + p2TOffset,
			TY:         1,
		},
		{
			Key:        P2B3,
			Img:        btnBlue,
			ImgPressed: btnBluePressed,
			TX:         5 + p2TOffset,
			TY:         1,
		},
		{
			Key:        P2B4,
			Img:        btnBlue,
			ImgPressed: btnBluePressed,
			TX:         3 + p2TOffset,
			TY:         2,
		},
		{
			Key:        P2B5,
			Img:        btnBlue,
			ImgPressed: btnBluePressed,
			TX:         4 + p2TOffset,
			TY:         2,
		},
		{
			Key:        P2B6,
			Img:        btnBlue,
			ImgPressed: btnBluePressed,
			TX:         5 + p2TOffset,
			TY:         2,
		},
	}
}

func drawButtons(g *Game, s *ebiten.Image) {
	TXOff := TileSize / 2
	TYOff := int(FTileSize * 2.5)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(TXOff-(TileSize/2)), float64(TYOff-(TileSize/2)))
	s.DrawImage(buttonBg, op)

	for _, btn := range g.btns {
		op.GeoM.Reset()
		op.GeoM.Translate(float64(btn.TX*TileSize+TXOff), float64(btn.TY*TileSize+TYOff))
		if btn.IsPressed {
			s.DrawImage(btn.ImgPressed, op)
		} else {
			s.DrawImage(btn.Img, op)
		}
	}
}

func updateButtons(g *Game) {
	for i := range g.btns {
		g.btns[i].IsPressed = false
		if ebiten.IsKeyPressed(g.btns[i].Key) {
			g.btns[i].IsPressed = true
		}
	}
}
