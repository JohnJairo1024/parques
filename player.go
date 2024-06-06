package main

import "github.com/fogleman/gg"

type Player struct {
    Name   string
    Pieces []*Piece
    Color  string // Color de las piezas del jugador
}

func NewPlayer(name, color string) *Player {
    pieces := make([]*Piece, 4)
    for i := range pieces {
        pieces[i] = &Piece{Position: -1} // Initial position (-1 means not on the board)
    }
    return &Player{
        Name:   name,
        Pieces: pieces,
        Color:  color,
    }
}

func (p *Player) Draw(dc *gg.Context) {
    dc.SetHexColor(p.Color)
    for _, piece := range p.Pieces {
        if piece.Position != -1 {
            x, y := getPositionCoords(piece.Position)
            dc.DrawCircle(float64(x), float64(y), 10)
            dc.Fill()
        }
    }
}

func getPositionCoords(position int) (int, int) {
    // Esta función debe retornar las coordenadas X, Y en el tablero para la posición dada
    // Aquí debes definir las coordenadas exactas para tu tablero
    return 100 + position*10, 100 + position*10 // Ejemplo
}
