package helper

import "time"

//ConvertStringToTimeStamp function to convert time string to time.Time
//you can use layout for parsing the time
func ConvertStringToTimeStamp(dataTime string) time.Time {
	layout := "2006-01-02"
	t, err := time.Parse(layout, dataTime)

	if err != nil {
		return time.Time{}
	}

	return t
}