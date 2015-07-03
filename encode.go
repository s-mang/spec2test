package spec2test

import (
	"encoding/json"
	"encoding/xml"
	"io"
)

type Encoder interface {
	Encode(v interface{}) error
}

type Decoder interface {
	Decode(v interface{}) error
}

func Encode(contentType string, v interface{}, w io.Writer) error {
	var encoder Encoder
	switch contentType {
	case "application/json":
		encoder = Encoder(json.NewEncoder(w))

	case "text/xml", "application/xml":
		encoder = Encoder(xml.NewEncoder(w))

	default:
		return ErrorContentTypeNotSupported
	}

	return encoder.Encode(v)
}

func Decode(contentType string, src io.Reader, dest interface{}) error {
	var decoder Decoder
	switch contentType {
	case "application/json":
		decoder = Decoder(json.NewDecoder(src))

	case "text/xml", "application/xml":
		decoder = Decoder(xml.NewDecoder(src))

	default:
		return ErrorContentTypeNotSupported
	}

	return decoder.Decode(dest)
}
