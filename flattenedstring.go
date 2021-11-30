package golabview



import (

	"math"

)



/**************************************************************************************************/

/* Flatten To Byte Slice public function (Only Big Endian)

/**************************************************************************************************/



//Flatten data into slice and return the write size, 0 is error

func Flatten(SliceToWrite []byte, element interface{}) int {

	switch data := element.(type) {

	case bool:

		return flattenBool(SliceToWrite, data)

	case uint8:

		return flattenU8(SliceToWrite, data)

	case int8:

		return flattenInt8(SliceToWrite, data)

	case uint16:

		return flattenU16(SliceToWrite, data)

	case int16:

		return flattenInt16(SliceToWrite, data)

	case uint32:

		return flattenU32(SliceToWrite, data)

	case int32:

		return flattenInt32(SliceToWrite, data)

	case uint64:

		return flattenU64(SliceToWrite, data)

	case int64:

		return flattenInt64(SliceToWrite, data)

	case float32:

		return flattenFloat32(SliceToWrite, data)

	case float64:

		return flattenFloat64(SliceToWrite, data)

	case []bool:

		return flatten1DBool(SliceToWrite, data, true)

	case []uint8:

		return flatten1DU8(SliceToWrite, data, true)

	case []int8:

		return flatten1DInt8(SliceToWrite, data, true)

	case []uint16:

		return flatten1DU16(SliceToWrite, data, true)

	case []int16:

		return flatten1DInt16(SliceToWrite, data, true)

	case []uint32:

		return flatten1DU32(SliceToWrite, data, true)

	case []int32:

		return flatten1DInt32(SliceToWrite, data, true)

	case []uint64:

		return flatten1DU64(SliceToWrite, data, true)

	case []int64:

		return flatten1DInt64(SliceToWrite, data, true)

	case []float32:

		return flatten1DFloat32(SliceToWrite, data, true)

	case []float64:

		return flatten1DFloat64(SliceToWrite, data, true)

	case [][]bool:

		return flatten2DBool(SliceToWrite, data)

	case [][]uint8:

		return flatten2DU8(SliceToWrite, data)

	case [][]int8:

		return flatten2DInt8(SliceToWrite, data)

	case [][]uint16:

		return flatten2DU16(SliceToWrite, data)

	case [][]int16:

		return flatten2DInt16(SliceToWrite, data)

	case [][]uint32:

		return flatten2DU32(SliceToWrite, data)

	case [][]int32:

		return flatten2DInt32(SliceToWrite, data)

	case [][]uint64:

		return flatten2DU64(SliceToWrite, data)

	case [][]int64:

		return flatten2DInt64(SliceToWrite, data)

	case [][]float32:

		return flatten2DFloat32(SliceToWrite, data)

	case [][]float64:

		return flatten2DFloat64(SliceToWrite, data)

	case string:

		return flattenString(SliceToWrite, data)

	default:

		// Unknown Type

		return 0

	}

}



/**************************************************************************************************/

/* Unflatten To Byte Slice public function

/**************************************************************************************************/



//Unflatten data from slice and return it

