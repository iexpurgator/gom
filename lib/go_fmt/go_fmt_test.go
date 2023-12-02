package go_fmt

import (
	"io"
	"testing"
)

func TestStdoutWriter_Println(t *testing.T) {
	type fields struct {
		writeOut  io.Writer
		writeErr  io.Writer
		printable bool
		colormode bool
	}
	type args struct {
		a []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantN   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := StdoutWriter{
				writeOut:  tt.fields.writeOut,
				writeErr:  tt.fields.writeErr,
				printable: tt.fields.printable,
				colormode: tt.fields.colormode,
			}
			gotN, err := w.Println(tt.args.a...)
			if (err != nil) != tt.wantErr {
				t.Errorf("StdoutWriter.Println() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("StdoutWriter.Println() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestStdoutWriter_Printf(t *testing.T) {
	type fields struct {
		writeOut  io.Writer
		writeErr  io.Writer
		printable bool
		colormode bool
	}
	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantN   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := StdoutWriter{
				writeOut:  tt.fields.writeOut,
				writeErr:  tt.fields.writeErr,
				printable: tt.fields.printable,
				colormode: tt.fields.colormode,
			}
			gotN, err := w.Printf(tt.args.format, tt.args.a...)
			if (err != nil) != tt.wantErr {
				t.Errorf("StdoutWriter.Printf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("StdoutWriter.Printf() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
