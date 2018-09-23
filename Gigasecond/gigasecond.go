package gigasecond

import "time"

//GigaSecond ...
const GigaSecond = 1e9

//AddGigasecond calculates a Giga second after a specified date
func AddGigasecond(t time.Time) time.Time {

	return t.Add(time.Duration(time.Second * GigaSecond))
}
