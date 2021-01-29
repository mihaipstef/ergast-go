package ergast

import (
    "encoding/json"
    "github.com/mihaipstef/ergast-go/helpers/request"
    "github.com/mihaipstef/ergast-go/helpers/common"
)

type Constructor = common.Constructor

type ConstructorsResponse struct {
    MRData struct {
        common.MRData
        ConstructorTable struct {
            ConstructorID string `json:"constructorId,omitempty"`
            Constructors  []Constructor `json:"Constructors"`
        } `json:"ConstructorTable"`
    } `json:"MRData"`
}

type ConstructorsRequest struct {
    request.Request
}

type ConstructorRequest struct {
    request.Request
}

func NewConstructorsRequest(content_type request.ResponseContentType) (*ConstructorsRequest) {
    req := ConstructorsRequest{Request: request.Request{ContentType: content_type}}
    req.Method = "constructors"
    return &req
}

func NewConstructorRequest(constructor string, content_type request.ResponseContentType) (*ConstructorRequest) {
    req := ConstructorRequest{Request: request.Request{ContentType: content_type}}
    req.Method = "constructors/"+constructor
    return &req
}

func (req* ConstructorRequest) Get() (*Constructor, error) {
    constructors, err := getConstructors(req)
    if err != nil {
        return nil, err
    }
    if len(constructors) < 1 {
        return nil, &common.ErgastNotFoundError{}
    }
    return &constructors[0], nil
}

func (req* ConstructorsRequest) Get(filters ...request.FilterFunc) ([]Constructor, error) {
    return getConstructors(req, filters...)
}

func getConstructors(req request.IRequest, filters ...request.FilterFunc) ([]Constructor, error){
    body, err := req.GetBody(filters...)
    if err != nil {
        return nil, err
    }
    constructors := parseConstructorsResponse(body)
    return constructors, nil
}

func parseConstructorsResponse(body []byte) ([]Constructor) {
    var constructors ConstructorsResponse
    json.Unmarshal(body, &constructors)
    return constructors.MRData.ConstructorTable.Constructors
}