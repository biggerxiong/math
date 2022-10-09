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
				X: util.StringMustToFloat(values[1]),
				Y: util.StringMustToFloat(values[2]),
			},
			Cap: util.StringMustToDecimal(values[3]),
		}
		ret = append(ret, &row)
	})

	return ret, err
}
