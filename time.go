package labview



import (

	"fmt"

	"math"

	"time"

)



/**************************************************************************************************/

/* Labview Time Constant

/**************************************************************************************************/


//To do : Create a goof fomarting function to replace this

//TimeRecLabview get string format

const (

	StringOfTimestamp        = 0 // string of full Timestamp yyyy mm dd HH MM SS

	StringOfDate             = 1 // String of Date yyyy mm dd

	StringOfDateHour         = 2 // String of Date yyyy mm dd HH

	StringOfHourMinuteSecond = 3 // String of Hour and minute HH MM SS

	StringOfHourMinute       = 4 // String of Hour and minute HH MM

	StringOfHour             = 5 // String of Hour HH

)



const (

	p9 = 1000000000

)



/**************************************************************************************************/

/* Labview Time External Function

/**************************************************************************************************/



// SetTime of Go time to Labview TimeRec

func (ltr *TimeRecLabview) SetTime(timeToWrite time.Time) {

	*ltr = TimeToTimeRec(timeToWrite)

}



// UnflattenFromByteSlice TimeRecLabview

func (ltr *TimeRecLabview) UnflattenFromByteSlice(SliceToRead []byte) {

	var data interface{}

	SliceToRead, data = UnflattenTimeRecLabview(SliceToRead)

	timeRec := data.(TimeRecLabview)

	*ltr = timeRec

}



//TimeRecToTime convert Labview TimeRec to Go time

func TimeRecToTime(TimeRec TimeRecLabview) (timeGo time.Time) {

	timeGo = time.Date(

		int(TimeRec.Year),

		time.Month(TimeRec.Month),

		int(TimeRec.DayOfMonth),

		int(TimeRec.Hour),

		int(TimeRec.Minute),

		int(TimeRec.Second),

		int(TimeRec.FractionalSecond*p9),

		time.FixedZone("UTC", 0),

	)



	return timeGo

}



//TimeRecToDouble convert Labview TimeRec to Labview double timestamp

func TimeRecToDouble(TimeRec TimeRecLabview) float64 {

	timetoConvert := TimeRecToTime(TimeRec)

	return TimeToDouble(timetoConvert)

}



//TimeToTimeRec convert Go time to Labview TimeRec

func TimeToTimeRec(timeToConvert time.Time) (TimeRec TimeRecLabview) {

	TimeRec.FractionalSecond = float64(timeToConvert.Nanosecond()) / p9

	TimeRec.Second = int32(timeToConvert.Second())

	TimeRec.Minute = int32(timeToConvert.Minute())

	TimeRec.Hour = int32(timeToConvert.Hour())

	TimeRec.DayOfMonth = int32(timeToConvert.Day())

	TimeRec.Month = int32(timeToConvert.Month())

	TimeRec.Year = int32(timeToConvert.Year())

	TimeRec.DayOfWeek = int32(timeToConvert.Weekday())

	TimeRec.DayOfYear = int32(timeToConvert.YearDay())

	TimeRec.Dst = 0



	return TimeRec

}



//TimeToTimeStampLabview convert Go time to Labview Timestamp

func TimeToTimeStampLabview(timeToConvert time.Time) (TimeToTimeStampLabview TimeStampLabview) {

	dblTime := TimeToDouble(timeToConvert)

	second := math.Trunc(dblTime)

	fractional := dblTime - second

	fractional = math.Abs(fractional)

	TimeToTimeStampLabview.Fraction = uint64(fractional * 18446744073709551600)

	TimeToTimeStampLabview.Second = int64(second)



	return TimeToTimeStampLabview

}



//TimeToDouble convert Go time to Labview double Timestamp

func TimeToDouble(timeToConvert time.Time) float64 {

	labviewTime := time.Date(1904, 1, 1, 0, 0, 0, 0, time.FixedZone("UTC", 0))

	timeDuration := timeToConvert.Sub(labviewTime)

	return float64(timeDuration.Nanoseconds()) / p9

}



//DoubleToTimeRec Convert Labviewdouble Timestamp to Labview TimeRec

func DoubleToTimeRec(timeToConvert float64) (TimeRec TimeRecLabview) {

	TimeRec.SetTime(DoubleToTime(timeToConvert))

	return TimeRec

}



//DoubleToTime Convert Labview double Timestamp to go Time

func DoubleToTime(timeToConvert float64) (timeGo time.Time) {

	labviewTime := time.Date(1904, 1, 1, 0, 0, 0, 0, time.FixedZone("UTC", 0))

	timeGo = labviewTime.Add(time.Nanosecond * time.Duration(timeToConvert*p9))

	return timeGo

}



// GetString get string from TimeRecLabview

func (ltr *TimeRecLabview) GetString(formatTime int) string {

	var stringToReturn string



	switch formatTime {

	default:

		stringToReturn = fmt.Sprintf("%04d_%02d_%02d_%02d_%02d_%02d", ltr.Year, ltr.Month, ltr.DayOfMonth, ltr.Hour, ltr.Minute, ltr.Second)

	case StringOfDate:

		stringToReturn = fmt.Sprintf("%04d_%02d_%02d", ltr.Year, ltr.Month, ltr.DayOfMonth)

	case StringOfDateHour:

		stringToReturn = fmt.Sprintf("%04d_%02d_%02d_%02d", ltr.Year, ltr.Month, ltr.DayOfMonth, ltr.Hour)

	case StringOfHourMinuteSecond:

		stringToReturn = fmt.Sprintf("%04d_%02d_%02d", ltr.Hour, ltr.Minute, ltr.Second)

	case StringOfHourMinute:

		stringToReturn = fmt.Sprintf("%02d_%02d", ltr.Hour, ltr.Minute)

	case StringOfHour:

		stringToReturn = fmt.Sprintf("%02d", ltr.Hour)

	}



	return stringToReturn

}