func Unflatten(SliceToRead []byte, element interface{}) (SliceReaded []byte, elementReaded interface{}) {

	var size, row, column int

	elementReaded = element

	switch element.(type) {

	case bool:

		return unflattenBool(SliceToRead)

	case uint8:

		return unflattenU8(SliceToRead)

	case uint16:

		return unflattenU16(SliceToRead)

	case uint32:

		return unflattenU32(SliceToRead)

	case uint64:

		return unflattenU64(SliceToRead)

	case int8:

		return unflattenInt8(SliceToRead)

	case int16:

		return unflattenInt16(SliceToRead)

	case int32:

		return unflattenInt32(SliceToRead)

	case int64:

		return unflattenInt64(SliceToRead)

	case float32:

		return unflattenFloat32(SliceToRead)

	case float64:

		return unflattenFloat64(SliceToRead)

	case []bool:

		SliceToRead, size = UnflattenSize(SliceToRead)

		return unflatten1DBool(SliceToRead, size)

	case []uint8:

		SliceToRead, size = UnflattenSize(SliceToRead)

		return unflatten1DU8(SliceToRead, size)

	case []uint16:

		SliceToRead, size = UnflattenSize(SliceToRead)

		return unflatten1DU16(SliceToRead, size)

	case []uint32:

		SliceToRead, size = UnflattenSize(SliceToRead)

		return unflatten1DU32(SliceToRead, size)

	case []uint64:

		SliceToRead, size = UnflattenSize(SliceToRead)

		return unflatten1DU64(SliceToRead, size)

	case []int8:

		SliceToRead, size = UnflattenSize(SliceToRead)

		return unflatten1DInt8(SliceToRead, size)

	case []int16:

		SliceToRead, size = UnflattenSize(SliceToRead)

		return unflatten1DInt16(SliceToRead, size)

	case []int32:

		SliceToRead, size = UnflattenSize(SliceToRead)

		return unflatten1DInt32(SliceToRead, size)

	case []int64:

		SliceToRead, size = UnflattenSize(SliceToRead)

		return unflatten1DInt64(SliceToRead, size)

	case []float32:

		SliceToRead, size = UnflattenSize(SliceToRead)

		return unflatten1DFloat32(SliceToRead, size)

	case []float64:

		SliceToRead, size = UnflattenSize(SliceToRead)

		return unflatten1DFloat64(SliceToRead, size)

	case [][]bool:

		SliceToRead, row = UnflattenSize(SliceToRead)

		SliceToRead, column = UnflattenSize(SliceToRead)

		return unflatten2DBool(SliceToRead, row, column)

	case [][]uint8:

		SliceToRead, row = UnflattenSize(SliceToRead)

		SliceToRead, column = UnflattenSize(SliceToRead)

		return unflatten2DU8(SliceToRead, row, column)

	case [][]uint16:

		SliceToRead, row = UnflattenSize(SliceToRead)

		SliceToRead, column = UnflattenSize(SliceToRead)

		return unflatten2DU16(SliceToRead, row, column)

	case [][]uint32:

		SliceToRead, row = UnflattenSize(SliceToRead)

		SliceToRead, column = UnflattenSize(SliceToRead)

		return unflatten2DU32(SliceToRead, row, column)

	case [][]uint64:

		SliceToRead, row = UnflattenSize(SliceToRead)

		SliceToRead, column = UnflattenSize(SliceToRead)

		return unflatten2DU64(SliceToRead, row, column)

	case [][]int8:

		SliceToRead, row = UnflattenSize(SliceToRead)

		SliceToRead, column = UnflattenSize(SliceToRead)

		return unflatten2DInt8(SliceToRead, row, column)

	case [][]int16:

		SliceToRead, row = UnflattenSize(SliceToRead)

		SliceToRead, column = UnflattenSize(SliceToRead)

		return unflatten2DInt16(SliceToRead, row, column)

	case [][]int32:

		SliceToRead, row = UnflattenSize(SliceToRead)

		SliceToRead, column = UnflattenSize(SliceToRead)

		return unflatten2DInt32(SliceToRead, row, column)

	case [][]int64:

		SliceToRead, row = UnflattenSize(SliceToRead)

		SliceToRead, column = UnflattenSize(SliceToRead)

		return unflatten2DInt64(SliceToRead, row, column)

	case [][]float32:

		SliceToRead, row = UnflattenSize(SliceToRead)

		SliceToRead, column = UnflattenSize(SliceToRead)

		return unflatten2DFloat32(SliceToRead, row, column)

	case [][]float64:

		SliceToRead, row = UnflattenSize(SliceToRead)

		SliceToRead, column = UnflattenSize(SliceToRead)

		return unflatten2DFloat64(SliceToRead, row, column)

	case string:

		SliceToRead, size = UnflattenSize(SliceToRead)

		return unflattenString(SliceToRead, size)

	case []string:

		SliceToRead, size = UnflattenSize(SliceToRead)

		return unflatten1DString(SliceToRead, size)

	default:

		// Unknown Type

	}

	return SliceToRead, element

}



/**************************************************************************************************/

/* Flastten to Slice functions (Big Endian Only)

/**************************************************************************************************/



//FlattenSize write slice size before data slice

func FlattenSize(SliceToWrite []byte, element int) int {

	size := uint32(element)

	if len(SliceToWrite) >= 4 {

		SliceToWrite[3] = byte(size)

		SliceToWrite[2] = byte((size >> 8))

		SliceToWrite[1] = byte((size >> 16))

		SliceToWrite[0] = byte((size >> 24))

		return 4

	}

	return 0

}



//flattenBool flatten boolean into slice and return the write size, 0 is error

func flattenBool(SliceToWrite []byte, element bool) int {

	if len(SliceToWrite) > 0 {

		if element {

			SliceToWrite[0] = 1

		} else {

			SliceToWrite[0] = 0

		}

		return 1

	}

	return 0

}



