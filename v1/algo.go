package v1

import (
	"main/config"
	"main/intf"
	"main/model"
	"main/v1/pq"
	"math"
	"sort"

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

	logrus.Tracef("a.IndexToEdge: %v", a.IndexToEdge)
	logrus.Tracef("a.IndexToNode: %v", a.IndexToNode)
	logrus.Tracef("a.IndexToStreet: %v", a.IndexToStreet)
	logrus.Tracef("a.IndexToUpStreams: %v", a.IndexToUpStreams)
	logrus.Tracef("a.IndexToMidStreams: %v", a.IndexToMidStreams)
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
	ret := make([]int, 0)
	ns := make([]*model.Node, 0)

	for _, node := range nodes {
		if !n.isNear(pointer, node) {
			continue
		}

		ns = append(ns, node)
	}

	sort.Slice(ns, func(i, j int) bool {
		return pointer.GetPoint().Distance(ns[i].GetPoint()) < pointer.GetPoint().Distance(ns[j].GetPoint())
	})

	for i := 0; i < len(ns) && i < config.GetConfig().NearCount; i++ {
		ret = append(ret, ns[i].ID)
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
	GDij         G
}

func NewAlgo(m *Models) *Algo {
	a := &Algo{
		Models:      *m,
		NearNodeIDs: NearNodeIDs{},
	}
	a.Init()
	return a
}

func (a *Algo) Dijkstra(start int, maxNode int) map[int]float64 {
	g := a.GEdges
	// init
	dist := make(map[int]float64)
	for k := 1; k <= maxNode; k++ {
		dist[k] = MaxDis
	}
	dist[start] = 0

	for j := 1; j <= maxNode; j++ {
		if _, ok := g[start][j]; ok {
			dist[j] = g[start][j]
		}
	}

	// dijkstra
	visited := make(map[int]bool)
	visited[start] = true
	for i := 1; i <= maxNode; i++ {
		min := MaxDis
		minIndex := -1
		for j := 1; j <= maxNode; j++ {
			if _, ok := visited[j]; !ok && dist[j] < min {
				min = dist[j]
				minIndex = j
			}
		}

		if minIndex == -1 {
			break
		}

		logrus.Tracef("relax: minIndex: %v", minIndex)
		visited[minIndex] = true
		for j := 1; j <= maxNode; j++ {
			_, ok1 := g[minIndex][j]
			_, ok2 := dist[j]
			if ok1 && ok2 && dist[j] > dist[minIndex]+g[minIndex][j] {
				dist[j] = dist[minIndex] + g[minIndex][j]
				logrus.Tracef("relax: j: %v, new: %f", j, dist[j])

			}
		}
	}

	return dist
}

func (a *Algo) gGetMinDis(tos []int, dij map[int]float64) float64 {
	minDis := MaxDis

	for _, to := range tos {
		if dis, ok := dij[to]; ok && dis < minDis {
			minDis = dis
		}
	}

	return minDis
}

// initGMidToStreet 初始化中游到街道的最短距离
func (a *Algo) initGMidToStreet() {
	g := make(map[int]map[int]float64)

	for _, mid := range a.MidStreams {
		midNears := a.MidStreamNearNodeIDs[mid.ID]
		if len(midNears) == 0 {
			continue
		}

		if _, ok := g[mid.ID]; !ok {
			g[mid.ID] = make(map[int]float64)
		}

		for _, near := range midNears {
			dij, ok := a.GDij[near]
			if !ok {
				continue
			}

			for _, street := range a.Streets {
				streetNears := a.StreetNearNodeIDs[street.ID]
				if len(streetNears) == 0 {
					continue
				}

				if dis := a.gGetMinDis(streetNears, dij); dis < MaxDis {
					g[mid.ID][street.ID] = dis
				}
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

		if _, ok := a.GEdges[edge.From][edge.To]; !ok {
			a.GEdges[edge.From][edge.To] = edge.Dis
			a.GEdges[edge.To][edge.From] = edge.Dis
		} else {
			a.GEdges[edge.From][edge.To] = math.Min(a.GEdges[edge.From][edge.To], edge.Dis)
			a.GEdges[edge.To][edge.From] = math.Min(a.GEdges[edge.To][edge.From], edge.Dis)
		}
	}

	logrus.Tracef("a.GEdges: %v", a.GEdges)
}

func (a *Algo) initGMidToStreetByMaxMul() {
	g := a.GMidToStreet

	for _, street := range a.Streets {
		for _, mid := range a.MidStreams {
			midToStreetMaxDis := mid.Point.Distance(&street.Point) * config.GetConfig().MaxDisMul

			if _, ok := g[mid.ID]; !ok {
				g[mid.ID] = make(map[int]float64)
			}

			if _, ok := g[mid.ID][street.ID]; !ok {
				g[mid.ID][street.ID] = midToStreetMaxDis
			} else {
				g[mid.ID][street.ID] = math.Min(g[mid.ID][street.ID], midToStreetMaxDis)
			}
		}
	}
}

func (a *Algo) initGMidDij() {
	g := make(map[int]map[int]float64)
	for _, mid := range a.MidStreams {
		midNears := a.MidStreamNearNodeIDs[mid.ID]
		if len(midNears) == 0 {
			continue
		}

		for _, near := range midNears {
			dij := a.Dijkstra(near, len(a.Nodes))
			g[near] = dij
		}
	}
	a.GDij = g
}

func (a *Algo) Init() {
	a.initNearNodeIDS()
	a.initIndexs()
	a.initGEdge()
	a.initGMidDij()
	a.initGMidToStreet()
	a.initGMidToStreetByMaxMul()
}

func (a *Algo) RunUpToMid() *Ans {
	ans := &Ans{Path: make([]*PathEdge, 0)}

	for _, mid := range a.MidStreams {
		midNear := a.MidStreamNearNodeIDs[mid.ID]
		path := make(Path, 0)

		for _, near := range midNear {
			dij := a.GDij[near]

			for _, up := range a.UpStreams {
				upNear := a.UpStreamNearNodeIDs[up.ID]

				for _, near2 := range upNear {
					if dis, ok := dij[near2]; ok {
						path = append(path, &PathEdge{
							From: buildUpStreamID(near2),
							To:   buildMidStreamID(near),
							Val:  mid.OriCap.Sub(mid.Cap),
							Dis:  dis})
					}
				}
			}
		}

		for _, up := range a.UpStreams {
			disMul := mid.Point.Distance(&up.Point) * config.GetConfig().MaxDisMul
			path = append(path, &PathEdge{
				From: buildUpStreamID(up.ID),
				To:   buildMidStreamID(mid.ID),
				Val:  mid.OriCap.Sub(mid.Cap),
				Dis:  disMul})
		}

		sort.Slice(path, func(i, j int) bool {
			return path[i].Dis < path[j].Dis
		})
		ans.Path = append(ans.Path, path[0])
	}

	return ans
}

func (a *Algo) RunMidToStreet() *Ans {
	ans := &Ans{Path: make([]*PathEdge, 0, 1)}
	queue := pq.NewPriorityQueue()

	// 初始化队列
	for _, mid := range a.MidStreams {
		for _, street := range a.Streets {
			if _, ok := a.GMidToStreet[mid.ID][street.ID]; !ok {
				continue
			}

			queue.Push(&pq.Item{From: mid.ID, To: street.ID, Priority: a.GMidToStreet[mid.ID][street.ID]})
		}
	}
	logrus.Debugf("queue: %+v", queue)

	for queue.Len() > 0 {
		item := queue.Pop()
		logrus.Debugf("range: get item: %v", item)

		street := a.IndexToStreet[item.To]
		logrus.Tracef("range: get street: %v", street)
		if street.Cap.IsZero() {
			continue
		}

		mid := a.IndexToMidStreams[item.From]
		logrus.Tracef("range: get mid: %v", mid)
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

		ans.AddEdge(item.From, item.To, val, item.Priority)
	}

	return ans
}

// RunMidToStreetCars 类 dijkstra 算法求最小生成树的变种
func (a *Algo) RunMidToStreetCars() *Ans {
	dis := make(map[int]float64, len(a.Streets))
	pre := make(map[int]int, len(a.Streets))
	next := make(map[int]int, len(a.Streets))
	count := make(map[int]int, len(a.Streets)) // street count
	sum := make(map[int]decimal.Decimal, len(a.Streets))
	visited := make(map[int]bool, len(a.Streets))
	start := make(map[int]int, len(a.Streets))

	// init
	for _, street := range a.Streets {
		pre[street.ID] = -1
		next[street.ID] = street.ID
		count[street.ID] = 1
		sum[street.ID] = street.OriCap
		dis[street.ID] = MaxDis

		for _, mid := range a.MidStreams {
			if _, ok := a.GMidToStreet[mid.ID][street.ID]; ok {
				if dis[street.ID] > a.GMidToStreet[mid.ID][street.ID] {
					dis[street.ID] = a.GMidToStreet[mid.ID][street.ID]
					start[street.ID] = mid.ID
				}
			}
		}
	}

	logrus.Tracef("ori dis: %v", dis)
	logrus.Tracef("ori sum: %v", sum)

	for i := 0; i < len(a.Streets); i++ {
		minDis := math.MaxFloat64
		minID := -1
		for _, street := range a.Streets {
			if visited[street.ID] {
				continue
			}

			if dis[street.ID] < minDis {
				minDis = dis[street.ID]
				minID = street.ID
			}
		}

		if minID == -1 {
			break
		}

		logrus.Tracef("minID: %v", minID)

		visited[minID] = true
		for _, street := range a.Streets {
			if visited[street.ID] {
				continue
			}

			if dis[street.ID] > dis[minID]+a.IndexToStreet[minID].Point.Distance(&street.Point) &&
				sum[street.ID].Add(sum[minID]).LessThanOrEqual(config.GetConfig().MaxCarCapDecimal) {

				dis[street.ID] = dis[minID] + a.IndexToStreet[minID].Point.Distance(&street.Point)
				pre[street.ID] = minID
				next[minID] = street.ID
				count[street.ID] = count[minID] + 1
				sum[street.ID] = sum[minID].Add(sum[street.ID])

				if count[street.ID] >= config.GetConfig().MaxStreetPerCar {
					visited[street.ID] = true
				}
				if sum[street.ID].GreaterThanOrEqual(config.GetConfig().MaxCarCapDecimal) {
					visited[street.ID] = true
				}
				break
			}
		}
	}

	logrus.Debugf("dis: %+v", dis)
	logrus.Debugf("pre: %+v", pre)
	logrus.Debugf("next: %+v", next)
	logrus.Debugf("count: %+v", count)
	logrus.Debugf("sum: %+v", sum)
	logrus.Debugf("start: %+v", start)

	ans := &Ans{Cars: make(Cars, 0, 1)}
	for _, street := range a.Streets {
		// count 为 1 的是起始点
		if count[street.ID] != 1 {
			continue
		}

		crossStreets := a.GetCarPath(street.ID, next)
		end := crossStreets[len(crossStreets)-1]
		ans.AddCarInfo(start[street.ID], crossStreets, sum[end], dis[end])
	}

	return ans
}

func (a *Algo) GetCarPath(start int, next map[int]int) []int {
	path := make([]int, 0, 1)
	path = append(path, start)

	for next[start] != start {
		start = next[start]
		path = append(path, start)
	}

	return path
}
