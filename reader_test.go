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

func TestReader_ReadDifferentFieldsPerRecord(t *testing.T) {
	r, err := OpenReader("testdata/diff_fields.csv", nil)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	_, err = r.Read()
	if err != nil {
		t.Fatal(err)
	}

	_, err = r.Read()
	if err == nil {
		t.Fatal("error should be non-nil")
	}
}

func TestReader_ReadDifferentFieldsPerRecord2(t *testing.T) {
	r, err := OpenReader("testdata/diff_fields.csv", &Option{FieldsPerRecord: -1})
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	_, err = r.Read()
	if err != nil {
		t.Fatal(err)
	}

	rec, err := r.Read()
	if err != nil {
		t.Fatal(err)
	}

	if err := assertStrArrayEqual([]string{"か"}, rec); err != nil {
		t.Fatal(err)
	}
}

func assertStrArrayEqual(want, got []string) error {
	if len(want) != len(got) {
		return fmt.Errorf("len: want %d, but got %d", len(want), len(got))
	}

	for i := 0; i < len(want); i++ {
		if want[i] != got[i] {
			return fmt.Errorf("index %d: want <%s>, but got <%s>", i, want[i], got[i])
		}
	}

	return nil
}
