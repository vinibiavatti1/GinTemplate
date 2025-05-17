package main

import (
	"app/engine"
)

func main() {
	e := engine.Setup("dev.env")
	e.Run()
}
