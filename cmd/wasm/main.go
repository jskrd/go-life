package main

import (
	"github.com/jskrd/go-life/internal/models"
)

func main() {
	done := make(chan struct{})

	game := models.NewGame()
	game.Start()

	<-done
}