//flattenU8 write U8 into slice and return the write size, 0 is error

func flattenU8(SliceToWrite []byte, element uint8) int {

	if len(SliceToWrite) >= 1 {

		SliceToWrite[0] = element

		return 1

	}

	return 0

}



//flattenU16 write U16 into slice and return the write size, 0 is error

func flattenU16(SliceToWrite []byte, element uint16) int {

	if len(SliceToWrite) >= 2 {

		SliceToWrite[1] = byte(element)

		SliceToWrite[0] = byte((element >> 8))

		return 2

	}

	return 0

}



//flattenU32 write U32 into slice and return the write size, 0 is error

func flattenU32(SliceToWrite []byte, element uint32) int {

	if len(SliceToWrite) >= 4 {

		SliceToWrite[3] = byte(element)

		SliceToWrite[2] = byte((element >> 8))

		SliceToWrite[1] = byte((element >> 16))

		SliceToWrite[0] = byte((element >> 24))

		return 4

	}

	return 0

}



//flattenU64 write Int64 into slice and return the write size, 0 is error

func flattenU64(SliceToWrite []byte, element uint64) int {

	if len(SliceToWrite) >= 8 {

		SliceToWrite[7] = byte(element)

		SliceToWrite[6] = byte((element >> 8))

		SliceToWrite[5] = byte((element >> 16))

		SliceToWrite[4] = byte((element >> 24))

		SliceToWrite[3] = byte((element >> 32))

		SliceToWrite[2] = byte((element >> 40))

		SliceToWrite[1] = byte((element >> 48))

		SliceToWrite[0] = byte((element >> 56))

		return 8

	}

	return 0

}



//flattenInt8 write U8 into slice and return the write size, 0 is error

func flattenInt8(SliceToWrite []byte, element int8) int {

	if len(SliceToWrite) >= 1 {

		SliceToWrite[0] = byte(element)

		return 1

	}

	return 0

}



//flattenInt16 write Int16 into slice and return the write size, 0 is error

func flattenInt16(SliceToWrite []byte, element int16) int {

	if len(SliceToWrite) >= 2 {

		SliceToWrite[1] = byte(element)

		SliceToWrite[0] = byte((element >> 8))

		return 2

	}

	return 0

}



//flattenInt32 write Int32 into slice and return the write size, 0 is error

func flattenInt32(SliceToWrite []byte, element int32) int {

	if len(SliceToWrite) >= 4 {

		SliceToWrite[3] = byte(element)

		SliceToWrite[2] = byte((element >> 8))

		SliceToWrite[1] = byte((element >> 16))

		SliceToWrite[0] = byte((element >> 24))

		return 4

	}

	return 0

}



//flattenInt64 write Int64 into slice and return the write size, 0 is error

func flattenInt64(SliceToWrite []byte, element int64) int {

	if len(SliceToWrite) >= 8 {

		SliceToWrite[7] = byte(element)

		SliceToWrite[6] = byte((element >> 8))

		SliceToWrite[5] = byte((element >> 16))

		SliceToWrite[4] = byte((element >> 24))

		SliceToWrite[3] = byte((element >> 32))

		SliceToWrite[2] = byte((element >> 40))

		SliceToWrite[1] = byte((element >> 48))

		SliceToWrite[0] = byte((element >> 56))

		return 8

	}

	return 0

}



//flattenFloat32 write float32 into slice and return the write size, 0 is error

func flattenFloat32(SliceToWrite []byte, element float32) int {

	if len(SliceToWrite) >= 4 {

		b := math.Float32bits(element)

		SliceToWrite[3] = byte(b)

		SliceToWrite[2] = byte((b >> 8))

		SliceToWrite[1] = byte((b >> 16))

		SliceToWrite[0] = byte((b >> 24))

		return 4

	}

	return 0

}



//flattenFloat64 write float64 into slice and return the write size, 0 is error

func flattenFloat64(SliceToWrite []byte, element float64) int {

	if len(SliceToWrite) >= 8 {

		b := math.Float64bits(element)

		SliceToWrite[7] = byte(b)

		SliceToWrite[6] = byte((b >> 8))

		SliceToWrite[5] = byte((b >> 16))

		SliceToWrite[4] = byte((b >> 24))

		SliceToWrite[3] = byte((b >> 32))

		SliceToWrite[2] = byte((b >> 40))

		SliceToWrite[1] = byte((b >> 48))

		SliceToWrite[0] = byte((b >> 56))

		return 8

	}

	return 0

}



