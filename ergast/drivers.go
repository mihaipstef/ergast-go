package ergast

import (
    "encoding/json"
    "github.com/mihaipstef/ergast-go/helpers/request"
    "github.com/mihaipstef/ergast-go/helpers/common"
)

type Driver = common.Driver

type DriversResponse struct {
    MRData struct {
        common.MRData
        DriverTable struct {
            DriverID        string `json:"driverId,omitempty"`
            Drivers  []Driver `json:"Drivers"`
        } `json:"DriverTable"`
    } `json:"MRData"`
}

type DriversRequest struct {
    request.Request
}

type DriverRequest struct {
    request.Request
}

func NewDriversRequest(content_type request.ResponseContentType) (*DriversRequest) {
    req := DriversRequest{Request: request.Request{ContentType: content_type}}
    req.Method = "drivers"
    return &req
}

func NewDriverRequest(driver string, content_type request.ResponseContentType) (*DriverRequest) {
    req := DriverRequest{Request: request.Request{ContentType: content_type}}
    req.Method = "drivers/"+driver
    return &req
}

func (req* DriverRequest) Get() (*Driver, error) {
    drivers, err := getDrivers(req)
    if err != nil {
        return nil, err
    }
    if len(drivers) < 1 {
        return nil, &common.ErgastNotFoundError{}
    }
    return &drivers[0], nil
}

func (req* DriversRequest) Get(filters ...request.FilterFunc) ([]Driver, error) {
    return getDrivers(req, filters...)
}

func getDrivers(req request.IRequest, filters ...request.FilterFunc) ([]Driver, error){
    body, err := req.GetBody(filters...)
    if err != nil {
        return nil, err
    }
    drivers := parseDriversResponse(body)
    return drivers, nil
}

func parseDriversResponse(body []byte) ([]Driver) {
    var drivers DriversResponse
    json.Unmarshal(body, &drivers)
    return drivers.MRData.DriverTable.Drivers
}