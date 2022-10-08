package reader

import (
	"main/model"
	"main/util"
)

func ReadStreets(path string) ([]*model.Street, error) {
	r, err := NewReader(path)
	if err != nil {
		return nil, err
	}

	ret := make([]*model.Street, 0)
	err = r.ForeachRows(func(i int, values []string) {
		row := model.Street{
			ID:            util.MustToInt(values[0]),
			BuildingCount: util.MustToInt(values[1]),
			FamilyCount:   util.MustToInt(values[2]),
			PeopleCount:   util.MustToInt(values[3]),
			X:             values[4],
			Y:             values[5],
			StreetIndex:   values[6],
			BelongTo:      values[7],
		}
		ret = append(ret, &row)
	})

	return ret, err
}