//flatten1DBool write 1d Boold Slice to Bytes and return the write size, 0 is error

func flatten1DBool(SliceToWrite []byte, element []bool, writeSize bool) int {

	var index int

	size := len(element)

	if writeSize {

		index += FlattenSize(SliceToWrite, size)

	}

	for _, value := range element {

		index += flattenBool(SliceToWrite[index:], value)

	}

	return index

}



//flatten1DU8 write uint8 into slice and return the write size, 0 is error

func flatten1DU8(SliceToWrite []byte, element []uint8, writeSize bool) int {

	var index int

	size := len(element)

	if writeSize {

		index += FlattenSize(SliceToWrite, size)

	}

	for _, value := range element {

		index += flattenU8(SliceToWrite[index:], value)

	}

	return index

}



//flatten1DInt8 write int8 into slice and return the write size, 0 is error

func flatten1DInt8(SliceToWrite []byte, element []int8, writeSize bool) int {

	var index int

	size := len(element)

	if writeSize {

		index += FlattenSize(SliceToWrite, size)

	}

	for _, value := range element {

		index += flattenInt8(SliceToWrite[index:], value)

	}

	return index

}



//flatten1DU16 write uint16 into slice and return the write size, 0 is error

func flatten1DU16(SliceToWrite []byte, element []uint16, writeSize bool) int {

	var index int

	size := len(element)

	if writeSize {

		index += FlattenSize(SliceToWrite, size)

	}

	for _, value := range element {

		index += flattenU16(SliceToWrite[index:], value)

	}

	return index

}



//flatten1DInt16 write int16 into slice and return the write size, 0 is error

func flatten1DInt16(SliceToWrite []byte, element []int16, writeSize bool) int {

	var index int

	size := len(element)

	if writeSize {

		index += FlattenSize(SliceToWrite, size)

	}

	for _, value := range element {

		index += flattenInt16(SliceToWrite[index:], value)

	}

	return index

}



//flatten1DU32 write uint32 into slice and return the write size, 0 is error

func flatten1DU32(SliceToWrite []byte, element []uint32, writeSize bool) int {

	var index int

	size := len(element)

	if writeSize {

		index += FlattenSize(SliceToWrite, size)

	}

	for _, value := range element {

		index += flattenU32(SliceToWrite[index:], value)

	}

	return index

}



//flatten1DInt32 write int32 into slice and return the write size, 0 is error

func flatten1DInt32(SliceToWrite []byte, element []int32, writeSize bool) int {

	var index int

	size := len(element)

	if writeSize {

		index += FlattenSize(SliceToWrite, size)

	}

	for _, value := range element {

		index += flattenInt32(SliceToWrite[index:], value)

	}

	return index

}



//flatten1DU64 write uint64 into slice and return the write size, 0 is error

func flatten1DU64(SliceToWrite []byte, element []uint64, writeSize bool) int {

	var index int

	size := len(element)

	if writeSize {

		index += FlattenSize(SliceToWrite, size)

	}

	for _, value := range element {

		index += flattenU64(SliceToWrite[index:], value)

	}

	return index

}



//flatten1DInt64 write int64 into slice and return the write size, 0 is error

func flatten1DInt64(SliceToWrite []byte, element []int64, writeSize bool) int {

	var index int

	size := len(element)

	if writeSize {

		index += FlattenSize(SliceToWrite, size)

	}

	for _, value := range element {

		index += flattenInt64(SliceToWrite[index:], value)

	}

	return index

}



//flatten1DFloat32 write float32 into slice and return the write size, 0 is error

func flatten1DFloat32(SliceToWrite []byte, element []float32, writeSize bool) int {

	var index int

	size := len(element)

	if writeSize {

		index += FlattenSize(SliceToWrite, size)

	}

	for _, value := range element {

		index += flattenFloat32(SliceToWrite[index:], value)

	}

	return index

}



//flatten1DFloat64 write float64 into slice and return the write size, 0 is error

func flatten1DFloat64(SliceToWrite []byte, element []float64, writeSize bool) int {

	var index int

	size := len(element)

	if writeSize {

		index += FlattenSize(SliceToWrite, size)

	}

	for _, value := range element {

		index += flattenFloat64(SliceToWrite[index:], value)

	}

	return index

}



