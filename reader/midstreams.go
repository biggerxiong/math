package reader

import (
	"main/model"
	"main/util"
)

func ReadMidStreams(path string) ([]*model.MidStream, error) {
	r, err := NewReader(path)
	if err != nil {
		return nil, err
	}

	ret := make([]*model.MidStream, 0)
	err = r.ForeachRows(func(i int, values []string) {
		row := model.MidStream{
			ID:  util.MustToInt(values[0]),
			X:   values[1],
			Y:   values[2],
			Cap: util.MustToInt(values[3]),
		}
		ret = append(ret, &row)
	})

	return ret, err
}
