package main

type Points int

const (
    // Qualifying
    Q1_Finish = 1
    Q2_Finish = 2
    Q3_Finish = 3
    Qualified_Ahead = 2
    Did_Not_Qualified = -5
    Quali_Disqualified = -10
    Qualified_P1 = 10
    Qualified_P2 = 9
    Qualified_P3 = 8
    Qualified_P4 = 7
    Qualified_P5 = 6
    Qualified_P6 = 5
    Qualified_P7 = 4
    Qualified_P8 = 3
    Qualified_P9 = 2
    Qualified_P10 = 1

    // Race
    Finished = 1
    Position_Gained = 2  // per position, max 10 pts
    Finished_Ahead = 3
    Fastest_Lap = 5
    Top10_Position_Lost = -2 // per position, max -10 pts
    Other_Position_Lost = -1 // per position, max -5 pts
    Not_Classified = -15
    Race_Disqualified = -20
    Finished_P1 = 25
    Finished_P2 = 18
    Finished_P3 = 15
    Finished_P4 = 12
    Finished_P5 = 10
    Finished_P6 = 8
    Finished_P7 = 6
    Finished_P8 = 4
    Finished_P9 = 2
    Finished_P10 = 1
)
