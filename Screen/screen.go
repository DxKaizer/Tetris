package screen

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

// colors es una lista de colores que se pueden usar para las piezas de Tetris.
var colors = []termbox.Attribute{
	termbox.ColorBlack,
	termbox.ColorDarkGray,
	termbox.ColorLightRed,
	termbox.ColorMagenta,
	termbox.ColorLightGreen,
	termbox.ColorBlue,
	termbox.ColorYellow,
	termbox.ColorGreen,
}

// gameScreen es una estructura que representa la pantalla del juego.
type gameScreen struct {
}

// RenderAsci imprime el tablero del juego en la consola.
func (g *gameScreen) RenderAsci(board [][]int) {
	fmt.Println("\n===========")
	for _, e := range board {
		for _, num := range e {
			if num > 0 {
				fmt.Print("X") // Imprime una "X" para cada bloque de una pieza.
			} else {
				fmt.Println(" ") // Imprime un espacio para cada bloque vacío.
			}
		}
		fmt.Println("") // Imprime una nueva línea al final de cada línea del tablero.
	}
}

// Render dibuja el tablero del juego en la terminal usando la biblioteca termbox.
func (g *gameScreen) Render(board [][]int) {
	offsety := 3   // El desplazamiento vertical del tablero en la terminal.
	offsetx := 3   // El desplazamiento horizontal del tablero en la terminal.
	cellWidth := 2 // El ancho de cada bloque en la terminal.

	// Limpia la terminal.
	termbox.Clear(termbox.ColorCyan, termbox.ColorGreen)

	// Dibuja cada bloque del tablero en la terminal.
	for y, e := range board {
		for x, num := range e {
			color := colors[num] // El color del bloque.
			for i := 0; i < cellWidth; i++ {
				// Dibuja el bloque en la terminal.
				termbox.SetCell(offsetx+cellWidth*x+i, offsety+y, ' ', color, color)
			}
		}
	}

	// Actualiza la terminal para mostrar los cambios.
	termbox.Flush()
}

// New crea y devuelve una nueva instancia de gameScreen.
func New() *gameScreen {
	return &gameScreen{}
}