func flatten2DBool(SliceToWrite []byte, element [][]bool) int {

	var index int

	size := len(element)

	index += FlattenSize(SliceToWrite, size) // Write nb array elements

	temp := index

	for i, value := range element {

		index += flatten1DBool(SliceToWrite[index:], value, i == 0)

	}

	if temp == index {

		index += FlattenSize(SliceToWrite[index:], 0) // Write nb array elements

	}

	return index

}



func flatten2DU8(SliceToWrite []byte, element [][]uint8) int {

	var index int

	size := len(element)

	index += FlattenSize(SliceToWrite, size) // Write nb array elements

	temp := index

	for i, value := range element {

		index += flatten1DU8(SliceToWrite[index:], value, i == 0)

	}

	if temp == index {

		index += FlattenSize(SliceToWrite[index:], 0) // Write nb array elements

	}

	return index

}



//flatten2DInt8 write float32 into slice and return the write size, 0 is error

func flatten2DInt8(SliceToWrite []byte, element [][]int8) int {

	var index int

	size := len(element)

	index += FlattenSize(SliceToWrite, size) // Write nb array elements

	temp := index

	for i, value := range element {

		index += flatten1DInt8(SliceToWrite[index:], value, i == 0)

	}

	if temp == index {

		index += FlattenSize(SliceToWrite[index:], 0) // Write nb array elements

	}

	return index

}



//flatten2DU16 write float32 into slice and return the write size, 0 is error

func flatten2DU16(SliceToWrite []byte, element [][]uint16) int {

	var index int

	size := len(element)

	index += FlattenSize(SliceToWrite, size) // Write nb array elements

	temp := index

	for i, value := range element {

		index += flatten1DU16(SliceToWrite[index:], value, i == 0)

	}

	if temp == index {

		index += FlattenSize(SliceToWrite[index:], 0) // Write nb array elements

	}

	return index

}



//flatten2DInt16 write float32 into slice and return the write size, 0 is error

func flatten2DInt16(SliceToWrite []byte, element [][]int16) int {

	var index int

	size := len(element)

	index += FlattenSize(SliceToWrite, size) // Write nb array elements

	temp := index

	for i, value := range element {

		index += flatten1DInt16(SliceToWrite[index:], value, i == 0)

	}

	if temp == index {

		index += FlattenSize(SliceToWrite[index:], 0) // Write nb array elements

	}

	return index

}



//flatten2DU32 write uint32 into slice and return the write size, 0 is error

func flatten2DU32(SliceToWrite []byte, element [][]uint32) int {

	var index int

	size := len(element)

	index += FlattenSize(SliceToWrite, size) // Write nb array elements

	temp := index

	for i, value := range element {

		index += flatten1DU32(SliceToWrite[index:], value, i == 0)

	}

	if temp == index {

		index += FlattenSize(SliceToWrite[index:], 0) // Write nb array elements

	}

	return index

}



//flatten2DInt32 write int32 into slice and return the write size, 0 is error

func flatten2DInt32(SliceToWrite []byte, element [][]int32) int {

	var index int

	size := len(element)

	index += FlattenSize(SliceToWrite, size) // Write nb array elements

	temp := index

	for i, value := range element {

		index += flatten1DInt32(SliceToWrite[index:], value, i == 0)

	}

	if temp == index {

		index += FlattenSize(SliceToWrite[index:], 0) // Write nb array elements

	}

	return index

}



//flatten2DU64 write uint64 into slice and return the write size, 0 is error

func flatten2DU64(SliceToWrite []byte, element [][]uint64) int {

	var index int

	size := len(element)

	index += FlattenSize(SliceToWrite, size) // Write nb array elements

	temp := index

	for i, value := range element {

		index += flatten1DU64(SliceToWrite[index:], value, i == 0)

	}

	if temp == index {

		index += FlattenSize(SliceToWrite[index:], 0) // Write nb array elements

	}

	return index

}



//flatten2DInt64 write int64 into slice and return the write size, 0 is error

func flatten2DInt64(SliceToWrite []byte, element [][]int64) int {

	var index int

	size := len(element)

	index += FlattenSize(SliceToWrite, size) // Write nb array elements

	temp := index

	for i, value := range element {

		index += flatten1DInt64(SliceToWrite[index:], value, i == 0)

	}

	if temp == index {

		index += FlattenSize(SliceToWrite[index:], 0) // Write nb array elements

	}

	return index

}



//flatten2DFloat32 write float32 into slice and return the write size, 0 is error

