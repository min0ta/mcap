package log

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

type LoggerMode int

const (
	FileMode LoggerMode = iota
	ConsoleMode
	BothMode
)

type Logger struct {
	Mode     LoggerMode
	Fallback LoggerMode
	out      func(a any) error
	Close    func() error // only for sigterm/sigstop handling
}

func New(mode LoggerMode, filepath string) *Logger {
	var out func(a any) error
	var close func() error
	switch mode {

	case FileMode:
		out, close = createFileLogger(filepath)

	case ConsoleMode:
		out = func(a any) error {
			_, err := fmt.Println(a)
			if err != nil {
				return err
			}
			return nil
		}
		close = func() error { return nil }

	case BothMode:
		var fileLogger func(a any) error
		fileLogger, close = createFileLogger(filepath)

		out = func(a any) error {
			err := fileLogger(a)
			if err != nil {
				return err
			}
			_, err = fmt.Println(a)
			if err != nil {
				return err
			}
			return nil
		}

	default:
		log.Fatal("UNKNOWN LOGGER MODE", mode, "\n MODE CAN BE ONLY 'FILE', 'CONSOLE' OR 'BOTH'")

	}

	return &Logger{
		Mode:  mode,
		out:   out,
		Close: close,
	}
}

func (l *Logger) WriteString(s string) {
	l.out(s)
}
func (l *Logger) WriteFormat(s string, args ...any) {
	l.out(fmt.Sprintf(s, args...))
}

func createFileLogger(loggerPath string) (func(a any) error, func() error) {
	file, err := os.OpenFile(loggerPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if errors.Is(err, os.ErrNotExist) {
		file, err = os.Create(loggerPath)
		if err != nil {
			log.Fatalf("FAILED TO CREATE LOG FILE! ERROR %v", err)
		}
	} else if err != nil {
		log.Fatalf("FAILED TO OPEN LOG FILE! ERROR: %v", err)
	}
	return func(a any) error {
		json, err := json.Marshal(a)
		if err != nil {
			return err
		}
		file.WriteString("---\n" + string(json) + "\n" + time.Now().String() + "\n---\n\n")
		return nil
	}, file.Close
}
