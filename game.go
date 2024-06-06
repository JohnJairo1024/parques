package main

import (
    "github.com/fogleman/gg"
)

type Game struct {
    Board   *Board
    Players []*Player
    Dice    *Dice
    turn    int
}

func NewGame() *Game {
    players := []*Player{
        NewPlayer("Player 1", "#FF0000"),
        NewPlayer("Player 2", "#00FF00"),
        NewPlayer("Player 3", "#0000FF"),
        NewPlayer("Player 4", "#FFFF00"),
    }
    return &Game{
        Players: players,
        Dice:    NewDice(),
        turn:    0,
    }
}

func (g *Game) Draw(dc *gg.Context) {
    for _, player := range g.Players {
        player.Draw(dc)
    }
}

func (g *Game) Start() {
    const S = 1024

    // Crear el contexto de dibujo
    dc := gg.NewContext(S, S)

    // Cargar la imagen del tablero
    imgFile, err := os.Open("tablero.png")
    if err != nil {
        log.Fatalf("failed to open image: %v", err)
    }
    defer imgFile.Close()

    img, err := png.Decode(imgFile)
    if err != nil {
        log.Fatalf("failed to decode image: %v", err)
    }

    // Dibujar la imagen del tablero
    dc.DrawImage(img, 0, 0)

    // Dibujar jugadores
    g.Draw(dc)

    // Guardar la imagen final
    dc.SavePNG("output.png")
}
