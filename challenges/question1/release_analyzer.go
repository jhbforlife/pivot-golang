package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
)

type ReleaseAnalyzer interface {
	GetReleaseStats() ([]ReleaseStats, error)
	GetReleaseQuality() (ReleaseQuality, error)
	GetReleaseHistory() ([]string, error)
}

type ReleaseStats struct {
	releaseID    string
	minQueryTime float64
	avgQueryTime float64
	maxQueryTime float64
}

type ReleaseQuality struct {
	bestReleaseID  string
	worstReleaseID string
}

type Analyzer struct{}

func NewAnalyzer() *Analyzer {
	a := Analyzer{}
	return &a
}

func (a *Analyzer) GetReleaseStats() ([]ReleaseStats, error) {
	resp := getResp()
	fmt.Println(resp[0])
	respByVersion := sortRespVersion(resp)
	fmt.Println(respByVersion["77de68daecd823babbb58edb1c8e14d7106e83bb"])
	queryStatsByVersion := getStats(respByVersion)
	fmt.Println(queryStatsByVersion["77de68daecd823babbb58edb1c8e14d7106e83bb"])
	bestWorst := getBestWorst(queryStatsByVersion)
	fmt.Println(bestWorst)
	return []ReleaseStats{}, errors.New("not implemented")
}

func (a *Analyzer) GetReleaseQuality() (ReleaseQuality, error) {
	return ReleaseQuality{}, errors.New("not implemented")
}

func (a *Analyzer) GetReleaseHistory() ([]string, error) {
	return nil, errors.New("not implemented")
}

type event struct {
	Timestamp int    `json:"timestamp"`
	Version   string `json:"version"`
	QueryTime int    `json:"query_time"`
}

type eventByVersion map[string][]event

type queryStats struct {
	minimum int
	average float64
	maximum int
}

type queryStatsByVersion map[string]queryStats

func getResp() []event {
	es := []event{}
	bs, readErr := os.ReadFile("events.json")
	checkErr(readErr)
	unmarshalErr := json.Unmarshal(bs, &es)
	checkErr(unmarshalErr)
	return es
}

func sortRespVersion(resp []event) eventByVersion {
	vm := eventByVersion{}
	for _, v := range resp {
		vm[v.Version] = append(vm[v.Version], v)
	}
	return vm
}

func getStats(vs eventByVersion) queryStatsByVersion {
	qm := queryStatsByVersion{}
	for k, v := range vs {
		var min int
		var avg float64
		var max int
		for i, e := range v {
			if e.QueryTime < min || i == 0 {
				min = e.QueryTime
			} else if e.QueryTime > max {
				max = e.QueryTime
			}
			avg += float64(e.QueryTime)
		}
		avg = math.Round((avg / float64(len(v)) * 100)) / 100
		qm[k] = queryStats{minimum: min, average: avg, maximum: max}
	}
	return qm
}

func getBestWorst(qm queryStatsByVersion) []string {
	ss := []string{}
	// for k, v := range qm {

	// }
	return ss
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
