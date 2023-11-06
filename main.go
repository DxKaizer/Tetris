package main

import (
	"math/rand"
	screen "tetris/Screen"
	"tetris/Tetris"
	"time"

	"github.com/nsf/termbox-go"
)

func main() {
	// Define la velocidad de animación.
	const animationSpeed = 50 * time.Millisecond

	// Inicializa el generador de números aleatorios.
	rand.Seed(time.Now().UnixNano())

	// Inicializa la biblioteca termbox para la entrada y salida de la terminal.
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	// Asegúrate de cerrar termbox al salir.
	defer termbox.Close()

	// Crea un canal para los eventos de termbox.
	evenQueue := make(chan termbox.Event)
	go func() {
		for {
			// Espera a que ocurra un evento y luego lo envía al canal.
			evenQueue <- termbox.PollEvent()
		}
	}()

	// Crea un temporizador para controlar la velocidad de animación.
	ticker := time.NewTimer(time.Duration(animationSpeed))

	// Crea una nueva instancia del juego Tetris.
	game := Tetris.New()

	// Crea una nueva instancia de la pantalla del juego.
	src := screen.New()

	// Bucle principal del juego.
	for {
		select {
		case ev := <-evenQueue:
			// Maneja los eventos de teclado.
			if ev.Type == termbox.EventKey {
				switch {
				case ev.Key == termbox.KeyArrowUp:
					game.Rotate() // Rota la pieza.
				case ev.Key == termbox.KeyArrowDown:
					game.SpeedUp() // Acelera la caída de la pieza.
				case ev.Key == termbox.KeyArrowLeft:
					game.MoveLeft() // Mueve la pieza a la izquierda.
				case ev.Key == termbox.KeyArrowRight:
					game.MoveRight() // Mueve la pieza a la derecha.
				case ev.Key == termbox.KeySpace:
					game.Fall() // Hace caer la pieza rápidamente.
				case ev.Ch == 'q':
					return // Sale del juego.
				case ev.Ch == 's':
					game.Start() // Inicia el juego.
				}
			}
		case <-ticker.C:
			// Renderiza el tablero del juego.
			src.Render(game.GetBoard())
			// Reinicia el temporizador.
			ticker.Reset(animationSpeed)
		case <-game.FallSpeed.C:
			// Ejecuta el bucle principal del juego.
			game.GameLoop()
		}
	}
}
