/*
Package gigasecond provides a function for adding one gigasecond to a supplied time.
*/
package gigasecond

import "time"

const gigasecond time.Duration = 1000000000 * time.Second

/*
AddGigasecond adds one gigasecond to a time parameter.
*/
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigasecond)
}
