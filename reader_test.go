package csv

import (
	"fmt"
	"golang.org/x/text/encoding/japanese"
	"io"
	"testing"
)

func TestReader_Read(t *testing.T) {
	r, err := OpenReader("testdata/utf8.csv", nil)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	rec, err := r.Read()
	if err != nil {
		t.Fatal(err)
	}

	if err := assertStrArrayEqual([]string{"あ", "い"}, rec); err != nil {
		t.Fatal(err)
	}

	if _, err := r.Read(); err != io.EOF {
		t.Fatal("err should be io.EOF")
	}
}

func TestReader_ReadShiftJIS(t *testing.T) {
	r, err := OpenReader("testdata/sjis.csv", &Option{Encoding: japanese.ShiftJIS})
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	rec, err := r.Read()
	if err != nil {
		t.Fatal(err)
	}

	if err := assertStrArrayEqual([]string{"あ", "い"}, rec); err != nil {
		t.Fatal(err)
	}

	if _, err := r.Read(); err != io.EOF {
		t.Fatal("err should be io.EOF")
	}
}

func assertStrArrayEqual(x, y []string) error {
	if len(x) != len(y) {
		return fmt.Errorf("array length should be equal: len(x)=%d, len(y)=%d", len(x), len(y))
	}

	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			return fmt.Errorf("element at index %d should be equal: x[%d]=%s, y[%d]=%s", i, i, x[i], i, y[i])
		}
	}

	return nil
}
