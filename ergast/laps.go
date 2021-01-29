package ergast

import (
    "strconv"
    "encoding/json"
    "github.com/mihaipstef/ergast-go/helpers/request"
    "github.com/mihaipstef/ergast-go/helpers/common"
)

type Timing struct {
    DriverID string             `json:"driverId"`
    Position int                `json:"position,string"`
    Time     common.LapDuration `json:"time,string"`
}

type Lap struct {
    Number  int         `json:"number,string"`
    Timings []Timing    `json:"Timings"`
}

type RaceLaps struct {
    common.Race
    Laps    []Lap   `json:"Laps"`
}

type LapsResponse struct {
    MRData struct {
        common.MRData
        RaceTable struct {
            Season string       `json:"season"`
            Round  string       `json:"round"`
            Races  []RaceLaps   `json:"Races"`
        } `json:"RaceTable"`
    } `json:"MRData"`
}

type LapsRequest struct {
    request.Request
}

func NewLapsRequest(year string, round int, lap int, content_type request.ResponseContentType) (*LapsRequest) {
    req := LapsRequest{Request: request.Request{ContentType: content_type}}
    req.AddFilters(ByYear(year), ByRound(round))
    req.Method = "laps/"+strconv.Itoa(lap)
    return &req
}

func (req* LapsRequest) Get(filters ...request.FilterFunc) (*RaceLaps, error) {
    body, err := req.GetBody(filters...)
    if err != nil {
        return nil, err
    }
    race, err := parseLapsResponse(body)
    return race, err
}

func parseLapsResponse(body []byte) (*RaceLaps, error) {
    var laps LapsResponse
    json.Unmarshal(body, &laps)
    if len(laps.MRData.RaceTable.Races) < 1 {
        return nil, &common.ErgastNotFoundError{}
    }
    return &laps.MRData.RaceTable.Races[0], nil
}