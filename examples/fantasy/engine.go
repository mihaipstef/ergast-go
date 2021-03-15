package main

import (
    "github.com/mihaipstef/ergast-go/ergast"
    "errors"
)

type TeamMate struct {
    Position int
    Grid int
    DriverID string
}

type LeaderBoard map[string][]int

type fantasyEngine struct {
    team_mates map[string]TeamMate
    driver_table LeaderBoard
    constructor_table LeaderBoard
    finished_points []int
    qualification_points []int
    no_of_rounds int
}

func MakeResultProcessor(no_of_rounds int) (func(int, *ergast.Result)(*LeaderBoard, *LeaderBoard, error)) {
    engine := fantasyEngine{team_mates: make(map[string]TeamMate), driver_table: make(LeaderBoard), constructor_table: make(LeaderBoard)}

    engine.finished_points = []int{Finished_P1, Finished_P2, Finished_P3, Finished_P4, Finished_P5, Finished_P6, Finished_P7, Finished_P8, Finished_P9, Finished_P10}
    engine.qualification_points = []int{Qualified_P1, Qualified_P2, Qualified_P3, Qualified_P4, Qualified_P5, Qualified_P6, Qualified_P7, Qualified_P8, Qualified_P9, Qualified_P10}
    engine.no_of_rounds = no_of_rounds

    return func(round int, result *ergast.Result)(*LeaderBoard, *LeaderBoard, error) {
        err := engine.processRaceResults(round, result);
        return &engine.driver_table, &engine.constructor_table, err
    }
}

func max(a, b int) (int) {
    if a > b {
        return a
    }
    return b
}

func (engine *fantasyEngine) processRaceResults(round int, result *ergast.Result) (error) {

    if round < 0 || round >= engine.no_of_rounds {
        return errors.New("Round is not in range")
    }

    if _, found := engine.driver_table[result.Driver.DriverID]; !found {
        engine.driver_table[result.Driver.DriverID] = make([]int, engine.no_of_rounds)
    }

    if _, found := engine.constructor_table[result.Constructor.ConstructorID]; !found {
        engine.constructor_table[result.Constructor.ConstructorID] = make([]int, engine.no_of_rounds)
    }

    // Team mates

    if team_mate, found := engine.team_mates[result.Constructor.ConstructorID]; found {
        // compare team mates race result and update points
        first_team_mate := team_mate.DriverID
        if result.Position < team_mate.Position {
            // swap
            first_team_mate = result.Driver.DriverID
        }
        // update points
        engine.driver_table[first_team_mate][round] += Finished_Ahead;

        // compare team mates qualification result and update points
        first_team_mate = team_mate.DriverID
        if result.Grid < team_mate.Grid {
            // swap
            first_team_mate = result.Driver.DriverID
        }
        // update points
        engine.driver_table[first_team_mate][round] += Qualified_Ahead;

    } else {
        // save for team mates comparison
        engine.team_mates[result.Constructor.ConstructorID] = TeamMate{Position:result.Position, Grid:result.Grid, DriverID:result.Driver.DriverID}
    }

    //Qaualification
    if result.Grid == 0 {
        // did not qualified
        engine.driver_table[result.Driver.DriverID][round] += Did_Not_Qualified

    } else if result.Grid >= 1 && result.Grid <= 10 {
        // Q1
        engine.driver_table[result.Driver.DriverID][round] += engine.qualification_points[result.Grid-1] + Q3_Finish
        engine.constructor_table[result.Constructor.ConstructorID][round] += engine.qualification_points[result.Grid-1] + Q3_Finish
    } else if result.Grid >= 1 && result.Grid <= 15 {
        // Q2
        engine.driver_table[result.Driver.DriverID][round] += Q2_Finish
        engine.constructor_table[result.Constructor.ConstructorID][round] += Q2_Finish
    } else {
        // Q3
        engine.driver_table[result.Driver.DriverID][round] += Q1_Finish
        engine.constructor_table[result.Constructor.ConstructorID][round] += Q1_Finish
    }

    // Finished
    if result.Position >= 1 && result.Position <= 10 {
        engine.driver_table[result.Driver.DriverID][round] += engine.finished_points[result.Position-1]
        engine.constructor_table[result.Constructor.ConstructorID][round] += engine.finished_points[result.Position-1]
    }

    if result.PositionText == "R" {
        engine.driver_table[result.Driver.DriverID][round] += Not_Classified
    } else {
        engine.driver_table[result.Driver.DriverID][round] += Finished
    }

    // Position gain/lost
    position_offset := result.Position - result.Grid
    if position_offset < 0 {
        // position gain
        position_offset = max(5, -position_offset)
        engine.driver_table[result.Driver.DriverID][round] += position_offset * Position_Gained
        engine.constructor_table[result.Constructor.ConstructorID][round] += position_offset * Position_Gained
    } else {
        // position lost
        position_offset = max(5, position_offset)
        if result.Position <= 10 {
            engine.driver_table[result.Driver.DriverID][round] += position_offset * Top10_Position_Lost
            engine.constructor_table[result.Constructor.ConstructorID][round] += position_offset * Top10_Position_Lost
        } else {
            engine.driver_table[result.Driver.DriverID][round] += position_offset * Other_Position_Lost
            engine.constructor_table[result.Constructor.ConstructorID][round] += position_offset * Other_Position_Lost
        }
    }

    return nil

}