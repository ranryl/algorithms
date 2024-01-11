package geohash

import (
	"fmt"
	"testing"
)

func TestBase32ToIndex(t *testing.T) {
	data := []struct {
		k byte
		v int
	}{
		{k: '0', v: 0},
		{k: '9', v: 9},
		{k: 'b', v: 10},
		{k: 'j', v: 17},
		{k: 'm', v: 19},
		{k: 't', v: 25},
		{k: 'z', v: 31},
	}
	for _, value := range data {
		result := Base32ToIndex(value.k)
		if value.v != result {
			t.Errorf("err data: %d, should be %d", result, value.v)
		}
	}
}
func TestGeoService(t *testing.T) {
	gsv := GeoService{
		root: &TrieNode{},
	}
	hashValue := gsv.Add(122.90900, 58.0009932, "北京", 12)
	fmt.Println(hashValue)
	fmt.Println(gsv.ListByPrefix(hashValue[0:8]))
}
