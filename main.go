package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func start() {
	gtk.Init(nil)

	builder, err := gtk.BuilderNewFromFile("ui.glade")
	if err != nil {
		log.Fatalln(err)
	}

	obj, err := builder.GetObject("win")
	if err != nil {
		log.Fatalln(err)
	}

	win := obj.(*gtk.Window)

	win.ShowAll()
}

func main() {
	start()
}
