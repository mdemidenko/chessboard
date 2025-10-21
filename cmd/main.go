package main

import (
	"chessboard/internal/board"
	"chessboard/internal/validator"
	"fmt"
	"log"
	"os"
)

func main() {
	var size int

	if len(os.Args) > 1 {
		validatedSize, err := validator.ValidateSize(os.Args[1])
		if err != nil {
			log.Fatalf("Ошибка валидации: %v", err)
		}
		size = validatedSize
	} else {
		size = 8 // значение по умолчанию
	}

	fmt.Printf("Генерируем доску %dx%d:\n", size, size)
	chessBoard := board.GenerateBoard(size)
	fmt.Println(chessBoard)
}
