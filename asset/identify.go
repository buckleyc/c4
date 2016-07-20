package asset

import (
	"io"
)

func Identify(src io.Reader) (*ID, error) {
	e := NewIDEncoder()
	_, err := io.Copy(e, src)
	if err != nil {
		return nil, err
	}
	return e.ID(), nil
}
