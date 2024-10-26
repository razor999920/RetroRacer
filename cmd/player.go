package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

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

func (p *Player) Update() error {
	speed := float64(300 / ebiten.TPS())
	verticalPosition := p.position.Y
	horitontalPosition := p.position.X

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		verticalPosition += speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		verticalPosition -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		horitontalPosition -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		horitontalPosition += speed
	}

	// Make sure the player is not going out of the frame
	if verticalPosition <= 0 || verticalPosition+float64(p.avatar.Bounds().Dy()) >= ScreenHeight || horitontalPosition <= 0 || horitontalPosition+float64(p.avatar.Bounds().Dx()) >= ScreenWidth {
		return nil
	}

	// Update the car's position
	p.position.X = horitontalPosition
	p.position.Y = verticalPosition

	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	// options.GeoM.Rotate(90.0 * math.Pi / 180.0)
	options.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(PlayerAvatar, options)
}
