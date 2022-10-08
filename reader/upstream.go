package reader

import (
	"main/model"
	"main/util"
)

func ReadUpStreams(path string) ([]*model.UpStream, error) {
	r, err := NewReader(path)
	if err != nil {
		return nil, err
	}

	ret := make([]*model.UpStream, 0)
	err = r.ForeachRows(func(i int, values []string) {
		row := model.UpStream{
			ID:  util.MustToInt(values[0]),
			X:   values[1],
			Y:   values[2],
			Cap: util.MustToInt(values[3]),
		}
		ret = append(ret, &row)
	})

	return ret, err
}
