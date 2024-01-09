package bitmap

import (
	"testing"
)

func TestBitMap(t *testing.T) {
	bitmap := NewBitMap(10000)
	bitmap.Set(100)
	t.Log(bitmap.IsExist(100))
	bitmap.Set(200)
	t.Log(bitmap.IsExist(200))
	bitmap.Set(1)
	t.Log(bitmap.IsExist(1))
	bitmap.Del(1)
	t.Log(bitmap.IsExist(1))
}
