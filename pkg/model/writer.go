package model

import (
	"io"
)

type Writer interface {
	io.Writer
	io.Closer
}
