package v1

import (
	"fmt"
	"main/config"
	"main/model"
	"os"
	"testing"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

func init() {
	err := config.InitConfig("../config/config_test.toml")
	if err != nil {
		panic(err)
	}
	fmt.Printf("config: %+v\n", config.GetConfig())

	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(false)
	logrus.SetFormatter(&nested.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"component", "category"},
	})
}

func TestAlgo(t *testing.T) {
	pStreet1 := model.Point{X: 2, Y: 1}
	pStreet2 := model.Point{X: 2, Y: 2}
	pStreet3 := model.Point{X: 2, Y: 3}

	pMid1 := model.Point{X: 1, Y: 1}
	pMid2 := model.Point{X: 1, Y: 2}

	streets := []*model.Street{
		{
			ID:          1,
			Point:       pStreet1,
			PeopleCount: 4,
			Cap:         decimal.NewFromInt(4),
		},
		{
			ID:          2,
			Point:       pStreet2,
			PeopleCount: 8,
			Cap:         decimal.NewFromInt(8),
		},
		{
			ID:          3,
			Point:       pStreet3,
			PeopleCount: 10,
			Cap:         decimal.NewFromInt(10),
		},
	}

	midStreams := []*model.MidStream{
		{
			ID:    1,
			Point: pMid1,
			Cap:   decimal.NewFromInt(100),
		},
		{
			ID:    2,
			Point: pMid2,
			Cap:   decimal.NewFromInt(200),
		},
	}

	nodes := []*model.Node{
		{ID: 1, Point: pStreet1},
		{ID: 2, Point: pStreet2},
		{ID: 3, Point: pStreet3},
		{ID: 4, Point: pMid1},
		{ID: 5, Point: pMid2},
	}

	edges := []*model.Edge{
		{ID: 1, From: 1, To: 4, Dis: 400},
		{ID: 2, From: 4, To: 2, Dis: 500},
		{ID: 3, From: 2, To: 5, Dis: 600},
		{ID: 4, From: 5, To: 3, Dis: 700},
	}

	algo := NewAlgo(&Models{
		Edges:      edges,
		Nodes:      nodes,
		Streets:    streets,
		UpStreams:  nil,
		MidStreams: midStreams,
	})

	ans := algo.RunMidToStreet()
	fmt.Printf("ans: %+v\n", ans)
}

func TestAlgo2(t *testing.T) {
	pStreet1 := model.Point{X: 2, Y: 1}
	pStreet2 := model.Point{X: 2, Y: 2}
	pStreet3 := model.Point{X: 2, Y: 3}

	pMid1 := model.Point{X: 1, Y: 1}
	pMid2 := model.Point{X: 1, Y: 2}

	streets := []*model.Street{
		{
			ID:          1,
			Point:       pStreet1,
			PeopleCount: 4,
			Cap:         decimal.NewFromInt(4),
		},
		{
			ID:          2,
			Point:       pStreet2,
			PeopleCount: 8,
			Cap:         decimal.NewFromInt(8),
		},
		{
			ID:          3,
			Point:       pStreet3,
			PeopleCount: 10,
			Cap:         decimal.NewFromInt(10),
		},
	}

	midStreams := []*model.MidStream{
		{
			ID:    1,
			Point: pMid1,
			Cap:   decimal.NewFromInt(10),
		},
		{
			ID:    2,
			Point: pMid2,
			Cap:   decimal.NewFromInt(20),
		},
	}

	nodes := []*model.Node{
		{ID: 1, Point: model.Point{X: pStreet1.X + 0.00001, Y: pStreet1.Y}},
		{ID: 2, Point: model.Point{X: pStreet2.X + 0.00001, Y: pStreet2.Y}},
		{ID: 3, Point: model.Point{X: pStreet3.X + 0.00001, Y: pStreet3.Y}},
		{ID: 4, Point: model.Point{X: pMid1.X + 0.00001, Y: pMid1.Y}},
		{ID: 5, Point: model.Point{X: pMid2.X + 0.00001, Y: pMid2.Y}},

		{ID: 6, Point: model.Point{X: pStreet1.X + 0.00002, Y: pStreet1.Y}},
		{ID: 7, Point: model.Point{X: pStreet2.X + 0.00002, Y: pStreet2.Y}},
		{ID: 8, Point: model.Point{X: pStreet3.X + 0.00002, Y: pStreet3.Y}},
	}

	edges := []*model.Edge{
		{ID: 1, From: 1, To: 4, Dis: 4},
		{ID: 2, From: 4, To: 2, Dis: 5},
		{ID: 3, From: 2, To: 5, Dis: 6},
		{ID: 4, From: 5, To: 3, Dis: 7},

		{ID: 5, From: 6, To: 4, Dis: 40},
		{ID: 2, From: 4, To: 7, Dis: 50},
		{ID: 3, From: 7, To: 5, Dis: 60},
		{ID: 4, From: 5, To: 8, Dis: 70},
	}

	algo := NewAlgo(&Models{
		Edges:      edges,
		Nodes:      nodes,
		Streets:    streets,
		UpStreams:  nil,
		MidStreams: midStreams,
	})

	ans := algo.RunMidToStreet()
	fmt.Printf("ans: %+v\n", ans)
}

