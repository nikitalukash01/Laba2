package calc

import (
	//"bufio"
	"io"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}