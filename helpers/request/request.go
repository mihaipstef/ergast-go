package request

import (
    "net/http"
    "strings"
    "io/ioutil"
)

type FilterType int

const (
    QUERY FilterType = iota
    PARAMETER
)

type FilterFunc func ()(string, string, FilterType)

type IRequest interface {
    GetBody(args ...FilterFunc) ([]byte, error)
}

type ResponseContentType int

const (
    JSON ResponseContentType = iota
    XML
)

var contentTypeToString = []string{"json","xml"}

type Request struct {
    Query map[string]string
    Parameters map[string]string
    ContentType ResponseContentType
    Method string
}

func (req *Request) buildParameters() string {
    var builder strings.Builder
    builder.WriteString("?")
    for key, value := range req.Parameters {
        builder.WriteString(key + "=")
        builder.WriteString(value)
        builder.WriteString("&")
    }
    return builder.String()
}

const BaseUrl = "http://ergast.com/api/f1"

func (req *Request) url() string {
    var builder strings.Builder
    builder.WriteString(BaseUrl)
    builder.WriteString(req.buildQuery())
    builder.WriteString("/")
    builder.WriteString(req.Method)
    builder.WriteString(".")
    builder.WriteString(contentTypeToString[req.ContentType])
    builder.WriteString(req.buildParameters())
    return builder.String()
}

func (req *Request) buildQuery() (string) {
    var builder strings.Builder
    if year, ok := req.Query["year"]; ok {
        builder.WriteString("/" + year)
    }
    if round, ok := req.Query["round"]; ok {
        builder.WriteString("/" + round)
    }

    for key, value := range req.Query {
        if key == "year" || key == "round" {
            continue
        }
        builder.WriteString("/")
        builder.WriteString(key)
        builder.WriteString("/")
        builder.WriteString(value)
    }
    return builder.String()
}

func (req *Request) getRaw() (res *http.Response, err error) {
    url := req.url()
    res, err = http.Get(url);
    return res, err
}

func (req *Request) GetBodyNoFilters() (body []byte, err error) {
    res, err := req.getRaw()
    if err != nil {
        return nil, err
    }
    body, err = ioutil.ReadAll(res.Body)
    return body, err
}

func (req *Request) AddFilters(args ...FilterFunc) () {
    for i := 0; i < len(args); i++ {
        key, value, t := args[i]()
        if t == QUERY {
            req.addQuery(key, value)
        } else if t == PARAMETER {
            req.addParameter(key, value)
        }
    }
}

func (req *Request) GetBody(args ...FilterFunc) (body []byte, err error) {
    req.AddFilters(args...)
    return req.GetBodyNoFilters()
}

func (req *Request) addQuery(key string, value string) {
    if req.Query == nil {
        req.Query = make(map[string]string)
    }
    req.Query[key] = value
}

func (req *Request) addParameter(key string, value string) {
    if req.Parameters == nil {
        req.Parameters = make(map[string]string)
    }
    req.Parameters[key] = value
}