func flatten2DFloat32(SliceToWrite []byte, element [][]float32) int {

	var index int

	size := len(element)

	index += FlattenSize(SliceToWrite, size) // Write nb array elements

	temp := index

	for i, value := range element {

		index += flatten1DFloat32(SliceToWrite[index:], value, i == 0)

	}

	if temp == index {

		index += FlattenSize(SliceToWrite[index:], 0) // Write nb array elements

	}

	return index

}



//flatten2DFloat64 write float32 into slice and return the write size, 0 is error

func flatten2DFloat64(SliceToWrite []byte, element [][]float64) int {

	var index int

	size := len(element)

	index += FlattenSize(SliceToWrite, size) // Write nb array elements

	temp := index

	for i, value := range element {

		index += flatten1DFloat64(SliceToWrite[index:], value, i == 0)

	}

	if temp == index {

		index += FlattenSize(SliceToWrite[index:], 0) // Write nb array elements

	}

	return index

}



//flattenString flatten string into slice and return the write size, 0 is error

func flattenString(SliceToWrite []byte, element string) int {

	var index int

	index = FlattenSize(SliceToWrite, len(element))

	if index == 4 {

		SliceToWrite = SliceToWrite[index:]

		for i, v := range []byte(element) {

			SliceToWrite[i] = v

		}

		index += len(element)

	}

	return index

}



/**************************************************************************************************/

/* Slice unflatten functions (Big Endian Only)

/**************************************************************************************************/



func UnflattenSize(SliceToRead []byte) (SliceReaded []byte, element int) {

	var temp uint32

	if len(SliceToRead) >= 4 {



		temp = uint32(SliceToRead[3])

		temp += uint32(SliceToRead[2]) << 8

		temp += uint32(SliceToRead[1]) << 16

		temp += uint32(SliceToRead[0]) << 24

		SliceToRead = SliceToRead[4:]

	}

	return SliceToRead, int(temp)

}



func unflattenBool(SliceToRead []byte) (SliceReaded []byte, element bool) {

	if len(SliceToRead) >= 1 {

		element = SliceToRead[0] != 0

		SliceToRead = SliceToRead[1:]

	}

	return SliceToRead, element

}



func unflattenU8(SliceToRead []byte) (SliceReaded []byte, element uint8) {

	if len(SliceToRead) >= 1 {

		element = uint8(SliceToRead[0])

		SliceToRead = SliceToRead[1:]

	}

	return SliceToRead, element

}



func unflattenInt8(SliceToRead []byte) (SliceReaded []byte, element int8) {

	if len(SliceToRead) >= 1 {

		element = int8(SliceToRead[0])

		SliceToRead = SliceToRead[1:]

	}

	return SliceToRead, element

}



func unflattenU16(SliceToRead []byte) (SliceReaded []byte, element uint16) {

	if len(SliceToRead) >= 2 {

		element = uint16(SliceToRead[1])

		element += uint16(SliceToRead[0]) << 8

		SliceToRead = SliceToRead[2:]

	}

	return SliceToRead, element

}



func unflattenInt16(SliceToRead []byte) (SliceReaded []byte, element int16) {

	if len(SliceToRead) >= 2 {

		element = int16(SliceToRead[1])

		element += int16(SliceToRead[0]) << 8

		SliceToRead = SliceToRead[2:]

	}

	return SliceToRead, element

}



func unflattenU32(SliceToRead []byte) (SliceReaded []byte, element uint32) {

	if len(SliceToRead) >= 4 {

		element = uint32(SliceToRead[3])

		element += uint32(SliceToRead[2]) << 8

		element += uint32(SliceToRead[1]) << 16

		element += uint32(SliceToRead[0]) << 24

		SliceToRead = SliceToRead[4:]

	}

	return SliceToRead, element

}



func unflattenInt32(SliceToRead []byte) (SliceReaded []byte, element int32) {

	if len(SliceToRead) >= 4 {

		element = int32(SliceToRead[3])

		element += int32(SliceToRead[2]) << 8

		element += int32(SliceToRead[1]) << 16

		element += int32(SliceToRead[0]) << 24

		SliceToRead = SliceToRead[4:]

	}

	return SliceToRead, element

}



func unflattenInt64(SliceToRead []byte) (SliceReaded []byte, element int64) {

	if len(SliceToRead) >= 8 {

		element = int64(SliceToRead[7])

		element += int64(SliceToRead[6]) << 8

		element += int64(SliceToRead[5]) << 16

		element += int64(SliceToRead[4]) << 24

		element += int64(SliceToRead[3]) << 32

		element += int64(SliceToRead[2]) << 40

		element += int64(SliceToRead[1]) << 48

		element += int64(SliceToRead[0]) << 56

		SliceToRead = SliceToRead[8:]

	}

	return SliceToRead, element

}