func TestAlgo3(t *testing.T) {
	pStreet1 := model.Point{X: 2, Y: 1}
	pStreet2 := model.Point{X: 2, Y: 2}
	pStreet3 := model.Point{X: 2, Y: 3}
	pStreet4 := model.Point{X: 2, Y: 6}

	pMid1 := model.Point{X: 1, Y: 1}
	pMid2 := model.Point{X: 1, Y: 2}

	streets := []*model.Street{
		{
			ID:          1,
			Point:       pStreet1,
			PeopleCount: 4,
			Cap:         decimal.NewFromInt(4),
		},
		{
			ID:          2,
			Point:       pStreet2,
			PeopleCount: 8,
			Cap:         decimal.NewFromInt(8),
		},
		{
			ID:          3,
			Point:       pStreet3,
			PeopleCount: 10,
			Cap:         decimal.NewFromInt(10),
		},
		{
			ID:          4,
			Point:       pStreet4,
			PeopleCount: 8,
			Cap:         decimal.NewFromInt(8),
		},
	}

	midStreams := []*model.MidStream{
		{
			ID:    1,
			Point: pMid1,
			Cap:   decimal.NewFromInt(10),
		},
		{
			ID:    2,
			Point: pMid2,
			Cap:   decimal.NewFromInt(20),
		},
	}

	nodes := []*model.Node{
		{ID: 1, Point: pStreet1},
		{ID: 2, Point: pStreet2},
		{ID: 3, Point: pStreet3},
		{ID: 4, Point: pMid1},
		{ID: 5, Point: pMid2},
	}

	edges := []*model.Edge{
		{ID: 1, From: 1, To: 4, Dis: 4},
		{ID: 2, From: 4, To: 2, Dis: 5},
		{ID: 3, From: 2, To: 3, Dis: 6},
		{ID: 4, From: 5, To: 3, Dis: 7},
	}

	algo := NewAlgo(&Models{
		Edges:      edges,
		Nodes:      nodes,
		Streets:    streets,
		UpStreams:  nil,
		MidStreams: midStreams,
	})

	ans := algo.RunMidToStreet()
	fmt.Printf("ans: %+v\n", ans)
}

func TestRunMidToStreetCars(t *testing.T) {
	pStreet1 := model.Point{X: 2, Y: 1}
	pStreet2 := model.Point{X: 2, Y: 2}
	pStreet3 := model.Point{X: 2, Y: 3}

	pMid1 := model.Point{X: 1, Y: 1}
	pMid2 := model.Point{X: 1, Y: 2}

	streets := []*model.Street{
		{
			ID:          1,
			Point:       pStreet1,
			PeopleCount: 4,
			Cap:         decimal.NewFromInt(4),
			OriCap:      decimal.NewFromInt(4),
		},
		{
			ID:          2,
			Point:       pStreet2,
			PeopleCount: 8,
			Cap:         decimal.NewFromInt(8),
			OriCap:      decimal.NewFromInt(8),
		},
		{
			ID:          3,
			Point:       pStreet3,
			PeopleCount: 10,
			Cap:         decimal.NewFromInt(10),
			OriCap:      decimal.NewFromInt(10),
		},
	}

	midStreams := []*model.MidStream{
		{
			ID:    1,
			Point: pMid1,
			Cap:   decimal.NewFromInt(10),
		},
		{
			ID:    2,
			Point: pMid2,
			Cap:   decimal.NewFromInt(20),
		},
	}

	nodes := []*model.Node{
		{ID: 1, Point: pStreet1},
		{ID: 2, Point: pStreet2},
		{ID: 3, Point: pStreet3},
		{ID: 4, Point: pMid1},
		{ID: 5, Point: pMid2},
	}

	edges := []*model.Edge{
		{ID: 1, From: 1, To: 4, Dis: 300},
		{ID: 2, From: 4, To: 2, Dis: 500},
		{ID: 3, From: 2, To: 5, Dis: 600},
		{ID: 4, From: 5, To: 3, Dis: 700},
	}

	algo := NewAlgo(&Models{
		Edges:      edges,
		Nodes:      nodes,
		Streets:    streets,
		UpStreams:  nil,
		MidStreams: midStreams,
	})

	ans := algo.RunMidToStreetCars()
	fmt.Printf("ans: %+v\n", ans.Cars)
}
