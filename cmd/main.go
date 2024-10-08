package main

import (
	_ "image/png"
	"log"
	"math"
	assets "terminalRacer"

	"github.com/hajimehoshi/ebiten/v2"
)

var CarSprite *ebiten.Image

func init() {
	var err error

	CarSprite, err = assets.LoadImage("ui/static/gameModels/car.png")
	if err != nil {
		log.Fatalf("Failed to load car sprite: %v", err)
	}
}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Rotate(90.0 * math.Pi / 180.0)
	options.GeoM.Translate(150, 100)

	screen.DrawImage(CarSprite, options)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	game := &Game{}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
