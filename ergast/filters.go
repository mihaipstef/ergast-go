package ergast

import (
    "strconv"
    "github.com/mihaipstef/ergast-go/helpers/request"
)

func Limit(limit int) (request.FilterFunc) {
    return func () (string, string, request.FilterType) {
        return "limit", strconv.Itoa(limit), request.PARAMETER
    }
}

func Offset(offset int) (request.FilterFunc) {
    return func () (string, string, request.FilterType) {
        return "offset", strconv.Itoa(offset), request.PARAMETER
    }
}

func ByYear(year string) (request.FilterFunc) {
    return func () (string, string, request.FilterType) {
        return "year", year, request.QUERY
    }
}

func ByRound(round int) (request.FilterFunc) {
    return func () (string, string, request.FilterType) {
        return "round", strconv.Itoa(round), request.QUERY
    }
}

func ByDriver(driver string) (request.FilterFunc) {
    return func () (string, string, request.FilterType) {
        return "drivers", driver, request.QUERY
    }
}

func ByConstructor(constructor string) (request.FilterFunc) {
    return func () (string, string, request.FilterType) {
        return "constructors", constructor, request.QUERY
    }
}