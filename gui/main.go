package gui

import "github.com/webview/webview"

func Start() {
	window := webview.New(true)
	defer window.Destroy()
	window.SetTitle("6502 Tools")
	window.SetSize(1280, 720, webview.HintNone)
	window.Navigate("http://localhost:2543")
	window.Run()
}
