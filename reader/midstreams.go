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
			ID: util.StringMustToInt(values[0]),
			Point: model.Point{
				X: values[1],
				Y: values[2],
			},
			Cap: util.StringMustToInt(values[3]),
		}
		ret = append(ret, &row)
	})

	return ret, err
}
