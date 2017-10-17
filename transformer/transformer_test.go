package transformer

import (
	"bytes"
	"testing"
)

func TestTransformer(t *testing.T) {
	tests := []struct {
		Input          []byte
		InputEncoding  string
		Output         []byte
		OutputEncoding string
		Failed         bool
	}{
		{
			Input:          []byte(""),
			InputEncoding:  "utf-8",
			Output:         []byte(""),
			OutputEncoding: "utf-8",
			Failed:         false,
		},
		{
			Input:          []byte("Gar\xe7on !"),
			InputEncoding:  "Windows1252",
			Output:         []byte("Garçon !"),
			OutputEncoding: "utf-8",
			Failed:         false,
		},
		{
			Input:          []byte("Garçon !"),
			InputEncoding:  "utf-8",
			Output:         []byte("Gar\xe7on !"),
			OutputEncoding: "windows-1252",
			Failed:         false,
		},
	}

	for i, tt := range tests {
		src := bytes.NewBuffer(tt.Input)
		var dst bytes.Buffer

		tr, err := New(src, &dst, tt.InputEncoding, tt.OutputEncoding)
		if err != nil {
			if !tt.Failed {
				t.Fatalf("#%d: expected to fail", i)
			}
		}
		if tt.Failed {
			t.Fatalf("#%d: expected to succeed", i)
		}

		err = tr.Transform()
		if err != nil {
			if !tt.Failed {
				t.Fatalf("#%d: expected to fail", i)
			}
		}
		if tt.Failed {
			t.Fatalf("#%d: expected to succeed", i)
		}

		if want, have := tt.Output, dst.Bytes(); !bytes.Equal(want, have) {
			t.Fatalf("#%d: expected %q, got %q", i, want, have)
		}
	}
}
