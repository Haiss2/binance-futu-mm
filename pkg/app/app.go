package app

import (
	"io"
	"os"

	"github.com/urfave/cli"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewApp creates a new cli App instance with common flags pre-loaded.
func NewApp() *cli.App {
	app := cli.NewApp()
	app.Flags = NewAppFlag()
	app.Flags = append(app.Flags, PostgresSQLFlags()...)
	return app
}

func NewLogger(c *cli.Context) (*zap.Logger, zap.AtomicLevel) {
	var writers = []io.Writer{os.Stdout}

	w := io.MultiWriter(writers...)

	atom := zap.NewAtomicLevelAt(zap.DebugLevel)
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.RFC3339TimeEncoder
	config.CallerKey = "caller"
	encoder := zapcore.NewConsoleEncoder(config)
	cc := zap.New(zapcore.NewCore(encoder, zapcore.AddSync(w), atom), zap.AddCaller())
	return cc, atom
}
