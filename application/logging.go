package application

import (
	"fmt"
	"os"
)

func (app *App) LogErrorf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	app.logger.Printf("[Error]: %s\n", msg)
}

func (app *App) LogInfo(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	app.logger.Printf("[Info]: %s\n", msg)
}

func (app *App) LogFatalf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	app.logger.Printf("[Error]: %s\n", msg)
	os.Exit(1)
}
