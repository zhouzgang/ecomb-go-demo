package localtime

import "time"

var Timezone = "UTC"

var Location = time.UTC

func SetLocation(timezone string) error {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return err
	}
	Location = loc
	Timezone = timezone
	return nil
}

func Now() time.Time {
	return time.Now().In(Location)
}


