package main

import (
	_ "image/png"
	"log"
	assets "terminalRacer"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

var PlayerAvatar *ebiten.Image

func init() {
	var err error

	PlayerAvatar, err = assets.LoadImage("ui/static/gameModels/car.png")
	if err != nil {
		log.Fatalf("Failed to load car sprite: %v", err)
	}
}

type Vector struct {
	X float64
	Y float64
}

type Game struct {
	player *Player
}

func (g *Game) Update() error {
	return g.player.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	player := NewPlayer()

	// Initilize game
	game := &Game{
		player: player,
		// attackTimer : NewTimet(5 *time.Second)
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
