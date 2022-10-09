package v1

import (
	"main/config"
	"main/intf"
	"main/model"
	"main/v1/pq"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

const MaxDis = 10000000.0

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

func (a *Algo) initIndexs() {
	a.IndexToEdge = make(map[int]*model.Edge)
	a.IndexToNode = make(map[int]*model.Node)
	a.IndexToStreet = make(map[int]*model.Street)
	a.IndexToUpStreams = make(map[int]*model.UpStream)
	a.IndexToMidStreams = make(map[int]*model.MidStream)

	for _, edge := range a.Edges {
		a.IndexToEdge[edge.ID] = edge
	}

	for _, node := range a.Nodes {
		a.IndexToNode[node.ID] = node
	}

	for _, street := range a.Streets {
		a.IndexToStreet[street.ID] = street
	}

	for _, upStream := range a.UpStreams {
		a.IndexToUpStreams[upStream.ID] = upStream
	}

	for _, midStream := range a.MidStreams {
		a.IndexToMidStreams[midStream.ID] = midStream
	}

	logrus.Debugf("a.IndexToEdge: %v", a.IndexToEdge)
	logrus.Debugf("a.IndexToNode: %v", a.IndexToNode)
	logrus.Debugf("a.IndexToStreet: %v", a.IndexToStreet)
	logrus.Debugf("a.IndexToUpStreams: %v", a.IndexToUpStreams)
	logrus.Debugf("a.IndexToMidStreams: %v", a.IndexToMidStreams)
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

type G map[int]map[int]float64 // A -> B 's min distance

type Algo struct {
	Index
	Models
	NearNodeIDs

	GMidToStreet G
	GEdges       G
}

func NewAlgo(m *Models) *Algo {
	a := &Algo{
		Models:      *m,
		NearNodeIDs: NearNodeIDs{},
	}
	a.Init()
	return a
}

func (a *Algo) gGetMinDis(froms []int, tos []int) float64 {
	minDis := MaxDis

	for _, from := range froms {
		for _, to := range tos {
			if dis, ok := a.GEdges[from][to]; ok && dis < minDis {
				minDis = dis
			}
		}
	}

	return minDis
}

// initGMidToStreet 初始化街道到中游的最短距离
func (a *Algo) initGMidToStreet() {
	g := make(map[int]map[int]float64)

	for _, street := range a.Streets {
		streetNears := a.StreetNearNodeIDs[street.ID]
		if len(streetNears) == 0 {
			continue
		}

		if _, ok := g[street.ID]; !ok {
			g[street.ID] = make(map[int]float64)
		}

		for _, mid := range a.MidStreams {
			midNears := a.MidStreamNearNodeIDs[mid.ID]
			if len(midNears) == 0 {
				continue
			}

			if dis := a.gGetMinDis(streetNears, midNears); dis < MaxDis {
				g[street.ID][mid.ID] = dis
			}
		}
	}

	a.GMidToStreet = g

	logrus.Debugf("a.GMidToStreet: %v", a.GMidToStreet)
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
			a.UpStreamNearNodeIDs[upStream.ID] = ids
		}
	}

	for _, midStream := range a.MidStreams {
		ids := a.GetNearNodeIds(midStream, a.Nodes)
		if len(ids) > 0 {
			a.MidStreamNearNodeIDs[midStream.ID] = ids
		}
	}

	logrus.Debugf("a.StreetNearNodeIDs: %v", a.StreetNearNodeIDs)
	logrus.Debugf("a.UpStreamNearNodeIDs: %v", a.UpStreamNearNodeIDs)
	logrus.Debugf("a.MidStreamNearNodeIDs: %v", a.MidStreamNearNodeIDs)
}

func (a *Algo) initGEdge() {
	a.GEdges = make(map[int]map[int]float64)
	for _, edge := range a.Edges {
		if _, ok := a.GEdges[edge.From]; !ok {
			a.GEdges[edge.From] = make(map[int]float64)
		}
		if _, ok := a.GEdges[edge.To]; !ok {
			a.GEdges[edge.To] = make(map[int]float64)
		}

		a.GEdges[edge.From][edge.To] = edge.Dis
		a.GEdges[edge.To][edge.From] = edge.Dis
	}

	logrus.Debugf("a.GEdges: %v", a.GEdges)
}

func (a *Algo) Init() {
	a.initNearNodeIDS()
	a.initIndexs()
	a.initGEdge()
	a.initGMidToStreet()
}

func (a *Algo) Run() *Ans {
	ans := &Ans{Path: make([]*PathEdge, 0, 1)}
	queue := pq.NewPriorityQueue()

	// 初始化队列
	for _, mid := range a.MidStreams {
		for _, street := range a.Streets {
			if _, ok := a.GMidToStreet[street.ID][mid.ID]; !ok {
				continue
			}

			queue.Push(&pq.Item{From: mid.ID, To: street.ID, Priority: a.GMidToStreet[street.ID][mid.ID]})
		}
	}
	logrus.Debugf("queue: %+v", queue)

	for queue.Len() > 0 {
		item := queue.Pop()
		logrus.Debugf("range: get item: %v", item)

		street := a.IndexToStreet[item.To]
		logrus.Debugf("range: get street: %v", street)
		if street.Cap.IsZero() {
			continue
		}

		mid := a.IndexToMidStreams[item.From]
		logrus.Debugf("range: get mid: %v", mid)
		if mid.Cap.IsZero() {
			continue
		}

		var val decimal.Decimal
		if mid.Cap.LessThanOrEqual(street.Cap) {
			val = mid.Cap
			street.Cap = street.Cap.Sub(mid.Cap)
			mid.Cap = decimal.Zero
		} else {
			val = street.Cap
			mid.Cap = mid.Cap.Sub(street.Cap)
			street.Cap = decimal.Zero
		}

		ans.AddEdge(item.From, item.To, val)
	}

	return ans
}
