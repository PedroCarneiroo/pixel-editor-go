package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// variaveis globais
var mapa [32][32]color.RGBA
var tamanho_quadrado = 20 // 640 dividido por 32

type Joguinho struct {
	// nao precisa de nada aqui pq to usando variavel global
}

func (j *Joguinho) Update() error {
	// desenhar com clique esquerdo
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		pos_x := x / tamanho_quadrado
		pos_y := y / tamanho_quadrado

		// verificando se nao ta clicando fora da tela pra nao dar erro
		if pos_x >= 0 {
			if pos_x < 32 {
				if pos_y >= 0 {
					if pos_y < 32 {
						mapa[pos_x][pos_y] = color.RGBA{0, 0, 0, 255}
					}
				}
			}
		}
	}

	// apagar com clique direito (copiei e colei de cima e mudei a cor)
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		x, y := ebiten.CursorPosition()
		pos_x := x / tamanho_quadrado
		pos_y := y / tamanho_quadrado

		if pos_x >= 0 && pos_x < 32 && pos_y >= 0 && pos_y < 32 {
			mapa[pos_x][pos_y] = color.RGBA{255, 255, 255, 255}
		}
	}

	// botao C limpa a tela toda
	if ebiten.IsKeyPressed(ebiten.KeyC) {
		for i := 0; i < 32; i++ {
			for k := 0; k < 32; k++ {
				mapa[i][k] = color.RGBA{255, 255, 255, 255}
			}
		}
	}

	return nil
}

func (j *Joguinho) Draw(tela *ebiten.Image) {
	// desenha os quadradinhos
	for i := 0; i < 32; i++ {
		for k := 0; k < 32; k++ {
			cor_atual := mapa[i][k]
			vector.DrawFilledRect(
				tela,
				float32(i*tamanho_quadrado),
				float32(k*tamanho_quadrado),
				float32(tamanho_quadrado),
				float32(tamanho_quadrado),
				cor_atual,
				true,
			)
		}
	}

	// faz as linhas cinzas por cima
	cor_linha := color.RGBA{200, 200, 200, 255}
	for cont := 0; cont <= 32; cont++ {
		// linha pra baixo
		vector.StrokeLine(tela, float32(cont*20), 0, float32(cont*20), 640, 1, cor_linha, true)
		// linha pro lado
		vector.StrokeLine(tela, 0, float32(cont*20), 640, float32(cont*20), 1, cor_linha, true)
	}
}

func (j *Joguinho) Layout(largura, altura int) (int, int) {
	return 640, 640
}

func main() {
	// pinta tudo de branco antes de começar
	for a := 0; a < 32; a++ {
		for b := 0; b < 32; b++ {
			mapa[a][b] = color.RGBA{255, 255, 255, 255}
		}
	}

	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("meu paint")

	meuJogo := &Joguinho{}
	err := ebiten.RunGame(meuJogo)

	if err != nil {
		log.Fatal(err)
	}
}