func unflattenU64(SliceToRead []byte) (SliceReaded []byte, element uint64) {

	if len(SliceToRead) >= 8 {

		element = uint64(SliceToRead[7])

		element += uint64(SliceToRead[6]) << 8

		element += uint64(SliceToRead[5]) << 16

		element += uint64(SliceToRead[4]) << 24

		element += uint64(SliceToRead[3]) << 32

		element += uint64(SliceToRead[2]) << 40

		element += uint64(SliceToRead[1]) << 48

		element += uint64(SliceToRead[0]) << 56

		SliceToRead = SliceToRead[8:]

	}

	return SliceToRead, element

}



func unflattenFloat32(SliceToRead []byte) (SliceReaded []byte, element float32) {

	if len(SliceToRead) >= 4 {

		temp := uint32(SliceToRead[3])

		temp += uint32(SliceToRead[2]) << 8

		temp += uint32(SliceToRead[1]) << 16

		temp += uint32(SliceToRead[0]) << 24

		element = math.Float32frombits(temp)

		SliceToRead = SliceToRead[4:]

	}

	return SliceToRead, element

}



func unflattenFloat64(SliceToRead []byte) (SliceReaded []byte, element float64) {

	if len(SliceToRead) >= 8 {

		temp := uint64(SliceToRead[7])

		temp += uint64(SliceToRead[6]) << 8

		temp += uint64(SliceToRead[5]) << 16

		temp += uint64(SliceToRead[4]) << 24

		temp += uint64(SliceToRead[3]) << 32

		temp += uint64(SliceToRead[2]) << 40

		temp += uint64(SliceToRead[1]) << 48

		temp += uint64(SliceToRead[0]) << 56

		element = math.Float64frombits(temp)

		SliceToRead = SliceToRead[8:]

	}

	return SliceToRead, element

}



func unflatten1DString(SliceToRead []byte, size int) (SliceReaded []byte, element []string) {

	element = make([]string, size)

	for i := range element {

		SliceToRead, size = UnflattenSize(SliceToRead)

		SliceToRead, element[i] = unflattenString(SliceToRead, size)

	}

	return SliceToRead, element

}



func unflatten1DBool(SliceToRead []byte, size int) (SliceReaded []byte, element []bool) {

	element = make([]bool, size)

	for i := range element {

		SliceToRead, element[i] = unflattenBool(SliceToRead)

	}

	return SliceToRead, element

}



func unflatten1DU8(SliceToRead []byte, size int) (SliceReaded []byte, element []uint8) {

	element = make([]uint8, size)

	for i := range element {

		SliceToRead, element[i] = unflattenU8(SliceToRead)

	}

	return SliceToRead, element

}



func unflatten1DU16(SliceToRead []byte, size int) (SliceReaded []byte, element []uint16) {

	element = make([]uint16, size)

	for i := range element {

		SliceToRead, element[i] = unflattenU16(SliceToRead)

	}

	return SliceToRead, element

}



func unflatten1DU32(SliceToRead []byte, size int) (SliceReaded []byte, element []uint32) {

	element = make([]uint32, size)

	for i := range element {

		SliceToRead, element[i] = unflattenU32(SliceToRead)

	}

	return SliceToRead, element

}



func unflatten1DU64(SliceToRead []byte, size int) (SliceReaded []byte, element []uint64) {

	element = make([]uint64, size)

	for i := range element {

		SliceToRead, element[i] = unflattenU64(SliceToRead)

	}

	return SliceToRead, element

}



func unflatten1DInt8(SliceToRead []byte, size int) (SliceReaded []byte, element []int8) {

	element = make([]int8, size)

	for i := range element {

		SliceToRead, element[i] = unflattenInt8(SliceToRead)

	}

	return SliceToRead, element

}



func unflatten1DInt16(SliceToRead []byte, size int) (SliceReaded []byte, element []int16) {

	element = make([]int16, size)

	for i := range element {

		SliceToRead, element[i] = unflattenInt16(SliceToRead)

	}

	return SliceToRead, element

}



func unflatten1DInt32(SliceToRead []byte, size int) (SliceReaded []byte, element []int32) {

	element = make([]int32, size)

	for i := range element {

		SliceToRead, element[i] = unflattenInt32(SliceToRead)

	}

	return SliceToRead, element

}



