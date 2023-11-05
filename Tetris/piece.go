package Tetris

import (
	"math"
	"math/rand"
)

// piece representa una pieza de Tetris.
type piece struct {
	shape     []vector // Los bloques que componen la forma de la pieza.
	color     int      // El color de la pieza.
	canRotate bool     // Si la pieza puede ser rotada.
}

// pieces es una lista de todas las posibles piezas de Tetris.
var pieces = []piece{
	// Cada pieza se define por su forma, color y si puede ser rotada.
	// La forma se define como una lista de vectores, donde cada vector representa un bloque en la pieza.
	// Aquí se definen las formas clásicas de las piezas de Tetris: L, I, O, T, S, Z, y J.
	// ...
}

// rotateBack rota la pieza en sentido antihorario.
func (p *piece) rotateBack() {
	ang := math.Pi / 2 * 3 // Ángulo de rotación (270 grados en radianes).
	p.rotateWithAngle(ang)
}

// rotate rota la pieza en sentido horario.
func (p *piece) rotate() {
	ang := math.Pi / 2 // Ángulo de rotación (90 grados en radianes).
	p.rotateWithAngle(ang)
}

// rotateWithAngle rota la pieza un cierto ángulo.
func (p *piece) rotateWithAngle(ang float64) {
	if !p.canRotate {
		return // Si la pieza no puede ser rotada, no hacemos nada.
	}
	angle := math.Pi / 2
	cos := int(math.Round(math.Cos(angle))) // Calcula el coseno del ángulo.
	sin := int(math.Round(math.Sin(angle))) // Calcula el seno del ángulo.

	// Rota cada bloque en la pieza.
	for i, e := range p.shape {
		ny := e.y*cos - e.x*sin
		nx := e.y*sin - e.x*cos

		p.shape[i] = vector{ny, nx} // Actualiza la posición del bloque.
	}
}

// randomPiece genera una pieza aleatoria.
func randomPiece() piece {
	idx := rand.Intn(len(pieces)-1) + 1 // Genera un índice aleatorio.
	pc := pieces[idx]                   // Obtiene la pieza en el índice aleatorio.
	return piece{
		shape:     append([]vector(nil), pc.shape...), // Crea una copia de la forma de la pieza.
		canRotate: pc.canRotate,                       // Copia la propiedad canRotate.
		color:     pc.color,                           // Copia el color.
	}
}
