package ergast

import (
    "encoding/json"
    "github.com/mihaipstef/ergast-go/helpers/request"
    "github.com/mihaipstef/ergast-go/helpers/common"
)

type Race = common.Race

type SchedulesResponse struct {
    MRData struct {
        common.MRData
        RaceTable struct {
            Season string `json:"season,omitempty"`
            Races  []Race `json:"Races"`
        } `json:"RaceTable"`
    } `json:"MRData"`
}

type SchedulesRequest struct {
    request.Request
}

func NewSchedulesRequest(content_type request.ResponseContentType) (*SchedulesRequest) {
    req := SchedulesRequest{Request: request.Request{ContentType: content_type}}
    req.Method = "races"
    return &req
}

func (req* SchedulesRequest) Get(filters ...request.FilterFunc) ([]Race, error) {
    body, err := req.GetBody(filters...)
    if err != nil {
        return nil, err
    }
    races := parseSchedulesResponse(body)
    return races, nil
}

func parseSchedulesResponse(body []byte) ([]Race) {
    var schedule SchedulesResponse
    json.Unmarshal(body, &schedule)
    return schedule.MRData.RaceTable.Races
}