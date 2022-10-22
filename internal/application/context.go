package application

import (
	"github.com/rs/zerolog"
	"matrix-alertmanager/internal/config"
	"matrix-alertmanager/internal/matrix"
	"matrix-alertmanager/internal/queue"
)

type Context struct {
	Config  config.Config
	Clients map[string]matrix.Matrix
	Queue   *queue.Queue
	Logic   Logic
	Logger  zerolog.Logger
}
