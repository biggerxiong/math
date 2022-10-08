package v1

import (
	"main/config"
	"main/intf"
	"main/model"
)

type Models struct {
	Edges   []*model.Edge
	Nodes   []*model.Node
	Streets []*model.Street

	UpStreams  []*model.UpStream
	MidStreams []*model.MidStream
}

type Index struct {
	IndexToEdge       map[int]*model.Edge
	IndexToNode       map[int]*model.Node
	IndexToStreet     map[int]*model.Street
	IndexToUpStreams  map[int]*model.UpStream
	IndexToMidStreams map[int]*model.MidStream
}

func (i Index) InitFromModels(m Models) {

}

type NearNodeIDs struct {
	StreetNearNodeIDs    map[int][]int
	UpStreamNearNodeIDs  map[int][]int
	MidStreamNearNodeIDs map[int][]int
}

func (n NearNodeIDs) isNear(pointer intf.Pointer, node *model.Node) bool {
	return pointer.GetPoint().Distance(node.GetPoint()) < config.GetConfig().MinDis
}

func (n NearNodeIDs) GetNearNodeIds(pointer intf.Pointer, nodes []*model.Node) []int {
	var ret []int
	for _, node := range nodes {
		if !n.isNear(pointer, node) {
			continue
		}

		if len(ret) == 0 {
			ret = make([]int, 0, 1)
		}
		ret = append(ret, node.ID)
	}

	return ret
}

type G struct {
	MinDistance map[int]map[int]string // A -> B 's min distance
}

type Algo struct {
	Models
	NearNodeIDs
}

func NewAlgo(m *Models) *Algo {
	a := &Algo{
		Models:      *m,
		NearNodeIDs: NearNodeIDs{},
	}
	a.Init()
	return a
}

func (a *Algo) initNearNodeIDS() {
	a.StreetNearNodeIDs = make(map[int][]int)
	a.UpStreamNearNodeIDs = make(map[int][]int)
	a.MidStreamNearNodeIDs = make(map[int][]int)

	for _, street := range a.Streets {
		ids := a.GetNearNodeIds(street, a.Nodes)
		if len(ids) > 0 {
			a.StreetNearNodeIDs[street.ID] = ids
		}
	}

	for _, upStream := range a.UpStreams {
		ids := a.GetNearNodeIds(upStream, a.Nodes)
		if len(ids) > 0 {
			a.StreetNearNodeIDs[upStream.ID] = ids
		}
	}

	for _, midStream := range a.MidStreams {
		ids := a.GetNearNodeIds(midStream, a.Nodes)
		if len(ids) > 0 {
			a.StreetNearNodeIDs[midStream.ID] = ids
		}
	}
}

func (a *Algo) Init() {
	a.initNearNodeIDS()
}
