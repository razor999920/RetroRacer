package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Track struct {
	avatar *ebiten.Image

	leftBoundary  float64
	rightBoundary float64

	avatarWidth float64
}

func NewTrack() *Track {
	avatar := GrassAvatar

	const (
		grassWidthPercentage = 0.25
		roadWidthPercentage  = 0.5
	)

	grassWidth := float64(ScreenWidth) * grassWidthPercentage
	leftBoundary := grassWidth
	rightBoundary := float64(ScreenWidth) - grassWidth

	return &Track{
		avatar:        avatar,
		leftBoundary:  leftBoundary,
		rightBoundary: rightBoundary,
		avatarWidth:   grassWidth,
	}
}

func (t *Track) Update() error {
	return nil
}

func (t *Track) Draw(screen *ebiten.Image) {
	// Left Track
	t.drawGrass(screen, 0)
	// Right Track
	t.drawGrass(screen, float64(ScreenWidth)-t.avatarWidth)
}

// Draw Grass texture
func (t *Track) drawGrass(screen *ebiten.Image, startX float64) {
	grassW := float64(t.avatar.Bounds().Dx())
	grassH := float64(t.avatar.Bounds().Dy())

	repeatY := float64(ScreenHeight) / grassH
	repeatX := t.leftBoundary / grassW

	for y := 0; y <= int(repeatY); y++ {
		for x := 0; x <= int(repeatX); x++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(startX+float64(x)*grassW, float64(y)*grassH)
			screen.DrawImage(t.avatar, op)
		}
	}
}

func (t *Track) IsOutOfBounds(x float64) bool {
	return x < t.leftBoundary || x > t.rightBoundary
}

func (t *Track) GetBoundaries() (left, right float64) {
	return t.leftBoundary, t.rightBoundary
}
