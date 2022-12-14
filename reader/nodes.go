package reader

import (
	"main/model"
	"main/util"
)

func ReadNodes(path string) ([]*model.Node, error) {
	r, err := NewReader(path)
	if err != nil {
		return nil, err
	}

	ret := make([]*model.Node, 0)
	err = r.ForeachRows(func(i int, values []string) {
		row := model.Node{
			ID: util.StringMustToInt(values[0]),
			Point: model.Point{
				X: util.StringMustToFloat(values[1]),
				Y: util.StringMustToFloat(values[2]),
			},
		}
		ret = append(ret, &row)
	})

	return ret, err
}
