package file

import (
	"io"
)

func Get(r io.Reader) ([]byte, error) {
	return io.ReadAll(r)
}

func Save(
	w io.Writer,
	s string,
) error {
	_, err := io.WriteString(w, s)
	return err
}
