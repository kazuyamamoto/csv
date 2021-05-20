package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

// Reader は CSV ファイルのリーダー。 io.Closer を実装する。
type Reader struct {
	reader *csv.Reader
	file   *os.File
}

// Option は CSV ファイルを読み取るためのオプション群。
type Option struct {
	// encoding/csv/Reader.FieldsPerRecord を参照。
	FieldsPerRecord int

	// CSV ファイルの文字エンコーディング。
	Encoding encoding.Encoding
}

// OpenReader は file で指定した CSV ファイルを開き、リーダーを作成して返す。
// opt != nil の場合、指定したオプションをリーダーにセットする。
func OpenReader(file string, opt *Option) (*Reader, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, errorf("%w", err)
	}

	if opt == nil {
		return &Reader{reader: csv.NewReader(f), file: f}, nil
	}

	var r io.Reader = f
	if opt.Encoding != nil {
		r = transform.NewReader(f, opt.Encoding.NewDecoder())
	}
	c := csv.NewReader(r)
	c.FieldsPerRecord = opt.FieldsPerRecord
	return &Reader{reader: c, file: f}, nil
}

// Read は1レコードを読み取る。読み取るレコードが無い場合 nil と io.EOF を返す。
func (r *Reader) Read() ([]string, error) {
	rec, err := r.reader.Read()
	if err == nil {
		return rec, nil
	} else if err == io.EOF {
		return nil, err
	} else {
		return nil, errorf("%w", err)
	}
}

// Close は CSV ファイルを閉じる。
func (r *Reader) Close() error {
	if err := r.file.Close(); err != nil {
		return errorf("%w", err)
	}
	return nil
}

func errorf(format string, v ...interface{}) error {
	return fmt.Errorf("csv: "+format, v...)
}
