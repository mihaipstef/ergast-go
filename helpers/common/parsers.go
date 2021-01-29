package common

import (
    "time"
    "strings"
)

type LapDuration struct {
    time.Duration
}

func (d *LapDuration) UnmarshalJSON(data []byte) (err error) {
    s := strings.Trim(string(data), `"`)
    s = strings.Replace(s, ":", "m", 1)
    s = strings.Replace(s, ".", "s", 1)
    s += "ms"
    d.Duration, err = time.ParseDuration(s)
    return err
}