package main

import (
    "image"
    "image/png"
    "log"
    "os"

    "github.com/fogleman/gg"
)

func main() {
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

    // Guardar la imagen final para verificar
    dc.SavePNG("output.png")
}
