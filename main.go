package main

import (
	"github.com/nikl85/goresttest/app"
	"github.com/nikl85/goresttest/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
