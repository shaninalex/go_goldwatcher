package main

import (
	"log"
	"net/http"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type Config struct {
	App            fyne.App
	InfoLog        *log.Logger
	ErrorLog       *log.Logger
	MainWindow     fyne.Window
	PriceContainer *fyne.Container
	HTTPClient     *http.Client
}

var (
	myApp Config
)

func main() {
	// create fyne Application
	fyneApp := app.NewWithID("shaninpersonal.xyz.goldwatcher.preferences")
	myApp.App = fyneApp
	myApp.HTTPClient = &http.Client{}

	// create our loggers
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "Error\t", log.Ldate|log.Ltime|log.Lshortfile)

	// open a connection to the database

	// create a database repository

	// create and size a fyne window
	myApp.MainWindow = fyneApp.NewWindow("GoldWatcher")
	myApp.MainWindow.Resize(fyne.NewSize(770, 410))
	myApp.MainWindow.SetFixedSize(true)
	myApp.MainWindow.SetMaster() // this is the main window of the application

	myApp.makeUI()

	// show and run the application
	myApp.MainWindow.ShowAndRun()
}
