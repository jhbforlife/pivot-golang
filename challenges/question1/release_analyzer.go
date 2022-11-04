// Given Code
package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"sort"
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
	// Given code
	// return nil, errors.New("not implemented")
	// Added code
	events := getEvents()
	eventsByRelease := sortEventsByRelease(events)
	return getStats(eventsByRelease), nil
}

func (a *Analyzer) GetReleaseQuality() (ReleaseQuality, error) {
	// Given code
	// return nil, errors.New("not implemented")
	// Added code
	eventsByRelease, err := a.GetReleaseStats()
	checkErr(err)
	return getBestWorst(eventsByRelease), nil
}

func (a *Analyzer) GetReleaseHistory() ([]string, error) {
	// Given code
	// return nil, errors.New("not implemented")
	// Added code
	eventsByRelease, err := a.GetReleaseStats()
	checkErr(err)
	return getHistory(eventsByRelease), nil
}

// Added code
type event struct {
	Timestamp float64 `json:"timestamp"`
	Version   string  `json:"version"`
	QueryTime float64 `json:"query_time"`
}

type eventsByRelease map[string][]event

func getEvents() []event {
	es := []event{}
	bs, readErr := os.ReadFile("events.json")
	checkErr(readErr)
	unmarshalErr := json.Unmarshal(bs, &es)
	checkErr(unmarshalErr)
	return es
}

func sortEventsByRelease(events []event) eventsByRelease {
	vm := eventsByRelease{}
	for _, v := range events {
		vm[v.Version] = append(vm[v.Version], v)
	}
	return vm
}

func getStats(rs eventsByRelease) []ReleaseStats {
	qs := []ReleaseStats{}
	for k, v := range rs {
		var min float64
		var avg float64
		var max float64
		for i, e := range v {
			if e.QueryTime < min || i == 0 {
				min = e.QueryTime
			} else if e.QueryTime > max {
				max = e.QueryTime
			}
			avg += float64(e.QueryTime)
		}
		avg = math.Round((avg / float64(len(v)) * 100)) / 100
		qs = append(qs, ReleaseStats{k, min, avg, max})
	}
	return qs
}

func getBestWorst(qs []ReleaseStats) ReleaseQuality {
	var best float64
	var worst float64
	var rq ReleaseQuality
	for i, e := range qs {
		avg := e.avgQueryTime
		if i == 0 {
			best = avg
			worst = avg
			rq.worstReleaseID = e.releaseID
			rq.bestReleaseID = e.releaseID
		} else {
			if avg < best {
				best = avg
				rq.bestReleaseID = e.releaseID
			} else if avg > worst {
				worst = avg
				rq.worstReleaseID = e.releaseID
			}
		}
	}
	return rq
}

func getHistory(events []ReleaseStats) []string {
	hs := []string{}
	sort.Slice(events, func(i, j int) bool {
		return events[i].releaseID < events[j].releaseID
	})
	for _, v := range events {
		hs = append(hs, v.releaseID)
	}
	fmt.Println(events)
	return hs
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
