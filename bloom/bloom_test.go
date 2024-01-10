package bloom

import (
	"fmt"
	"testing"

	"github.com/spaolacci/murmur3"
)

func TestBloom(t *testing.T) {
	bloom := NewBloom(1000000, []BloomFilterHash{murmur3.Sum64})
	data := [][]byte{[]byte("aaaa"), []byte("bbbb"), []byte("ccccc"), []byte("ddddd"), []byte("aaaa"), []byte("bbbb")}
	for _, v := range data {
		isExist := bloom.IsExsit(v)
		bloom.Add(v)
		if !isExist {
			fmt.Printf("%s is not exists\n", v)
		} else {
			fmt.Printf("%s is exists\n", v)
		}
	}
}
