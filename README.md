# golabview
Connect your go code with your labview code

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

	SliceToRead, data = labview.Unflatten(SliceToRead, cal.CalA)
	cal.CalA = data.([][]float32)

	SliceToRead, data = labview.Unflatten(SliceToRead, cal.CalB)
	cal.CalB = data.([][]float32)

	SliceToRead, data = labview.Unflatten(SliceToRead, cal.Date)
	cal.Date = data.(float64)

	return SliceToRead
}

// FlattenToByteSlice Calibration and return writed size
func (cal Calibration) FlattenToByteSlice(sliceToWrite []byte) int {
	var index int
	index += labview.Flatten(sliceToWrite[index:], cal.CalA)
	index += labview.Flatten(sliceToWrite[index:], cal.CalB)
	index += labview.Flatten(sliceToWrite[index:], cal.Date)
	return index
}

```
