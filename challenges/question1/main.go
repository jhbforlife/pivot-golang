package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
)

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

func main() {
	resp := getResp()
	fmt.Println(resp[0])
	respByVersion := sortRespVersion(resp)
	fmt.Println(respByVersion["77de68daecd823babbb58edb1c8e14d7106e83bb"])
	queryStatsByVersion := getStats(respByVersion)
	fmt.Println(queryStatsByVersion["77de68daecd823babbb58edb1c8e14d7106e83bb"])
}

func getResp() []event {
	rs := []event{}
	bs, readErr := os.ReadFile("events.json")
	checkErr(readErr)
	unmarshalErr := json.Unmarshal(bs, &rs)
	checkErr(unmarshalErr)
	return rs
}

func sortRespVersion(resp []event) eventByVersion {
	vs := eventByVersion{}
	for _, v := range resp {
		vs[v.Version] = append(vs[v.Version], v)
	}
	return vs
}

func getStats(vs eventByVersion) queryStatsByVersion {
	qs := queryStatsByVersion{}
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
		qs[k] = queryStats{minimum: min, average: avg, maximum: max}
	}
	return qs
}

// func getBestWorst() {

// }

// func getReleaseHistory() {

// }

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
