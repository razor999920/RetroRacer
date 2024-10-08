package assets

import (
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed ui/static/gameModels/*
var gameAssets embed.FS

func LoadImage(name string) (*ebiten.Image, error) {
	f, err := gameAssets.Open(name)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img), nil
}
