package apptweak

import (
	"encoding/json"
	"io"
)

func Bind(body io.ReadCloser, v interface{}) error {
	defer body.Close()
	err := json.NewDecoder(body).Decode(v)
	return err
}
