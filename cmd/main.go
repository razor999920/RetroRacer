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
var GrassAvatar *ebiten.Image

func init() {
	var err error

	PlayerAvatar, err = assets.LoadImage("ui/static/gameModels/car.png")
	if err != nil {
		log.Fatalf("Failed to load car sprite: %v", err)
	}

	GrassAvatar, err = assets.LoadImage("ui/static/gameModels/grass.png")
	if err != nil {
		log.Fatalf("Failed to load glass sprite: %v", err)
	}
}

type Vector struct {
	X float64
	Y float64
}

type Game struct {
	track  *Track
	player *Player
}

func (g *Game) Update() error {
	// Track
	err := g.track.Update()
	// Player
	err = g.player.Update()

	if err != nil {
		return err
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Track
	g.track.Draw(screen)
	// Player
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	// Initilize game
	game := &Game{
		track:  NewTrack(),
		player: NewPlayer(),
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
