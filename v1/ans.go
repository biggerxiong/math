package v1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type PathEdge struct {
	From string
	To   string
	Val  decimal.Decimal
	Dis  float64
}

func (p *PathEdge) String() string {
	return fmt.Sprintf("%s -> %s: cap:%s, dis:%f", p.From, p.To, p.Val.String(), p.Dis)
}

func (p *PathEdge) ToStrArr() []string {
	return []string{p.From, p.To, p.Val.String(), strconv.FormatFloat(p.Dis, 'f', 4, 64)}
}

type Path []*PathEdge

func (p Path) AddEdge(from, to string, val decimal.Decimal, dis float64) Path {
	return append(p, &PathEdge{From: from, To: to, Val: val, Dis: dis})
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
	return "street-" + strconv.Itoa(id)
}

func buildMidStreamID(id int) string {
	return "mid-" + strconv.Itoa(id)
}

func buildUpStreamID(id int) string {
	return "up-" + strconv.Itoa(id)
}

func (a *Ans) AddEdge(from, to int, val decimal.Decimal, dis float64) {
	a.Path = a.Path.AddEdge(buildMidStreamID(from), buildStreetID(to), val, dis)
	logrus.Debugf("ans: add edge: %v", a.Path[len(a.Path)-1])
}
