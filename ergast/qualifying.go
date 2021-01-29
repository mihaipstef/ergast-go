package ergast

import (
    "encoding/json"
    "github.com/mihaipstef/ergast-go/helpers/request"
    "github.com/mihaipstef/ergast-go/helpers/common"
)

type QualifyingResult struct {
    common.Result
    Q1  common.LapDuration  `json:"Q1,string,omitempty"`
    Q2  common.LapDuration  `json:"Q2,string,omitempty"`
    Q3  common.LapDuration  `json:"Q3,string,omitempty"`
}

type Qualifying struct {
    common.Race
    Results    []QualifyingResult  `json:"QualifyingResults"`
}

type QualifyingResultsResponse struct {
    MRData struct {
        common.MRData
        RaceTable struct {
            Season string  `json:"season"`
            Round  int     `json:"round,string"`
            Qualifying  []Qualifying  `json:"Races"`
        } `json:"RaceTable"`
    } `json:"MRData"`
}

type QualifyingResultsRequest struct {
    request.Request
}

func NewQualifyingResultsRequest(year string, round int, content_type request.ResponseContentType) (*QualifyingResultsRequest) {
    req := QualifyingResultsRequest{Request: request.Request{ContentType: content_type}}
    req.AddFilters(ByYear(year), ByRound(round))
    req.Method = "qualifying"
    return &req
}

func (req* QualifyingResultsRequest) Get(filters ...request.FilterFunc) (*Qualifying, error) {
    body, err := req.GetBody(filters...)
    if err != nil {
        return nil, err
    }
    qualis := parseQualifyingResultsResponse(body)
    if len(qualis) < 1 {
        return nil, &common.ErgastNotFoundError{}
    }
    return &qualis[0], nil
}

func parseQualifyingResultsResponse(body []byte) ([]Qualifying) {
    var results QualifyingResultsResponse
    json.Unmarshal(body, &results)
    return results.MRData.RaceTable.Qualifying
}