package main

import (
	"app/engine"
)

func main() {
	e := engine.Setup("prod.env")
	e.Run()
}
