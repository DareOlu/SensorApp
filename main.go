package main

import (
	_ "embed"
	"time"

	streamdata "github.com/mm/front-side-v005/backend"
	broker "github.com/mm/front-side-v005/backend/broker"
	"github.com/wailsapp/wails"
)

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func main() {
	//stats := &simul.Stats{}
	broker.Connectbroker()
	time.Sleep(1200 * time.Millisecond)

	streamdata := &streamdata.Stats{}
	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "front-side-v005",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(streamdata)
	app.Run()
}