func unflatten1DInt64(SliceToRead []byte, size int) (SliceReaded []byte, element []int64) {

	element = make([]int64, size)

	for i := range element {

		SliceToRead, element[i] = unflattenInt64(SliceToRead)

	}

	return SliceToRead, element

}



func unflatten1DFloat32(SliceToRead []byte, size int) (SliceReaded []byte, element []float32) {

	element = make([]float32, size)

	for i := range element {

		SliceToRead, element[i] = unflattenFloat32(SliceToRead)

	}

	return SliceToRead, element

}



func unflatten1DFloat64(SliceToRead []byte, size int) (SliceReaded []byte, element []float64) {

	element = make([]float64, size)

	for i := range element {

		SliceToRead, element[i] = unflattenFloat64(SliceToRead)

	}

	return SliceToRead, element

}



func unflatten2DBool(SliceToRead []byte, row int, columns int) (SliceReaded []byte, element [][]bool) {

	element = make([][]bool, row)

	for i := range element {

		SliceToRead, element[i] = unflatten1DBool(SliceToRead, columns)

	}

	return SliceToRead, element

}



func unflatten2DU8(SliceToRead []byte, row int, columns int) (SliceReaded []byte, element [][]uint8) {

	element = make([][]uint8, row)

	for i := range element {

		SliceToRead, element[i] = unflatten1DU8(SliceToRead, columns)

	}

	return SliceToRead, element

}



func unflatten2DU16(SliceToRead []byte, row int, columns int) (SliceReaded []byte, element [][]uint16) {

	element = make([][]uint16, row)

	for i := range element {

		SliceToRead, element[i] = unflatten1DU16(SliceToRead, columns)

	}

	return SliceToRead, element

}



func unflatten2DU32(SliceToRead []byte, row int, columns int) (SliceReaded []byte, element [][]uint32) {

	element = make([][]uint32, row)

	for i := range element {

		SliceToRead, element[i] = unflatten1DU32(SliceToRead, columns)

	}

	return SliceToRead, element

}



func unflatten2DU64(SliceToRead []byte, row int, columns int) (SliceReaded []byte, element [][]uint64) {

	element = make([][]uint64, row)

	for i := range element {

		SliceToRead, element[i] = unflatten1DU64(SliceToRead, columns)

	}

	return SliceToRead, element

}



func unflatten2DInt8(SliceToRead []byte, row int, columns int) (SliceReaded []byte, element [][]int8) {

	element = make([][]int8, row)

	for i := range element {

		SliceToRead, element[i] = unflatten1DInt8(SliceToRead, columns)

	}

	return SliceToRead, element

}



func unflatten2DInt16(SliceToRead []byte, row int, columns int) (SliceReaded []byte, element [][]int16) {

	element = make([][]int16, row)

	for i := range element {

		SliceToRead, element[i] = unflatten1DInt16(SliceToRead, columns)

	}

	return SliceToRead, element

}



func unflatten2DInt32(SliceToRead []byte, row int, columns int) (SliceReaded []byte, element [][]int32) {

	element = make([][]int32, row)

	for i := range element {

		SliceToRead, element[i] = unflatten1DInt32(SliceToRead, columns)

	}

	return SliceToRead, element

}



func unflatten2DInt64(SliceToRead []byte, row int, columns int) (SliceReaded []byte, element [][]int64) {

	element = make([][]int64, row)

	for i := range element {

		SliceToRead, element[i] = unflatten1DInt64(SliceToRead, columns)

	}

	return SliceToRead, element

}



func unflatten2DFloat32(SliceToRead []byte, row int, columns int) (SliceReaded []byte, element [][]float32) {

	element = make([][]float32, row)

	for i := range element {

		SliceToRead, element[i] = unflatten1DFloat32(SliceToRead, columns)

	}

	return SliceToRead, element

}



func unflatten2DFloat64(SliceToRead []byte, row int, columns int) (SliceReaded []byte, element [][]float64) {

	element = make([][]float64, row)

	for i := range element {

		SliceToRead, element[i] = unflatten1DFloat64(SliceToRead, columns)

	}

	return SliceToRead, element

}



func unflattenString(SliceToRead []byte, size int) (SliceReaded []byte, element string) {

	element = string(SliceToRead[:size])

	SliceToRead = SliceToRead[size:]

	return SliceToRead, element

}



/**************************************************************************************************/

/* Slice flatten special struct functions (Big Endian Only)

/**************************************************************************************************/
