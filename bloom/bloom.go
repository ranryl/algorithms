package bloom

import "algorithms/bitmap"

type Bloom struct {
	bitMap   bitmap.BitMap
	hashList []BloomFilterHash
}

type BloomFilterHash func(data []byte) uint64

func NewBloom(lens uint64, hashFuncs []BloomFilterHash) *Bloom {
	return &Bloom{
		bitMap:   *bitmap.NewBitMap(lens),
		hashList: hashFuncs,
	}
}

func (b *Bloom) Add(data []byte) {
	bitMapLens := b.bitMap.Len()
	for _, v := range b.hashList {
		hashValue := v(data) % bitMapLens
		b.bitMap.Set(hashValue)
	}
}
func (b *Bloom) IsExsit(data []byte) bool {
	bitMapLens := b.bitMap.Len()
	for _, v := range b.hashList {
		hashValue := v(data) % bitMapLens
		if !b.bitMap.IsExist(hashValue) {
			return false
		}
	}
	return true
}
