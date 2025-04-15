package logging

import (
	"io"
)

type InitParams struct {
	Level  Level
	Writer io.Writer
}
