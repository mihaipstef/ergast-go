package ergast

import (
    "encoding/json"
    "github.com/mihaipstef/ergast-go/helpers/request"
    "github.com/mihaipstef/ergast-go/helpers/common"
)

type Season struct {
    Season string `json:"season,omitempty"`
    URL    string `json:"url,omitempty"`
}

type SeasonsResponse struct {
    MRData struct {
        common.MRData
        SeasonTable struct {
            ConstructorId string `json:"constructorId,omitempty"`
            DriverId string `json:"driverId,omitempty"`
            Seasons []Season `json:"Seasons,omitempty"`
        } `json:"SeasonTable,omitempty"`
    } `json:"MRData,omitempty"`
}

type SeasonsRequest struct {
    request.Request
}

func NewSeasonsRequest(content_type request.ResponseContentType) (*SeasonsRequest) {
    req := SeasonsRequest{Request: request.Request{ContentType: content_type}}
    req.Method = "seasons"
    return &req
}

func (req* SeasonsRequest) Get(filters ...request.FilterFunc) ([]Season, error) {
    body, err := req.GetBody(filters...)
    if err != nil {
        return nil, err
    }
    seasons := parseSeasonsResponse(body)
    return seasons, nil
}

func parseSeasonsResponse(body []byte) ([]Season) {
    var seasons SeasonsResponse
    json.Unmarshal(body, &seasons)
    return seasons.MRData.SeasonTable.Seasons
}