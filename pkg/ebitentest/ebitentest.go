package ebitentest

import (
	"embed"
	"fmt"
	"image/color"
	"log"
	randv2 "math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed assets/*
var embedFS embed.FS

var BGColor = color.NRGBA{R: 191, G: 251, B: 191, A: 255}

var screenWidth int = 1280
var screenHeight int = 720

var spriteImg = func() *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFileSystem(embedFS, "assets/sprite.png")
	if err != nil {
		log.Fatal(err)
	}
	return img
}()

var rand = func() *randv2.Rand {
	s := randv2.NewPCG(0, 0)
	return randv2.New(s)
}()

type Sprite struct {
	X      float64
	Y      float64
	XVel   float64
	YVel   float64
	XScale float64
	YScale float64
	XSize  float64
	YSize  float64
}

type Game struct {
	tickCnt int
	sprs    []Sprite
}

func (g *Game) Init() {
	g.tickCnt = 0
	g.sprs = make([]Sprite, 2500)
	for i := range g.sprs {
		sp := &g.sprs[i]
		sp.XScale = rand.Float64()
		if sp.XScale < 0.1 {
			sp.XScale = 0.1
		}
		sp.XSize = 64 * sp.XScale
		sp.YScale = rand.Float64()
		if sp.YScale < 0.1 {
			sp.YScale = 0.1
		}
		sp.YSize = 64 * sp.YScale
		sp.X = rand.Float64() * (float64(screenWidth) - sp.XSize)
		sp.Y = rand.Float64() * (float64(screenHeight) - sp.YSize)
		sp.XVel = rand.Float64()*5 - 2.5
		sp.YVel = rand.Float64()*5 - 2.5
	}
}

func (g *Game) Draw(s *ebiten.Image) {
	s.Fill(BGColor)
	op := &ebiten.DrawImageOptions{}
	for _, sp := range g.sprs {
		op.GeoM.Reset()
		op.GeoM.Scale(sp.XScale, sp.YScale)
		op.GeoM.Translate(sp.X, sp.Y)

		s.DrawImage(spriteImg, op)
	}
	ebitenutil.DebugPrint(s, fmt.Sprintf("FPS: %v", ebiten.ActualFPS()))
}

func (g *Game) Update() error {
	g.tickCnt++
	if g.tickCnt > (60 * 15) {
		return fmt.Errorf("Closing app after 15 seconds")
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) || ebiten.IsKeyPressed(ebiten.KeyQ) {
		return fmt.Errorf("Closing app from keyboard exit")
	}

	for i := range g.sprs {
		sp := &g.sprs[i]
		sp.X += sp.XVel
		sp.Y += sp.YVel

		if sp.X < 0-sp.XSize {
			sp.X = float64(screenWidth)
		} else if sp.X > float64(screenWidth) {
			sp.X = 0 - sp.XSize
		}

		if sp.Y < 0-sp.YSize {
			sp.Y = float64(screenHeight)
		} else if sp.Y > float64(screenHeight) {
			sp.Y = 0 - sp.YSize
		}
	}
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func StartEbitenTest() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Kiosk Test")
	ebiten.SetTPS(60)
	ebiten.SetVsyncEnabled(true)
	g := &Game{}
	g.Init()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
