package main

import "github.com/poeticalcode/gworkflow/internal/engine"

func main() {
	wfe := engine.Default()
	wfe.Deploy(nil, "")
}
