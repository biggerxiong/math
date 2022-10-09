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

type CarInfo struct {
	ID int

	StartMidID   int
	CrossStreets []int

	CapSum decimal.Decimal
	Mile   float64
}

func (c *CarInfo) String() string {
	return fmt.Sprintf("CarInfo{ID:%d, StartMidID:%d, CrossStreets:%v, CapSum:%s, Mile:%f}",
		c.ID, c.StartMidID, c.CrossStreets, c.CapSum.String(), c.Mile)
}

func (c *CarInfo) ToStrArr() []string {
	return []string{
		strconv.Itoa(c.ID),
		buildMidStreamID(c.StartMidID),
		strings.Join(c.ToPathStrings(), ","),
		c.CapSum.String(),
		strconv.FormatFloat(c.Mile, 'f', 4, 64),
	}
}

func (c *CarInfo) ToPathStrings() []string {
	ss := make([]string, 0, len(c.CrossStreets)+1)

	for _, streetID := range c.CrossStreets {
		ss = append(ss, buildStreetID(streetID))
	}
	return ss
}

type Cars []*CarInfo

type Ans struct {
	Path Path
	Cars Cars
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

func (a *Ans) AddCarInfo(start int, crossStreets []int, capSum decimal.Decimal, mile float64) {
	a.Cars = append(a.Cars, &CarInfo{
		ID:           len(a.Cars) + 1,
		StartMidID:   start,
		CrossStreets: crossStreets,
		CapSum:       capSum,
		Mile:         mile,
	})
	logrus.Debugf("ans: add car: %v", a.Cars[len(a.Cars)-1])
}