// SetFromString get string from TimeRecLabview

func (ltr *TimeRecLabview) SetFromString(timeString string, formatTime int) bool {

	var err error



	success := true

	switch formatTime {

	default:

		_, err = fmt.Sscanf(timeString, "%04d_%02d_%04d_%02d_%02d_%02d", ltr.Year, ltr.Month, ltr.DayOfMonth, ltr.Hour, ltr.Minute, ltr.Second)

	case StringOfDate:

		_, err = fmt.Sscanf(timeString, "%04d_%02d_%02d", ltr.Year, ltr.Month, ltr.DayOfMonth)

	case StringOfDateHour:

		_, err = fmt.Sscanf(timeString, "%04d_%02d_%02d_02d", ltr.Year, ltr.Month, ltr.DayOfMonth, ltr.Hour)

	case StringOfHourMinuteSecond:

		_, err = fmt.Sscanf(timeString, "%04d_%02d_%02d", ltr.Hour, ltr.Minute, ltr.Second)

	case StringOfHourMinute:

		_, err = fmt.Sscanf(timeString, "%02d_%02d", ltr.Hour, ltr.Minute)

	case StringOfHour:

		_, err = fmt.Sscanf(timeString, "%02d", ltr.Hour)

	}

	if err != nil {

		success = false

	}

	return success

}



//TestTimeConversion check conversion function

func TestTimeConversion() {

	//01.05.2020 - 08:12:45.01234579



	timestampDoubleTest := 3671165565.01234579



	timestampTime := DoubleToTime(timestampDoubleTest)



	timestampTimeRec := TimeToTimeRec(timestampTime)



	timestampDouble := TimeRecToDouble(timestampTimeRec)



	timestampTimeRec2 := DoubleToTimeRec(timestampDouble)



	timestampTime2 := TimeRecToTime(timestampTimeRec2)



	timestampDouble2 := TimeToDouble(timestampTime2)



	if timestampDouble2 == timestampDoubleTest { // should use a tolerance but it's work on compuer, on target too?



	} else {



	}



}



/**************************************************************************************************/

/* Labview Time Struct

/**************************************************************************************************/



//TimeRecLabview Labview TimeStamp

type TimeRecLabview struct {

	FractionalSecond float64 `json:"fractional second"`

	Second           int32   `json:"second"`

	Minute           int32   `json:"minute"`

	Hour             int32   `json:"hour"`

	DayOfMonth       int32   `json:"day of month"`

	Month            int32   `json:"month"`

	Year             int32   `json:"year"`

	DayOfWeek        int32   `json:"day of week"`

	DayOfYear        int32   `json:"day of year"`

	Dst              int32   `json:"DST"`

}



//TimeStampLabview struct

type TimeStampLabview struct {

	Second   int64

	Fraction uint64

}



/**************************************************************************************************/

/* Labview Time Struct Flatten & UnFlatten

/**************************************************************************************************/



// FlattenToByteSlice flatten TimeStamp to byte slice

func (ltr *TimeRecLabview) FlattenToByteSlice(dataSlice []byte) int {

	var index int

	index += Flatten(dataSlice[index:], ltr.FractionalSecond)

	index += Flatten(dataSlice[index:], ltr.Second)

	index += Flatten(dataSlice[index:], ltr.Minute)

	index += Flatten(dataSlice[index:], ltr.Hour)

	index += Flatten(dataSlice[index:], ltr.DayOfMonth)

	index += Flatten(dataSlice[index:], ltr.Month)

	index += Flatten(dataSlice[index:], ltr.Year)

	index += Flatten(dataSlice[index:], ltr.DayOfWeek)

	index += Flatten(dataSlice[index:], ltr.DayOfYear)

	index += Flatten(dataSlice[index:], ltr.Dst)

	return index

}



//UnflattenTimeRecLabview unflatten TimeRec

func UnflattenTimeRecLabview(SliceToRead []byte) (SliceReaded []byte, ltr TimeRecLabview) {

	var data interface{}



	SliceToRead, data = Unflatten(SliceToRead, ltr.FractionalSecond)

	ltr.FractionalSecond = data.(float64)



	SliceToRead, data = Unflatten(SliceToRead, ltr.Second)

	ltr.Second = data.(int32)



	SliceToRead, data = Unflatten(SliceToRead, ltr.Minute)

	ltr.Minute = data.(int32)



	SliceToRead, data = Unflatten(SliceToRead, ltr.Hour)

	ltr.Hour = data.(int32)



	SliceToRead, data = Unflatten(SliceToRead, ltr.DayOfMonth)

	ltr.DayOfMonth = data.(int32)



	SliceToRead, data = Unflatten(SliceToRead, ltr.Month)

	ltr.Month = data.(int32)



	SliceToRead, data = Unflatten(SliceToRead, ltr.Year)

	ltr.Year = data.(int32)



	SliceToRead, data = Unflatten(SliceToRead, ltr.DayOfWeek)

	ltr.DayOfWeek = data.(int32)



	SliceToRead, data = Unflatten(SliceToRead, ltr.DayOfYear)

	ltr.DayOfYear = data.(int32)



	SliceToRead, data = Unflatten(SliceToRead, ltr.Dst)

	ltr.Dst = data.(int32)

	return SliceToRead, ltr

}
