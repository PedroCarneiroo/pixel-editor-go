package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 640
	screenHeight = 640
	gridSize     = 32
	cellSize     = screenWidth / gridSize
)

type App struct {
	canvas [gridSize][gridSize]color.RGBA
}

func NewApp() *App {
	app := &App{}
	app.ClearCanvas()
	return app
}

func (a *App) ClearCanvas() {
	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			a.canvas[x][y] = color.RGBA{255, 255, 255, 255}
		}
	}
}

func (a *App) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		cx, cy := mx/cellSize, my/cellSize

		if cx >= 0 && cx < gridSize && cy >= 0 && cy < gridSize {
			a.canvas[cx][cy] = color.RGBA{0, 0, 0, 255}
		}
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		mx, my := ebiten.CursorPosition()
		cx, cy := mx/cellSize, my/cellSize

		if cx >= 0 && cx < gridSize && cy >= 0 && cy < gridSize {
			a.canvas[cx][cy] = color.RGBA{255, 255, 255, 255}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyC) {
		a.ClearCanvas()
	}

	return nil
}

func (a *App) Draw(screen *ebiten.Image) {
	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			c := a.canvas[x][y]
			vector.DrawFilledRect(
				screen,
				float32(x*cellSize),
				float32(y*cellSize),
				float32(cellSize),
				float32(cellSize),
				c,
				true,
			)
		}
	}

	gridColor := color.RGBA{200, 200, 200, 255}
	for i := 0; i <= gridSize; i++ {
		vector.StrokeLine(screen, float32(i*cellSize), 0, float32(i*cellSize), float32(screenHeight), 1, gridColor, true)
		vector.StrokeLine(screen, 0, float32(i*cellSize), float32(screenWidth), float32(i*cellSize), 1, gridColor, true)
	}
}

func (a *App) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Editor de Pixel Art em Go")

	if err := ebiten.RunGame(NewApp()); err != nil {
		log.Fatal(err)
	}
}
