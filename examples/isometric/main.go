package main

import (
	"github.com/split-cube-studios/ardent"
	"github.com/split-cube-studios/ardent/engine"
)

const (
	w, h = 854, 480
)

func main() {
	// create new game instance
	game := ardent.NewGame(
		"Isometric",
		w,
		h,
		engine.FlagResizable,
		// use Ebiten backend
		ardent.EBITEN,
		// tick function
		func() {},
		// layout function
		func(ow, oh int) (int, int) {
			return w, h
		},
	)

	// get component factory
	component := game.Component()

	atlas, _ := component.NewAtlasFromAssetPath("tiles.asset")

	data := [2][][]int{
		{
			{2, 1, 1, 1, 1},
			{1, 1, 1, 2, 1},
			{1, 1, 2, 1, 1},
			{2, 1, 1, 1, 1},
			{1, 1, 1, 2, 2},
		},
		{
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 3, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		},
	}

	mapper := map[int]engine.Image{
		1: atlas.GetImage("grass_1"),
		2: atlas.GetImage("grass_2"),
		3: atlas.GetImage("tree"),
	}

	tilemap := component.NewTilemap(128, data, mapper)
	camera := component.NewCamera()

	camera.LookAt(64, 128)

	renderer := component.NewIsoRenderer()
	renderer.SetTilemap(tilemap)
	renderer.SetCamera(camera)

	game.AddIsoRenderer(renderer)
	game.Run()
}