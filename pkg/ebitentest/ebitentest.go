package ebitentest

import (
	"embed"
	"fmt"
	"image/color"
	"log"
	randv2 "math/rand/v2"
	"os"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed assets/*
var embedFS embed.FS

var BGColor = color.NRGBA{R: 0, G: 0, B: 0, A: 255}

var (
	MaxSprites        = 10000
	TPS               = 60
	ScreenWidth       = 480 // 16 tiles
	ScreenHeight      = 256 // 8 tiles
	WindowWidth       = 480
	WindowHeight      = 256
	VSyncEnabled      = true
	FullScreenEnabled = true
	WindowTitle       = "Ebiten Test"
	TileSize          = 32
	FTileSize         = float64(32)
)

var rand = func() *randv2.Rand {
	s := randv2.NewPCG(10, 1)
	return randv2.New(s)
}()

type Game struct {
	sprCnt     int
	startHoldF int
	sprs       []Sprite
	btns       []Button
}

func (g *Game) Init() {
	g.sprCnt = 1
	g.startHoldF = 0
	initSprites(g)
	initButtons(g)
}

func (g *Game) Draw(s *ebiten.Image) {
	s.Fill(BGColor)
	drawSprites(g, s)
	drawButtons(g, s)
	ebitenutil.DebugPrint(s, fmt.Sprintf("%v sprites; FPS: %v", g.sprCnt, ebiten.ActualFPS()))
}

func (g *Game) Update() error {

	if ebiten.IsKeyPressed(Esc) {
		//ebiten.IsKeyPressed(ebiten.Key1) {
		return fmt.Errorf("Closing app from Escape exit")
	}

	if ebiten.IsKeyPressed(P1Start) {
		g.startHoldF++
		if g.startHoldF > 120 {
			return fmt.Errorf("Closing app from P1 Start exit")
		}
	} else {
		g.startHoldF = 0
	}

	if ebiten.IsKeyPressed(P1Up) {
		g.sprCnt++
	} else if ebiten.IsKeyPressed(P1Down) {
		g.sprCnt--
	} else if ebiten.IsKeyPressed(P1Right) {
		g.sprCnt += 10
	} else if ebiten.IsKeyPressed(P1Left) {
		g.sprCnt -= 10
	}
	if g.sprCnt > MaxSprites {
		g.sprCnt = MaxSprites
	} else if g.sprCnt < 1 {
		g.sprCnt = 1
	}
	updateSprites(g)
	updateButtons(g)
	return nil
}

func getSafeENVInt(key string, def int) int {
	envStr := os.Getenv(key)
	if envStr != "" {
		num, err := strconv.Atoi(envStr)
		if err == nil {
			fmt.Println(key, num)
			return num
		}
	}
	return def
}

func updateSettingsFromENV() {
	MaxSprites = getSafeENVInt("EBITENTEST_MAX_SPRITES", MaxSprites)
	TPS = getSafeENVInt("EBITENTEST_TPS", TPS)
	fullScreenEnabledInt := getSafeENVInt("EBITENTEST_FULL_SCREEN_ENABLED", 1)
	if fullScreenEnabledInt == 0 {
		FullScreenEnabled = false
	}
	ScreenWidth = getSafeENVInt("EBITENTEST_SCREEN_WIDTH", ScreenWidth)
	ScreenHeight = getSafeENVInt("EBITENTEST_SCREEN_HEIGHT", ScreenHeight)
	WindowWidth = getSafeENVInt("EBITENTEST_WINDOW_WIDTH", WindowWidth)
	WindowHeight = getSafeENVInt("EBITENTEST_WINDOW_HEIGHT", WindowHeight)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func StartEbitenTest() {
	updateSettingsFromENV()
	ebiten.SetWindowSize(WindowWidth, WindowHeight)
	ebiten.SetWindowTitle(WindowTitle)
	ebiten.SetTPS(TPS)
	ebiten.SetVsyncEnabled(VSyncEnabled)
	ebiten.SetFullscreen(FullScreenEnabled)
	g := &Game{}
	g.Init()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
