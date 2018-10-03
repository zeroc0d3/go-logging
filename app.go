package main

import (
	"bytes"
	"io"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/dimiro1/banner"
	_ "github.com/dimiro1/banner/autoload"
	"github.com/mattn/go-colorable"
)

const LOG_FILE = "/tmp/mylog.log"

func InitLogo() {
	isEnabled := true
	isColorEnabled := true
	banner.Init(colorable.NewColorableStdout(), isEnabled, isColorEnabled, bytes.NewBufferString("GO-Boilerplate {{ .AnsiColor.Green }}(Running){{ .AnsiColor.Default }}\n"))
}

func main() {
	InitLogo()
	println("Log file saved to:", LOG_FILE)

	// create the logger
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}

	fName, err := os.OpenFile(LOG_FILE, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		logger.Fatal(err)
	}
	defer fName.Close()
	// multiwriter simultaneously
	logger.SetOutput(io.MultiWriter(os.Stdout, fName))

    // YOUR RUNNING APPLICATION
    // ---> here
}

