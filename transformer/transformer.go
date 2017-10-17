package transformer

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/unicode"
)

type Transformer struct {
	src io.Reader
	dst io.Writer
}

func New(r io.Reader, w io.Writer, inputEncoding, outputEncoding string) (*Transformer, error) {
	src, err := encodingReader(r, inputEncoding)
	if err != nil {
		return nil, err
	}
	dst, err := encodingWriter(w, outputEncoding)
	if err != nil {
		return nil, err
	}
	return &Transformer{src: src, dst: dst}, nil
}

func (t *Transformer) Transform() error {
	_, err := io.Copy(t.dst, t.src)
	return err
}

func encodingReader(r io.Reader, encoding string) (io.Reader, error) {
	switch strings.ToLower(encoding) {
	default:
		return nil, fmt.Errorf("unknown encoding %q", encoding)
	case "utf-8", "utf8":
		return unicode.UTF8.NewDecoder().Reader(r), nil
	case "windows1252", "windows-1252":
		return charmap.Windows1252.NewDecoder().Reader(r), nil
	case "iso88591", "iso-8859-1", "iso8859-1", "iso_8859_1", "iso8859_1":
		return charmap.ISO8859_1.NewDecoder().Reader(r), nil
	case "iso88592", "iso-8859-2", "iso8859-2", "iso_8859_2", "iso8859_2":
		return charmap.ISO8859_2.NewDecoder().Reader(r), nil
	case "iso88593", "iso-8859-3", "iso8859-3", "iso_8859_3", "iso8859_3":
		return charmap.ISO8859_3.NewDecoder().Reader(r), nil
	case "iso88594", "iso-8859-4", "iso8859-4", "iso_8859_4", "iso8859_4":
		return charmap.ISO8859_4.NewDecoder().Reader(r), nil
	case "iso88595", "iso-8859-5", "iso8859-5", "iso_8859_5", "iso8859_5":
		return charmap.ISO8859_5.NewDecoder().Reader(r), nil
	}
}

func encodingWriter(w io.Writer, encoding string) (io.Writer, error) {
	switch strings.ToLower(encoding) {
	default:
		return nil, fmt.Errorf("unknown encoding %q", encoding)
	case "utf-8", "utf8":
		return unicode.UTF8.NewEncoder().Writer(w), nil
	case "windows1252", "windows-1252":
		return charmap.Windows1252.NewEncoder().Writer(w), nil
	case "iso88591", "iso-8859-1", "iso8859-1", "iso_8859_1", "iso8859_1":
		return charmap.ISO8859_1.NewEncoder().Writer(w), nil
	case "iso88592", "iso-8859-2", "iso8859-2", "iso_8859_2", "iso8859_2":
		return charmap.ISO8859_2.NewEncoder().Writer(w), nil
	case "iso88593", "iso-8859-3", "iso8859-3", "iso_8859_3", "iso8859_3":
		return charmap.ISO8859_3.NewEncoder().Writer(w), nil
	case "iso88594", "iso-8859-4", "iso8859-4", "iso_8859_4", "iso8859_4":
		return charmap.ISO8859_4.NewEncoder().Writer(w), nil
	case "iso88595", "iso-8859-5", "iso8859-5", "iso_8859_5", "iso8859_5":
		return charmap.ISO8859_5.NewEncoder().Writer(w), nil
	}
}
