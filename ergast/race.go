package ergast

import (
    "encoding/json"
    "github.com/mihaipstef/ergast-go/helpers/request"
    "github.com/mihaipstef/ergast-go/helpers/common"
)

type Time struct {
    Millis string `json:"millis,omitempty"`
    Time   string `json:"time"`
}

type AverageSpeed struct {
    Units string `json:"units"`
    Speed string `json:"speed"`
}

type FastestLap struct {
    Rank         int            `json:"rank,string"`
    Lap          int            `json:"lap,string"`
    Time         Time           `json:"Time"`
    AverageSpeed AverageSpeed   `json:"AverageSpeed"`
}

type Result struct {
    common.Result
    Points       int            `json:"points,string"`
    Grid         int            `json:"grid,string"`
    Laps         int            `json:"laps,string"`
    Status       string         `json:"status"`
    Time         Time           `json:"Time,omitempty"`
    FastestLap   FastestLap     `json:"FastestLap"`
}

type RaceResult struct {
    common.Race
    Results  []Result `json:"Results"`
}

type RaceResultsResponse struct {
    MRData struct {
        common.MRData
        RaceTable struct {
            Season string           `json:"season"`
            Round  int              `json:"round,string"`
            Races  []RaceResult     `json:"Races"`
        } `json:"RaceTable"`
    } `json:"MRData"`
}

type RaceResultsRequest struct {
    request.Request
}

func NewRaceResultsRequest(year string, round int,content_type request.ResponseContentType) (*RaceResultsRequest) {
    req := RaceResultsRequest{Request: request.Request{ContentType: content_type}}
    req.AddFilters(ByYear(year), ByRound(round))
    req.Method = "results"
    return &req
}

func (req* RaceResultsRequest) Get(filters ...request.FilterFunc) (*RaceResult, error) {
    body, err := req.GetBody(filters...)
    if err != nil {
        return nil, err
    }
    races := parseRaceResultsResponse(body)
    if len(races) < 1 {
        return nil, &common.ErgastNotFoundError{}
    }
    return &races[0], nil
}

func parseRaceResultsResponse(body []byte) ([]RaceResult) {
    var results RaceResultsResponse
    json.Unmarshal(body, &results)
    return results.MRData.RaceTable.Races
}