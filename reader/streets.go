package reader

import (
	"main/config"
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
			ID:            util.StringMustToInt(values[0]),
			BuildingCount: util.StringMustToInt(values[1]),
			FamilyCount:   util.StringMustToInt(values[2]),
			PeopleCount:   util.StringMustToInt(values[3]),
			Point: model.Point{
				X: values[4],
				Y: values[5],
			},
			StreetIndex: values[6],
			BelongTo:    values[7],
		}

		row.Cap = util.IntMustToDecimal(row.PeopleCount).
			Mul(util.StringMustToDecimal(config.GetConfig().FoodsPerPerson)).
			String()
		ret = append(ret, &row)
	})

	return ret, err
}
