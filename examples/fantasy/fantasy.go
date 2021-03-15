package main

import (
    "github.com/mihaipstef/ergast-go/ergast"
    "os"
    "fmt"
    "strconv"
)

func printLeaderboardCsv(leaderboard *LeaderBoard) () {
    for leader, scores := range *leaderboard {
        leader_scores := leader + ","
        for _, score := range scores {
            leader_scores += strconv.Itoa(score)
            leader_scores += ","
        }
        fmt.Println(leader_scores)
    }
}

func main() {

    if len(os.Args) < 2 {
        fmt.Println("Call: fantasy <year>")
        return
    }

    year := os.Args[1]

    races_req := ergast.NewSchedulesRequest(ergast.JSON)
    races, err_races := races_req.Get(ergast.ByYear(year))
    if err_races != nil {
        fmt.Println("Failed to get the races table")
        return
    }

    var round_results *ergast.RaceResult
    var err error
    var driver_table *LeaderBoard
    var constructor_table *LeaderBoard

    process := MakeResultProcessor(len(races))

    for _, race := range races {
        round_results, err = ergast.NewRaceResultsRequest(year, race.Round, ergast.JSON).Get()
        if err != nil {
            fmt.Println("Failed to get the results for round ", race.Round)
            return
        }

        for _, result := range round_results.Results {
            driver_table, constructor_table, err = process(race.Round - 1, &result)
            if err != nil {
                fmt.Println("Something went wrong in %d/%d", race.Round, len(races))
                return
            }
        }
    }

    // Print results
    printLeaderboardCsv(driver_table)
    printLeaderboardCsv(constructor_table)
}