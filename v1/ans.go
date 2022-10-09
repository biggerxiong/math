package v1

import (
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type PathEdge struct {
	From string
	To   string
	Val  decimal.Decimal
}

func (p *PathEdge) String() string {
	return p.From + "->" + p.To + ":" + p.Val.String()
}

type Path []*PathEdge

func (p Path) AddEdge(from, to string, val decimal.Decimal) Path {
	return append(p, &PathEdge{From: from, To: to, Val: val})
}

type Ans struct {
	Path Path
}

func (a *Ans) String() string {
	ss := make([]string, 0, len(a.Path))
	for _, edge := range a.Path {
		ss = append(ss, edge.String())
	}
	return "Ans{" + "Path: " + "[\n" + strings.Join(ss, ",\n") + "\n]" + "}"
}

func buildStreetID(id int) string {
	return "street" + strconv.Itoa(id)
}

func buildMidStreamID(id int) string {
	return "mid" + strconv.Itoa(id)
}

func (a *Ans) AddEdge(from, to int, val decimal.Decimal) {
	a.Path = a.Path.AddEdge(buildMidStreamID(from), buildStreetID(to), val)
	logrus.Debugf("ans: add edge: %v", a.Path[len(a.Path)-1])
}
