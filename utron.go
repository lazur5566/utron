package utron

import (
	"io"

	"github.com/lazur5566/utron/app"
)

// NewApp creates a new bare-bone utron application. To use the MVC components, you should call
// the Init method before serving requests.
func NewApp(logOut io.Writer) *app.App {
	return app.NewApp(logOut)
}

// NewMVC creates a new MVC utron app. If cfg is passed, it should be a directory to look for
// the configuration files. The App returned is initialized.
func NewMVC(logOut io.Writer, cfg ...string) (*app.App, error) {
	return app.NewMVC(logOut, cfg...)
}
