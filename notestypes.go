package domigo

import (
	"time"

	"lukechampine.com/uint128"
)

type Boolean = bool
type Byte = byte
type Integer = int16
type Long = int32
type Single = float32
type Double = float64
type Currency = float64
type String = string
type Variant = uint128.Uint128
type Time = time.Time

type primitiveType interface {
	Boolean | Byte | Integer | Long | Single | Double | String | Variant | Time
}
