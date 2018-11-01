package convert

import (
	"strconv"
	"time"
)

// UnixTimeMSToString converts a string representing unix time
// in ms to a human-readable string.
// No error checking is performed.
func UnixTimeMSToString(t string) string {
	msec, _ := strconv.ParseInt(t, 10, 64)
	tm := time.Unix(msec/1000, 0)
	return tm.String()
}
