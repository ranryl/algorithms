package geohash

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	fmt.Println(Encode(122.90900, 58.0009932, 12))
	fmt.Println(Encode(50.90900, 18.0009932, 12))
	fmt.Println(Encode(10.90900, 32.0009932, 12))
	fmt.Println(Encode(18.9923900, -18.009932, 12))
	fmt.Println(Encode(-50.90900, 18.0009932, 12))
}
