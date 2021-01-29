package ergast

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "time"
    "github.com/mihaipstef/ergast-go/helpers/common"
    "github.com/mihaipstef/ergast-go/helpers/request"
)

func TestCircuits(t *testing.T) {
    circuit, _ := NewCircuitRequest("portimao", request.JSON).Get()
    assert.Equal(t, "portimao", circuit.CircuitID, "")
    req := NewCircuitsRequest(request.JSON)
    circuits, _ := req.Get(ByDriver("raikkonen"), ByYear("2020"))
    assert.Equal(t, 14, len(circuits), "")
}

func TestConstructors(t *testing.T) {
    constructor, _ := NewConstructorRequest("alfa", request.JSON).Get()
    assert.Equal(t, "alfa", constructor.ConstructorID, "")
    req := NewConstructorsRequest(request.JSON)
    constructors, _ := req.Get(ByDriver("raikkonen"), ByYear("2020"))
    assert.Equal(t, 1, len(constructors), "")
    assert.Equal(t, "alfa", constructors[0].ConstructorID, "")
}

func TestDrivers(t *testing.T) {
    req := NewDriverRequest("raikkonen", request.JSON)
    driver, _ := req.Get()
    assert.Equal(t, "raikkonen", driver.DriverID, "")
    reqs := NewDriversRequest(request.JSON)
    drivers, _ := reqs.Get(ByConstructor("alfa"), ByYear("2020"))
    assert.Equal(t, 2, len(drivers), "")
    assert.Equal(t, "GIO", drivers[0].Code, "")
    assert.Equal(t, "RAI", drivers[1].Code, "")
    drivers, _ = reqs.Get(Limit(1))
    assert.Equal(t, 1, len(drivers), "")
}

func TestLaps(t *testing.T) {
    race, _ := NewLapsRequest("2020", 1, 1, request.JSON).Get(ByDriver("raikkonen"))
    assert.Equal(t, "red_bull_ring", race.Circuit.CircuitID, "")
    expected_lap_time, _ := time.ParseDuration("1m20s781ms")
    assert.Equal(t, common.LapDuration{Duration: expected_lap_time}, race.Laps[0].Timings[0].Time, "")
}

func TestResults(t *testing.T) {
    race, _ := NewRaceResultsRequest("2008", 5, request.JSON).Get(ByDriver("raikkonen"))
    assert.Equal(t, 3, race.Results[0].Position, "")
    assert.Equal(t, 58, race.Results[0].Laps, "")
    assert.Equal(t, 4, race.Results[0].Grid, "")
    quali, _ := NewQualifyingResultsRequest("2008", 5, request.JSON).Get(ByDriver("raikkonen"))
    assert.Equal(t, 4, quali.Results[0].Position, "")
    expected_lap_time, _ := time.ParseDuration("1m27s936ms")
    assert.Equal(t, common.LapDuration{Duration: expected_lap_time}, quali.Results[0].Q3, "")
}

func TestSchedules(t *testing.T) {
    req := NewSchedulesRequest(request.JSON)
    races, _ := req.Get(ByDriver("raikkonen"), ByYear("2019"))
    assert.Equal(t, 21, len(races), "")
    assert.Equal(t, "2019-03-17", races[0].Date, "")
}

func TestSeasons(t *testing.T) {
    req := NewSeasonsRequest(request.JSON)
    seasons, _ := req.Get(ByDriver("hakkinen"), Limit(20))
    assert.Equal(t, 11, len(seasons), "")
}

func TestStandings(t *testing.T) {
    req := NewConstructorStandingsRequest(request.JSON)
    constructors, _ := req.Get(ByConstructor("alfa"), ByYear("2020"), ByRound(15))
    assert.Equal(t, 8, constructors[0].ConstructorStandings[0].Position, "")
    req_driver := NewDriverStandingsRequest(request.JSON)
    drivers, _ := req_driver.Get(ByDriver("raikkonen"), ByYear("2020"), ByRound(15))
    assert.Equal(t, 16, drivers[0].DriverStandings[0].Position, "")
}