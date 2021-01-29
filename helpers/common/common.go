package common

type MRData struct {
    Xmlns       string  `json:"xmlns,omitempty"`
    Series      string  `json:"series,omitempty"`
    URL         string  `json:"url,omitempty"`
    Limit       int     `json:"limit,string,omitempty"`
    Offset      int     `json:"offset,string,omitempty"`
    Total       int     `json:"total,string,omitempty"`
}

type Location struct {
    Lat      float32    `json:"lat,string"`
    Long     float32    `json:"long,string"`
    Locality string     `json:"locality,omitempty"`
    Country  string     `json:"country,omitempty"`
}

type Circuit struct {
    CircuitID   string   `json:"circuitId,omitempty"`
    URL         string   `json:"url,omitempty"`
    CircuitName string   `json:"circuitName,omitempty"`
    Location    Location `json:"Location,omitempty"`
}

type Race struct {
    Season   string     `json:"season,omitempty"`
    Round    int        `json:"round,string,omitempty"`
    URL      string     `json:"url,omitempty"`
    RaceName string     `json:"raceName,omitempty"`
    Circuit  Circuit    `json:"Circuit,omitempty"`
    Date     string     `json:"date,omitempty"`
    Time     string     `json:"time,omitempty"`
}

type Driver struct {
    DriverID        string `json:"driverId,omitempty"`
    PermanentNumber int    `json:"permanentNumber,string,omitempty"`
    Code            string `json:"code,omitempty"`
    URL             string `json:"url,omitempty"`
    GivenName       string `json:"givenName,omitempty"`
    FamilyName      string `json:"familyName,omitempty"`
    DateOfBirth     string `json:"dateOfBirth,omitempty"`
    Nationality     string `json:"nationality,omitempty"`
}

type Constructor struct {
    ConstructorID string `json:"constructorId,omitempty"`
    URL           string `json:"url,omitempty"`
    Name          string `json:"name,omitempty"`
    Nationality   string `json:"nationality,omitempty"`
}

type Result struct {
    Number       int             `json:"number,string,omitempty"`
    Position     int             `json:"position,string,omitempty"`
    PositionText string          `json:"positionText,omitempty"`
    Driver       Driver          `json:"Driver"`
    Constructor  Constructor     `json:"Constructor"`
}

type ErgastNotFoundError struct {}

func (e *ErgastNotFoundError) Error() string {
    return "ErgastNotFoundError"
}