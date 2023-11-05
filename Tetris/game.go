package Tetris

import "time"

// Constantes para la altura y anchura del tablero del juego.
const (
	BOARD_HEIGHT = 16
	BOARD_WIDTH  = 10
)

// Constantes para los estados del juego.
const (
	gameInit gameState = iota // Juego no iniciado
	gamePlay                  // Juego en progreso
	gameOver                  // Juego terminado
)

// gameState representa el estado actual del juego.
type gameState int

// game representa el estado actual del juego de Tetris.
type game struct {
	board     [][]int     // El tablero del juego
	piece     piece       // La pieza actual que se está jugando
	position  vector      // La posición actual de la pieza en el tablero
	state     gameState   // El estado actual del juego
	FallSpeed *time.Timer // Temporizador para controlar la velocidad de caída de las piezas
}

// blockOnBoardByPosition calcula la posición de un bloque en el tablero.
func (g *game) blockOnBoardByPosition(v vector) vector {
	py := g.position.y + v.y
	px := g.position.x + v.x
	return vector{py, px}
}

// colision comprueba si la pieza actual colisiona con los bordes del tablero o con otras piezas.
func (g *game) colision() bool {
	for _, v := range g.piece.shape {
		pos := g.blockOnBoardByPosition(v)

		// Comprueba si la pieza está fuera de los límites del tablero.
		if pos.x < 0 || pos.x >= BOARD_WIDTH || pos.y < 0 || pos.y >= BOARD_HEIGHT {
			return true
		}
		// Comprueba si la pieza colisiona con otra pieza en el tablero.
		if pos.y < len(g.board) && pos.x < len(g.board[pos.y]) && g.board[pos.y][pos.x] > 0 {
			return true
		}
	}
	return false
}

// MoveLeft mueve la pieza actual hacia la izquierda si es posible.
func (g *game) MoveLeft() {
	g.moveIfPossible(vector{0, -1})
}

// MoveRight mueve la pieza actual hacia la derecha si es posible.
func (g *game) MoveRight() {
	g.moveIfPossible(vector{0, 1})
}

// SpeedUp aumenta la velocidad de caída de la pieza actual.
func (g *game) SpeedUp() {
	if g.state != gamePlay {
		return
	}
	g.FallSpeed.Reset(50)
}

// Rotate rota la pieza actual si es posible.
func (g *game) Rotate() {
	if g.state != gamePlay {
		return
	}
	g.piece.rotate()
	if g.colision() {
		g.piece.rotateBack()
	}
}

// Fall hace caer la pieza actual si es posible.
func (g *game) Fall() {
	if g.state != gamePlay {
		return
	}
	for {
		if g.moveIfPossible(vector{1, 0}) {
			g.FallSpeed.Reset(1 * time.Millisecond)
			return
		}
	}
}

// moveIfPossible mueve la pieza actual en la dirección especificada si es posible.
func (g *game) moveIfPossible(v vector) bool {
	if g.state != gamePlay {
		return false
	}
	g.position.x += v.x
	g.position.y += v.y
	if g.colision() {
		g.position.x -= v.x
		g.position.y -= v.y
		return false
	}
	return true
}

// lockpiece bloquea la pieza actual en su posición en el tablero.
func (g *game) lockpiece() {
	g.board = g.GetBoard()
}

// removeLines elimina todas las líneas completas del tablero.
func (g *game) removeLines() {
	el := make([]int, BOARD_WIDTH)
	for i := 0; i < BOARD_WIDTH; i++ {
		el[i] = 0
	}
	emptyLine := [][]int{el}
	newBoard := make([][]int, len(g.board))
	copy(newBoard, g.board)

	for y := 0; y < BOARD_HEIGHT; y++ {
		fullLine := true
		for x := 0; x < BOARD_WIDTH; x++ {
			if g.board[y][x] == 0 {
				fullLine = false
				break
			}
		}
		if fullLine {
			newBoard = append(emptyLine, newBoard[:y]...)
			newBoard = append(newBoard, newBoard[y+1:]...)
		}
	}
	g.board = newBoard
}

// GameLoop es el bucle principal del juego que se ejecuta en cada tick del juego.
func (g *game) GameLoop() {
	if !g.moveIfPossible(vector{1, 0}) {
		g.lockpiece()
		g.removeLines()
		g.getPiece()
		if g.colision() {
			g.FallSpeed.Stop()
			g.state = gameOver
			return
		}

	}
	g.resetFallSpeed()
}

// GetBoard devuelve el estado actual del tablero con la pieza actual.
func (g *game) GetBoard() [][]int {
	cBoard := make([][]int, len(g.board))
	for i := range g.board {
		cBoard[i] = make([]int, len(g.board[i]))
		copy(cBoard[i], g.board[i])
	}

	for _, v := range g.piece.shape {
		pos := g.blockOnBoardByPosition(v)
		if pos.y >= 0 && pos.y < len(cBoard) && pos.x >= 0 && pos.x < len(cBoard[pos.y]) {
			cBoard[pos.y][pos.x] = g.piece.color
		}
	}
	return cBoard
}

// getPiece obtiene una nueva pieza aleatoria y la coloca en la posición inicial en el tablero.
func (g *game) getPiece() {
	g.piece = randomPiece()
	g.position = vector{0, BOARD_WIDTH / 2}
}

// resetFallSpeed restablece la velocidad de caída de la pieza actual.
func (g *game) resetFallSpeed() {
	g.FallSpeed.Reset(700 * time.Millisecond)
}

// Start inicia el juego si no está en progreso.
func (g *game) Start() {
	if g.state == gamePlay {
		return
	}
	g.state = gamePlay
	g.getPiece()
	g.resetFallSpeed()
}

// init inicializa el estado del juego.
func (g *game) init() {
	g.board = make([][]int, BOARD_HEIGHT)
	for y := 0; y < BOARD_HEIGHT; y++ {
		g.board[y] = make([]int, BOARD_WIDTH)
		for x := 0; x < BOARD_WIDTH; x++ {
			g.board[y][x] = 0
		}
	}
	g.position = vector{1, BOARD_WIDTH / 2}
	g.piece = pieces[0]
	g.FallSpeed = time.NewTimer(time.Duration(1000 * time.Second))
	g.FallSpeed.Stop()
	g.state = gameInit
}

// New crea y devuelve una nueva instancia del juego.
func New() *game {
	g := &game{}
	g.init()
	return g
}
