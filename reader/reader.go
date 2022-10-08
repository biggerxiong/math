package reader

import (
	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"
)

var DefaultSheet = "Sheet1"

type XlsxReader struct {
	path string
	f    *excelize.File
}

func NewReader(path string) (*XlsxReader, error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return nil, err
	}

	return &XlsxReader{path: path, f: f}, nil
}

func (x *XlsxReader) ForeachRows(f func(i int, values []string)) error {
	rows, err := x.f.GetRows(DefaultSheet)
	if err != nil {
		return errors.Wrap(err, "get rows error")
	}

	for i, row := range rows {
		if i == 0 {
			continue
		}
		
		f(i, row)
	}

	return nil
}
