package main

import "github.com/hajimehoshi/ebiten/v2"

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

func (p *Player) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	// options.GeoM.Rotate(90.0 * math.Pi / 180.0)
	options.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(PlayerAvatar, options)
}
