package imgcat

import (
	"encoding/base64"
	"io"
	"strings"

	"github.com/pkg/errors"
)

func Copy(w io.Writer, r io.Reader) error {
	pr, pw := io.Pipe()
	header := strings.NewReader("\033]1337;File=inline=1:")
	footer := strings.NewReader("\a\n")

	go func() {
		defer pw.Close()
		wc := base64.NewEncoder(base64.StdEncoding, pw)
		_, err := io.Copy(wc, r)
		if err != nil {
			pw.CloseWithError(errors.Wrap(err, "could not encode image"))
			return
		}

		if err = wc.Close(); err != nil {
			pw.CloseWithError(errors.Wrap(err, "Unable to close base64 encoder"))
			return
		}
	}()

	_, err := io.Copy(w, io.MultiReader(header, pr, footer))

	return err
}

// NewWriter to cat the image on terminal
func NewWriter(w io.Writer) io.WriteCloser {
	pr, pw := io.Pipe()
	wc := &writer{
		pw,
		make(chan bool),
	}
	go func() {
		defer close(wc.done)
		err := Copy(w, pr)
		pr.CloseWithError(err)
	}()
	return wc
}

type writer struct {
	pw   *io.PipeWriter
	done chan bool
}

func (w *writer) Write(data []byte) (int, error) {
	return w.pw.Write(data)
}

func (w *writer) Close() error {
	if err := w.pw.Close(); err != nil {
		return err
	}

	<-w.done
	return nil
}
