package go_fmt

import (
	"fmt"
	"io"

	"github.com/mattn/go-colorable"
)

var (
	Console Writer
)

func init() {
	Console = &StdoutWriter{
		writeOut:  colorable.NewColorableStdout(),
		writeErr:  colorable.NewColorableStderr(),
		printable: true,
		colormode: true,
	}
}

type Writer interface {
	Println(a ...interface{}) (n int, err error)
	Printf(format string, a ...interface{}) (n int, err error)
	Errorln(a ...interface{}) (n int, err error)
	Errorf(format string, a ...interface{}) (n int, err error)
	SetPrintable()
	UnsetPrintable()
	SetPrintcolor()
	UnsetPrintcolor()
}

type StdoutWriter struct {
	writeOut  io.Writer
	writeErr  io.Writer
	printable bool
	colormode bool
}

func (w StdoutWriter) Println(a ...interface{}) (n int, err error) {
	if !w.printable {
		return 0, nil
	}

	return fmt.Fprintln(w.writeOut, a...)
}

func (w StdoutWriter) Printf(format string, a ...interface{}) (n int, err error) {
	if !w.printable {
		return 0, nil
	}

	return fmt.Fprintf(w.writeOut, format, a...)
}

func (w StdoutWriter) Errorln(a ...interface{}) (n int, err error) {
	if !w.printable {
		return 0, nil
	}

	return fmt.Fprintln(w.writeErr, a...)
}

func (w StdoutWriter) Errorf(format string, a ...interface{}) (n int, err error) {
	if !w.printable {
		return 0, nil
	}

	return fmt.Fprintf(w.writeErr, format, a...)
}

func (w *StdoutWriter) SetPrintable() {
	w.printable = true
}

func (w *StdoutWriter) UnsetPrintable() {
	w.printable = false
}

func (w *StdoutWriter) SetPrintcolor() {
	w.colormode = true
	w.writeOut = colorable.NewColorableStdout()
	w.writeErr = colorable.NewColorableStderr()
}

func (w *StdoutWriter) UnsetPrintcolor() {
	w.colormode = false
	w.writeOut = colorable.NewNonColorable(w.writeOut)
	w.writeErr = colorable.NewNonColorable(w.writeErr)
}
