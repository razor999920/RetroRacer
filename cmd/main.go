package main

import (
	_ "image/png"
	"log"
	assets "terminalRacer"

	"github.com/hajimehoshi/ebiten/v2"
)

var PlayerSprite *ebiten.Image

func init() {
	var err error

	PlayerSprite, err = assets.LoadImage("ui/static/gameModels/car.png")
	if err != nil {
		log.Fatalf("Failed to load car sprite: %v", err)
	}
}

type Vector struct {
	X float64
	Y float64
}
type Game struct {
	playerPosition Vector
}

func (g *Game) Update() error {
	speed := float64(300 / ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.playerPosition.Y += speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.playerPosition.Y -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.playerPosition.X -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.playerPosition.X += speed
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	// options.GeoM.Rotate(90.0 * math.Pi / 180.0)
	options.GeoM.Translate(g.playerPosition.X, g.playerPosition.Y)

	screen.DrawImage(PlayerSprite, options)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	game := &Game{
		playerPosition: Vector{X: 100, Y: 100},
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
