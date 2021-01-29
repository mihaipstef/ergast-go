package ergast

import (
    "encoding/json"
    "github.com/mihaipstef/ergast-go/helpers/request"
    "github.com/mihaipstef/ergast-go/helpers/common"
)

type Standing struct {
    Position        int     `json:"position,string,omitempty"`
    PositionText    string  `json:"positionText,omitempty"`
    Points          int     `json:"points,string,omitempty"`
    Wins            int     `json:"wins,string,omitempty"`
}

type DriverStanding struct {
    Standing
    Driver          common.Driver           `json:"Driver,omitempty"`
    Constructors    []common.Constructor    `json:"Constructors,omitempty"`
}

type ConstructorStanding struct {
    Standing
    Constructor    common.Constructor `json:"Constructors,omitempty"`
}

type Standings struct {
    Season                  string                  `json:"season,omitempty"`
    Round                   int                     `json:"round,string,omitempty"`
    DriverStandings         []DriverStanding        `json:"DriverStandings,omitempty"`
    ConstructorStandings    []ConstructorStanding   `json:"ConstructorStandings,omitempty"`
}

type StandingsResponse struct {
    MRData struct {
        common.MRData
        StandingsTable struct {
            StandingsList  []Standings `json:"StandingsLists"`
        } `json:"StandingsTable"`
    } `json:"MRData"`
}

type DriverStandingsRequest struct {
    request.Request
}

type ConstructorStandingsRequest struct {
    request.Request
}

func NewDriverStandingsRequest(content_type request.ResponseContentType) (*DriverStandingsRequest) {
    req := DriverStandingsRequest{Request: request.Request{ContentType: content_type}}
    req.Method = "driverStandings"
    return &req
}

func NewConstructorStandingsRequest(content_type request.ResponseContentType) (*ConstructorStandingsRequest) {
    req := ConstructorStandingsRequest{Request: request.Request{ContentType: content_type}}
    req.Method = "constructorStandings"
    return &req
}

func (req* DriverStandingsRequest) Get(filters ...request.FilterFunc) ([]Standings, error) {
    return get(req, filters...)
}

func (req* ConstructorStandingsRequest) Get(filters ...request.FilterFunc) ([]Standings, error) {
    return get(req, filters...)
}

func get(req request.IRequest, filters ...request.FilterFunc) ([]Standings, error){
    body, err := req.GetBody(filters...)
    if err != nil {
        return nil, err
    }
    standings := parseStandingsResponse(body)
    return standings, nil
}

func parseStandingsResponse(body []byte) ([]Standings) {
    var standings StandingsResponse
    json.Unmarshal(body, &standings)
    return standings.MRData.StandingsTable.StandingsList
}