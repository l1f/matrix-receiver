package application

import (
	"github.com/rs/zerolog"
	"matrix-alertmanager/internal/config"
	"matrix-alertmanager/internal/matrix"
)

type Context struct {
	Config  config.Config
	Clients map[string]matrix.Matrix
	Logic   Logic
	Logger  zerolog.Logger
}
