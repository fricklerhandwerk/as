// +build 386 arm

package as_test

import (
	"math"
	"testing"

	"as"
)

func TestInt32(t *testing.T) {
	assertNoError(t, as.Int32, int8(math.MinInt8))
	assertNoError(t, as.Int32, int8(math.MaxInt8))

	assertNoError(t, as.Int32, int16(math.MinInt16))
	assertNoError(t, as.Int32, int16(math.MaxInt16))

	assertNoError(t, as.Int32, int32(math.MinInt32))
	assertNoError(t, as.Int32, int32(math.MaxInt32))

	assertError(t, as.Int32, int64(math.MinInt64))
	assertError(t, as.Int32, int64(math.MaxInt64))

	assertNoError(t, as.Int32, int(math.MinInt32))
	assertNoError(t, as.Int32, int(math.MaxInt32))

	assertNoError(t, as.Int32, uint8(0))
	assertNoError(t, as.Int32, uint8(math.MaxUint8))

	assertNoError(t, as.Int32, uint16(0))
	assertNoError(t, as.Int32, uint16(math.MaxUint16))

	assertNoError(t, as.Int32, uint32(0))
	assertError(t, as.Int32, uint32(math.MaxUint32))

	assertNoError(t, as.Int32, uint64(0))
	assertError(t, as.Int32, uint64(math.MaxUint64))

	assertNoError(t, as.Int32, uint(0))
	assertError(t, as.Int32, uint(math.MaxUint32))
}