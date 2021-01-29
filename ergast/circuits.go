package ergast

import (
    "encoding/json"
    "github.com/mihaipstef/ergast-go/helpers/request"
    "github.com/mihaipstef/ergast-go/helpers/common"
)

type Circuit = common.Circuit

type CircuitsResponse struct {
    MRData struct {
        common.MRData
        CircuitTable struct {
            Circuits  []Circuit `json:"Circuits"`
        } `json:"CircuitTable"`
    } `json:"MRData"`
}

type CircuitsRequest struct {
    request.Request
}

type CircuitRequest struct {
    request.Request
}

func NewCircuitsRequest(content_type request.ResponseContentType) (*CircuitsRequest) {
    req := CircuitsRequest{Request: request.Request{ContentType: content_type}}
    req.Method = "circuits"
    return &req
}

func NewCircuitRequest(circuit string, content_type request.ResponseContentType) (*CircuitRequest) {
    req := CircuitRequest{Request: request.Request{ContentType: content_type}}
    req.Method = "circuits/"+circuit
    return &req
}

func (req* CircuitRequest) Get() (*Circuit, error) {
    circuits, err := getCircuits(req)
    if err != nil {
        return nil, err
    }
    if len(circuits) < 1 {
        return nil, &common.ErgastNotFoundError{}
    }
    return &circuits[0], nil
}

func (req* CircuitsRequest) Get(filters ...request.FilterFunc) ([]Circuit, error) {
    return getCircuits(req, filters...)
}

func getCircuits(req request.IRequest, filters ...request.FilterFunc) ([]Circuit, error){
    body, err := req.GetBody(filters...)
    if err != nil {
        return nil, err
    }
    circuits := parseCircuitsResponse(body)
    return circuits, nil
}

func parseCircuitsResponse(body []byte) ([]Circuit) {
    var circuits CircuitsResponse
    json.Unmarshal(body, &circuits)
    return circuits.MRData.CircuitTable.Circuits
}
