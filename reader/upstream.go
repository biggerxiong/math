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
			ID: util.StringMustToInt(values[0]),
			Point: model.Point{
				X: util.StringMustToFloat(values[1]),
				Y: util.StringMustToFloat(values[2]),
			},
			Cap: util.StringMustToDecimal(values[3]),
		}
		ret = append(ret, &row)
	})

	return ret, err
}
