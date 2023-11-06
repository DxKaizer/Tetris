package Tetris

import "time"

const (
	BOARD_HEIGHT = 16
	BOARD_WIDTH  = 10
)
const (
	gameInit gameState = iota
	gamePlay
	gameOver
)

type gameState int

type game struct {
	board     [][]int
	piece     piece
	position  vector
	state     gameState
	FallSpeed *time.Timer
}

func (g *game) blockOnBoardByPosition(v vector) vector {
	py := g.position.y + v.y
	px := g.position.x + v.x
	return vector{py, px}

}

func (g *game) colision() bool {
	for _, v := range g.piece.shape {
		pos := g.blockOnBoardByPosition(v)

		if pos.x < 0 || pos.x >= BOARD_WIDTH || pos.y < 0 || pos.y >= BOARD_HEIGHT {
			return true
		}
		if pos.y < len(g.board) && pos.x < len(g.board[pos.y]) && g.board[pos.y][pos.x] > 0 {
			return true
		}
	}
	return false
}

func (g *game) MoveLeft() {
	g.moveIfPossible(vector{0, -1})
}
func (g *game) MoveRight() {
	g.moveIfPossible(vector{0, 1})
}

func (g *game) SpeedUp() {
	if g.state != gamePlay {
		return
	}
	g.FallSpeed.Reset(50)
}
func (g *game) Rotate() {
	if g.state != gamePlay {
		return
	}
	g.piece.rotate()
	if g.colision() {
		g.piece.rotateBack()
	}

}
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

func (g *game) lockpiece() {
	g.board = g.GetBoard()
}

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

func (g *game) getPiece() {
	g.piece = randomPiece()
	g.position = vector{0, BOARD_WIDTH / 2}
}
func (g *game) resetFallSpeed() {
	g.FallSpeed.Reset(700 * time.Millisecond)
}
func (g *game) Start() {
	if g.state == gamePlay {
		return
	}
	g.state = gamePlay
	g.getPiece()
	g.resetFallSpeed()
}

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

func New() *game {
	g := &game{}
	g.init()
	return g
}
