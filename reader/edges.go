package reader

import (
	"main/model"
	"main/util"
)

func ReadEdges(path string) ([]*model.Edge, error) {
	r, err := NewReader(path)
	if err != nil {
		return nil, err
	}

	ret := make([]*model.Edge, 0)
	err = r.ForeachRows(func(i int, values []string) {
		row := model.Edge{
			ID:   util.StringMustToInt(values[0]),
			From: util.StringMustToInt(values[1]),
			To:   util.StringMustToInt(values[2]),
			Dis:  util.StringMustToFloat(values[3]),
		}
		ret = append(ret, &row)
	})

	return ret, err
}
