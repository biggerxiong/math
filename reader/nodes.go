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
			ID: util.MustToInt(values[0]),
			X:  values[1],
			Y:  values[2],
		}
		ret = append(ret, &row)
	})

	return ret, err
}
