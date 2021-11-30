# golabview

Connect your go code with your LabVIEW code ! 
This module is a bridge between LabVIEW and Golang.

# Functions 

## Flatten / Unflatten 
This module provide a solution to flatten and unflatten data from and to LabVIEW from a Golang project.
For the moment it is a manual method. An improvement could be to use the reflection.

Example : 
```

/**************************************************************************************************/
/* Calibration Struct
/**************************************************************************************************/

//Calibration struct
type Calibration struct {
	CalA [][]float32
	CalB [][]float32
	Date float64
}

/**************************************************************************************************/
/* Calibration Struct Flatten & UnFlatten
/**************************************************************************************************/

// UnflattenFromByteSlice Calibration return rest of readed slice
func (cal *Calibration) UnflattenFromByteSlice(SliceToRead []byte) []byte {
	var data interface{}

	SliceToRead, data = golabview.Unflatten(SliceToRead, cal.CalA)
	cal.CalA = data.([][]float32)

	SliceToRead, data = golabview.Unflatten(SliceToRead, cal.CalB)
	cal.CalB = data.([][]float32)

	SliceToRead, data = golabview.Unflatten(SliceToRead, cal.Date)
	cal.Date = data.(float64)

	return SliceToRead
}

// FlattenToByteSlice Calibration and return writed size
func (cal Calibration) FlattenToByteSlice(sliceToWrite []byte) int {
	var index int
	index += golabview.Flatten(sliceToWrite[index:], cal.CalA)
	index += golabview.Flatten(sliceToWrite[index:], cal.CalB)
	index += golabview.Flatten(sliceToWrite[index:], cal.Date)
	return index
}

```

## Time conversion

This module provide a solution to convert Golang Time and LabView Time.

Example : 
```
	timestampDoubleTest := golabview.TimeToDouble(time.Now())
	timestampTime := golabview.DoubleToTime(timestampDoubleTest)
	timestampTimeRec := golabview.TimeToTimeRec(timestampTime)
	timestampDouble := golabview.TimeRecToDouble(timestampTimeRec)
	timestampTimeRec2 := golabview.DoubleToTimeRec(timestampDouble)
	timestampTime2 := golabview.TimeRecToTime(timestampTimeRec2)
	timestampDouble2 := golabview.TimeToDouble(timestampTime2)

```
