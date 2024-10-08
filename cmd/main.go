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

type Player struct {
	position Vector
	avatar   *ebiten.Image
}

func NewPlayer() *Player {
	avatar := PlayerAvatar

	bounds := avatar.Bounds()
	horizontalPosition := float64(bounds.Dx()) / 2
	verticalPosition := float64(bounds.Dy())

	position := Vector{
		X: ScreenWidth/2 - horizontalPosition,
		Y: ScreenHeight - verticalPosition,
	}

	return &Player{
		avatar:   avatar,
		position: position,
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	// options.GeoM.Rotate(90.0 * math.Pi / 180.0)
	options.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(PlayerAvatar, options)

}

func (p *Player) Update() error {
	speed := float64(300 / ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.position.Y += speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.position.Y -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.position.X -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.position.X += speed
	}

	return nil
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
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
