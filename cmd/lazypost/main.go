package main

import (
	"log"

	"github.com/jublx/lazypost/internal/ui"
)

func main() {
	p := ui.NewProgram()
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
