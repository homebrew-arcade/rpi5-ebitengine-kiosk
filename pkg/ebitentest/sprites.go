package ebitentest

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	X        float64
	Y        float64
	XVel     float64
	YVel     float64
	XScale   float64
	YScale   float64
	XSize    float64
	YSize    float64
	Rotation float64
	Img      *ebiten.Image
}

func initSprites(g *Game) {
	g.sprs = make([]Sprite, MaxSprites)
	for i := range g.sprs {
		sp := &g.sprs[i]
		sp.XScale = rand.Float64()
		if sp.XScale < 0.1 {
			sp.XScale = 0.1
		}
		sp.XSize = FTileSize * sp.XScale
		sp.YScale = rand.Float64()
		if sp.YScale < 0.2 {
			sp.YScale = 0.2
		}
		sp.YSize = FTileSize * sp.YScale
		sp.X = rand.Float64() * (float64(ScreenWidth) - sp.XSize)
		sp.Y = rand.Float64() * (float64(ScreenHeight) - sp.YSize)
		sp.XVel = rand.Float64()*5 - 2.5
		sp.YVel = rand.Float64()*5 - 2.5
		sp.Rotation = rand.Float64() * 6.28318530718
		sp.Img = imgs[rand.IntN(len(imgs)-1)]
	}
}

func drawSprites(g *Game, s *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	for i := range g.sprCnt {
		sp := &g.sprs[i]
		op.GeoM.Reset()
		op.GeoM.Scale(sp.XScale, sp.YScale)
		op.GeoM.Translate(-FTileSize*sp.XScale/2, -FTileSize*sp.XScale/2)
		op.GeoM.Rotate(sp.Rotation)
		op.GeoM.Translate(sp.X+(FTileSize*sp.XScale/2), sp.Y+(FTileSize*sp.XScale/2))
		s.DrawImage(sp.Img, op)
	}
}

func updateSprites(g *Game) {
	for i := range g.sprCnt {
		sp := &g.sprs[i]
		sp.X += sp.XVel
		sp.Y += sp.YVel

		if sp.X < 0-sp.XSize {
			sp.X = float64(ScreenWidth)
		} else if sp.X > float64(ScreenWidth) {
			sp.X = 0 - sp.XSize
		}

		if sp.Y < 0-sp.YSize {
			sp.Y = float64(ScreenHeight)
		} else if sp.Y > float64(ScreenHeight) {
			sp.Y = 0 - sp.YSize
		}
	}
}
